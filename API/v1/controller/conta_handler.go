package controller

import (
	"controle_pessoal_de_financas/API/v1/dao"
	"controle_pessoal_de_financas/API/v1/helper"
	"controle_pessoal_de_financas/API/v1/model/conta"
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
		listaContas, err = dao.CarregaTiposContaAtiva(db)
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
		err = DefineHeaderRetorno(w, SetHeaderJSON, verif, status, errors.New("Somente administradores podem consultar um tipo de conta inativa"))
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
