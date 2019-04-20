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
	porta := fmt.Sprintf(":%s", configuracoes["porta"])
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

	// testes()
}

// func testes() {
// 	////////////////////////////////////////////////////////////////
// 	// TESTE dao.CarregaPessoa
// 	////////////////////////////////////////////////////////////////
// 	var db = dao.GetDB()
// 	listaPessoas, err := dao.CarregaPessoas(db)
// 	fmt.Printf("MAIN\nerro: %v\ntipo: %T\n", err, listaPessoas)
// 	for n, p := range listaPessoas {
// 		fmt.Printf("%3d [%T]: %v\n", n, p, p)
// 	}
// 	///////////////////////////////////////////////////////////////

// 	///////////////////////////////////////////////////////////////
// 	// TESTE dao.AdicionaPessoa
// 	///////////////////////////////////////////////////////////////
// 	var db = dao.GetDB()
// 	p, _ := pessoa.GetPessoaTest()
// 	p.Cpf = "38674832680"
// 	p.Usuario = "teste_inclusao"
// 	p, err := dao.AdicionaPessoa(db, p)
// 	fmt.Println(p, err)
// 	///////////////////////////////////////////////////////////////

// 	///////////////////////////////////////////////////////////////
// 	// TESTE dao.ProcuraPessoa
// 	///////////////////////////////////////////////////////////////
// 	var db = dao.GetDB()
// 	cpf := "38674832680"
// 	p, err := dao.ProcuraPessoa(db, cpf)
// 	fmt.Println(p, err)
// 	//////////////////////////////////////////////////////////////

// 	//////////////////////////////////////////////////////////////
// 	// TESTE dao.RemovePessoa
// 	//////////////////////////////////////////////////////////////
// 	db := dao.GetDB()
// 	cpf := "38674832680"
// 	err := dao.RemovePessoa(db, cpf)
// 	fmt.Println(err)
// 	//////////////////////////////////////////////////////////////

// }
