package controller

import (
	"controle_pessoal_de_financas/API/v1/dao"
	"controle_pessoal_de_financas/API/v1/helper"
	"controle_pessoal_de_financas/API/v1/logger"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func PessoaIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var status int

	token, err := helper.GetToken(r, MySigningKey)
	if err != nil {
		status = http.StatusInternalServerError
		defineStatusEmRetornoELog(w, status, err)

		return
	}

	usuarioToken, emailToken, err := helper.GetClaims(token)
	if err != nil {
		status = http.StatusInternalServerError
		defineStatusEmRetornoELog(w, status, err)

		return
	}

	listaPessoas, err := dao.CarregaPessoasSimples(db)
	if err != nil {
		status = http.StatusInternalServerError
		defineStatusEmRetornoELog(w, status, err)

		return
	}

	p, err := listaPessoas.ProcuraPessoaPorUsuario(usuarioToken)
	if err != nil || p.Email != emailToken {
		status = http.StatusInternalServerError
		defineStatusEmRetornoELog(w, status, err)

		return
	}

	status = http.StatusOK
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(listaPessoas); err != nil {
		logger.GeraLogFS(fmt.Sprintf("[%d] %s", status, err.Error()), time.Now())

		return
	}

	logger.GeraLogFS(fmt.Sprintf("[%d] %s[%d elementos]", status, "Enviando listagem de pessoas", len(listaPessoas)), time.Now())

}

// func UsuarioShow(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	nomeUsuario := vars["usuarioNome"]

// 	token := helper.GetToken(r, MySigningKey)

// 	_, usuarioToken, err := helper.GetClaims(token)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	if usuarioToken == nomeUsuario {
// 		usuarioEncontrado, err := usuario.DaoProcuraUsuario("nome", nomeUsuario)
// 		if err == nil {
// 			log.Println(err)
// 			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 			w.WriteHeader(http.StatusOK)

// 			if err := json.NewEncoder(w).Encode(*usuarioEncontrado); err != nil {
// 				log.Println(err)
// 			}
// 		} else {
// 			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 			w.WriteHeader(http.StatusNotFound)
// 		}
// 	}
// }

// func UsuarioShowAdmin(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	adminUsuario := vars["adminNome"]
// 	nomeUsuario := vars["usuarioNome"]

// 	token := helper.GetToken(r, MySigningKey)

// 	_, usuarioToken, err := helper.GetClaims(token)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	if usuarioToken == adminUsuario {
// 		usuarioEncontrado, err := usuario.DaoProcuraUsuario("nome", nomeUsuario)
// 		usuarioEncontrado.Senha = ""
// 		if err == nil {
// 			log.Println(err)
// 			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 			w.WriteHeader(http.StatusOK)

// 			if err := json.NewEncoder(w).Encode(*usuarioEncontrado); err != nil {
// 				log.Println(err)
// 			}
// 		} else {
// 			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 			w.WriteHeader(http.StatusNotFound)
// 		}
// 	}
// }

// func UsuarioCreate(w http.ResponseWriter, r *http.Request) {
// 	token := helper.GetToken(r, MySigningKey)

// 	admin, _, err := helper.GetClaims(token)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	if admin {
// 		var novoUsuario usuario.Usuario

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

// 		u, err := usuario.DaoAdicionaUsuario(novoUsuario)
// 		if err != nil {
// 			log.Println(err)
// 			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 			w.WriteHeader(http.StatusNotFound)
// 		} else {
// 			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 			w.WriteHeader(http.StatusCreated)
// 			if err := json.NewEncoder(w).Encode(u); err != nil {
// 				log.Println(err)
// 			}
// 		}
// 	}
// }

// func UsuarioRemove(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	nomeUsuario := vars["usuarioNome"]

// 	token := helper.GetToken(r, MySigningKey)

// 	admin, _, err := helper.GetClaims(token)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	if admin {
// 		err := usuario.DaoRemoveUsuario(nomeUsuario)
// 		if err != nil {
// 			log.Println(err)
// 			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 			w.WriteHeader(http.StatusNotFound)
// 		} else {
// 			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 			w.WriteHeader(http.StatusOK)
// 		}
// 	}
// }

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
