package controller

import (
	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/dao"
	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/helper"
	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/model/conta"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// ContaIndex é um handler/controller que responde a rota '[GET] /contas' e retorna StatusOK(200) e uma listagem de contas de acordo com o tipo de usuário(admin/comum) caso o TOKEN informado for válido e o usuário associado ao token for cadastrado na API/DB. Caso ocorra algum erro, retorna StatusInternalServerError(500). Quando solicitado como usuário comum, retorna somente contas ativas, enquanto que como administrador, retorna todos os registros de contas
func ContaIndex(w http.ResponseWriter, r *http.Request) {
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

	_, err = dao.ProcuraPessoaPorUsuario(db, usuarioToken)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	var listaContas conta.Contas
	if admin {
		listaContas, err = dao.CarregaContas(db)
		err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
		if err != nil {
			return
		}
	} else {
		listaContas, err = dao.CarregaContasAtiva(db)
		err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
		if err != nil {
			return
		}
	}

	status = http.StatusOK // 200
	funcao := "ContaIndex"
	DefineHeaderRetornoDados(
		w,
		SetHeaderJSON,
		status,
		listaContas,
		funcao,
		"Listagem de contas",
		"Enviando listagem de contas")
}

// ContaShow é um handler/controller que responde a rota '[GET] /contas/{conta}' e retorna StatusOK(200) e os dados da conta(nome) solicitada caso o TOKEN informado for válido e o usuário associado ao token for cadastrado na API/DB. Caso ocorra algum erro, retorna StatusInternalServerError(500)
func ContaShow(w http.ResponseWriter, r *http.Request) {
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

	usuarioDB, err := dao.ProcuraPessoaPorUsuario(db, usuarioToken)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	contaEncontrada, err := dao.ProcuraConta(db, contaRota)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	if !contaEncontrada.Estado {
		verif := !(admin && usuarioDB.Administrador)
		err = DefineHeaderRetorno(w, SetHeaderJSON, verif, status, errors.New("Somente administradores podem consultar uma conta inativa"))
		if err != nil {
			return
		}
	}

	status = http.StatusOK // 200
	funcao := "ContaShow"
	DefineHeaderRetornoDado(
		w,
		SetHeaderJSON,
		status,
		contaEncontrada,
		funcao,
		fmt.Sprintf("Dados de conta '%s'", contaEncontrada.Nome),
		fmt.Sprintf("Enviando dados de conta '%s'", contaEncontrada.Nome))
}

// ContaCreate é um handler/controller que responde a rota '[POST] /contas' e retorna StatusCreated(201) e os dados do tipa criada através das informações informadas via JSON(body) caso o TOKEN informado for válido e o usuário associado ao token for cadastrado na API/DB. Caso ocorra algum erro, retorna StatusInternalServerError(500) ou StatusUnprocessableEntity(422) caso as informações no JSON não corresponderem ao formato {"nome":"?",  "nome_tipo_conta":"?", "codigo":"?", "conta_pai":"?", "comentario":"?"}
func ContaCreate(w http.ResponseWriter, r *http.Request) {
	var status = http.StatusInternalServerError
	var contaFromJSON conta.Conta
	var novaConta *conta.Conta

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

	err = json.Unmarshal(body, &contaFromJSON)
	status = http.StatusUnprocessableEntity // 422
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	novaConta, err = conta.NewConta(
		contaFromJSON.Nome, contaFromJSON.NomeTipoConta, contaFromJSON.Codigo, contaFromJSON.ContaPai, contaFromJSON.Comentario)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	c, err := dao.AdicionaConta(db, novaConta)
	status = http.StatusInternalServerError // 500
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	status = http.StatusCreated // 201
	funcao := "ContaCreate"
	DefineHeaderRetornoDado(
		w,
		SetHeaderJSON,
		status,
		c,
		funcao,
		fmt.Sprintf("Dados da conta '%s'", c.Nome),
		fmt.Sprintf("Enviando dados da conta '%s'", c.Nome))
}

// ContaRemove é um handler/controller que responde a rota '[DELETE] /contas/{Conta}' e retorna StatusOK(200) e uma mensagem de confirmação caso o TOKEN informado for válido, o usuário associado ao token for cadastrado na API/DB e seja um administrador, que a conta informada esteja cadastrado no BD. Caso ocorra algum erro, retorna StatusInternalServerError(500)
func ContaRemove(w http.ResponseWriter, r *http.Request) {
	var status = http.StatusInternalServerError // 500

	vars := mux.Vars(r)
	contaRemocao := vars["conta"]

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

	err = dao.RemoveConta(db, contaRemocao)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	status = http.StatusOK // 200
	funcao := "ContaRemove"
	DefineHeaderRetornoDado(
		w,
		SetHeaderJSON,
		status,
		contaRemocao,
		funcao,
		fmt.Sprintf("Removido conta '%s'", contaRemocao),
		fmt.Sprintf("Enviando resposta de remoção de conta '%s'", contaRemocao))
}

// ContaAlter é um handler/controller que responde a rota '[PUT] /contas/{conta}' e retorna StatusOK(200) e uma mensagem de confirmação com os dados da conta alterada caso o TOKEN informado for válido, o usuário associado ao token for cadastrado na API/DB e a conta informada na rota existir. Caso ocorra algum erro, retorna StatusInternalServerError(500) ou StatusUnprocessableEntity(422), caso o JSON não seguir o formato {["nome":"?",]  "nome_tipo_conta":"?", "codigo":"?", "conta_pai":"?", "comentario":"?"}, sendo campo nome opcional, ou StatusNotModified(304) caso ocorra algum erro na alteração do BD. Quando não for informado nome, esse campo não será alterado
func ContaAlter(w http.ResponseWriter, r *http.Request) {
	var status = http.StatusInternalServerError // 500
	var contaFromJSON conta.Conta

	vars := mux.Vars(r)
	contaAlteracao := vars["conta"]

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

	err = json.Unmarshal(body, &contaFromJSON)
	status = http.StatusUnprocessableEntity // 422
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	if len(contaFromJSON.Nome) == 0 {
		contaFromJSON.Nome = contaAlteracao
	}

	c, err := dao.AlteraConta(
		db,
		contaAlteracao,
		&contaFromJSON)
	status = http.StatusNotModified // 304
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	status = http.StatusOK // 200
	funcao := "ContaAlter"
	DefineHeaderRetornoDado(
		w,
		SetHeaderJSON,
		status,
		c,
		funcao,
		fmt.Sprintf("Novos dados de conta '%s'", c.Nome),
		fmt.Sprintf("Enviando novos dados de conta '%s'", c.Nome))
}

// ContaEstado é um handler/controller que responde a rota '[PUT] /contas/{conta}/estado' e retorna StatusOK(200) e uma mensagem de confirmação com os dados da conta alterada caso o TOKEN informado for válido, o usuário associado ao token for cadastrado na API/DB e a conta informada na rota existir. Somente usuários ADMINISTRADORES podem ATIVAR contas, USUÁRIO COMUNS podem somente INATIVAR. Caso ocorra algum erro, retorna StatusInternalServerError(500), StatusUnprocessableEntity(422), caso o JSON não seguir o formato {"estado": ?}, StatusNotModified(304) caso ocorra algum erro na alteração do BD ou StatusNotFound(404) caso a conta informada na rota não existir
func ContaEstado(w http.ResponseWriter, r *http.Request) {
	var status = http.StatusInternalServerError // 500
	var estadoConta estado

	vars := mux.Vars(r)
	contaAlteracao := vars["conta"]

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

	err = json.Unmarshal(body, &estadoConta)
	status = http.StatusUnprocessableEntity // 422
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	usuarioComum := !(admin && usuarioDB.Administrador)
	if usuarioComum {
		verif := estadoConta.Estado

		status = http.StatusInternalServerError // 500
		err = DefineHeaderRetorno(w, SetHeaderJSON, verif, status, errors.New("Somente administradores podem ativar uma conta"))
		if err != nil {
			return
		}
	}

	contaDBAlteracao, err := dao.ProcuraConta(db, contaAlteracao)
	status = http.StatusNotFound // 404
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	var alteraEstado func(*sql.DB, string) (*conta.Conta, error)
	if estadoConta.Estado {
		alteraEstado = dao.AtivaConta
	} else {
		alteraEstado = dao.InativaConta
	}
	c, err := alteraEstado(db, contaDBAlteracao.Nome)
	status = http.StatusNotModified // 304
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	status = http.StatusOK // 200
	funcao := "ContaEstado"
	DefineHeaderRetornoDado(
		w,
		SetHeaderJSON,
		status,
		c,
		funcao,
		fmt.Sprintf("Novos dados de conta '%s'", c.Nome),
		fmt.Sprintf("Enviando novos dados de conta '%s'", c.Nome))
}
