package lancamento

import (
	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/model/pessoa"
	"fmt"
	"reflect"
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
	il.Altera("99999999999", time.Now().Local(), "99999", "Nova descrição de Lançamento")
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

func TestAltera(t *testing.T) {
	l := GetLancamentoTest()
	id := l.ID
	data := time.Date(2001, time.January, 2, 23, 59, 58, 0, new(time.Location))

	err := l.Altera("98989898998", data, "5678B", "Pgto de conta de energia")
	if err != nil {
		t.Error(err, l)
	}

	if l.ID != id {
		t.Error("ID modificado após usar método Altera", id, "!=", l.ID)
	}

	err = l.Altera("123.456.789-12", data, "5678B", "Pgto de conta de energia")
	if err.Error() != "Tamanho de campo CPF inválido[14 caracter(es)]" {
		t.Error(err, l)
	}

	err = l.Altera("", data, "5678B", "Pgto de conta de energia")
	if err.Error() != "Tamanho de campo CPF inválido[0 caracter(es)]" {
		t.Error(err, l)
	}

	err = l.Altera("98989898998", data, "1234567890abc1234567", "Pgto de conta de energia")
	if err.Error() != "Tamanho de campo Número inválido[20 caracter(es)]" {
		t.Error(err, l)
	}

	err = l.Altera("98989898998", data, "", "Pgto de conta de energia")
	if err != nil {
		t.Error(err, l)
	}

	err = l.Altera("98989898998", data, "5678B", "Descrição de pagamento de conta com tamanho muito grande para dar erro em teste unitário de modelo Lancamento... ... ... .. . . . . ")
	if err.Error() != "Tamanho de campo Descrição inválido[135 caracter(es)]" {
		t.Error(err, l)
	}

	err = l.Altera("98989898998", data, "5678B", "")
	if err.Error() != "Tamanho de campo Descrição inválido[0 caracter(es)]" {
		t.Error(err, l)
	}
}

func TestAlteraCampos(t *testing.T) {
	data := time.Date(2001, time.January, 2, 23, 59, 58, 0, new(time.Location))
	campos := map[string]interface{}{
		"cpf":       "98989898998",
		"data":      data,
		"numero":    "5678B",
		"descricao": "Pgto de conta de energia"}

	l := GetLancamentoTest()
	id := l.ID
	err := l.AlteraCampos(campos)
	if err != nil {
		t.Error(err, l)
	}

	if l.ID != id {
		t.Error("ID modificado após usar método AlteraCampos", id, "!=", l.ID)
	}

	campos["cpf"] = "123.456.789-12"
	campos["numero"] = "5678B"
	campos["descricao"] = "Pgto de conta de energia"
	err = l.AlteraCampos(campos)
	if err.Error() != "Tamanho de campo CPF inválido[14 caracter(es)]" {
		t.Error(err, l)
	}

	campos["cpf"] = ""
	campos["numero"] = "5678B"
	campos["descricao"] = "Pgto de conta de energia"
	err = l.AlteraCampos(campos)
	if err.Error() != "Tamanho de campo CPF inválido[0 caracter(es)]" {
		t.Error(err, l)
	}

	campos["cpf"] = "98989898998"
	campos["numero"] = "1234567890abc1234567"
	campos["descricao"] = "Pgto de conta de energia"
	err = l.AlteraCampos(campos)
	if err.Error() != "Tamanho de campo Número inválido[20 caracter(es)]" {
		t.Error(err, l)
	}

	campos["cpf"] = "98989898998"
	campos["numero"] = ""
	campos["descricao"] = "Pgto de conta de energia"
	err = l.AlteraCampos(campos)
	if err != nil {
		t.Error(err, l)
	}

	campos["cpf"] = "98989898998"
	campos["numero"] = "5678B"
	campos["descricao"] = "Descrição de pagamento de conta com tamanho muito grande para dar erro em teste unitário de modelo Lancamento... ... ... .. . . . . "
	err = l.AlteraCampos(campos)
	if err.Error() != "Tamanho de campo Descrição inválido[135 caracter(es)]" {
		t.Error(err, l)
	}

	campos["cpf"] = "98989898998"
	campos["numero"] = "5678B"
	campos["descricao"] = ""
	err = l.AlteraCampos(campos)
	if err.Error() != "Tamanho de campo Descrição inválido[0 caracter(es)]" {
		t.Error(err, l)
	}

	campos["cpf"] = "12345678912"
	campos["data"] = "01/01/2010"
	campos["numero"] = "5678B"
	campos["descricao"] = "Pgto de conta de energia"
	err = l.AlteraCampos(campos)
	if err.Error() != "Data inválida 01/01/2010[string]" {
		t.Error(err, l)
	}
}

func TestAtivaInativa(t *testing.T) {
	l := GetLancamentoTest()

	l.Inativa()
	if l.Estado != false {
		t.Error("Erro em função lancamento.Inativa, atributo Estado inválido", l)
	}

	l.Ativa()
	if l.Estado != true {
		t.Error("Erro em função lancamento.Ativa, atributo Estado inválido", l)
	}
}

func TestProcuraLancamento(t *testing.T) {
	cpf := "36925814712"
	id := 789456

	l1 := GetLancamentoTest()
	l1.CpfPessoa = cpf
	l1b := GetLancamentoTest()
	l1b.CpfPessoa = cpf

	l2 := GetLancamentoTest()
	l2.ID = id

	lancamentos := Lancamentos{l1, l2, l1b}

	l3, err := lancamentos.ProcuraLancamentoID(id)
	if !reflect.DeepEqual(l3, l2) {
		t.Error(err, l3)
	}

	l4, err := lancamentos.ProcuraLancamentoID(0)
	if err == nil {
		t.Error("Erro em função lancamento.ProcuraLancamentoID, retonou Lancamento para ID de lancamento inexistente", l4)
	}

	ls5, err := lancamentos.ProcuraLancamentoCPF(cpf)
	if !reflect.DeepEqual(ls5[0], l1) {
		t.Error(err, ls5[0], l1)
	}

	ls6, err := lancamentos.ProcuraLancamentoCPF(cpf)
	if len(ls6) != 2 {
		t.Error("Erro em função lancamento.ProcuraLancamentoCPF, retonou quantidade de Lancamentos diferente do informado", len(ls6), "!=", 2)
	}

	ls7, err := lancamentos.ProcuraLancamentoCPF("")
	if err == nil {
		t.Error("Erro em função lancamento.ProcuraLancamentoCPF, retonou Lancamento para CPF de lancamento inexistente", ls7)
	}
}

func TestLen(t *testing.T) {
	l1 := GetLancamentoTest()
	l2 := GetLancamentoTest()

	lancamentos := Lancamentos{l1, l2}

	if lancamentos.Len() != 2 {
		t.Error("Erro em função lancamento.Len, retorna quantidade de elementos diferente do real", len(lancamentos))
	}
}
