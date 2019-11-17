package lancamento

import (
	"controle_pessoal_de_financas/API/v1/model/pessoa"
	"fmt"
	"testing"
	"time"
)

func TestILancamento(t *testing.T) {
	l1 := &Lancamento{}
	l2 := new(Lancamento)
	lancamentos := Lancamentos{l1, l2}

	for _, l := range lancamentos {
		var il ILancamento = l
		l3, ok := il.(*Lancamento)
		if !ok {
			t.Error(l3)
		}
	}

	defer func() {
		if r := recover(); r != nil {
			t.Error(r)
		}
	}()

	var il ILancamento
	il = GetLancamentoTest()
	fmt.Println(il.String(), il.Repr())
	il.VerificaAtributos()
	il.Altera(99999, "99999999999", time.Now().Local(), "99999", "Nova descrição de Lançamento")
	campos := map[string]interface{}{
		"id":        99990,
		"cpf":       "99999999990",
		"data":      time.Now().Local(),
		"numero":    "99990",
		"descricao": "Mais nova descrição de Lançamento"}
	il.AlteraCampos(campos)
	il.Ativa()
	il.Inativa()
}

func TestConverteParaLancamento(t *testing.T) {
	l1 := &Lancamento{}
	l2 := new(Lancamento)
	lancamentos := Lancamentos{l1, l2}

	for _, l := range lancamentos {
		nl, err := converterParaLancamento(l)
		if err != nil {
			t.Error(nl)
		}
	}
}

func TestNew(t *testing.T) {
	pessoa, _ := pessoa.GetPessoaTest()
	id := 99999
	data := time.Now().Local()
	numero := "99999"
	descricao := "Pagamento de conta de telefone do mês 9 de 2000"
	estadoPadrao := true

	l := New(id, pessoa.Cpf, data, numero, descricao)

	if l.ID != id {
		t.Error("Erro em função lancamento.New, atributo ID", id)
	}

	if l.CpfPessoa != pessoa.Cpf {
		t.Error("Erro em função lancamento.New, atributo CpfPessoa", l)
	}

	if l.Data != data {
		t.Error("Erro em função lancamento.New, atributo Data", l)
	}

	if l.Numero != numero {
		t.Error("Erro em função lancamento.New, atributo Numero", l)
	}

	if l.Descricao != descricao {
		t.Error("Erro em função lancamento.New, atributo Descricao", l)
	}

	if l.DataCriacao.Unix() != data.Unix() {
		t.Error("Erro em função lancamento.New, atributo DataCriacao", l)
	}

	if l.DataModificacao.Unix() != data.Unix() {
		t.Error("Erro em função lancamento.New, atributo DataModificacao", l)
	}

	if l.Estado != estadoPadrao {
		t.Error("Erro em função lancamento.New, atributo Estado", l)
	}
}

func TestNewLancamento(t *testing.T) {
	pessoa, _ := pessoa.GetPessoaTest()
	id := 99999
	data := time.Now().Local()
	numero := "99999"
	descricao := "Pagamento de conta de telefone do mês 9 de 2000"
	estadoPadrao := true

	l, err := NewLancamento(id, pessoa.Cpf, data, numero, descricao)
	if err != nil {
		t.Error(err, l)
	}

	if l.ID != id {
		t.Error("Erro em função lancamento.New, atributo ID", id)
	}

	if l.CpfPessoa != pessoa.Cpf {
		t.Error("Erro em função lancamento.New, atributo CpfPessoa", l)
	}

	if l.Data != data {
		t.Error("Erro em função lancamento.New, atributo Data", l)
	}

	if l.Numero != numero {
		t.Error("Erro em função lancamento.New, atributo Numero", l)
	}

	if l.Descricao != descricao {
		t.Error("Erro em função lancamento.New, atributo Descricao", l)
	}

	if l.DataCriacao.Unix() != data.Unix() {
		t.Error("Erro em função lancamento.New, atributo DataCriacao", l)
	}

	if l.DataModificacao.Unix() != data.Unix() {
		t.Error("Erro em função lancamento.New, atributo DataModificacao", l)
	}

	if l.Estado != estadoPadrao {
		t.Error("Erro em função lancamento.New, atributo Estado", l)
	}

	descricao = ""
	l1, err := NewLancamento(id, pessoa.Cpf, data, numero, descricao)
	if err.Error() != "Tamanho de campo Descrição inválido[0 caracter(es)]" {
		t.Error("Erro em função lancamento.NewLancamento, não retornou o erro esperado", l1, err)
	}

	descricao = "Descrição de lançamento muito longa para dar erro em teste unitário de modelo Lancamento ..... .... ..... .... ... ... ... . . .  . ."
	l1, err = NewLancamento(id, pessoa.Cpf, data, numero, descricao)
	if err.Error() != "Tamanho de campo Descrição inválido[137 caracter(es)]" {
		t.Error("Erro em função lancamento.NewLancamento, não retornou o erro esperado", l1, err)
	}

	numero = "123456789123456789123"
	l1, err = NewLancamento(id, pessoa.Cpf, data, numero, descricao)
	if err.Error() != "Tamanho de campo Número inválido[21 caracter(es)]" {
		t.Error("Erro em função lancamento.NewLancamento, não retornou o erro esperado", l1, err)
	}

	cpf := ""
	l1, err = NewLancamento(id, cpf, data, numero, descricao)
	if err.Error() != "Tamanho de campo CPF inválido[0 caracter(es)]" {
		t.Error("Erro em função lancamento.NewLancamento, não retornou o erro esperado", l1, err)
	}

	cpf = "123.456.789.10"
	l1, err = NewLancamento(id, cpf, data, numero, descricao)
	if err.Error() != "Tamanho de campo CPF inválido[14 caracter(es)]" {
		t.Error("Erro em função lancamento.NewLancamento, não retornou o erro esperado", l1, err)
	}
}

func TestGetLancamentoTest(t *testing.T) {
	data := time.Date(2000, 2, 1, 12, 30, 0, 0, new(time.Location))
	id := 9999
	cpf := "12345678910"
	numero := "1234A"
	descricao := "Pgto conta energia"
	dataCriacao := data
	dataModificacao := data
	estado := true

	l := GetLancamentoTest()

	if l.ID != id {
		t.Error("Erro em função conta.GetContaTest, atributo ID", l)
	}

	if l.CpfPessoa != cpf {
		t.Error("Erro em função conta.GetContaTest, atributo CpfPessoa", l)
	}

	if l.Numero != numero {
		t.Error("Erro em função conta.GetContaTest, atributo Numero", l)
	}

	if l.Descricao != descricao {
		t.Error("Erro em função conta.GetContaTest, atributo Descricao", l)
	}

	if l.Data.Unix() != data.Unix() {
		t.Error("Erro em função conta.GetContaTest, atributo Data", l)
	}

	if l.DataCriacao.Unix() != dataCriacao.Unix() {
		t.Error("Erro em função conta.GetContaTest, atributo DataCriacao", l)
	}

	if l.DataModificacao.Unix() != dataModificacao.Unix() {
		t.Error("Erro em função conta.GetContaTest, atributo DataModificacao", l)
	}

	if l.Estado != estado {
		t.Error("Erro em função conta.GetContaTest, atributo Estado", l)
	}
}

func TestString(t *testing.T) {
	l := GetLancamentoTest()

	if l.String() != "9999	12345678910	01/02/2000 12:30:00	1234A	Pgto conta energia	01/02/2000 12:30:00	01/02/2000 12:30:00	ativo" {
		t.Error("Erro em função lancamento.String", l.String())
	}
}

func TestRepr(t *testing.T) {
	l := GetLancamentoTest()

	if l.Repr() != "9999	12345678910	2000-02-01 12:30:00 +0000 UTC	1234A	Pgto conta energia	2000-02-01 12:30:00 +0000 UTC	2000-02-01 12:30:00 +0000 UTC	true" {
		t.Error("Erro em função lancamento.Repr", l.Repr())
	}
}
