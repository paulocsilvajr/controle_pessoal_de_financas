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

// MySigningKey é a chave usada para assinar o token da API
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

// ReturnJSON é a struct base para o retorno em formato JSON
type ReturnJSON struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
}

// ReturnTokenJSON é uma struct baseada em ReturnJSON com um campo para o token em string
type ReturnTokenJSON struct {
	ReturnJSON
	Token string `json:"token"`
}

// ReturnData é uma struct baseada em ReturnJSON com dois campos adicionais: Count, para a quantidade de registros, e Data, para dados em uma interface
type ReturnData struct {
	ReturnJSON
	Count int         `json:"count"`
	Data  interface{} `json:"data"`
}

// Dados é uma interface que obriga a criação do método Len() para obter a quantidade de elementos contidos dentro do slice
type Dados interface {
	Len() int
}

// Index é um handler/controller que responde a rota '[GET] /' e retorna StatusOK(200) e as rotas da API em um JSON, somente se o token informado no header for válido
func Index(w http.ResponseWriter, r *http.Request) {
	status := http.StatusOK
	rotas := config.Rotas

	funcao := "Index"
	DefineHeaderRetornoDados(
		w,
		SetHeaderJSON,
		status,
		rotas,
		funcao,
		"Rotas de API",
		"Enviando rotas de API")
}

// Login é um handler/controller que responde a rota '[POST] /login/{usuario}' e retorna StatusOK(200) e o token em string(com dados: usuário, email, tipo e validade codificados) caso o usuário e senha informados via JSON(body) estiverem corretos. Não é necessário ter um token válido para consultar essa rota. Retorna StatusUnprocessableEntity(422) caso o JSON(body) for inválido, StatusNotAcceptable(406) caso o usuário/senha informado estiver incorreto ou não existir, StatusNotFound(404) caso o usuário estiver inativo
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
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	usuarioEncontrado, err := dao.ProcuraPessoaPorUsuario(db, nomeUsuario)
	status = http.StatusNotAcceptable // 406
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	status = http.StatusNotFound // 404
	verif := usuarioEncontrado.Estado == false
	err = DefineHeaderRetorno(w, SetHeaderJSON, verif, status, errors.New("Usuário inativo"))
	if err != nil {
		return
	}

	senhaHash := usuarioEncontrado.Senha
	senhaInformadaHash := helper.GetSenhaSha256(usuarioInformado.Senha)
	usuarioBD := usuarioEncontrado.Usuario
	verif = !(senhaHash == senhaInformadaHash && usuarioBD == usuarioInformado.Usuario)
	status = http.StatusNotAcceptable // 406
	err = DefineHeaderRetorno(w, SetHeaderJSON, verif, status, errors.New("Usuário ou Senha inválida"))
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

	SetHeaderJSON(w)
	status = http.StatusOK
	funcao := "Login"
	msg := fmt.Sprintf("%s: Token com duração de %d segundos", funcao, intSegundos)
	defineStatusETokenEmRetornoELog(w, status, msg, tokenString)
}

// TokenValido é um handler/controller que responde a rota '[GET] /token' e retorna StatusOK(200) e uma mensagem caso o token passado no cabeçalho da requisição seja válido(dentro do período de validade). Usado principalmente para requisitar um novo token no cliente caso esteja vencido
func TokenValido(w http.ResponseWriter, r *http.Request) {
	SetHeaderJSON(w)

	funcao := "TokenValido"
	msg := fmt.Sprintf("%s: Token válido", funcao)
	defineStatusEMensagemEmRetornoELog(w, http.StatusOK, msg)
}

// API é um handler/controller que responde a rota '[GET] /API' e retorna StatusOK(200) e uma mensagem informado a quantidade de rotas disponíveis na API e outros detalhes. Não é necessário ter um token válido para consultar essa rota
func API(w http.ResponseWriter, r *http.Request) {
	status := http.StatusOK
	DefineHeaderRetornoDado(w,
		SetHeaderJSON,
		status,
		map[string]string{"API OnLine": fmt.Sprintf("%.2d rota(s) cadastrada(s)", len(config.Rotas))},
		"API",
		"API Online, faça o login pela rota [GET] '/login/{usuario}' e consulte todas as rotas em [GET] '/'",
		"API Online")
}

// DefineHeaderRetorno define o header(de acordo com função informada), o retorno(StatusCode) e um erro, caso a verif(verificação boleana informada) seja verdadeira. Retorna erro=nil somente se a verificação seja falsa. Usado para parar a execução de um handler caso ocorra um erro e necessite dar retorno ao cliente da API. Gera o retorno para o cliente da API e o log do sistema(tela e arquivo)
func DefineHeaderRetorno(w http.ResponseWriter, header func(w http.ResponseWriter), verif bool, status int, err error) error {
	if verif {
		header(w)

		defineStatusEmRetornoELog(w, status, err)

		return err
	}
	return nil
}

// DefineHeaderRetornoDados define o header(de acordo com função informada), retorno(StatusCode), dados para o retorno(que implemente a interface Dados), o nome da função que o invocou, a mensagem de retorno e de log. Retorno um erro caso ocorra algum problema no processo de gerar o log e retorno. Gera o retorno para o cliente da API e o log do sistema(tela e arquivo). Usado geralmente no final dos handler para enviar vários registros para o cliente da API
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

// DefineHeaderRetornoDado define o header(de acordo com função informada), retorno(StatusCode), um único dado para o retorno(tipo genérico interface{}), o nome da função que o invocou, a mensagem de retorno e de log. Retorno um erro caso ocorra algum problema no processo de gerar o log e retorno. Gera o retorno para o cliente da API e o log do sistema(tela e arquivo). Usado geralmente no final dos handler para enviar um único registro para o cliente da API
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
	retorno := &ReturnJSON{
		StatusCode: status,
		Message:    msg,
	}

	return json.NewEncoder(w).Encode(retorno)
}

func retornoStatusMsgToken(w http.ResponseWriter, status int, msg, token string) error {
	retorno := new(ReturnTokenJSON)
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
