package main

import (
	"controle_pessoal_de_financas/API/v1/config"
	"controle_pessoal_de_financas/API/v1/helper"
	"controle_pessoal_de_financas/API/v1/route"
	"fmt"
	"log"
	"net/http"
)

func main() {
	helper.CriarDiretorioSeNaoExistir("config")

	configuracoes := config.AbrirConfiguracoes()
	porta := helper.FormatarPorta(configuracoes["porta"])
	host := configuracoes["host"]
	protocolo := configuracoes["protocolo"]

	router := route.NewRouter()

	var endereco string
	if host != "localhost" {
		endereco = host + porta
	} else {
		endereco = porta
	}

	fmt.Printf("Acesse o servidor/API pelo endereço: %s://%s%s\n", protocolo, host, porta)
	fmt.Printf("ou pelo IP: %s://%s%s\n\n[CTRL + c] para sair\n\n", protocolo, helper.GetLocalIP(), porta)

	if protocolo == "https" {
		certFile := "keys/new.cert.cert"
		keyFile := "keys/new.cert.key"
		log.Fatal(http.ListenAndServeTLS(endereco, certFile, keyFile, router))
	} else {
		log.Fatal(http.ListenAndServe(endereco, router))
	}
}
