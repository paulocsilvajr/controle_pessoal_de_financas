package detalhe_lancamento

import (
	"controle_pessoal_de_financas/API/v1/model/conta"
	"controle_pessoal_de_financas/API/v1/model/lancamento"
	"fmt"
	"testing"
)

func TestIDetalheLancamento(t *testing.T) {
	dl1 := &DetalheLancamento{}
	dl2 := new(DetalheLancamento)
	detalheLancamentos := DetalheLancamentos{dl1, dl2}

	for _, dl := range detalheLancamentos {
		var idl IDetalheLancamento = dl
		dl3, ok := idl.(*DetalheLancamento)
		if !ok {
			t.Error(dl3)
		}
	}

	defer func() {
		if r := recover(); r != nil {
			t.Error(r)
		}
	}()

	var idl IDetalheLancamento
	idl = GetDetalheLancamentoCTest()
	fmt.Println(idl.String(), idl.Repr())
	idl.VerificaAtributos()
	idl.Altera("Nome conta em detalhe lancamento", 20, 100)
	campos := map[string]string{
		"nomeConta": "Detalhe lancamento 02",
		"credito":   "125",
		"debito":    "25.00"}
	idl.AlteraCampos(campos)
	idl.DebitoToStr()
	idl.CreditoToStr()
}

func TestConverterParaDetalheLancamento(t *testing.T) {
	dl1 := &DetalheLancamento{}
	dl2 := new(DetalheLancamento)
	detalheLancamentos := DetalheLancamentos{dl1, dl2}

	for _, dl := range detalheLancamentos {
		ndl, err := converteParaDetalheLancamento(dl)
		if err != nil {
			t.Error(ndl)
		}
	}
}

func TestNew(t *testing.T) {
	lancamento := lancamento.GetLancamentoTest()
	conta := conta.GetContaTest()
	debito := 200.
	credito := 18.8

	dl := New(lancamento.ID, conta.Nome, debito, credito)

	if dl.IDLancamento != lancamento.ID {
		t.Error("Erro em função detalhe_lancamento.New, atributo ID")
	}

	if dl.NomeConta != conta.Nome {
		t.Error("Erro em função detalhe_lancamento.New, atributo NomeConta")
	}

	if dl.Debito != debito {
		t.Error("Erro em função detalhe_lancamento.New, atributo Debito")
	}

	if dl.Credito != credito {
		t.Error("Erro em função detalhe_lancamento.New, atributo Credito")
	}

	dl1 := NewC(lancamento.ID, conta.Nome, credito)
	if dl1.Debito != 0.0 {
		t.Error("Erro em função detalhe_lancamento.NewC, atributo Debito")
	}

	if dl1.Credito != credito {
		t.Error("Erro em função detalhe_lancamento.NewC, atributo Credito")
	}

	dl2 := NewD(lancamento.ID, conta.Nome, debito)
	if dl2.Debito != debito {
		t.Error("Erro em função detalhe_lancamento.NewD, atributo Debito")
	}

	if dl2.Credito != 0.0 {
		t.Error("Erro em função detalhe_lancamento.NewD, atributo Credito")
	}
}
