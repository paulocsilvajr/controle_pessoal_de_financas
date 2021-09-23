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
	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/model/tipo_conta"
	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/route"
)

func main() {
	helper.CriarDiretorioSeNaoExistir("config")

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
		certFile := "keys/new.cert.cert"
		keyFile := "keys/new.cert.key"
		log.Fatal(http.ListenAndServeTLS(endereco, certFile, keyFile, router))
	} else {
		log.Fatal(http.ListenAndServe(endereco, router))
	}
}

func criarUsuarioAdminInicial() {
	db := dao.GetDB02()
	defer dao.CloseDB(db)

	admin := pessoa.New("00000000000", "Administrador", "admin", "admin", "meuemail@email.com")

	_, err := dao.ProcuraPessoaPorUsuario02(db, admin.Usuario)
	var novaPessoa *pessoa.Pessoa
	if err != nil {
		novaPessoa, err = dao.AdicionaPessoaAdmin02(db, admin)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Novo usuário ADMINISTRADOR:", novaPessoa)
		}
	} else {
		confirmacao := inputString("Usuário admin já existe, deseja resetar para a senha padrão?[s/N]: ")
		if strings.ToLower(confirmacao) == "s" {
			novaPessoa, err = dao.AlteraPessoaPorUsuario02(db, admin.Usuario, admin)
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Println("Novo usuário ADMINISTRADOR:", novaPessoa)
			}
		}
	}

	err = dao.CloseDB(db)
	if err != nil {
		log.Fatal(err)
	}
}

func criarContasPadrao() {
	db := dao.GetDB02()
	defer dao.CloseDB(db)

	var tiposConta tipo_conta.TiposConta
	tc1, _ := tipo_conta.NewTipoConta("banco", "saque", "deposito")
	tc2, _ := tipo_conta.NewTipoConta("carteira", "gastar", "receber")
	tc3, _ := tipo_conta.NewTipoConta("despesa", "desconto", "despesa")
	tc4, _ := tipo_conta.NewTipoConta("cartão de crédito", "cobrar", "pagamento")
	tc5, _ := tipo_conta.NewTipoConta("receita", "receita", "cobrar")
	tc6, _ := tipo_conta.NewTipoConta("ativo", "diminuir", "aumentar")
	tc7, _ := tipo_conta.NewTipoConta("passivo", "aumentar", "diminuir")
	tc8, _ := tipo_conta.NewTipoConta("líquido", "aumentar", "diminuir")

	tiposConta = append(tiposConta, tc1, tc2, tc3, tc4, tc5, tc6, tc7, tc8)
	var err error

	for _, tc := range tiposConta {
		_, err = dao.AdicionaTipoConta02(db, tc)
		if err != nil {
			// fmt.Printf("ERRO ao adicionar o tipo conta padrão '%s'[%s]", tc.Nome, err)
			continue
		}
		fmt.Printf("Adicionado tipo conta padrão '%s'", tc.Nome)
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
  -r, --routes        exibe as métodos/rotas cadastradas na API
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
