package controller

import (
	"controle_pessoal_de_financas/API/v1/dao"
	"controle_pessoal_de_financas/API/v1/helper"
	"controle_pessoal_de_financas/API/v1/logger"
	"controle_pessoal_de_financas/API/v1/model/pessoa"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func PessoaIndex(w http.ResponseWriter, r *http.Request) {
	var status int

	SetHeaderJson(w)

	token, err := helper.GetToken(r, MySigningKey)
	if err != nil {
		status = http.StatusInternalServerError
		defineStatusEmRetornoELog(w, status, err)

		return
	}

	usuarioToken, emailToken, admin, err := helper.GetClaims(token)
	if err != nil {
		status = http.StatusInternalServerError
		defineStatusEmRetornoELog(w, status, err)

		return
	}

	var listaPessoas pessoa.PessoasI
	if admin {
		listaPessoas, err = dao.CarregaPessoas(db)
		if err != nil {
			status = http.StatusInternalServerError
			defineStatusEmRetornoELog(w, status, err)

			return
		}
	} else {
		listaPessoas, err = dao.CarregaPessoasSimples(db)
		if err != nil {
			status = http.StatusInternalServerError
			defineStatusEmRetornoELog(w, status, err)

			return
		}
	}

	p, err := listaPessoas.ProcuraPessoaPorUsuario(usuarioToken)
	if err != nil || p.GetEmail() != emailToken {
		status = http.StatusInternalServerError
		defineStatusEmRetornoELog(w, status, err)

		return
	}

	status = http.StatusOK

	funcao := "PessoaIndex"
	msg := fmt.Sprintf("%s: Listagem de pessoas", funcao)
	err = retornoData(w, status, msg, listaPessoas.Len(), listaPessoas)

	logger.GeraLogFS(
		fmt.Sprintf("[%d] %s: Enviando listagem de pessoas[%d elementos] %v",
			status,
			funcao,
			listaPessoas.Len(),
			err),
		time.Now())
}

func PessoaShow(w http.ResponseWriter, r *http.Request) {
	var status int
	var msg string

	SetHeaderJson(w)

	vars := mux.Vars(r)
	usuarioRota := vars["usuario"]

	token, err := helper.GetToken(r, MySigningKey)
	if err != nil {
		status = http.StatusInternalServerError
		defineStatusEmRetornoELog(w, status, err)

		return
	}

	usuarioToken, _, _, err := helper.GetClaims(token)
	if usuarioToken != usuarioRota || err != nil {
		status = http.StatusInternalServerError
		defineStatusEmRetornoELog(w, status, err)

		return
	}

	usuarioEncontrado, err := dao.ProcuraPessoaPorUsuario(db, usuarioRota)
	if err != nil {
		status = http.StatusInternalServerError
		defineStatusEmRetornoELog(w, status, err)

		return
	}

	status = http.StatusOK

	funcao := "PessoaShow"
	msg = fmt.Sprintf("%s: Dados de usuário %s", funcao, usuarioEncontrado.Usuario)
	quant := 1
	err = retornoData(w, status, msg, quant, usuarioEncontrado)

	logger.GeraLogFS(
		fmt.Sprintf("[%d] %s: Enviando dados de pessoa %s[%d elementos] %v",
			status,
			funcao,
			usuarioEncontrado.Usuario,
			quant,
			err),
		time.Now())
}
