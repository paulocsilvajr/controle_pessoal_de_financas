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
