package main

import (
	"bufio"
	"controle_pessoal_de_financas/API/v1/config"
	"controle_pessoal_de_financas/API/v1/dao"
	"controle_pessoal_de_financas/API/v1/helper"
	"controle_pessoal_de_financas/API/v1/model/pessoa"
	"controle_pessoal_de_financas/API/v1/route"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
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

	admin, _ := pessoa.NewPessoaAdmin("00000000000", "Administrador", "admin", "admin", "meuemail@email.com")

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
			criarUsuarioAdminInicial()
		case "--help", "-h":
			fmt.Printf(`Uso: %s [ -h | --help | -i | --init ]
Inicia a API do "Controle Pessoa de Finanças"
Argumentos:
  -i, --init         cria o usuário administrador inicial "admin" com senha "admin"
  -h, --help         exibe essa ajuda
`, args[0])
			os.Exit(1)
		}
	}
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
