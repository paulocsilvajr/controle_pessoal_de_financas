package controller

import (
	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/dao"
	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/helper"
	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/model/tipo_conta"
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

// TipoContaIndex é um handler/controller que responde a rota '[GET] /tipos_conta' e retorna StatusOK(200) e uma listagem de tipos de conta de acordo com o tipo de usuário(admin/comum) caso o TOKEN informado for válido e o usuário associado ao token for cadastrado na API/DB. Caso ocorra algum erro, retorna StatusInternalServerError(500). Quando solicitado como usuário comum, retorna somente tipos de conta ativos, enquanto que como administrador, retorna todos os registros de tipo de conta
func TipoContaIndex(w http.ResponseWriter, r *http.Request) {
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

	var listaTiposConta tipo_conta.TiposConta
	if admin {
		listaTiposConta, err = dao.CarregaTiposConta(db)
		err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
		if err != nil {
			return
		}
	} else {
		listaTiposConta, err = dao.CarregaTiposContaAtiva(db)
		err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
		if err != nil {
			return
		}
	}

	status = http.StatusOK // 200
	funcao := "TipoContaIndex"
	DefineHeaderRetornoDados(
		w,
		SetHeaderJSON,
		status,
		listaTiposConta,
		funcao,
		"Listagem de tipos de conta",
		"Enviando listagem de tipos de conta")
}

// TipoContaShow é um handler/controller que responde a rota '[GET] /tipos_conta/{tipoConta}' e retorna StatusOK(200) e os dados do tipo de conta(nome) solicitada caso o TOKEN informado for válido e o usuário associado ao token for cadastrado na API/DB. Caso ocorra algum erro, retorna StatusInternalServerError(500)
func TipoContaShow(w http.ResponseWriter, r *http.Request) {
	var status = http.StatusInternalServerError // 500

	vars := mux.Vars(r)
	tipoContaRota := vars["tipoConta"]

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

	tipoContaEncontrada, err := dao.ProcuraTipoConta(db, tipoContaRota)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	if !tipoContaEncontrada.Estado {
		verif := !(admin && usuarioDB.Administrador)
		err = DefineHeaderRetorno(w, SetHeaderJSON, verif, status, errors.New("Somente administradores podem consultar um tipo de conta inativa"))
		if err != nil {
			return
		}
	}

	status = http.StatusOK // 200
	funcao := "TipoContaShow"
	DefineHeaderRetornoDado(
		w,
		SetHeaderJSON,
		status,
		tipoContaEncontrada,
		funcao,
		fmt.Sprintf("Dados de tipo de conta '%s'", tipoContaEncontrada.Nome),
		fmt.Sprintf("Enviando dados de tipo de conta '%s'", tipoContaEncontrada.Nome))
}

// TipoContaCreate é um handler/controller que responde a rota '[POST] /tipos_conta' e retorna StatusCreated(201) e os dados do tipo de conta criada através das informações informadas via JSON(body) caso o TOKEN informado for válido e o usuário associado ao token for cadastrado na API/DB. Caso ocorra algum erro, retorna StatusInternalServerError(500) ou StatusUnprocessableEntity(422) caso as informações no JSON não corresponderem ao formato {"nome":"?",  "descricao_debito":"?", "descricao_credito":"?"}
func TipoContaCreate(w http.ResponseWriter, r *http.Request) {
	var status = http.StatusInternalServerError
	var tipoContaFromJSON tipo_conta.TipoConta
	var novoTipoConta *tipo_conta.TipoConta

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

	err = json.Unmarshal(body, &tipoContaFromJSON)
	status = http.StatusUnprocessableEntity // 422
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	novoTipoConta, err = tipo_conta.NewTipoConta(
		tipoContaFromJSON.Nome,
		tipoContaFromJSON.DescricaoDebito,
		tipoContaFromJSON.DescricaoCredito)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	t, err := dao.AdicionaTipoConta(db, novoTipoConta)
	status = http.StatusInternalServerError // 500
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	status = http.StatusCreated // 201
	funcao := "TipoContaCreate"
	DefineHeaderRetornoDado(
		w,
		SetHeaderJSON,
		status,
		t,
		funcao,
		fmt.Sprintf("Dados de tipo de conta '%s'", t.Nome),
		fmt.Sprintf("Enviando dados de tipo de conta '%s'", t.Nome))
}

// TipoContaRemove é um handler/controller que responde a rota '[DELETE] /tipos_conta/{tipoConta}' e retorna StatusOK(200) e uma mensagem de confirmação caso o TOKEN informado for válido, o usuário associado ao token for cadastrado na API/DB e seja um administrador, que o tipo de conta informado seja cadastrado no BD. Caso ocorra algum erro, retorna StatusInternalServerError(500)
func TipoContaRemove(w http.ResponseWriter, r *http.Request) {
	var status = http.StatusInternalServerError // 500

	vars := mux.Vars(r)
	tipoContaRemocao := vars["tipoConta"]

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

	err = dao.RemoveTipoConta(db, tipoContaRemocao)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	status = http.StatusOK // 200
	funcao := "TipoContaRemove"
	DefineHeaderRetornoDado(
		w,
		SetHeaderJSON,
		status,
		tipoContaRemocao,
		funcao,
		fmt.Sprintf("Removido tipo de conta '%s'", tipoContaRemocao),
		fmt.Sprintf("Enviando resposta de remoção de tipo de conta '%s'", tipoContaRemocao))
}

// TipoContaAlter é um handler/controller que responde a rota '[PUT] /tipos_conta/{tipoConta}' e retorna StatusOK(200) e uma mensagem de confirmação com os dados do tipo de conta alterado caso o TOKEN informado for válido, o usuário associado ao token for cadastrado na API/DB e o tipo de conta informado na rota existir. Caso ocorra algum erro, retorna StatusInternalServerError(500) ou StatusUnprocessableEntity(422), caso o JSON não seguir o formato {["nome":"?",]  "descricao_debito":"?", "descricao_credito":"?"}, sendo campo nome opcional, ou StatusNotModified(304) caso ocorra algum erro na alteração do BD. Quando não for informado nome, esse campo não será alterado
func TipoContaAlter(w http.ResponseWriter, r *http.Request) {
	var status = http.StatusInternalServerError // 500
	var tipoContaFromJSON tipo_conta.TipoConta

	vars := mux.Vars(r)
	tipoContaAlteracao := vars["tipoConta"]

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

	err = json.Unmarshal(body, &tipoContaFromJSON)
	status = http.StatusUnprocessableEntity // 422
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	if len(tipoContaFromJSON.Nome) == 0 {
		tipoContaFromJSON.Nome = tipoContaAlteracao
	}

	t, err := dao.AlteraTipoConta(
		db,
		tipoContaAlteracao,
		&tipoContaFromJSON)
	status = http.StatusNotModified // 304
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	status = http.StatusOK // 200
	funcao := "TipoContaAlter"
	DefineHeaderRetornoDado(
		w,
		SetHeaderJSON,
		status,
		t,
		funcao,
		fmt.Sprintf("Novos dados de tipo de conta '%s'", t.Nome),
		fmt.Sprintf("Enviando novos dados de tipo de conta '%s'", t.Nome))
}

// TipoContaEstado é um handler/controller que responde a rota '[PUT] /tipos_conta/{tipoConta}/estado' e retorna StatusOK(200) e uma mensagem de confirmação com os dados do tipo de conta alterada caso o TOKEN informado for válido, o usuário associado ao token for cadastrado na API/DB e o tipo de conta informado na rota existir. Somente usuários ADMINISTRADORES podem ATIVAR tipos de conta, USUÁRIO COMUNS podem somente INATIVAR. Caso ocorra algum erro, retorna StatusInternalServerError(500), StatusUnprocessableEntity(422), caso o JSON não seguir o formato {"estado": ?}, StatusNotModified(304) caso ocorra algum erro na alteração do BD ou StatusNotFound(404) caso o tipo de conta informado na rota não existir
func TipoContaEstado(w http.ResponseWriter, r *http.Request) {
	var status = http.StatusInternalServerError // 500
	var estadoTipoConta estado

	vars := mux.Vars(r)
	tipoContaAlteracao := vars["tipoConta"]

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

	err = json.Unmarshal(body, &estadoTipoConta)
	status = http.StatusUnprocessableEntity // 422
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	usuarioComum := !(admin && usuarioDB.Administrador)
	if usuarioComum {
		verif := estadoTipoConta.Estado

		status = http.StatusInternalServerError // 500
		err = DefineHeaderRetorno(w, SetHeaderJSON, verif, status, errors.New("Somente administradores podem ativar um tipo de conta"))
		if err != nil {
			return
		}
	}

	tipoContaDBAlteracao, err := dao.ProcuraTipoConta(db, tipoContaAlteracao)
	status = http.StatusNotFound // 404
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	var alteraEstado func(*sql.DB, string) (*tipo_conta.TipoConta, error)
	if estadoTipoConta.Estado {
		alteraEstado = dao.AtivaTipoConta
	} else {
		alteraEstado = dao.InativaTipoConta
	}
	t, err := alteraEstado(db, tipoContaDBAlteracao.Nome)
	status = http.StatusNotModified // 304
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	status = http.StatusOK // 200
	funcao := "TipoContaEstado"
	DefineHeaderRetornoDado(
		w,
		SetHeaderJSON,
		status,
		t,
		funcao,
		fmt.Sprintf("Novos dados de tipo de conta '%s'", t.Nome),
		fmt.Sprintf("Enviando novos dados de tipo de conta '%s'", t.Nome))
}
