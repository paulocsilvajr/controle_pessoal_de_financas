package controller

import (
	"controle_pessoal_de_financas/API/v1/dao"
	"controle_pessoal_de_financas/API/v1/helper"
	"controle_pessoal_de_financas/API/v1/model/pessoa"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func PessoaIndex(w http.ResponseWriter, r *http.Request) {
	var status = http.StatusInternalServerError

	token, err := helper.GetToken(r, MySigningKey)
	err = DefineHeaderRetorno(w, SetHeaderJson, err != nil, status, err)
	if err != nil {
		return
	}

	usuarioToken, emailToken, admin, err := helper.GetClaims(token)
	err = DefineHeaderRetorno(w, SetHeaderJson, err != nil, status, err)
	if err != nil {
		return
	}

	var listaPessoas pessoa.PessoasI
	if admin {
		listaPessoas, err = dao.CarregaPessoas(db)
		err = DefineHeaderRetorno(w, SetHeaderJson, err != nil, status, err)
		if err != nil {
			return
		}
	} else {
		listaPessoas, err = dao.CarregaPessoasSimples(db)
		err = DefineHeaderRetorno(w, SetHeaderJson, err != nil, status, err)
		if err != nil {
			return
		}
	}

	p, err := listaPessoas.ProcuraPessoaPorUsuario(usuarioToken)
	err = DefineHeaderRetorno(w, SetHeaderJson, err != nil, status, err)
	if err != nil {
		return
	}

	err = DefineHeaderRetorno(w, SetHeaderJson, p.GetEmail() != emailToken, status, errors.New("Email de token não confere com email de pessoa"))
	if err != nil {
		return
	}

	status = http.StatusOK
	funcao := "PessoaIndex"
	DefineHeaderRetornoDados(
		w,
		SetHeaderJson,
		status,
		listaPessoas,
		funcao,
		"Listagem de pessoas",
		"Enviando listagem de pessoas")
}

func PessoaShow(w http.ResponseWriter, r *http.Request) {
	var status = http.StatusInternalServerError

	vars := mux.Vars(r)
	usuarioRota := vars["usuario"]

	token, err := helper.GetToken(r, MySigningKey)
	err = DefineHeaderRetorno(w, SetHeaderJson, err != nil, status, err)
	if err != nil {
		return
	}

	usuarioToken, _, _, err := helper.GetClaims(token)
	err = DefineHeaderRetorno(w, SetHeaderJson, err != nil, status, err)
	if err != nil {
		return
	}

	verif := usuarioToken != usuarioRota
	err = DefineHeaderRetorno(w, SetHeaderJson, verif, status, errors.New("Usuário de token diferente do solicitado na rota"))
	if err != nil {
		return
	}

	usuarioEncontrado, err := dao.ProcuraPessoaPorUsuario(db, usuarioRota)
	err = DefineHeaderRetorno(w, SetHeaderJson, err != nil, status, err)
	if err != nil {
		return
	}

	status = http.StatusOK
	funcao := "PessoaShow"
	DefineHeaderRetornoDado(
		w,
		SetHeaderJson,
		status,
		usuarioEncontrado,
		funcao,
		fmt.Sprintf("Dados de pessoa %s", usuarioEncontrado.Usuario),
		fmt.Sprintf("Enviando dados de pessoa %s", usuarioEncontrado.Usuario))
}
