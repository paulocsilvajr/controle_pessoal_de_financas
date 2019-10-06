package tipo_conta

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestITipoConta(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error(r)
		}
	}()

	var t1 ITipoConta
	t1 = GetTipoContaTest()

	// Métodos obrigatórios da Interface ITipoConta
	_ = t1.String()
	t1.Repr()
	t1.VerificaAtributos()
	t1.Altera("Nome do tipo de conta", "Débito", "Crédito")

	campos := map[string]string{
		"nome":             "tipo conta 02",
		"descricaoDebito":  "débito 2",
		"DescricaoCredito": "crédito 2"}
	t1.AlteraCampos(campos)
	t1.Ativa()
	t1.Inativa()
}

func TestNew(t *testing.T) {
	nome := "Bancos em geral"
	debito := "saídas"
	credito := "entradas"
	data := time.Now().Local()

	t1 := New(nome, debito, credito)

	if t1.Nome != nome {
		t.Error("Erro em função tipo_conta.New, atributo Nome", t1)
	}

	if t1.DescricaoDebito != debito {
		t.Error("Erro em função tipo_conta.New, atributo DescricaoDebito", t1)
	}

	if t1.DescricaoCredito != credito {
		t.Error("Erro em função tipo_conta.New, atributo DescricaoCredito", t1)
	}

	if t1.DataCriacao.Unix() < data.Unix() {
		t.Error("Erro em função tipo_conta.New, atributo DataCriacao", t1)
	}

	if t1.DataModificacao.Unix() < data.Unix() {
		t.Error("Erro em função tipo_conta.New, atributo DataModificacao", t1)
	}

	if t1.Estado != true {
		t.Error("Erro em função tipo_conta.New, atributo Estado", t1)
	}
}

func TestNewTipoConta(t *testing.T) {
	nome := "Bancos em geral"
	debito := "saídas"
	credito := "entradas"
	data := time.Now().Local()

	t1, err := NewTipoConta(nome, debito, credito)

	if t1.Nome != nome {
		t.Error("Erro em função tipo_conta.NewTipoConta, atributo Nome", t1)
	}

	if t1.DescricaoDebito != debito {
		t.Error("Erro em função tipo_conta.NewTipoConta, atributo DescricaoDebito", t1)
	}

	if t1.DescricaoCredito != credito {
		t.Error("Erro em função tipo_conta.NewTipoConta, atributo DescricaoCredito", t1)
	}

	if t1.DataCriacao.Unix() < data.Unix() {
		t.Error("Erro em função tipo_conta.NewTipoConta, atributo DataCriacao", t1)
	}

	if t1.DataModificacao.Unix() < data.Unix() {
		t.Error("Erro em função tipo_conta.NewTipoConta, atributo DataModificacao", t1)
	}

	if t1.Estado != true {
		t.Error("Erro em função tipo_conta.NewTipoConta, atributo Estado", t1)
	}

	if err != nil {
		t.Error(err, t1)
	}

	credito = ""
	t2, err := NewTipoConta(nome, debito, credito)
	if err.Error() != "Descrição do crédito com tamanho inválido[0]" {
		t.Error("Erro em função tipo_conta.NewTipoConta, não retornou o erro esperado", t2, err)
	}

	debito = ""
	t2, err = NewTipoConta(nome, debito, credito)
	if err.Error() != "Descrição do débito com tamanho inválido[0]" {
		t.Error("Erro em função tipo_conta.NewTipoConta, não retornou o erro esperado", t2, err)
	}

	nome = ""
	t2, err = NewTipoConta(nome, debito, credito)
	if err.Error() != "Nome com tamanho inválido[0]" {
		t.Error("Erro em função tipo_conta.NewTipoConta, não retornou o erro esperado", t2, err)
	}

	descricao := "Descrição muito longa para crédito ou débito"
	t3, err := NewTipoConta("Tipo de conta 01", "débito", descricao)
	if err.Error() != fmt.Sprintf("Descrição do crédito com tamanho inválido[%d]", len(descricao)) {
		t.Error("Erro em função tipo_conta.NewTipoConta, não retornou o erro esperado", t3, err)
	}

	t3, err = NewTipoConta("Tipo de conta 01", descricao, descricao)
	if err.Error() != fmt.Sprintf("Descrição do débito com tamanho inválido[%d]", len(descricao)) {
		t.Error("Erro em função tipo_conta.NewTipoConta, não retornou o erro esperado", t3, err)
	}

	nome = "Nome muito longo que ultrapassa o máximo do tamanho da nome definido em tipo_conta"
	t3, err = NewTipoConta(nome, debito, credito)
	if err.Error() != fmt.Sprintf("Nome com tamanho inválido[%d]", len(nome)) {
		t.Error("Erro em função tipo_conta.NewTipoConta, não retornou o erro esperado", t3, err)
	}
}

func TestGetTipoContaTeste(t *testing.T) {
	nome := "banco teste 01"
	debito := "saque"
	credito := "depósito"
	criacao := time.Date(2000, 2, 1, 12, 30, 0, 0, new(time.Location))
	modificacao := criacao
	estado := true

	t1 := GetTipoContaTest()

	if t1.Nome != nome {
		t.Error("Erro em função tipo_conta.GetTipoContaTest, atributo Nome", t1)
	}

	if t1.DescricaoDebito != debito {
		t.Error("Erro em função tipo_conta.GetTipoContaTest, atributo DescricaoDebito", t1)
	}

	if t1.DescricaoCredito != credito {
		t.Error("Erro em função tipo_conta.GetTipoContaTest, atributo DescricaoCredito", t1)
	}

	if t1.DataCriacao.Unix() != criacao.Unix() {
		t.Error("Erro em função tipo_conta.GetTipoContaTest, atributo DataCriacao", t1)
	}

	if t1.DataModificacao.Unix() != modificacao.Unix() {
		t.Error("Erro em função tipo_conta.GetTipoContaTest, atributo DataModificacao", t1)
	}

	if t1.Estado != estado {
		t.Error("Erro em função tipo_conta.GetTipoContaTest, atributo Estado", t1)
	}
}

func TestString(t *testing.T) {
	t1 := GetTipoContaTest()

	if t1.String() != "banco teste 01	saque	depósito	01/02/2000 12:30:00	01/02/2000 12:30:00	ativo" {
		t.Error("Erro em função tipo_conta.String", t1.String())
	}
}

func TestRepr(t *testing.T) {
	t1 := GetTipoContaTest()

	if t1.Repr() != "banco teste 01	saque	depósito	2000-02-01 12:30:00 +0000 UTC	2000-02-01 12:30:00 +0000 UTC	true" {
		t.Error("Erro em função tipo_conta.String", t1.Repr())
	}
}

func TestAltera(t *testing.T) {
	t1 := GetTipoContaTest()

	err := t1.Altera("Tipo Conta 05", "-", "+")
	if err != nil {
		t.Error(err, t1)
	}

	err = t1.Altera("Nome muito longo que ultrapassa o máximo do tamanho da nome definido em tipo_conta", "-", "+")
	if err.Error() != "Nome com tamanho inválido[83]" {
		t.Error(err, t1)
	}

	err = t1.Altera("Tipo Conta 05", "Descrição muito longa para crédito ou débito", "+")
	if err.Error() != "Descrição do débito com tamanho inválido[48]" {
		t.Error(err, t1)
	}

	err = t1.Altera("Tipo Conta 05", "-", "Descrição muito longa para crédito ou débito")
	if err.Error() != "Descrição do crédito com tamanho inválido[48]" {
		t.Error(err, t1)
	}
}

func TestAlteraCampos(t *testing.T) {
	campos := map[string]string{
		"nome":             "Tipo Conta 05",
		"descricaoDebito":  "-",
		"DescricaoCredito": "+"}

	t1 := GetTipoContaTest()
	err := t1.AlteraCampos(campos)
	if err != nil {
		t.Error(err, t1)
	}

	campos["descricaoDebito"] = "Saída"
	campos["DescricaoCredito"] = "Entrada"
	err = t1.AlteraCampos(campos)
	if err != nil {
		t.Error(err, t1)
	}

	campos["nome"] = "Nome muito longo que ultrapassa o máximo do tamanho da nome definido em tipo_conta"
	campos["descricaoDebito"] = "Saída"
	campos["DescricaoCredito"] = "Entrada"
	err = t1.AlteraCampos(campos)
	if err.Error() != "Nome com tamanho inválido[83]" {
		t.Error(err, t1)
	}

	campos["nome"] = "Tipo Conta 05"
	campos["descricaoDebito"] = "Descrição muito longa para crédito ou débito"
	campos["DescricaoCredito"] = "Entrada"
	err = t1.AlteraCampos(campos)
	if err.Error() != "Descrição do débito com tamanho inválido[48]" {
		t.Error(err, t1)
	}

	campos["nome"] = "Tipo Conta 05"
	campos["descricaoDebito"] = "Saída"
	campos["descricaoCredito"] = "Descrição muito longa para crédito ou débito"
	err = t1.AlteraCampos(campos)
	if err.Error() != "Descrição do crédito com tamanho inválido[48]" {
		t.Error(err, t1)
	}
}

func TestAtiva(t *testing.T) {
	t1 := GetTipoContaTest()

	t1.Inativa()
	if t1.Estado != false {
		t.Error("Erro em função tipo_conta.Inativa, atributo Estado inválido", t1)
	}

	t1.Ativa()
	if t1.Estado != true {
		t.Error("Erro em função tipo_conta.Ativa, atributo estado inválido", t1)
	}
}

func TestProcuraTipoConta(t *testing.T) {
	t1 := GetTipoContaTest()
	t2 := GetTipoContaTest()

	tiposConta := TiposConta{t1, t2}

	t3, err := tiposConta.ProcuraTipoConta(t1.Nome)
	if !reflect.DeepEqual(t3, t1) {
		t.Error(err, t3)
	}

	t4, err := tiposConta.ProcuraTipoConta("NENHUM")
	if err == nil {
		t.Error("Erro em função tipo_conta.ProcuraTipoConta, retornou TipoConta para nome de tipo de conta inexistente", t4)
	}
}

func TestLen(t *testing.T) {
	t1 := GetTipoContaTest()
	t2 := GetTipoContaTest()

	tiposConta := TiposConta{t1, t2}

	if tiposConta.Len() != 2 {
		t.Error("Erro em função tipo_conta.Len, retorno quantidade de elemento diferente do real", len(tiposConta))
	}
}
