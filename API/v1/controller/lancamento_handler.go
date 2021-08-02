package controller

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/dao"
	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/helper"
	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/model/detalhe_lancamento"
	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/model/lancamento"

	"github.com/gorilla/mux"
)

// LancamentoPersJSON (lancamento personalizado em formato JSON) é a união dos modelos Lancamento e DetalheLancamento para simplificar o JSON enviado e recebido deste handler. Ao receber um JSON deste tipo, é desempacotado 1 lancamento e 2 detalhes lancamento e enviado para o BD. Para o envio, é feito o processo inverso
type LancamentoPersJSON struct {
	ID               int       `json:"id"`
	CpfPessoa        string    `json:"cpf_pessoa"`
	NomeContaOrigem  string    `json:"nome_conta_origem"`
	Data             time.Time `json:"data"`
	Numero           string    `json:"numero"`
	Descricao        string    `json:"descricao"`
	NomeContaDestino string    `json:"nome_conta_destino"`
	Debito           float64   `json:"debito"`
	Credito          float64   `json:"credito"`
	DataCriacao      time.Time `json:"data_criacao"`
	DataModificacao  time.Time `json:"data_modificacao"`
	Estado           bool      `json:"estado"`
}

// LancamentosPersJSON é um slice que representa um conjunto/lista de LancamentoPersJSON
type LancamentosPersJSON []*LancamentoPersJSON

// Len retorna a quantidade de itens(int) de LancamentoPersJSON no slice LancamentosPersJSON. Implementado por exigência da interface handler.Dados, usada nas função DefineHeaderRetornoDados
func (lpj LancamentosPersJSON) Len() int {
	return len(lpj)
}

// dados somente é usado para agrupar em um slice tipos diferentes nas funções deste handler
type dados []interface{}

// Len é um método obrigatório para cumprir com a interface handler.Dados
func (d dados) Len() int {
	return len(d)
}

// LancamentoIndex é um handler/controller que responde a rota '[GET] /lancamentos' e retorna StatusOK(200) e uma listagem de lancamentos de acordo com o tipo de usuário(admin/comum) caso o TOKEN informado for válido e o usuário associado ao token for cadastrado na API/DB. Caso ocorra algum erro, retorna StatusInternalServerError(500). Quando solicitado como usuário comum, retorna somente lancamentos ativos ref a esse usuário(cpf), enquanto que como administrador, retorna todos os registros de lancamentos ref ao usuário
func LancamentoIndex(w http.ResponseWriter, r *http.Request) {
	db := dao.GetDB()
	defer db.Close()

	var status = http.StatusInternalServerError // 500

	token, err := helper.GetToken(r, GetMySigningKey())
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	usuarioToken, _, admin, err := helper.GetClaims(token)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	pessoa, err := dao.ProcuraPessoaPorUsuario(db, usuarioToken)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	var listaLancamentosPersJSON LancamentosPersJSON
	if admin {
		listaLancamentosPersJSON, err = empacotaParaLancPersJSONPorCpf(db, pessoa.Cpf, true)
		err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
		if err != nil {
			return
		}
	} else {
		listaLancamentosPersJSON, err = empacotaParaLancPersJSONPorCpf(db, pessoa.Cpf, false)
		err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
		if err != nil {
			return
		}
	}

	status = http.StatusOK
	funcao := "LancamentoIndex"
	DefineHeaderRetornoDados(
		w,
		SetHeaderJSON,
		status,
		listaLancamentosPersJSON,
		funcao,
		"Listagem de lançamentos",
		"Enviando listagem de lançamentos")
}

// LancamentoShow é um handler/controller que responde a rota '[GET] /lancamentos/{lancamento}' e retorna StatusOK(200) e os dados do lancamento(ID) solicitado caso o TOKEN informado for válido e o usuário associado ao token for cadastrado na API/DB. Caso ocorra algum erro, retorna StatusInternalServerError(500)
func LancamentoShow(w http.ResponseWriter, r *http.Request) {
	db := dao.GetDB()
	defer db.Close()

	var status = http.StatusInternalServerError // 500

	vars := mux.Vars(r)
	lancamentoRota := vars["lancamento"]

	token, err := helper.GetToken(r, GetMySigningKey())
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	usuarioToken, _, admin, err := helper.GetClaims(token)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	usuarioDB, err := dao.ProcuraPessoaPorUsuario(db, usuarioToken)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	idLancamentoRota, err := converteParaIDLancamento(lancamentoRota)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	lancamentoEncontrado, err := dao.ProcuraLancamento(db, idLancamentoRota)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	if !lancamentoEncontrado.Estado {
		verif := !(admin && usuarioDB.Administrador)
		err = DefineHeaderRetorno(w, SetHeaderJSON, verif, status, errors.New("Somente administradores podem consultar um lancamento inativo"))
		if err != nil {
			return
		}
	}

	lancamentosConvertidos, err := empacotaParaLancamentoPersJSON(db, lancamentoEncontrado)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	lancamentoJSON := lancamentosConvertidos[0]
	status = http.StatusOK // 200
	funcao := "LancamentoShow"
	DefineHeaderRetornoDado(
		w,
		SetHeaderJSON,
		status,
		lancamentoJSON,
		funcao,
		fmt.Sprintf("Dados de lancamento '%d'", lancamentoJSON.ID),
		fmt.Sprintf("Enviando dados de lancamento '%d'", lancamentoJSON.ID))
}

// LancamentoCreate é um handler/controller que responde a rota '[POST] /lancamentos' e retorna StatusCreated(201) e os dados de um lancamento e seus detalhes lancamento criados através das informações informadas via JSON(body) caso o TOKEN informado for válido e o usuário associado ao token for cadastrado na API/DB. Caso ocorra algum erro, retorna StatusInternalServerError(500) ou StatusUnprocessableEntity(422) caso as informações no JSON não corresponderem ao formato {"cpf_pessoa":"?",  "nome_conta_origem":"?", "data":"?", "numero":"?", "descricao":"?", "nome_conta_destino":"?", "debito":?, "credito":?}
func LancamentoCreate(w http.ResponseWriter, r *http.Request) {
	db := dao.GetDB()
	defer db.Close()

	var status = http.StatusInternalServerError // 500
	var lancamentoFromJSON LancamentoPersJSON
	var novoLancamento *lancamento.Lancamento
	var novoDetalheLancamento1, novoDetalheLancamento2 *detalhe_lancamento.DetalheLancamento

	token, err := helper.GetToken(r, GetMySigningKey())
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	usuarioToken, _, _, err := helper.GetClaims(token)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	_, err = dao.ProcuraPessoaPorUsuario(db, usuarioToken)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, LimitData))
	if err != nil {
		log.Println(err)
	}

	if err := r.Body.Close(); err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(body, &lancamentoFromJSON)
	status = http.StatusUnprocessableEntity // 422
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	novoLancamento, novoDetalheLancamento1, novoDetalheLancamento2, err = desempacotaDeLancPersJSON(lancamentoFromJSON)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	l, err := dao.AdicionaLancamento(db, novoLancamento)
	status = http.StatusInternalServerError // 500
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	novoDetalheLancamento1.IDLancamento = l.ID
	dl1, err := dao.AdicionaDetalheLancamento(db, novoDetalheLancamento1)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	novoDetalheLancamento2.IDLancamento = l.ID
	dl2, err := dao.AdicionaDetalheLancamento(db, novoDetalheLancamento2)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	d := dados{l, dl1, dl2}
	status = http.StatusCreated // 201
	funcao := "LancamentoCreate"
	DefineHeaderRetornoDados(
		w,
		SetHeaderJSON,
		status,
		d,
		funcao,
		fmt.Sprintf("Dados de lancamento '%d' e seus detalhes lancamento", l.ID),
		fmt.Sprintf("Enviando dados de lancamento '%d' e seus detalhes lancamento", l.ID))

}

// LancamentoRemove é um handler/controller que responde a rota '[DELETE] /lancamentos/{lancamento}' e retorna StatusOK(200) e uma mensagem de confirmação caso o TOKEN informado for válido, o usuário associado ao token for cadastrado na API/DB e seja um administrador, que a conta informada esteja cadastrado no BD. Caso ocorra algum erro, retorna StatusInternalServerError(500)
func LancamentoRemove(w http.ResponseWriter, r *http.Request) {
	db := dao.GetDB()
	defer db.Close()

	var status = http.StatusInternalServerError // 500

	vars := mux.Vars(r)
	lancamentoRemocao := vars["lancamento"]

	token, err := helper.GetToken(r, GetMySigningKey())
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	usuarioToken, _, admin, err := helper.GetClaims(token)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	usuarioDB, err := dao.ProcuraPessoaPorUsuario(db, usuarioToken)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	verif := !(admin && usuarioDB.Administrador)
	err = DefineHeaderRetorno(w, SetHeaderJSON, verif, status, errors.New("Somente administradores podem usar essa rota"))
	if err != nil {
		return
	}

	idLancamentoRemocao, err := converteParaIDLancamento(lancamentoRemocao)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	err = dao.RemoveLancamento(db, idLancamentoRemocao)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	status = http.StatusOK // 200
	funcao := "LancamentoRemove"
	DefineHeaderRetornoDado(
		w,
		SetHeaderJSON,
		status,
		lancamentoRemocao,
		funcao,
		fmt.Sprintf("Removido lancamento '%s'", lancamentoRemocao),
		fmt.Sprintf("Enviando resposta de remoção de lancamento '%s'", lancamentoRemocao))
}

// LancamentoAlter é um handler/controller que responde a rota '[PUT] /lancamentos/{lancamento}' e retorna StatusOK(200) e uma mensagem de confirmação com os dados do lancamento alterado caso o TOKEN informado for válido, o usuário associado ao token for cadastrado na API/DB e o lancamento informado na rota existir. Caso ocorra algum erro, retorna StatusInternalServerError(500) ou StatusUnprocessableEntity(422), caso o JSON não seguir o formato {"cpf_pessoa":"?", "nome_conta_origem":"?", "data":"?", "numero":"?", "descricao":"?", "nome_conta_destino":"?", "debito":?, "credito":?}, ou StatusNotModified(304) caso ocorra algum erro na alteração do BD.
func LancamentoAlter(w http.ResponseWriter, r *http.Request) {
	db := dao.GetDB()
	defer db.Close()

	var status = http.StatusInternalServerError // 500
	var lancamentoFromJSON LancamentoPersJSON
	var lancamentoAlteracao *lancamento.Lancamento
	var detalheLancamentoAlteracao1, detalheLancamentoAlteracao2 *detalhe_lancamento.DetalheLancamento

	vars := mux.Vars(r)
	lancamentoAlteracaoID := vars["lancamento"]
	origemAlteracao := vars["origem"]
	destinoAlteracao := vars["destino"]

	token, err := helper.GetToken(r, GetMySigningKey())
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	usuarioToken, _, _, err := helper.GetClaims(token)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	_, err = dao.ProcuraPessoaPorUsuario(db, usuarioToken)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, LimitData))
	if err != nil {
		log.Println(err)
	}

	if err := r.Body.Close(); err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(body, &lancamentoFromJSON)
	status = http.StatusUnprocessableEntity // 422
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	lancamentoAlteracao, detalheLancamentoAlteracao1, detalheLancamentoAlteracao2, err = desempacotaDeLancPersJSON(lancamentoFromJSON)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	idLancamentoAlteracao, err := converteParaIDLancamento(lancamentoAlteracaoID)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	transacao, err := db.Begin()
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	l, err := dao.AlteraLancamento2(db, transacao, idLancamentoAlteracao, lancamentoAlteracao)
	status = http.StatusInternalServerError // 500
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		transacao.Rollback()
		return
	}

	detalheLancamentoAlteracao1.IDLancamento = l.ID
	dl1, err := dao.AlteraDetalheLancamento2(db, transacao, l.ID, origemAlteracao, detalheLancamentoAlteracao1)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		transacao.Rollback()
		return
	}

	detalheLancamentoAlteracao2.IDLancamento = l.ID
	dl2, err := dao.AlteraDetalheLancamento2(db, transacao, l.ID, destinoAlteracao, detalheLancamentoAlteracao2)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		transacao.Rollback()
		return
	}

	err = transacao.Commit()
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	d := dados{l, dl1, dl2}
	status = http.StatusOK // 200
	funcao := "LancamentoAlter"
	DefineHeaderRetornoDado(
		w,
		SetHeaderJSON,
		status,
		d,
		funcao,
		fmt.Sprintf("Novos dados de lancamento '%d'", l.ID),
		fmt.Sprintf("Enviando novos dados de lancamento '%d'", l.ID))
}

// LancamentoEstado é um handler/controller que responde a rota '[PUT] /lancamentos/{lancamento}/estado' e retorna StatusOK(200) e uma mensagem de confirmação com os dados do lancamento alterado caso o TOKEN informado for válido, o usuário associado ao token for cadastrado na API/DB e o lancamento informado na rota existir. Somente usuários ADMINISTRADORES podem ATIVAR contas, USUÁRIO COMUNS podem somente INATIVAR. Caso ocorra algum erro, retorna StatusInternalServerError(500), StatusUnprocessableEntity(422), caso o JSON não seguir o formato {"estado": ?}, StatusNotModified(304) caso ocorra algum erro na alteração do BD ou StatusNotFound(404) caso o lancamento informado na rota não existir
func LancamentoEstado(w http.ResponseWriter, r *http.Request) {
	db := dao.GetDB()
	defer db.Close()

	var status = http.StatusInternalServerError // 500
	var estadoLancamento estado

	vars := mux.Vars(r)
	lancamentoAlteracao := vars["lancamento"]

	token, err := helper.GetToken(r, GetMySigningKey())
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	usuarioToken, _, admin, err := helper.GetClaims(token)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	usuarioDB, err := dao.ProcuraPessoaPorUsuario(db, usuarioToken)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, LimitData))
	if err != nil {
		log.Println(err)
	}

	if err := r.Body.Close(); err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(body, &estadoLancamento)
	status = http.StatusUnprocessableEntity // 422
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	usuarioComum := !(admin && usuarioDB.Administrador)
	if usuarioComum {
		verif := estadoLancamento.Estado

		status = http.StatusInternalServerError // 500
		err = DefineHeaderRetorno(w, SetHeaderJSON, verif, status, errors.New("Somente administradores podem ativar um lancamento"))
		if err != nil {
			return
		}
	}

	idLancamentoAlteracao, err := converteParaIDLancamento(lancamentoAlteracao)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	lancamentoDBAlteracao, err := dao.ProcuraLancamento(db, idLancamentoAlteracao)
	status = http.StatusNotFound // 404
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	var alteraEstado func(*sql.DB, int) (*lancamento.Lancamento, error)
	if estadoLancamento.Estado {
		alteraEstado = dao.AtivaLancamento
	} else {
		alteraEstado = dao.InativaLancamento
	}
	l, err := alteraEstado(db, lancamentoDBAlteracao.ID)
	status = http.StatusNotModified // 304
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	status = http.StatusOK // 200
	funcao := "LancamentoEstado"
	DefineHeaderRetornoDado(
		w,
		SetHeaderJSON,
		status,
		l,
		funcao,
		fmt.Sprintf("Novos dados de lancamento '%d'", l.ID),
		fmt.Sprintf("Enviando novos dados de lancamento '%d'", l.ID))
}

// LancamentoPorConta é um handler/controller que responde a rota '[GET] /lancamentos_conta/{conta}' e retorna StatusOK(200) e uma listagem de lancamentos de acordo a conta informada e com o tipo de usuário(admin/comum) caso o TOKEN informado for válido e o usuário associado ao token for cadastrado na API/DB. Caso ocorra algum erro, retorna StatusInternalServerError(500). Quando solicitado como usuário comum, retorna somente lancamentos ativos ref a esse usuário(cpf), enquanto que como administrador, retorna todos os registros de lancamentos ref ao usuário
func LancamentoPorConta(w http.ResponseWriter, r *http.Request) {
	db := dao.GetDB()
	defer db.Close()

	var status = http.StatusInternalServerError // 500

	vars := mux.Vars(r)
	contaRota := vars["conta"]

	token, err := helper.GetToken(r, GetMySigningKey())
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	usuarioToken, _, admin, err := helper.GetClaims(token)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	pessoa, err := dao.ProcuraPessoaPorUsuario(db, usuarioToken)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	var listaLancamentosPersJSON LancamentosPersJSON
	if admin {
		listaLancamentosPersJSON, err = empacotaParaLancPersJSONPorCpfEConta(db, pessoa.Cpf, contaRota, true)
		err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
		if err != nil {
			return
		}
	} else {
		listaLancamentosPersJSON, err = empacotaParaLancPersJSONPorCpfEConta(db, pessoa.Cpf, contaRota, false)
		err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
		if err != nil {
			return
		}
	}

	status = http.StatusOK
	funcao := "LancamentoPorConta"
	DefineHeaderRetornoDados(
		w,
		SetHeaderJSON,
		status,
		listaLancamentosPersJSON,
		funcao,
		fmt.Sprintf("Listagem de lançamentos por conta %s", contaRota),
		fmt.Sprintf("Enviando listagem de lançamentos por conta %s", contaRota))
}

func desempacotaDeLancPersJSON(lancamentoFromJSON LancamentoPersJSON) (lanc *lancamento.Lancamento, detLancamento1, detLancamento2 *detalhe_lancamento.DetalheLancamento, err error) {
	lanc, err = lancamento.NewLancamento(
		0,
		lancamentoFromJSON.CpfPessoa,
		lancamentoFromJSON.Data,
		lancamentoFromJSON.Numero,
		lancamentoFromJSON.Descricao)
	if err != nil {
		return
	}

	detLancamento1, err = detalhe_lancamento.NewDetalheLancamento(
		0,
		lancamentoFromJSON.NomeContaOrigem,
		lancamentoFromJSON.Debito,
		lancamentoFromJSON.Credito)
	if err != nil {
		return
	}

	detLancamento2, err = detalhe_lancamento.NewDetalheLancamento(
		0,
		lancamentoFromJSON.NomeContaDestino,
		lancamentoFromJSON.Credito,
		lancamentoFromJSON.Debito)
	if err != nil {
		return
	}

	return
}

func empacotaParaLancPersJSONPorCpf(db *sql.DB, cpfPessoa string, todosOsRegistros bool) (listaLancamentosPersJSON LancamentosPersJSON, err error) {
	var listaLancamentos lancamento.Lancamentos
	if todosOsRegistros {
		listaLancamentos, err = dao.CarregaLancamentosPorCpf(db, cpfPessoa)
	} else {
		listaLancamentos, err = dao.CarregaLancamentosAtivoPorCpf(db, cpfPessoa)
	}
	if err != nil {
		return
	}

	return empacotaParaLancamentoPersJSON(db, listaLancamentos...)
}

func empacotaParaLancPersJSONPorCpfEConta(db *sql.DB, cpfPessoa, conta string, todosOsRegistros bool) (listaLancamentosPersJSON LancamentosPersJSON, err error) {
	var listaLancamentos lancamento.Lancamentos
	if todosOsRegistros {
		listaLancamentos, err = dao.CarregaLancamentosPorCpfEConta(db, cpfPessoa, conta)
	} else {
		listaLancamentos, err = dao.CarregaLancamentosAtivoPorCpfEConta(db, cpfPessoa, conta)
	}
	if err != nil {
		return
	}

	return empacotaParaLancamentoPersJSONPorConta(db, conta, listaLancamentos...)
}

func empacotaParaLancamentoPersJSON(db *sql.DB, lancamentos ...*lancamento.Lancamento) (listaLancamentosPersJSON LancamentosPersJSON, err error) {
	for _, lancamento := range lancamentos {
		lancPersJSON := new(LancamentoPersJSON)

		detalheLancamentos, err := dao.CarregaDetalheLancamentosPorIDLancamento(db, lancamento.ID)
		// todo o lancamento deve ter 2 detalhes lancamento obrigatóriamente. Caso contrário, o laço é quebrado e retorna o erro personalizado
		if err != nil && len(detalheLancamentos) != 2 {
			err = errors.New("Lancamentos devem ser obrigatoriamente em par, representando o valor em débito e crédito de cada conta")
			break
		}

		lancPersJSON.ID = lancamento.ID
		lancPersJSON.CpfPessoa = lancamento.CpfPessoa
		lancPersJSON.NomeContaOrigem = detalheLancamentos[0].NomeConta
		lancPersJSON.Data = lancamento.Data
		lancPersJSON.Numero = lancamento.Numero
		lancPersJSON.Descricao = lancamento.Descricao
		lancPersJSON.NomeContaDestino = detalheLancamentos[1].NomeConta
		lancPersJSON.Debito = detalheLancamentos[0].Debito
		lancPersJSON.Credito = detalheLancamentos[0].Credito
		lancPersJSON.DataCriacao = lancamento.DataCriacao
		lancPersJSON.DataModificacao = lancamento.DataModificacao
		lancPersJSON.Estado = lancamento.Estado

		listaLancamentosPersJSON = append(listaLancamentosPersJSON, lancPersJSON)
	}

	return
}

func empacotaParaLancamentoPersJSONPorConta(db *sql.DB, conta string, lancamentos ...*lancamento.Lancamento) (listaLancamentosPersJSON LancamentosPersJSON, err error) {
	for _, lancamento := range lancamentos {
		lancPersJSON := new(LancamentoPersJSON)

		detalheLancamentos, err := dao.CarregaDetalheLancamentosPorIDLancamento(db, lancamento.ID)
		// todo o lancamento deve ter 2 detalhes lancamento obrigatóriamente. Caso contrário, o laço é quebrado e retorna o erro personalizado
		if err != nil && len(detalheLancamentos) != 2 {
			err = errors.New("Lancamentos devem ser obrigatoriamente em par, representando o valor em débito e crédito de cada conta")
			break
		}

		lancPersJSON.ID = lancamento.ID
		lancPersJSON.CpfPessoa = lancamento.CpfPessoa
		lancPersJSON.NomeContaOrigem = detalheLancamentos[0].NomeConta
		lancPersJSON.Data = lancamento.Data
		lancPersJSON.Numero = lancamento.Numero
		lancPersJSON.Descricao = lancamento.Descricao
		lancPersJSON.NomeContaDestino = detalheLancamentos[1].NomeConta

		if conta == lancPersJSON.NomeContaOrigem {
			lancPersJSON.Debito = detalheLancamentos[0].Debito
			lancPersJSON.Credito = detalheLancamentos[0].Credito
		} else {
			lancPersJSON.Debito = detalheLancamentos[1].Debito
			lancPersJSON.Credito = detalheLancamentos[1].Credito
		}

		lancPersJSON.DataCriacao = lancamento.DataCriacao
		lancPersJSON.DataModificacao = lancamento.DataModificacao
		lancPersJSON.Estado = lancamento.Estado

		listaLancamentosPersJSON = append(listaLancamentosPersJSON, lancPersJSON)
	}

	return
}

func converteParaIDLancamento(codigo string) (numero int, err error) {
	numero, err = strconv.Atoi(codigo)
	if err != nil {
		err = fmt.Errorf("ID de lançamento informado inválido[%s]", codigo)
		return
	}
	return
}
