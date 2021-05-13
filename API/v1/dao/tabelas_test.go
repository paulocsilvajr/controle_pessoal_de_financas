package dao

import (
	"testing"

	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/model/pessoa"
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
	p1 := pessoa.Pessoa{
		Cpf:          "12345678910",
		NomeCompleto: "Teste 01",
		Usuario:      "teste01",
		Senha:        "123456",
		Email:        "teste01@email.com",
	}

	p2 := pessoa.Pessoa{
		Cpf:          "10987654321",
		NomeCompleto: "Teste 02",
		Usuario:      "teste02",
		Senha:        "654321",
		Email:        "teste02@email.com",
	}

	err := db2.Create(&p2).Error
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
	p1, p2 = pessoa.Pessoa{}, pessoa.Pessoa{}

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
