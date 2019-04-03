package controller

import (
	"controle_pessoal_de_financas/API/v1/config"
	"controle_pessoal_de_financas/API/v1/dao"
	"controle_pessoal_de_financas/API/v1/helper"
	"controle_pessoal_de_financas/API/v1/model/pessoa"
	"encoding/json"
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

// func Index(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")
// 	// O header ACESS-CONTROL-ALLOW-ORIGIN deve ser declarado em cada página
// 	// em que houver requisições à API.
// 	// Como a aplicação está sendo desenvolvida em Vue.js, é uma SPA(Single Page Application),
// 	// portando, é necessário declarar o header somente na página inicial.
// 	w.Header().Set("Access-Control-Allow-Origin", "*")

// 	// data := make(map[string]string)
// 	// data["title"] = "SCAK"

// 	tmpl, err := template.ParseFiles("www/index.html")
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	// tmpl.Execute(w, data)
// 	tmpl.Execute(w, nil)
// }

// func GetStatic(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	nomeStatic := vars["nomeStatic"]
// 	pathStatic := fmt.Sprintf("www/dist/%s", nomeStatic)

// 	if strings.Contains(nomeStatic, ".css") {
// 		w.Header().Set("Content-Type", "text/css; charset=utf-8")
// 	} else if strings.Contains(nomeStatic, ".js") {
// 		w.Header().Set("Content-Type", "text/javascript; charset=utf-8")
// 	} else {
// 		log.Printf("Arquivo %s inválido", nomeStatic)
// 		return
// 	}

// 	tmpl, err := template.ParseFiles(pathStatic)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	tmpl.Execute(w, nil)
// }

func Login(w http.ResponseWriter, r *http.Request) {
	var status int
	vars := mux.Vars(r)
	nomeUsuario := vars["usuario"]

	var usuarioInformado pessoa.Pessoa
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Println(err)
	}

	if err := r.Body.Close(); err != nil {
		log.Println(err)
	}

	if err := json.Unmarshal(body, &usuarioInformado); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		status = http.StatusUnprocessableEntity // 422
		w.WriteHeader(status)

		retornoStatus(w, status)
		// if err := json.NewEncoder(w).Encode(err); err != nil {
		// 	log.Println(err)
		// }
	}

	usuarioEncontrado, err := dao.ProcuraPessoaPorUsuario(db, nomeUsuario)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		status = http.StatusNotAcceptable // 406
		w.WriteHeader(status)

		retornoStatus(w, status)

		return
	}

	if usuarioEncontrado.Estado == false {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		status = http.StatusNotFound // 404
		w.WriteHeader(status)

		retornoStatus(w, status)

		return
	}

	senhaHash := usuarioEncontrado.Senha
	senhaInformadaHash := helper.GetSenhaSha256(usuarioInformado.Senha)
	if senhaHash != senhaInformadaHash {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		status = http.StatusNotAcceptable // 406
		w.WriteHeader(status)

		retornoStatus(w, status)

		return
	}

	token := jwt.New(jwt.SigningMethodHS256)

	duracao := config.AbrirConfiguracoes()["duracao_token"]
	intSegundos, err := strconv.Atoi(duracao)
	segundos := time.Duration(intSegundos)
	if err != nil {
		log.Println(err)
		segundos = 3600
	}

	claims := token.Claims.(jwt.MapClaims)
	claims["usuario"] = usuarioEncontrado.Usuario
	claims["email"] = usuarioEncontrado.Email
	claims["exp"] = time.Now().Add(time.Second * segundos).Unix()

	tokenString, _ := token.SignedString(MySigningKey)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(
		map[string]string{"token": tokenString},
	)
}

func TokenValido(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func retornoStatus(w http.ResponseWriter, status int) {
	json.NewEncoder(w).Encode(
		map[string]int{"status": status},
	)
}
