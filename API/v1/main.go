package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/jedib0t/go-pretty/table"
	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/config"
	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/dao"
	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/helper"
	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/route"
)

func main() {
	helper.CriarDiretorioAbs("config")

	configuracoes := config.AbrirConfiguracoes()
	porta := helper.FormatarPorta(configuracoes["porta"])
	host := configuracoes["host"]
	protocolo := configuracoes["protocolo"]

	if exitCode := verificaParametrosInicializacao(); exitCode >= 0 {
		os.Exit(exitCode)
	}

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
		dir, err := helper.GetDiretorioAbs()
		if err != nil {
			log.Fatal(err)
		}
		certFile := path.Join(dir, "keys/new.cert.cert")
		keyFile := path.Join(dir, "keys/new.cert.key")

		log.Fatal(http.ListenAndServeTLS(endereco, certFile, keyFile, router))
	} else {
		log.Fatal(http.ListenAndServe(endereco, router))
	}
}

func verificaParametrosInicializacao() int {
	args := os.Args

	if len(args) >= 2 {
		switch args[1] {
		case "--init", "-i":
			if err := dao.CreateDB(); err != nil {
				panic(err)
			}

			if err := dao.CriarTabelas(); err != nil {
				panic(err)
			}

			criarUsuarioAdminInicial()

			criarContasPadrao()

			fmt.Println("\nCriado banco de dados, tabelas e usuário Admin, se essas estruturas não existirem.\nReexecute a API sem parâmetros para reconhecer o banco de dados e iniciar o seu uso")
			return 0
		case "--routes", "-r":
			imprimeRotas()
			return 0
		case "--help", "-h":
			fmt.Printf(`Uso: %s [ -h | --help | -i | --init ]
Inicia a API do "Controle Pessoa de Finanças"
Argumentos:
  -i, --init         cria o banco de dados(de acordo com config.json), as tabelas e o usuário administrador inicial "admin" com senha "admin"
  -r, --routes       exibe as métodos/rotas cadastradas na API
  -h, --help         exibe essa ajuda
`, args[0])
			return 1
		default:
			fmt.Println("Argumento informado inválido. Use '-h' e '--help' para obter ajuda")
			return 1
		}
	}
	return -1
}

func imprimeRotas() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"TIPO", "ROTA"})
	for _, rota := range config.Rotas {
		t.AppendRow([]interface{}{rota.Tipo, rota.Rota})
		// t.AppendSeparator()
	}
	t.Render()

	fmt.Println(`Para mais detalhes de cada rota, logue na API pelo rota '/login/{usuario}' com o corpo '{"usuario":"string", "senha":"string"}', armazene o token associado ao seu usuário e consulte a rota '/'`)
}

func inputString(prompt string) string {
	// fonte: https://tutorialedge.net/golang/reading-console-input-golang/
	var caracterNovaLinha byte = '\n'

	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	texto, _ := reader.ReadString(caracterNovaLinha)
	texto = strings.Replace(texto, "\n", "", -1)
	return texto
}
