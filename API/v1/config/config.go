package config

import (
	"encoding/json"
	"log"
	"os"
)

// https://play.golang.org/p/6dX5SMdVtr

const arquivo = "config/config.json"

type rota struct {
	Tipo, Rota, Descricao string
}

type rotas map[string]rota

func (r rotas) Len() int {
	return len(r)
}

var Rotas = rotas{
	"API": rota{
		"GET",
		"/API",
		"",
	},
	"Index": rota{
		"GET",
		"/",
		"",
	},
	"Login": rota{
		"POST",
		"/login/{usuario}",
		`Body: {"usuario":"?",  "senha":"?"}`,
	},
	"TokenValido": rota{
		"GET",
		"/token",
		"",
	},
	"PessoaIndex": rota{
		"GET",
		"/pessoas",
		"",
	},
	"PessoaShow": rota{
		"GET",
		"/pessoas/{usuario}",
		"",
	},
	"PessoaShowAdmin": rota{
		"GET",
		"/pessoas/{usuarioAdmin}/{usuario}",
		"",
	},
	"PessoaCreate": rota{
		"POST",
		"/pessoas",
		`Body: {"cpf":"?",  "nome_completo":"?", "usuario":"?", "senha":"?", "email":"?"[, "administrador": ?]}`,
	},
	"PessoaRemove": rota{
		"DELETE",
		"/pessoas/{usuario}",
		"",
	},
	"PessoaAlter": rota{
		"PUT",
		"/pessoas/{usuario}",
		`Body: {"cpf":"?",  "nome_completo":"?", "usuario":"?", "senha":"?", "email":"?"}`,
	},
	"PessoaEstado": rota{
		"PUT",
		"/pessoas/{usuario}/estado",
		`Body: {"estado":"?"}`,
	},
	"PessoaAdmin": rota{
		"PUT",
		"/pessoas/{usuario}/admin",
		`Body: {"administrador":"?"}`,
	},
}

type Configuracoes map[string]string

func AbrirConfiguracoes() Configuracoes {
	decodeFile, err := os.Open(arquivo)
	if err != nil {
		criarConfigPadrao()

		decodeFile, err = os.Open(arquivo)
		if err != nil {
			log.Println(err)
		}
	}
	defer decodeFile.Close()

	decoder := json.NewDecoder(decodeFile)

	configuracoes := make(Configuracoes)

	decoder.Decode(&configuracoes)

	return configuracoes
}

func criarConfigPadrao() {
	configuracoes := make(Configuracoes)
	configuracoes["porta"] = "8085"
	configuracoes["host"] = "localhost"
	configuracoes["duracao_token"] = "3600"
	configuracoes["protocolo"] = "https"

	encodeFile, err := os.Create(arquivo)
	if err != nil {
		log.Println(err)
	}

	encoder := json.NewEncoder(encodeFile)

	if err := encoder.Encode(configuracoes); err != nil {
		log.Println(err)
	}
	encodeFile.Close()
}
