package controller

import (
	"controle_pessoal_de_financas/API/v1/config"
	"controle_pessoal_de_financas/API/v1/dao"
	"controle_pessoal_de_financas/API/v1/helper"
	"controle_pessoal_de_financas/API/v1/logger"
	"controle_pessoal_de_financas/API/v1/model/pessoa"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var (
	MySigningKey = []byte(`Remember, remember, the 5th of November.
The gunpowder treason and plot;
I know of no reason why the gunpowder treason
Should ever be forgot.`)
	db = dao.GetDB()
)

// LimitData define a quantidade máxima de bytes que o body de uma requisição suporta.
// Padrão: 1048576 Bytes == 1 MegaByte
const LimitData int64 = 1048576

type ReturnJson struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
}

type ReturnTokenJson struct {
	ReturnJson
	Token string `json:"token"`
}

type ReturnData struct {
	ReturnJson
	Count int         `json:"count"`
	Data  interface{} `json:"data"`
}

type Dados interface {
	Len() int
}

func Index(w http.ResponseWriter, r *http.Request) {
	status := http.StatusOK
	rotas := config.Rotas

	funcao := "Index"
	DefineHeaderRetornoDados(
		w,
		SetHeaderJson,
		status,
		rotas,
		funcao,
		"Rotas de API",
		"Enviando rotas de API")
}

func Login(w http.ResponseWriter, r *http.Request) {
	var status int
	vars := mux.Vars(r)
	nomeUsuario := vars["usuario"]

	var usuarioInformado pessoa.Pessoa
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, LimitData))
	if err != nil {
		log.Println(err)
	}

	if err := r.Body.Close(); err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(body, &usuarioInformado)
	status = http.StatusUnprocessableEntity // 422
	err = DefineHeaderRetorno(w, SetHeaderJson, err != nil, status, err)
	if err != nil {
		return
	}

	usuarioEncontrado, err := dao.ProcuraPessoaPorUsuario(db, nomeUsuario)
	status = http.StatusNotAcceptable // 406
	err = DefineHeaderRetorno(w, SetHeaderJson, err != nil, status, err)
	if err != nil {
		return
	}

	status = http.StatusNotFound // 404
	verif := usuarioEncontrado.Estado == false
	err = DefineHeaderRetorno(w, SetHeaderJson, verif, status, errors.New("Usuário inativo"))
	if err != nil {
		return
	}

	senhaHash := usuarioEncontrado.Senha
	senhaInformadaHash := helper.GetSenhaSha256(usuarioInformado.Senha)
	usuarioBD := usuarioEncontrado.Usuario
	verif = senhaHash != senhaInformadaHash || usuarioBD != usuarioInformado.Usuario
	status = http.StatusNotAcceptable // 406
	err = DefineHeaderRetorno(w, SetHeaderJson, verif, status, errors.New("Usuário ou Senha inválida"))
	if err != nil {
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)

	duracao := config.AbrirConfiguracoes()["duracao_token"]
	intSegundos, err := strconv.Atoi(duracao)
	if err != nil {
		log.Println(err)
		intSegundos = 3600
	}
	segundos := time.Duration(intSegundos)

	helper.SetClaims(
		token,
		segundos,
		usuarioEncontrado.Usuario,
		usuarioEncontrado.Email,
		usuarioEncontrado.Administrador)

	tokenString, _ := token.SignedString(MySigningKey)

	SetHeaderJson(w)
	status = http.StatusOK
	funcao := "Login"
	msg := fmt.Sprintf("%s: Token com duração de %d segundos", funcao, intSegundos)
	defineStatusETokenEmRetornoELog(w, status, msg, tokenString)
}

func TokenValido(w http.ResponseWriter, r *http.Request) {
	SetHeaderJson(w)

	funcao := "TokenValido"
	msg := fmt.Sprintf("%s: Token válido", funcao)
	defineStatusEMensagemEmRetornoELog(w, http.StatusOK, msg)
}

func DefineHeaderRetorno(w http.ResponseWriter, header func(w http.ResponseWriter), verif bool, status int, err error) error {
	if verif {
		header(w)

		defineStatusEmRetornoELog(w, status, err)

		return err
	}
	return nil
}

func DefineHeaderRetornoDados(
	w http.ResponseWriter,
	header func(w http.ResponseWriter),
	status int,
	dados Dados,
	funcao, msgRetorno, msgLog string) error {

	header(w)

	err := retornoData(
		w,
		status,
		fmt.Sprintf("%s: %s", funcao, msgRetorno),
		dados.Len(),
		dados)

	logger.GeraLogFS(
		fmt.Sprintf("[%d] %s: %s[%d elementos] %v",
			status,
			funcao,
			msgLog,
			dados.Len(),
			err),
		time.Now())

	return err
}

func DefineHeaderRetornoDado(
	w http.ResponseWriter,
	header func(w http.ResponseWriter),
	status int,
	dado interface{},
	funcao, msgRetorno, msgLog string) error {

	header(w)

	err := retornoData(
		w,
		status,
		fmt.Sprintf("%s: %s", funcao, msgRetorno),
		1,
		dado)

	logger.GeraLogFS(
		fmt.Sprintf("[%d] %s: %s[%d elemento] %v",
			status,
			funcao,
			msgLog,
			1,
			err),
		time.Now())

	return err
}

func retornoStatus(w http.ResponseWriter, status int) {
	json.NewEncoder(w).Encode(
		map[string]int{"status": status},
	)
}

func retornoStatusMsg(w http.ResponseWriter, status int, msg string) error {
	retorno := &ReturnJson{
		StatusCode: status,
		Message:    msg,
	}

	return json.NewEncoder(w).Encode(retorno)
}

func retornoStatusMsgToken(w http.ResponseWriter, status int, msg, token string) error {
	retorno := new(ReturnTokenJson)
	retorno.StatusCode = status
	retorno.Message = msg
	retorno.Token = token

	return json.NewEncoder(w).Encode(retorno)
}

func retornoData(w http.ResponseWriter, status int, msg string, count int, data interface{}) error {
	retorno := new(ReturnData)
	retorno.StatusCode = status
	retorno.Message = msg
	retorno.Count = count
	retorno.Data = data

	return json.NewEncoder(w).Encode(retorno)
}

func defineStatusEmRetornoELog(w http.ResponseWriter, status int, err error) error {
	return defineStatusEMensagemEmRetornoELog(w, status, err.Error())
}

func defineStatusEMensagemEmRetornoELog(w http.ResponseWriter, status int, msg string) error {
	w.WriteHeader(status) // w.WriteHeader deve vir SEMPRE antes de json.NewEncoder()

	logger.GeraLogFS(fmt.Sprintf("[%d] %s", status, msg), time.Now())

	return retornoStatusMsg(w, status, msg)
}

func defineStatusETokenEmRetornoELog(w http.ResponseWriter, status int, msg, token string) error {
	w.WriteHeader(status)

	logger.GeraLogFS(fmt.Sprintf("[%d] %s", status, msg), time.Now())

	return retornoStatusMsgToken(w, status, msg, token)
}
