package controller

import (
	"controle_pessoal_de_financas/API/v1/dao"
	"controle_pessoal_de_financas/API/v1/helper"
	"controle_pessoal_de_financas/API/v1/model/pessoa"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
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

	pessoaEncontrada, err := dao.ProcuraPessoaPorUsuario(db, usuarioRota)
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
		pessoaEncontrada,
		funcao,
		fmt.Sprintf("Dados de pessoa '%s'", pessoaEncontrada.Usuario),
		fmt.Sprintf("Enviando dados de pessoa '%s'", pessoaEncontrada.Usuario))
}

func PessoaShowAdmin(w http.ResponseWriter, r *http.Request) {
	var status = http.StatusInternalServerError

	vars := mux.Vars(r)
	usuarioAdmin := vars["usuarioAdmin"]
	usuario := vars["usuario"]

	token, err := helper.GetToken(r, MySigningKey)
	err = DefineHeaderRetorno(w, SetHeaderJson, err != nil, status, err)
	if err != nil {
		return
	}

	usuarioToken, _, admin, err := helper.GetClaims(token)
	err = DefineHeaderRetorno(w, SetHeaderJson, err != nil, status, err)
	if err != nil {
		return
	}

	verif := usuarioToken != usuarioAdmin
	err = DefineHeaderRetorno(w, SetHeaderJson, verif, status, errors.New("Usuário de token diferente do informado na rota"))
	if err != nil {
		return
	}

	usuarioDB, err := dao.ProcuraPessoaPorUsuario(db, usuarioAdmin)
	err = DefineHeaderRetorno(w, SetHeaderJson, err != nil, status, err)
	if err != nil {
		return
	}

	verif = !admin || !usuarioDB.Administrador
	err = DefineHeaderRetorno(w, SetHeaderJson, verif, status, errors.New("Somente administradores podem usar essa rota"))
	if err != nil {
		return
	}

	status = http.StatusNotFound
	pessoaEncontrada, err := dao.ProcuraPessoaPorUsuario(db, usuario)
	err = DefineHeaderRetorno(w, SetHeaderJson, err != nil, status, err)
	if err != nil {
		return
	}

	status = http.StatusOK
	funcao := "PessoaShowAdmin"
	DefineHeaderRetornoDado(
		w,
		SetHeaderJson,
		status,
		pessoaEncontrada,
		funcao,
		fmt.Sprintf("Dados de pessoa '%s'", pessoaEncontrada.Usuario),
		fmt.Sprintf("Enviando dados de pessoa '%s'", pessoaEncontrada.Usuario))
}

func PessoaCreate(w http.ResponseWriter, r *http.Request) {
	var status = http.StatusInternalServerError
	var pessoaFromJson pessoa.Pessoa
	var novaPessoa *pessoa.Pessoa

	token, err := helper.GetToken(r, MySigningKey)
	err = DefineHeaderRetorno(w, SetHeaderJson, err != nil, status, err)
	if err != nil {
		return
	}

	usuarioToken, _, admin, err := helper.GetClaims(token)
	err = DefineHeaderRetorno(w, SetHeaderJson, err != nil, status, err)
	if err != nil {
		return
	}

	usuarioDB, err := dao.ProcuraPessoaPorUsuario(db, usuarioToken)
	err = DefineHeaderRetorno(w, SetHeaderJson, err != nil, status, err)
	if err != nil {
		return
	}

	verif := !admin || !usuarioDB.Administrador
	err = DefineHeaderRetorno(w, SetHeaderJson, verif, status, errors.New("Somente administradores podem usar essa rota"))
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

	err = json.Unmarshal(body, &pessoaFromJson)
	status = http.StatusUnprocessableEntity // 422
	err = DefineHeaderRetorno(w, SetHeaderJson, err != nil, status, err)
	if err != nil {
		return
	}

	novaPessoa, err = pessoa.NewPessoa(
		pessoaFromJson.Cpf,
		pessoaFromJson.NomeCompleto,
		pessoaFromJson.Usuario,
		pessoaFromJson.Senha,
		pessoaFromJson.Email)
	err = DefineHeaderRetorno(w, SetHeaderJson, err != nil, status, err)
	if err != nil {
		return
	}
	adicionaPessoa := dao.AdicionaPessoa
	if pessoaFromJson.Administrador {
		adicionaPessoa = dao.AdicionaPessoaAdmin
	}

	p, err := adicionaPessoa(db, novaPessoa)
	status = http.StatusInternalServerError // 500
	err = DefineHeaderRetorno(w, SetHeaderJson, err != nil, status, err)
	if err != nil {
		return
	}

	status = http.StatusCreated // 201
	funcao := "PessoaCreate"
	DefineHeaderRetornoDado(
		w,
		SetHeaderJson,
		status,
		p,
		funcao,
		fmt.Sprintf("Dados de pessoa '%s'", p.Usuario),
		fmt.Sprintf("Enviando dados de pessoa '%s'", p.Usuario))
}

func PessoaRemove(w http.ResponseWriter, r *http.Request) {
	var status = http.StatusInternalServerError // 500

	vars := mux.Vars(r)
	usuarioRemocao := vars["usuario"]

	token, err := helper.GetToken(r, MySigningKey)
	err = DefineHeaderRetorno(w, SetHeaderJson, err != nil, status, err)
	if err != nil {
		return
	}

	usuarioToken, _, admin, err := helper.GetClaims(token)
	err = DefineHeaderRetorno(w, SetHeaderJson, err != nil, status, err)
	if err != nil {
		return
	}

	usuarioDB, err := dao.ProcuraPessoaPorUsuario(db, usuarioToken)
	err = DefineHeaderRetorno(w, SetHeaderJson, err != nil, status, err)
	if err != nil {
		return
	}

	verif := !admin || !usuarioDB.Administrador
	err = DefineHeaderRetorno(w, SetHeaderJson, verif, status, errors.New("Somente administradores podem usar essa rota"))
	if err != nil {
		return
	}

	verif = usuarioToken == usuarioRemocao
	err = DefineHeaderRetorno(w, SetHeaderJson, verif, status, fmt.Errorf("O usuário %s não pode remover a si mesmo", usuarioToken))
	if err != nil {
		return
	}

	err = dao.RemovePessoaPorUsuario(db, usuarioRemocao)
	err = DefineHeaderRetorno(w, SetHeaderJson, err != nil, status, err)
	if err != nil {
		return
	}

	status = http.StatusOK // 200
	funcao := "PessoaRemove"
	DefineHeaderRetornoDado(
		w,
		SetHeaderJson,
		status,
		usuarioRemocao,
		funcao,
		fmt.Sprintf("Removido pessoa '%s'", usuarioRemocao),
		fmt.Sprintf("Enviando resposta de remoção de pessoa '%s'", usuarioRemocao))

}

// func UsuarioAlter(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	nomeUsuario := vars["usuarioNome"]

// 	var novoUsuario usuario.Usuario

// 	token := helper.GetToken(r, MySigningKey)

// 	admin, usuarioToken, err := helper.GetClaims(token)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	if admin || nomeUsuario == usuarioToken {
// 		// io.LimitReader define limite para o tamanho do json
// 		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

// 		if err != nil {
// 			log.Println(err)
// 		}

// 		if err := r.Body.Close(); err != nil {
// 			log.Println(err)
// 		}

// 		if err := json.Unmarshal(body, &novoUsuario); err != nil {
// 			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 			w.WriteHeader(422) // unprocessable entity
// 			if err := json.NewEncoder(w).Encode(err); err != nil {
// 				log.Println(err)
// 			}
// 		}

// 		u, err := usuario.DaoAlteraUsuario(nomeUsuario, novoUsuario)
// 		if err != nil {
// 			log.Println(err)
// 			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 			w.WriteHeader(http.StatusNotModified)
// 		} else {
// 			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 			w.WriteHeader(http.StatusOK)
// 			if err := json.NewEncoder(w).Encode(u); err != nil {
// 				log.Println(err)
// 			}
// 		}
// 	}
// }
