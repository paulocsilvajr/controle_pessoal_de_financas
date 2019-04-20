package controller

import (
	"controle_pessoal_de_financas/API/v1/dao"
	"controle_pessoal_de_financas/API/v1/helper"
	"controle_pessoal_de_financas/API/v1/logger"
	"controle_pessoal_de_financas/API/v1/model/pessoa"
	"fmt"
	"log"
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
	retornoData(w, status, msg, listaPessoas.Len(), listaPessoas)
	if err != nil {
		logger.GeraLogFS(fmt.Sprintf("[%d] %s", status, err.Error()), time.Now())

		return
	}

	msg = fmt.Sprintf("%s: Enviando listagem de pessoas", funcao)
	logger.GeraLogFS(fmt.Sprintf("[%d] %s[%d elementos]", status, msg, listaPessoas.Len()), time.Now())
}

func PessoaShow(w http.ResponseWriter, r *http.Request) {
	var status int
	var msg string

	vars := mux.Vars(r)
	usuario := vars["usuario"]

	token, err := helper.GetToken(r, MySigningKey)
	usuarioToken, _, _, err := helper.GetClaims(token)
	if err != nil {
		log.Println(err)
	}

	SetHeaderJson(w)

	if usuarioToken != usuario {
		status = http.StatusNotFound
		msg = "Usuário não autenticado(token)"
		defineStatusEMensagemEmRetornoELog(w, status, msg)

		return
	}

	usuarioEncontrado, err := dao.ProcuraPessoaPorUsuario(db, usuario)
	if err != nil {
		status = http.StatusInternalServerError
		defineStatusEMensagemEmRetornoELog(w, status, err.Error())

		return
	}

	status = http.StatusOK
	funcao := "PessoaShow"
	msg = fmt.Sprintf("%s: Dados de usuário %s", funcao, usuarioEncontrado.Usuario)
	quant := 1
	retornoData(w, status, msg, quant, usuarioEncontrado)

	msg = fmt.Sprintf("%s: Enviando dados de pessoa %s", funcao, usuarioEncontrado.Usuario)
	logger.GeraLogFS(fmt.Sprintf("[%d] %s[%d elementos]", status, msg, quant), time.Now())
}
