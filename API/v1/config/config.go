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

var Rotas = map[string]rota{
	"Index": rota{
		"GET",
		"/",
		"",
	},
	"Login": rota{
		"POST",
		"/login/{usuario}",
		"",
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
