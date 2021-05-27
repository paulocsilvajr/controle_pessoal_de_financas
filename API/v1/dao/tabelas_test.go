package dao

import (
	"testing"
	"time"

	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/model/conta"
	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/model/detalhe_lancamento"
	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/model/lancamento"
	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/model/pessoa"
	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/model/tipo_conta"
	"gorm.io/gorm"
)

var db2 *gorm.DB

func init() {
	CreateDBParaTestes()

	db2 = GetDB02ParaTestes()
}

func TestCriarTabelaPessoa(t *testing.T) {
	err := CriarTabelaPessoa(db2)
	if err != nil {
		t.Error(err)
	}
}

func TestCRUDPessoa(t *testing.T) {
	// Criar - INSERT
	p1 := getTPessoa1()

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
	tc1 := getTTipoConta1()

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
	tc := getTTipoConta1()
	tc1 := &tc

	c := getTConta1(*tc1)
	c1 := &c

	c = getTConta2(tc, c)
	c2 := &c

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

func TestCriarTabelaLancamento(t *testing.T) {
	err := CriarTabelaLancamento(db2)
	if err != nil {
		t.Error(err)
	}
}

func TestCRUDTabelaLancamento(t *testing.T) {
	// Criar - INSERT
	p1 := getTPessoa1()

	err := db2.Create(&p1).Error
	if err != nil {
		t.Error(err)
	}

	l := getTLancamento1(p1)
	l1 := &l

	err = db2.Create(l1).Error
	if err != nil {
		t.Error(err)
	}

	// Alterar - UPDATE
	l1.Data = time.Date(2022, time.Month(12), 24, 0, 0, 0, 0, time.UTC)
	l1.Numero = setNullString("010101")
	l1.Descricao = "Lançamento 010101"

	err = db2.Save(l1).Error
	if err != nil {
		t.Error(err)
	}

	// Consultar - SELECT
	id := l1.ID
	l1 = nil

	err = db2.Where("id = ?", id).First(&l1).Error
	// t.Error(l1)
	if err != nil {
		t.Error(err)
	}

	// Remover - DELETE
	err = db2.Delete(l1).Error
	if err != nil {
		t.Error(err)
	}

	err = db2.Delete(p1).Error
	if err != nil {
		t.Error(err)
	}
}

func TestCriarTabelaDetalheLancamento(t *testing.T) {
	err := CriarTabelaDetalheLancamento(db2)
	if err != nil {
		t.Error(err)
	}
}

func TestCRUDTabelaDetalheLancamento(t *testing.T) {
	// Criar - INSERT
	p1 := getTPessoa1()
	err := db2.Create(&p1).Error
	if err != nil {
		t.Error(err)
	}

	l1 := getTLancamento1(p1)
	err = db2.Create(&l1).Error
	if err != nil {
		t.Error(err)
	}

	tc1 := getTTipoConta1()
	err = db2.Create(&tc1).Error
	if err != nil {
		t.Error(err)
	}

	c1 := getTConta1(tc1)
	err = db2.Create(&c1).Error
	if err != nil {
		t.Error(err)
	}

	c2 := getTConta2(tc1, c1)
	err = db2.Create(&c2).Error
	if err != nil {
		t.Error(err)
	}

	dl1 := detalhe_lancamento.TDetalheLancamento{
		IDLancamento: l1.ID,
		NomeConta:    c1.Nome,
		Debito:       setNullFloat64(0),
		Credito:      setNullFloat64(101),
	}

	dl2 := detalhe_lancamento.TDetalheLancamento{
		IDLancamento: l1.ID,
		NomeConta:    c2.Nome,
		Debito:       setNullFloat64(101),
		Credito:      setNullFloat64(0),
	}

	err = db2.Create(&dl1).Error
	if err != nil {
		t.Error(err)
	}

	err = db2.Create(&dl2).Error
	if err != nil {
		t.Error(err)
	}

	// Alterar - UPDATE
	dl1.Debito = setNullFloat64(200)
	dl1.Credito = setNullFloat64(0)

	dl2.Debito = dl1.Credito
	dl2.Credito = dl1.Debito

	err = db2.Save(&dl1).Error
	if err != nil {
		t.Error(err)
	}

	err = db2.Save(&dl2).Error
	if err != nil {
		t.Error(err)
	}

	// Consultar - SELECT
	id := dl1.IDLancamento
	dls := []detalhe_lancamento.TDetalheLancamento{}

	err = db2.Where("id_lancamento = ?", id).Find(&dls).Error
	// t.Error(dls, dl1, dl2)
	if err != nil {
		t.Error(err)
	}

	// Remover - DELETE
	err = db2.Delete(dl1).Error
	if err != nil {
		t.Error(err)
	}

	err = db2.Delete(dl2).Error
	if err != nil {
		t.Error(err)
	}

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

	err = db2.Delete(l1).Error
	if err != nil {
		t.Error(err)
	}

	err = db2.Delete(p1).Error
	if err != nil {
		t.Error(err)
	}

}

func getTPessoa1() pessoa.TPessoa {
	return pessoa.TPessoa{
		Cpf:          "12345678910",
		NomeCompleto: "Teste 01",
		Usuario:      "teste01",
		Senha:        "123456",
		Email:        "teste01@email.com",
	}
}

func getTPessoaAdmin1() pessoa.TPessoa {
	return pessoa.TPessoa{
		Cpf:          "00000000000",
		NomeCompleto: "Administrador 01",
		Usuario:      "admin01",
		Senha:        "123456",
		Email:        "admin01@email.com",
	}
}

func getTTipoConta1() tipo_conta.TTipoConta {
	return tipo_conta.TTipoConta{
		Nome:             "banco",
		DescricaoDebito:  "saque",
		DescricaoCredito: "depósito",
	}
}

func getTConta1(tc tipo_conta.TTipoConta) conta.TConta {
	return conta.TConta{
		Nome:          "Juros",
		NomeTipoConta: tc.Nome,
		Codigo:        setNullString("001"),
		Comentario:    setNullString("teste de conta 001 em banco"),
	}
}

func getTConta2(tc tipo_conta.TTipoConta, c conta.TConta) conta.TConta {
	return conta.TConta{
		Nome:          "Bradesco - juros",
		NomeTipoConta: tc.Nome,
		Codigo:        setNullString("002"),
		ContaPai:      setNullString(c.Nome),
		Comentario:    setNullString("teste de conta 002 em banco"),
	}
}

func getTLancamento1(p pessoa.TPessoa) lancamento.TLancamento {
	return lancamento.TLancamento{
		CpfPessoa: p.Cpf,
		Numero:    setNullString("101010"),
		Descricao: "Lançamento teste 101010",
	}
}
