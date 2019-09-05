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
