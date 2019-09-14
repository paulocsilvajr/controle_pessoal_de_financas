package conta

import (
	"controle_pessoal_de_financas/API/v1/model/tipo_conta"
	"testing"
	"time"
)

func TestIConta(t *testing.T) {
	c1 := &Conta{}
	c2 := new(Conta)

	var ic1 IConta = c1
	c3, ok := ic1.(*Conta)
	if !ok {
		t.Error(c3)
	}

	var ic2 IConta = c2
	c4, ok := ic2.(*Conta)
	if !ok {
		t.Error(c4)
	}

	defer func() {
		if r := recover(); r != nil {
			t.Error(r)
		}
	}()

	var ic3 IConta
	ic3 = GetContaTest()
	_ = ic3.String()
	ic3.Repr()
	ic3.VerificaAtributos()
	ic3.Altera("Novo nome de conta", "Outro tipo de conta", "003", c4.Nome, "Alterando conta")
	campos := map[string]string{
		"nome":          "Novo nome de conta 02",
		"nomeTipoConta": "Outro tipo de conta 02",
		"codigo":        "004",
		"contaPai":      "",
		"comentario":    "Alterando conta 02"}
	ic3.AlteraCampos(campos)
	ic3.Ativa()
	ic3.Inativa()
}

func TestConverteParaConta(t *testing.T) {
	c1 := &Conta{}
	c2 := new(Conta)

	c3, err := converteParaConta(c1)
	if err != nil {
		t.Error(c3)
	}

	c4, err := converteParaConta(c2)
	if err != nil {
		t.Error(c4)
	}
}

func TestNew(t *testing.T) {
	tipoConta := tipo_conta.GetTipoContaTest()
	nome := "Banco 02"
	codigo := "002"
	contaPai := ""
	comentario := "teste de conta 02"
	data := time.Now().Local()
	estadoPadrao := true

	c1 := New(nome, tipoConta.Nome, codigo, contaPai, comentario)

	if c1.Nome != nome {
		t.Error("Erro em função conta.New, atributo Nome", c1)
	}

	if c1.NomeTipoConta != tipoConta.Nome {
		t.Error("Erro em função conta.New, atributo NomeTipoConta", c1)
	}

	if c1.Codigo != codigo {
		t.Error("Erro em função conta.New, atributo Codigo", c1)
	}

	if c1.ContaPai != contaPai {
		t.Error("Erro em função conta.New, atributo ContaPai", c1)
	}

	if c1.Comentario != comentario {
		t.Error("Erro em função conta.New, atributo Comentario", c1)
	}

	if c1.DataCriacao.Unix() != data.Unix() {
		t.Error("Erro em função conta.New, atributo DataCriacao", c1)
	}

	if c1.DataModificacao.Unix() != data.Unix() {
		t.Error("Erro em função conta.New, atributo DataModificacao", c1)
	}

	if c1.Estado != estadoPadrao {
		t.Error("Erro em função conta.New, atributo Estado", c1)
	}
}

func TestNewConta(t *testing.T) {
	tipoConta := tipo_conta.GetTipoContaTest()
	nome := "Banco 02"
	codigo := "002"
	contaPai := ""
	comentario := "teste de conta 02"
	data := time.Now().Local()
	estadoPadrao := true

	c1, err := NewConta(nome, tipoConta.Nome, codigo, contaPai, comentario)

	if c1.Nome != nome {
		t.Error("Erro em função conta.New, atributo Nome", c1)
	}

	if c1.NomeTipoConta != tipoConta.Nome {
		t.Error("Erro em função conta.New, atributo NomeTipoConta", c1)
	}

	if c1.Codigo != codigo {
		t.Error("Erro em função conta.New, atributo Codigo", c1)
	}

	if c1.ContaPai != contaPai {
		t.Error("Erro em função conta.New, atributo ContaPai", c1)
	}

	if c1.Comentario != comentario {
		t.Error("Erro em função conta.New, atributo Comentario", c1)
	}

	if c1.DataCriacao.Unix() != data.Unix() {
		t.Error("Erro em função conta.New, atributo DataCriacao", c1)
	}

	if c1.DataModificacao.Unix() != data.Unix() {
		t.Error("Erro em função conta.New, atributo DataModificacao", c1)
	}

	if c1.Estado != estadoPadrao {
		t.Error("Erro em função conta.New, atributo Estado", c1)
	}

	if err != nil {
		t.Error(err, c1)
	}

	comentario = ""
	c2, err := NewConta(nome, tipoConta.Nome, codigo, contaPai, comentario)
	if err.Error() != "Tamanho de campo Comentário inválido[0 caracter(es)]" {
		t.Error("Erro em função conta.NewConta, nao retornou o erro esperado", c2, err)
	}

	comentario = "Descrição de conta muito longa para dar erro em teste unitário de modelo Conta ......... ........... ............. ............. ..................."
	c2, err = NewConta(nome, tipoConta.Nome, codigo, contaPai, comentario)
	if err.Error() != "Tamanho de campo Comentário inválido[151 caracter(es)]" {
		t.Error("Erro em função conta.NewConta, nao retornou o erro esperado", c2, err)
	}

	contaPai = "Nome de Conta Pai muito longa para dar erro em teste unitário de modelo Conta"
	c2, err = NewConta(nome, tipoConta.Nome, codigo, contaPai, comentario)
	if err.Error() != "Tamanho de campo Nome da Conta pai inválido[78 caracter(es)]" {
		t.Error("Erro em função conta.NewConta, nao retornou o erro esperado", c2, err)
	}

	codigo = "12345678901234567890"
	c2, err = NewConta(nome, tipoConta.Nome, codigo, contaPai, comentario)
	if err.Error() != "Tamanho de campo Código inválido[20 caracter(es)]" {
		t.Error("Erro em função conta.NewConta, nao retornou o erro esperado", c2, err)
	}

	codigo = "ABCD123456789-20000101"
	c2, err = NewConta(nome, tipoConta.Nome, codigo, contaPai, comentario)
	if err.Error() != "Tamanho de campo Código inválido[22 caracter(es)]" {
		t.Error("Erro em função conta.NewConta, nao retornou o erro esperado", c2, err)
	}

	nomeTipoConta := ""
	c2, err = NewConta(nome, nomeTipoConta, codigo, contaPai, comentario)
	if err.Error() != "Tamanho de campo Nome do Tipo da Conta inválido[0 caracter(es)]" {
		t.Error("Erro em função conta.NewConta, nao retornou o erro esperado", c2, err)
	}

	nomeTipoConta = "Nome de Tipo de Conta muito longa para dar erro em teste unitário de modelo Conta"
	c2, err = NewConta(nome, nomeTipoConta, codigo, contaPai, comentario)
	if err.Error() != "Tamanho de campo Nome do Tipo da Conta inválido[82 caracter(es)]" {
		t.Error("Erro em função conta.NewConta, nao retornou o erro esperado", c2, err)
	}

	nome = ""
	c2, err = NewConta(nome, nomeTipoConta, codigo, contaPai, comentario)
	if err.Error() != "Tamanho de campo Nome inválido[0 caracter(es)]" {
		t.Error("Erro em função conta.NewConta, nao retornou o erro esperado", c2, err)
	}

	nome = "Nome muito longo para dar erro em teste unitário de modelo Conta"
	c2, err = NewConta(nome, nomeTipoConta, codigo, contaPai, comentario)
	if err.Error() != "Tamanho de campo Nome inválido[65 caracter(es)]" {
		t.Error("Erro em função conta.NewConta, nao retornou o erro esperado", c2, err)
	}
}

func TestGetContaTeste(t *testing.T) {
	nome := "Ativos Teste 01"
	nomeTipoConta := "banco teste 01"
	codigo := "001"
	contaPai := ""
	comentario := "teste de Conta"
	criacao := time.Date(2000, 2, 1, 12, 30, 0, 0, new(time.Location))
	modificacao := criacao
	estado := true

	c1 := GetContaTest()

	if c1.Nome != nome {
		t.Error("Erro em função conta.GetContaTest, atributo Nome", c1)
	}

	if c1.NomeTipoConta != nomeTipoConta {
		t.Error("Erro em função conta.GetContaTest, atributo NomeTipoConta", c1)
	}

	if c1.Codigo != codigo {
		t.Error("Erro em função conta.GetContaTest, atributo Codigo", c1)
	}

	if c1.ContaPai != contaPai {
		t.Error("Erro em função conta.GetContaTest, atributo ContaPai", c1)
	}

	if c1.Comentario != comentario {
		t.Error("Erro em função conta.GetContaTest, atributo Comentario", c1)
	}

	if c1.DataCriacao.Unix() != criacao.Unix() {
		t.Error("Erro em função conta.GetContaTest, atributo DataCriacao", c1)
	}

	if c1.DataModificacao.Unix() != modificacao.Unix() {
		t.Error("Erro em função conta.GetContaTest, atributo DataModificacao", c1)
	}

	if c1.Estado != estado {
		t.Error("Erro em função conta.GetContaTest, atributo Estado", c1)
	}
}

func TestString(t *testing.T) {
	c1 := GetContaTest()

	if c1.String() != "Ativos Teste 01	banco teste 01	001		teste de Conta	01/02/2000 12:30:00	01/02/2000 12:30:00	ativo" {
		t.Error("Erro em função conta.String", c1.String())
	}
}

func TestRepr(t *testing.T) {
	c1 := GetContaTest()

	if c1.Repr() != "Ativos Teste 01	banco teste 01	001		teste de Conta	2000-02-01 12:30:00 +0000 UTC	2000-02-01 12:30:00 +0000 UTC	true" {
		t.Error("Erro em função conta.Repr", c1.Repr())
	}
}

func TestAltera(t *testing.T) {
	c1 := GetContaTest()

	err := c1.Altera("Ativos Teste 02", "banco teste 02", "002", "Ativos", "teste de Conta 02")
	if err != nil {
		t.Error(err, c1)
	}

	err = c1.Altera("Nome muito longo para dar erro em teste unitário de modelo Conta", "banco teste 02", "002", "Ativos", "teste de Conta 02")
	if err.Error() != "Tamanho de campo Nome inválido[65 caracter(es)]" {
		t.Error(err, c1)
	}

	err = c1.Altera("Ativos Teste 02", "Nome de Tipo de Conta muito longa para dar erro em teste unitário de modelo Conta", "002", "Ativos", "teste de Conta 02")
	if err.Error() != "Tamanho de campo Nome do Tipo da Conta inválido[82 caracter(es)]" {
		t.Error(err, c1)
	}

	err = c1.Altera("Ativos Teste 02", "banco teste 02", "ABCD123456789-20000101", "Ativos", "teste de Conta 02")
	if err.Error() != "Tamanho de campo Código inválido[22 caracter(es)]" {
		t.Error(err, c1)
	}

	err = c1.Altera("Ativos Teste 02", "banco teste 02", "002", "Nome de Conta Pai muito longa para dar erro em teste unitário de modelo Conta", "teste de Conta 02")
	if err.Error() != "Tamanho de campo Nome da Conta pai inválido[78 caracter(es)]" {
		t.Error(err, c1)
	}

	err = c1.Altera("Ativos Teste 02", "banco teste 02", "002", "Ativos", "Descrição de conta muito longa para dar erro em teste unitário de modelo Conta ......... ........... ............. ............. ...................")
	if err.Error() != "Tamanho de campo Comentário inválido[151 caracter(es)]" {
		t.Error(err, c1)
	}
}
