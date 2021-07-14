package dao

import (
	"fmt"
	"strings"
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

	p2 := &pessoa.TPessoa{
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
	p1 = new(pessoa.TPessoa)
	p2 = new(pessoa.TPessoa)

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
	tc1, tc2 = &tipo_conta.TTipoConta{}, tipo_conta.TTipoConta{}

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

	// ERROR: constraint "tipo_conta_fk" for relation "conta" already exists (SQLSTATE 42710)
	if err := verificaErroConstraintExists(err); err != nil {
		t.Error(err)
	}
}

func TestCRUDTabelaConta(t *testing.T) {
	// Criar - INSERT
	tc1 := getTTipoConta1()
	tx := db2.Create(tc1)
	err := tx.Error
	if err != nil {
		t.Error(err)
	}
	if err := VerificaQuantidadeRegistrosAfetados(tx, 1); err != nil {
		t.Error(err)
	}

	c1 := getTConta1(tc1)
	tx = db2.Create(c1)
	err = tx.Error
	if err != nil {
		t.Error(err)
	}
	if err := VerificaQuantidadeRegistrosAfetados(tx, 1); err != nil {
		t.Error(err)
	}

	c2 := getTConta2(tc1, c1)
	tx = db2.Create(c2)
	err = tx.Error
	if err != nil {
		t.Error(err)
	}
	if err := VerificaQuantidadeRegistrosAfetados(tx, 1); err != nil {
		t.Error(err)
	}

	// Alterar - UPDATE
	tx = db2.Model(c1).Update("nome", "Juros recebidos")
	err = tx.Error
	if err != nil {
		t.Error(err)
	}
	if err := VerificaQuantidadeRegistrosAfetados(tx, 1); err != nil {
		t.Error(err)
	}

	// porque foi alterado a chave primária de "c1", é necessário consultar "c2" em BD para pegar a entidade atualizada, com o campo conta_pai atualizado, para que ao fazer update não dê conflito de chave primária. Se não for obtido "c2" em consulta, o GORM tenta inserir um novo pelo método Save
	err = db2.First(c2).Error
	if err != nil {
		t.Error(err)
	}
	c2.Codigo = setNullString("002a")
	c2.Comentario = setNullString("alteração em conta 002 para conta 002a")
	tx = db2.Save(c2)
	err = tx.Error
	if err != nil {
		t.Error(err)
	}
	if err := VerificaQuantidadeRegistrosAfetados(tx, 1); err != nil {
		t.Error(err)
	}

	// Consultar - SELECT
	c1PK := c1.Nome
	c2PK := c2.Nome
	c3 := new(conta.TConta)
	c4 := new(conta.TConta)

	tx = db2.Where(fmt.Sprintf("%s = ?", GetPrimaryKeyConta()), c1PK).First(c3)
	err = tx.Error
	if err != nil {
		t.Error(err)
	}
	if err := VerificaQuantidadeRegistrosAfetados(tx, 1); err != nil {
		t.Error(err)
	}

	tx = db2.Where(fmt.Sprintf("%s = ?", GetPrimaryKeyConta()), c2PK).First(c4)
	err = tx.Error
	if err != nil {
		t.Error(err)
	}
	if err := VerificaQuantidadeRegistrosAfetados(tx, 1); err != nil {
		t.Error(err)
	}

	// Remover - DELETE
	tx = db2.Delete(c2)
	err = tx.Error
	if err != nil {
		t.Error(err)
	}
	if err := VerificaQuantidadeRegistrosAfetados(tx, 1); err != nil {
		t.Error(err)
	}

	tx = db2.Delete(c1)
	err = tx.Error
	if err != nil {
		t.Error(err)
	}
	if err := VerificaQuantidadeRegistrosAfetados(tx, 1); err != nil {
		t.Error(err)
	}

	tx = db2.Delete(tc1)
	err = tx.Error
	if err != nil {
		t.Error(err)
	}
	if err := VerificaQuantidadeRegistrosAfetados(tx, 1); err != nil {
		t.Error(err)
	}
}

func TestCriarTabelaLancamento(t *testing.T) {
	err := CriarTabelaLancamento(db2)

	// ERROR: constraint "pessoa_lancamento_fk" for relation "lancamento" already exists (SQLSTATE 42710)
	if err := verificaErroConstraintExists(err); err != nil {
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

	l1 := getTLancamento1(p1)

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

	// ERROR: constraint "conta_detalhe_lancamento_fk" for relation "detalhe_lancamento" already exists (SQLSTATE 42710)
	if err := verificaErroConstraintExists(err); err != nil {
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

func getTPessoa1() *pessoa.TPessoa {
	p, err := pessoa.NewPessoa("12345678910", "Teste 01", "teste01", "123456", "teste01@email.com")
	if err != nil {
		return nil
	}
	return ConvertePessoaParaTPessoa(p)
}

func getTPessoaAdmin1() *pessoa.TPessoa {
	p, err := pessoa.NewPessoaAdmin("00000000000", "Administrador 01", "admin01", "123456", "admin01@email.com")
	if err != nil {
		return nil
	}
	return ConvertePessoaParaTPessoa(p)
}

func getTTipoConta1() *tipo_conta.TTipoConta {
	tc, err := tipo_conta.NewTipoConta("banco", "saque", "depósito")
	if err != nil {
		return nil
	}

	return ConverteTipoContaParaTTipoConta(tc)
}

func getTConta1(tc *tipo_conta.TTipoConta) *conta.TConta {
	c, err := conta.NewConta("Juros", tc.Nome, "001", "", "teste de conta 001 em banco")
	if err != nil {
		return nil
	}
	return ConverteContaParaTConta(c)
}

func getTConta2(tc *tipo_conta.TTipoConta, cp *conta.TConta) *conta.TConta {
	c, err := conta.NewConta("Bradesco - juros", tc.Nome, "002", cp.Nome, "teste de conta 001 em banco")
	if err != nil {
		return nil
	}
	return ConverteContaParaTConta(c)
}

func getTLancamento1(p *pessoa.TPessoa) *lancamento.TLancamento {
	l, err := lancamento.NewLancamento02(p.Cpf, time.Now(), "101010", "Lançamento teste 101010")
	if err != nil {
		return nil
	}
	return ConverteLancamentoParaTLancamento(l)
}

func verificaErroConstraintExists(err error) error {
	if err != nil {
		erroConstraintExists := strings.Contains(err.Error(), "already exists")
		if !erroConstraintExists {
			return err
		}
	}
	return nil
}
