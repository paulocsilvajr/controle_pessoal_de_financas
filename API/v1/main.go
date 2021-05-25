package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/table"
	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/config"
	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/dao"
	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/helper"
	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/model/pessoa"
	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/route"
)

func main() {
	helper.CriarDiretorioSeNaoExistir("config")

	configuracoes := config.AbrirConfiguracoes()
	porta := helper.FormatarPorta(configuracoes["porta"])
	host := configuracoes["host"]
	protocolo := configuracoes["protocolo"]

	verificaParametrosInicializacao()

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

func criarUsuarioAdminInicial() {
	db := dao.GetDB()

	admin := pessoa.New("00000000000", "Administrador", "admin", "admin", "meuemail@email.com")

	_, err := dao.ProcuraPessoaPorUsuario(db, admin.Usuario)
	novaPessoa := new(pessoa.Pessoa)
	if err != nil {
		novaPessoa, err = dao.AdicionaPessoaAdmin(db, admin)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Novo usuário ADMINISTRADOR:", novaPessoa)
		}
	} else {
		confirmacao := inputString("Usuário admin já existe, deseja resetar para a senha padrão?[s/N]: ")
		if strings.ToLower(confirmacao) == "s" {
			novaPessoa, err = dao.AlteraPessoaPorUsuario(db, admin.Usuario, admin)
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Println("Novo usuário ADMINISTRADOR:", novaPessoa)
			}
		}
	}

	err = db.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func verificaParametrosInicializacao() {
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

			// criarUsuarioAdminInicial()

			fmt.Println("\nCriado banco de dados, tabelas e usuário Admin.\nReexecute a API sem parâmetros para reconhecer o banco de dados e iniciar o seu uso")
			os.Exit(0)
		case "--rotes", "-r":
			imprimeRotas()
			os.Exit(0)
		case "--help", "-h":
			fmt.Printf(`Uso: %s [ -h | --help | -i | --init ]
Inicia a API do "Controle Pessoa de Finanças"
Argumentos:
  -i, --init         cria o banco de dados(de acordo com config.json), as tabelas e o usuário administrador inicial "admin" com senha "admin"
  -r, --routes        exibe as métodos/rotas cadastradas na API
  -h, --help         exibe essa ajuda
`, args[0])
			os.Exit(1)
		}
	}
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
