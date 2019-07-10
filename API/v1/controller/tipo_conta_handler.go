package controller

import (
	"controle_pessoal_de_financas/API/v1/dao"
	"controle_pessoal_de_financas/API/v1/helper"
	"controle_pessoal_de_financas/API/v1/model/tipo_conta"
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

func TipoContaEstado(w http.ResponseWriter, r *http.Request) {
	var status = http.StatusInternalServerError // 500
	type estado struct {
		Estado bool `json:"estado"`
	}
	var estadoTipoConta estado

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

	err = json.Unmarshal(body, &estadoTipoConta)
	status = http.StatusUnprocessableEntity // 422
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
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
