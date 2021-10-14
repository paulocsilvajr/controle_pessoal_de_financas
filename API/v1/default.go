package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/dao"
	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/model/conta"
	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/model/pessoa"
	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/model/tipo_conta"
)

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
		fmt.Printf("Adicionado tipo conta padrão '%s'\n", tc.Nome)
	}

	var contas conta.Contas
	c1, _ := conta.NewConta("ativos", "ativo", "1", "", "")
	c2, _ := conta.NewConta("despesas", "despesa", "2", "", "")
	c3, _ := conta.NewConta("líquidos", "líquido", "3", "", "")
	c4, _ := conta.NewConta("passivos", "passivo", "4", "", "")
	c5, _ := conta.NewConta("receitas", "receita", "5", "", "")
	c6, _ := conta.NewConta("conta corrente", "banco", "6", "ativos", "")
	c7, _ := conta.NewConta("conta poupança", "banco", "7", "ativos", "")
	c8, _ := conta.NewConta("dinheiro em carteira", "carteira", "8", "ativos", "")
	c9, _ := conta.NewConta("banco teste 1", "banco", "9", "conta corrente", "")
	c10, _ := conta.NewConta("banco teste 2", "banco", "10", "conta poupança", "")
	c11, _ := conta.NewConta("internet", "despesa", "11", "despesas", "")
	c12, _ := conta.NewConta("telefone", "despesa", "12", "despesas", "")
	c13, _ := conta.NewConta("serviços", "despesa", "13", "despesas", "")
	c14, _ := conta.NewConta("eletricidade", "despesa", "14", "serviços", "")
	c15, _ := conta.NewConta("refeições fora", "despesa", "15", "despesas", "")
	c16, _ := conta.NewConta("computador", "despesa", "16", "despesas", "")

	contas = append(contas, c1, c2, c3, c4, c5, c6, c7, c8, c9, c10, c11, c12, c13, c14, c15, c16)
	for _, c := range contas {
		_, err = dao.AdicionaConta02(db, c)
		if err != nil {
			// fmt.Printf("ERRO ao adicionar o conta padrão '%s'[%s]", c.Nome, err)
			continue
		}
		fmt.Printf("Adicionado conta padrão '%s'\n", c.Nome)
	}
}
