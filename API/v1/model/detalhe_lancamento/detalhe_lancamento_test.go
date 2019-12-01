package detalhe_lancamento

import (
	"controle_pessoal_de_financas/API/v1/model/conta"
	"controle_pessoal_de_financas/API/v1/model/lancamento"
	"fmt"
	"reflect"
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

func TestNewDetalheLancamento(t *testing.T) {
	lancamento := lancamento.GetLancamentoTest()
	conta := conta.GetContaTest()
	debito := 200.
	credito := 18.8

	dl, err := NewDetalheLancamento(lancamento.ID, conta.Nome, debito, credito)
	if err != nil {
		t.Error(err, dl)
	}

	if dl.IDLancamento != lancamento.ID {
		t.Error("Erro em função detalhe_lancamento.NewDetalheLancamento, atributo ID")
	}

	if dl.NomeConta != conta.Nome {
		t.Error("Erro em função detalhe_lancamento.NewDetalheLancamento, atributo NomeConta")
	}

	if dl.Debito != debito {
		t.Error("Erro em função detalhe_lancamento.NewDetalheLancamento, atributo Debito")
	}

	if dl.Credito != credito {
		t.Error("Erro em função detalhe_lancamento.NewDetalheLancamento, atributo Credito")
	}

	dl1, err := NewDetalheLancamento(lancamento.ID, conta.Nome, 0, credito)
	if err != nil {
		t.Error(err, dl)
	}

	if dl1.Debito != 0.0 {
		t.Error("Erro em função detalhe_lancamento.NewDetalheLancamento, atributo Debito")
	}

	if dl1.Credito != credito {
		t.Error("Erro em função detalhe_lancamento.NewDetalheLancamento, atributo Credito")
	}

	dl2, err := NewDetalheLancamento(lancamento.ID, conta.Nome, debito, 0)
	if err != nil {
		t.Error(err, dl)
	}

	if dl2.Debito != debito {
		t.Error("Erro em função detalhe_lancamento.NewDetalheLancamento, atributo Debito")
	}

	if dl2.Credito != 0.0 {
		t.Error("Erro em função detalhe_lancamento.NewDetalheLancamento, atributo Credito")
	}

	credito = -20
	dl, err = NewDetalheLancamento(lancamento.ID, conta.Nome, debito, credito)
	if err.Error() != "O campo Credito deve ser >= 0" {
		t.Error("Erro em função detalhe_lancamento.NewLancamento, não retornou o erro esperado", dl, err)
	}

	debito = -30
	dl, err = NewDetalheLancamento(lancamento.ID, conta.Nome, debito, credito)
	if err.Error() != "O campo Debito deve ser >= 0" {
		t.Error("Erro em função detalhe_lancamento.NewLancamento, não retornou o erro esperado", dl, err)
	}

	nomeConta := ""
	dl, err = NewDetalheLancamento(lancamento.ID, nomeConta, debito, credito)
	if err.Error() != "Tamanho de campo NomeConta inválido[0 caracter(es)]" {
		t.Error("Erro em função detalhe_lancamento.NewLancamento, não retornou o erro esperado", dl, err)
	}

	nomeConta = "Descrição de nome de conta muito grande para gerar erro em função NewDetalheLancamento no parâmetro nomeConta do tipo string ... .. ... ..... .... .. .. . .. .. . . . . ."
	dl, err = NewDetalheLancamento(lancamento.ID, nomeConta, debito, credito)
	if err.Error() != "Tamanho de campo NomeConta inválido[175 caracter(es)]" {
		t.Error("Erro em função detalhe_lancamento.NewLancamento, não retornou o erro esperado", dl, err)
	}
}

func TestGetDetalheLancamentoTest(t *testing.T) {
	id := 9999
	nomeConta := "Detalhe de conta teste A"
	debito := 100.
	credito := 200.

	dl := GetDetalheLancamentoDTest()

	if dl.IDLancamento != id {
		t.Error("Erro em função detalhe_lancamento.GetDetalheLancamentoDTest, atributo IDLancamento", dl)
	}

	if dl.NomeConta != nomeConta {
		t.Error("Erro em função detalhe_lancamento.GetDetalheLancamentoDTest, atributo NomeConta", dl)
	}

	if dl.Debito != debito {
		t.Error("Erro em função detalhe_lancamento.GetDetalheLancamentoDTest, atributo Debito", dl)
	}

	if dl.Credito != 0.0 {
		t.Error("Erro em função detalhe_lancamento.GetDetalheLancamentoDTest, atributo Credito", dl)
	}

	dl = GetDetalheLancamentoCTest()

	if dl.Debito != 0.0 {
		t.Error("Erro em função detalhe_lancamento.GetDetalheLancamentoDTest, atributo Debito", dl)
	}

	if dl.Credito != credito {
		t.Error("Erro em função detalhe_lancamento.GetDetalheLancamentoDTest, atributo Credito", dl)
	}
}

func TestString(t *testing.T) {
	dl := GetDetalheLancamentoCTest()

	if dl.String() != "9998	Detalhe de conta teste B	0.000	200.000" {
		t.Error("Erro em função detalhe_lancamento.String", dl.String())
	}

	dl = GetDetalheLancamentoDTest()

	if dl.String() != "9999	Detalhe de conta teste A	100.000	0.000" {
		t.Error("Erro em função detalhe_lancamento.String", dl.String())
	}
}

func TestRepr(t *testing.T) {
	dl := GetDetalheLancamentoCTest()

	if dl.Repr() != "9998	Detalhe de conta teste B	0.000000	200.000000" {
		t.Error("Erro em função detalhe_lancamento.Repr", dl.Repr())
	}

	dl = GetDetalheLancamentoDTest()

	if dl.Repr() != "9999	Detalhe de conta teste A	100.000000	0.000000" {
		t.Error("Erro em função detalhe_lancamento.Repr", dl.Repr())
	}
}

func TestAltera(t *testing.T) {
	dl := GetDetalheLancamentoCTest()
	id := dl.IDLancamento

	err := dl.Altera("Teste de nome de conta", 100, 10)
	if err != nil {
		t.Error(err, dl)
	}

	if dl.IDLancamento != id {
		t.Error("IDLancamento modificado após usar método Altera", id, "!=", dl.IDLancamento)
	}

	err = dl.Altera("", 100, 10)
	if err.Error() != "Tamanho de campo NomeConta inválido[0 caracter(es)]" {
		t.Error(err, dl)
	}

	err = dl.Altera("Descrição de nome de conta muito grande para gerar erro em função NewDetalheLancamento no parâmetro nomeConta do tipo string ... .. ... ..... .... .. .. . .. .. . . . . .", 100, 10)
	if err.Error() != "Tamanho de campo NomeConta inválido[175 caracter(es)]" {
		t.Error(err, dl)
	}

	err = dl.Altera("Test de nome de conta", -20, 100)
	if err.Error() != "O campo Debito deve ser >= 0" {
		t.Error(err, dl)
	}

	err = dl.Altera("Test de nome de conta", 100, -20)
	if err.Error() != "O campo Credito deve ser >= 0" {
		t.Error(err, dl)
	}
}

func TestAlteraCampos(t *testing.T) {
	campos := map[string]string{
		"nomeConta": "Teste de nome de conta",
		"debito":    "100.00",
		"credito":   "10.00"}

	dl := GetDetalheLancamentoCTest()
	id := dl.IDLancamento
	err := dl.AlteraCampos(campos)
	if err != nil {
		t.Error(err, dl)
	}

	if dl.IDLancamento != id {
		t.Error("IDLancamento modificado após usar método AlteraCampos", id, "!=", dl.IDLancamento)
	}

	campos["nomeConta"] = ""
	campos["debito"] = "100.00"
	campos["credito"] = "10.00"
	err = dl.AlteraCampos(campos)
	if err.Error() != "Tamanho de campo NomeConta inválido[0 caracter(es)]" {
		t.Error(err, dl)
	}

	campos["nomeConta"] = "Descrição de nome de conta muito grande para gerar erro em função NewDetalheLancamento no parâmetro nomeConta do tipo string ... .. ... ..... .... .. .. . .. .. . . . . ."
	campos["debito"] = "100.00"
	campos["credito"] = "10.00"
	err = dl.AlteraCampos(campos)
	if err.Error() != "Tamanho de campo NomeConta inválido[175 caracter(es)]" {
		t.Error(err, dl)
	}

	campos["nomeConta"] = "Teste de nome de conta"
	campos["debito"] = "100,00"
	campos["credito"] = "10.00"
	err = dl.AlteraCampos(campos)
	if err.Error() != "Erro ao converter string para float64" {
		t.Error(err, dl)
	}

	campos["nomeConta"] = "Teste de nome de conta"
	campos["debito"] = "100.00"
	campos["credito"] = "10,00"
	err = dl.AlteraCampos(campos)
	if err.Error() != "Erro ao converter string para float64" {
		t.Error(err, dl)
	}

	campos["nomeConta"] = "Teste de nome de conta"
	campos["debito"] = "100"
	campos["credito"] = "10.00"
	err = dl.AlteraCampos(campos)
	if err != nil {
		t.Error(err, dl)
	}

	campos["nomeConta"] = "Teste de nome de conta"
	campos["debito"] = "100.00"
	campos["credito"] = "10"
	err = dl.AlteraCampos(campos)
	if err != nil {
		t.Error(err, dl)
	}

	campos["nomeConta"] = "Teste de nome de conta"
	campos["debito"] = "100.00"
	campos["credito"] = "-10.00"
	err = dl.AlteraCampos(campos)
	if err.Error() != "O campo credito deve ser >= 0" {
		t.Error(err, dl)
	}

	campos["nomeConta"] = "Teste de nome de conta"
	campos["debito"] = "-100.00"
	campos["credito"] = "10.00"
	err = dl.AlteraCampos(campos)
	if err.Error() != "O campo debito deve ser >= 0" {
		t.Error(err, dl)
	}
}

func TestCreditoToStr(t *testing.T) {
	dl := GetDetalheLancamentoCTest()
	if dl.CreditoToStr() != "200.000" {
		t.Error("Função CreditoToStr não retorna a string esperada", dl.CreditoToStr())
	}
}

func TestDebitoToStr(t *testing.T) {
	dl := GetDetalheLancamentoDTest()
	if dl.DebitoToStr() != "100.000" {
		t.Error("Função DebitoToStr não retorna a string esperada", dl.DebitoToStr())
	}
}

func TestProcuraDetalheLancamento(t *testing.T) {
	dl1 := GetDetalheLancamentoCTest()
	idLancamento1 := dl1.IDLancamento
	nomeConta1 := dl1.NomeConta

	dl2 := GetDetalheLancamentoDTest()
	idLancamento2 := dl2.IDLancamento
	nomeConta2 := dl2.NomeConta

	dl2b := GetDetalheLancamentoDTest()
	idLancamento2b := dl2b.IDLancamento
	nomeConta2b := dl2b.NomeConta

	detalheLancamentos := DetalheLancamentos{dl1, dl2}

	dl3, err := detalheLancamentos.ProcuraDetalheLancamento(idLancamento1, nomeConta1)
	if !reflect.DeepEqual(dl3, dl1) {
		t.Error(err, dl3)
	}

	dl3, err = detalheLancamentos.ProcuraDetalheLancamento(idLancamento2, nomeConta2)
	if !reflect.DeepEqual(dl3, dl2) {
		t.Error(err, dl3)
	}

	dl3, err = detalheLancamentos.ProcuraDetalheLancamento(0, "")
	if err.Error() != "DetalheLancamento [0:] informado não existe na listagem" {
		t.Error("Erro em função detalhe_lancamento.ProcuraDetalheLancamento, retornou DetalheLancamento para IDLancamento e NomeConta inexistente", dl3, err)
	}

	dl3, err = detalheLancamentos.ProcuraDetalheLancamento(idLancamento1, "")
	if err.Error() != "DetalheLancamento [9998:] informado não existe na listagem" {
		t.Error("Erro em função detalhe_lancamento.ProcuraDetalheLancamento, retornou DetalheLancamento para IDLancamento e NomeConta inexistente", dl3, err)
	}

	dl3, err = detalheLancamentos.ProcuraDetalheLancamento(0, nomeConta1)
	if err.Error() != "DetalheLancamento [0:Detalhe de conta teste B] informado não existe na listagem" {
		t.Error("Erro em função detalhe_lancamento.ProcuraDetalheLancamento, retornou DetalheLancamento para IDLancamento e NomeConta inexistente", dl3, err)
	}

	detalheLancamentos = append(detalheLancamentos, dl2b)
	dls1, err := detalheLancamentos.ProcuraDetalheLancamentosPorNomeConta(nomeConta1)
	if dls1.Len() != 1 {
		t.Error("Erro em função detalhe_lancamento.ProcuraDetalheLancamentosPorNomeConta, retornou mais resultado do que o esperado", len(dls1))
	}

	if !reflect.DeepEqual(dls1[0], dl1) {
		t.Error(err, dls1[0])
	}

	dls2, err := detalheLancamentos.ProcuraDetalheLancamentosPorNomeConta(nomeConta2b)
	if dls2.Len() != 2 {
		t.Error("Erro em função detalhe_lancamento.ProcuraDetalheLancamentosPorNomeConta, retornou mais resultado do que o esperado", len(dls2))
	}

	dls, err := detalheLancamentos.ProcuraDetalheLancamentosPorNomeConta("")
	if err.Error() != "Não foi encontrado nenhum DetalheLancamento com o NomeConta[] informado" {
		t.Error(err, dls)
	}

	dls3, err := detalheLancamentos.ProcuraDetalheLancamentosPorIDLancamento(idLancamento1)
	if dls3.Len() != 1 {
		t.Error("Erro em função detalhe_lancamento.ProcuraDetalheLancamentosPorIDLancamento, retornou mais resultado do que o esperado", len(dls3))
	}

	if !reflect.DeepEqual(dls3[0], dl1) {
		t.Error(err, dls3[0])
	}

	dls4, err := detalheLancamentos.ProcuraDetalheLancamentosPorIDLancamento(idLancamento2b)
	if dls4.Len() != 2 {
		t.Error("Erro em função detalhe_lancamento.ProcuraDetalheLancamentosPorIDLancamento, retornou mais resultado do que o esperado", len(dls4))
	}

	dls, err = detalheLancamentos.ProcuraDetalheLancamentosPorIDLancamento(0)
	if err.Error() != "Não foi encontrado nenhum DetalheLancamento com o IDLancamento[0] informado" {
		t.Error(err, dls)
	}
}

func TestLen(t *testing.T) {
	dl1 := GetDetalheLancamentoCTest()
	dl2 := GetDetalheLancamentoDTest()

	detalheLancamentos := DetalheLancamentos{dl1, dl2}

	if detalheLancamentos.Len() != 2 {
		t.Error("erro em função detalhelancamento.Len, retorna quantidade de elementos diferente do real", len(detalheLancamentos))
	}
}
