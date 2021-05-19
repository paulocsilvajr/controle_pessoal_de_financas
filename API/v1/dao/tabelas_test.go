package dao

import (
	"testing"

	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/model/conta"
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

func TestCriarTabelaConta(t *testing.T) {
	err := CriarTabelaConta(db2)
	if err != nil {
		t.Error(err)
	}
}

func TestCRUDTabelaConta(t *testing.T) {
	// Criar - INSERT
	tc1 := &tipo_conta.TTipoConta{
		Nome:             "banco",
		DescricaoDebito:  "saque",
		DescricaoCredito: "depósito",
	}

	c1 := &conta.TConta{
		Nome:          "Juros",
		NomeTipoConta: tc1.Nome,
		Codigo:        setNullString("001"),
		Comentario:    setNullString("teste de conta 001 em banco"),
	}

	c2 := &conta.TConta{
		Nome:          "Bradesco - juros",
		NomeTipoConta: tc1.Nome,
		Codigo:        setNullString("002"),
		ContaPai:      setNullString(c1.Nome),
		Comentario:    setNullString("teste de conta 002 em banco"),
	}

	err := db2.Create(tc1).Error
	if err != nil {
		t.Error(err)
	}

	err = db2.Create(c1).Error
	if err != nil {
		t.Error(err)
	}

	err = db2.Create(c2).Error
	if err != nil {
		t.Error(err)
	}

	// Alterar - UPDATE
	err = db2.Model(c1).Update("nome", "Juros recebidos").Error
	if err != nil {
		t.Error(err)
	}

	// porque foi alterado a chave primária de "c1", é necessário consultar "c2" em BD para pegar a entidade atualizada, com o campo conta_pai atualizado, para quando for fazer update não dê conflito de chave primária. Se não for obtido "c2" em consulta, o GORM tenta inserir um novo pelo método Save
	err = db2.First(c2).Error
	if err != nil {
		t.Error(err)
	}

	c2.Codigo = setNullString("002a")
	c2.Comentario = setNullString("alteração em conta 002 para conta 002a")

	err = db2.Save(c2).Error
	if err != nil {
		t.Error(err)
	}

	// Consultar - SELECT
	c1Nome := c1.Nome
	c2Nome := c2.Nome
	c1, c2 = nil, nil

	err = db2.Where("nome = ?", c1Nome).First(&c1).Error
	// t.Error(c1)
	if err != nil {
		t.Error(err)
	}

	err = db2.Where("nome = ?", c2Nome).First(&c2).Error
	// t.Error(c2)
	if err != nil {
		t.Error(err)
	}

	// Remover - DELETE
	err = db2.Delete(c1).Error
	if err != nil {
		t.Error(err)
	}

	err = db2.Delete(c2).Error
	if err != nil {
		t.Error(err)
	}

	err = db2.Delete(tc1).Error
	if err != nil {
		t.Error(err)
	}
}
