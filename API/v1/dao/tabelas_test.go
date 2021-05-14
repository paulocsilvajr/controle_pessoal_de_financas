package dao

import (
	"testing"

	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/model/pessoa"
	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/model/tipo_conta"
	"gorm.io/gorm"
)

var db2 *gorm.DB = GetDB02()

func TestCriarTabelaPessoa(t *testing.T) {
	err := CriarTabelaPessoa(db2)
	if err != nil {
		t.Error(err)
	}
}

func TestCRUDPessoa(t *testing.T) {
	// Criar - INSERT
	p1 := pessoa.TPessoa{
		Cpf:          "12345678910",
		NomeCompleto: "Teste 01",
		Usuario:      "teste01",
		Senha:        "123456",
		Email:        "teste01@email.com",
	}

	p2 := pessoa.TPessoa{
		Cpf:          "10987654321",
		NomeCompleto: "Teste 02",
		Usuario:      "teste02",
		Senha:        "654321",
		Email:        "teste02@email.com",
	}

	err := db2.Create(&p1).Error
	if err != nil {
		t.Error(err)
	}

	err = db2.Create(&p2).Error
	if err != nil {
		t.Error(err)
	}

	// Alterar - UPDATE
	p1.NomeCompleto = "Pessoa Teste 01"
	p1.Usuario = "t01"
	p1.Senha = "159753"

	err = db2.Save(&p1).Error
	if err != nil {
		t.Error(err)
	}

	// Consultar - SELECT
	p1Cpf := p1.Cpf
	p2Cpf := p2.Cpf
	p1, p2 = pessoa.TPessoa{}, pessoa.TPessoa{}

	err = db2.First(&p1, p1Cpf).Error
	// t.Error(&p1)
	if err != nil {
		t.Error(err)
	}

	err = db2.First(&p2, p2Cpf).Error
	// t.Error(&p2)
	if err != nil {
		t.Error(err)
	}

	// Remover - DELETE
	err = db2.Delete(&p1).Error
	if err != nil {
		t.Error(err)
	}

	err = db2.Delete(&p2).Error
	if err != nil {
		t.Error(err)
	}
}

func TestCriarTabelaTipoConta(t *testing.T) {
	err := CriarTabelaTipoConta(db2)
	if err != nil {
		t.Error(err)
	}
}

func TestCRUDTipoConta(t *testing.T) {
	// Criar - INSERT
	tc1 := tipo_conta.TTipoConta{
		Nome:             "banco",
		DescricaoDebito:  "saque",
		DescricaoCredito: "depósito",
	}

	tc2 := tipo_conta.TTipoConta{
		Nome:             "carteira",
		DescricaoDebito:  "gastar",
		DescricaoCredito: "receber",
	}

	err := db2.Create(&tc1).Error
	if err != nil {
		t.Error(err)
	}

	err = db2.Create(&tc2).Error
	if err != nil {
		t.Error(err)
	}

	// Alterar - UPDATE
	tc1.DescricaoDebito = "retiradas"

	err = db2.Save(&tc1).Error
	if err != nil {
		t.Error(err)
	}

	tc2.DescricaoDebito = "gastos"
	tc2.DescricaoCredito = "recebimentos"

	err = db2.Save(&tc2).Error
	if err != nil {
		t.Error(err)
	}

	// Consultar - SELECT
	tc1Nome := tc1.Nome
	tc2Nome := tc2.Nome
	tc1, tc2 = tipo_conta.TTipoConta{}, tipo_conta.TTipoConta{}

	err = db2.Where("nome = ?", tc1Nome).First(&tc1).Error
	// t.Error(&tc1)
	if err != nil {
		t.Error(err)
	}

	err = db2.Where("nome = ?", tc2Nome).First(&tc2).Error
	// t.Error(&tc2)
	if err != nil {
		t.Error(err)
	}

	// Remover - DELETE
	err = db2.Delete(&tc1).Error
	if err != nil {
		t.Error(err)
	}

	err = db2.Delete(&tc2).Error
	if err != nil {
		t.Error(err)
	}
}
