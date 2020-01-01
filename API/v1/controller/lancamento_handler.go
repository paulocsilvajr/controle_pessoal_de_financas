package controller

import (
	"controle_pessoal_de_financas/API/v1/dao"
	"controle_pessoal_de_financas/API/v1/helper"
	"controle_pessoal_de_financas/API/v1/model/detalhe_lancamento"
	"controle_pessoal_de_financas/API/v1/model/lancamento"
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

func LancamentoShow(w http.ResponseWriter, r *http.Request) {

}

// LancamentoCreate é um handler/controller que responde a rota '[POST] /lancamentos' e retorna StatusCreated(201) e os dados de um lancamento e seus detalhes lancamento criados através das informações informadas via JSON(body) caso o TOKEN informado for válido e o usuário associado ao token for cadastrado na API/DB. Caso ocorra algum erro, retorna StatusInternalServerError(500) ou StatusUnprocessableEntity(422) caso as informações no JSON não corresponderem ao formato {"cpf_pessoa":"?",  "nome_conta_origem":"?", "data":"?", "numero":"?", "descricao":"?", "nome_conta_destino":"?", "debito":?, "credito":?}
func LancamentoCreate(w http.ResponseWriter, r *http.Request) {
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

	idLancamentoRemocao, err := strconv.Atoi(lancamentoRemocao)
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

	for _, lancamento := range listaLancamentos {
		var lancPersJSON *LancamentoPersJSON

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
	if err != nil {
		return
	}

	return
}
