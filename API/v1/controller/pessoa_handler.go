package controller

import (
	"controle_pessoal_de_financas/API/v1/dao"
	"controle_pessoal_de_financas/API/v1/helper"
	"controle_pessoal_de_financas/API/v1/logger"
	"controle_pessoal_de_financas/API/v1/model/pessoa"
	"fmt"
	"net/http"
	"time"
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
