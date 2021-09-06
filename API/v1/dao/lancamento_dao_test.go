package dao

import (
	"testing"
	"time"

	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/model/lancamento"
	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/model/pessoa"
)

var (
	testLancamento01  *lancamento.Lancamento
	testPessoaAdmin02 *pessoa.Pessoa
)

func TestAdicionaLancamento02(t *testing.T) {
	var err error

	p := getPessoaAdmin2()
	testPessoaAdmin02, err = AdicionaPessoa02(db2, p)
	if err != nil {
		t.Error(err)
	}

	l := getLancamento2(testPessoaAdmin02)
	testLancamento01, err = AdicionaLancamento02(db2, l)
	if err != nil {
		t.Error(err)
	}

	cpfEsperado := l.CpfPessoa
	cpfObtido := testLancamento01.CpfPessoa
	id := testLancamento01.ID
	if cpfEsperado != cpfObtido {
		t.Errorf("CPF de pessoa em lancamento %d diferente do esperado. Esperado: '%s', obtido: '%s'", id, cpfEsperado, cpfObtido)
	}
}

func TestCarregaLancamentos02(t *testing.T) {
	lancamentos, err := CarregaLancamentos02(db2)
	if err != nil {
		t.Error(err)
	}

	quantLancamentos := len(lancamentos)
	quantEsperada := 1
	if quantLancamentos != quantEsperada {
		t.Errorf("consulta de Lancamentos retornou uma quantidade de registros diferente do esperado. Esperado: %d, obtido: %d", quantEsperada, quantLancamentos)
	}

	cpf := testLancamento01.CpfPessoa
	id := testLancamento01.ID
	for _, l := range lancamentos {
		if l.CpfPessoa != cpf {
			t.Errorf("registro de lancamento(CpfPessoa) encontrado diferente do esperado. Esperado: '%s', obtido: '%s'", cpf, l.CpfPessoa)
		}

		if l.ID != id {
			t.Errorf("registro de lancamento(ID) encontrado diferente do esperado. Esperado: %d, obtido: %d", id, l.ID)
		}
	}
}

func TestCarregaLancamentosPorCPF02(t *testing.T) {
	cpf := testLancamento01.CpfPessoa
	lancamentos, err := CarregaLancamentosPorCPF02(db2, cpf)
	if err != nil {
		t.Error(err)
	}

	quantLancamentos := len(lancamentos)
	quantEsperada := 1
	if quantLancamentos != quantEsperada {
		t.Errorf("consulta de Lancamentos retornou uma quantidade de registros diferente do esperado. Esperado: %d, obtido: %d", quantEsperada, quantLancamentos)
	}

	id := testLancamento01.ID
	for _, l := range lancamentos {
		if l.CpfPessoa != cpf {
			t.Errorf("registro de lancamento(CpfPessoa) encontrado diferente do esperado. Esperado: '%s', obtido: '%s'", cpf, l.CpfPessoa)
		}

		if l.ID != id {
			t.Errorf("registro de lancamento(ID) encontrado diferente do esperado. Esperado: %d, obtido: %d", id, l.ID)
		}
	}
}

func TestAlteraLancamento02(t *testing.T) {
	id := testLancamento01.ID
	novoNumero := "Ln1234"
	novaData := time.Now()
	novaDescricao := "NOVA descrição Lanc Ln1234"

	testLancamento01.Numero = novoNumero
	testLancamento01.Data = novaData
	testLancamento01.Descricao = novaDescricao

	transacao := db2.Begin()
	l, err := AlteraLancamento02(db2, transacao, id, testLancamento01)
	transacao.Commit()
	if err != nil {
		t.Error(err)
	}

	numero := l.Numero
	if numero != novoNumero {
		t.Errorf("alteração de lancamento com ID %d retornou um 'Número' diferente do esperado. Esperado: '%s', obtido: '%s'", id, novoNumero, numero)
	}

	data := l.Data
	if data.Unix() != novaData.Unix() {
		t.Errorf("alteração de lancamento com ID %d retornou uma 'Data' diferente do esperado. Esperado: '%s', obtido: '%s'", id, novaData, data)
	}

	descricao := l.Descricao
	if descricao != novaDescricao {
		t.Errorf("alteração de lancamento com ID %d retornou uma 'Descrição' diferente do esperado. Esperado: '%s', obtido: '%s'", id, novaDescricao, novoNumero)
	}
}

func TestInativaLancamento02(t *testing.T) {
	id := testLancamento01.ID
	l, err := InativaLancamento02(db2, id)
	if err != nil {
		t.Error(err)
	}

	if l != nil {
		idObtido := l.ID
		if idObtido != id {
			t.Errorf("inativação de lançamento retornou um lançamento com ID diferente do esperado. Esperado: %d, obtido: %d", id, idObtido)
		}

		estadoObtido := l.Estado
		estadoEsperado := false
		if estadoObtido != estadoEsperado {
			t.Errorf("inativação de lançamento retornou um lançamento com estado diferente do esperado. Esperado: %t, obtido: %t", estadoEsperado, estadoObtido)
		}
	} else {
		t.Errorf("func InativaLancamento02(db2, %d) retornou um lançamento nulo(nil)", id)
	}
}

func TestCarregaLancamentosInativo02(t *testing.T) {
	lancamentos, err := CarregaLancamentosInativo02(db2)
	if err != nil {
		t.Error(err)
	}

	quantObtida := len(lancamentos)
	quantEsperada := 1
	if quantEsperada != quantObtida {
		t.Errorf("consulta de lançamentos inativos retornou uma quantidade de registros incorreta. Esperado: %d, obtido: %d", quantEsperada, quantObtida)
	}
}

func TestCarregaLancamentosInativoPorCPF02(t *testing.T) {
	cpf := testLancamento01.CpfPessoa
	lancamentos, err := CarregaLancamentosInativoPorCPF02(db2, cpf)
	if err != nil {
		t.Error(err)
	}

	quantEsperada := 1
	quantObtida := lancamentos.Len()
	if quantEsperada != quantObtida {
		t.Errorf("consulta de lançamentos inativos por CPF '%s' retornou uma quantidade de registros incorreta. Esperado: %d, obtido: %d", cpf, quantEsperada, quantObtida)
	}

	for _, lanc := range lancamentos {
		cpfObtido := lanc.CpfPessoa
		if lanc.CpfPessoa != cpf {
			t.Errorf("consulta de lançamentos inativos por CPF '%s' retornou um cpf diferente do esperado. Esperado '%[1]s', obtido: '%s'", cpf, cpfObtido)
		}
	}
}

func TestAtivaLancamento02(t *testing.T) {
	id := testLancamento01.ID
	l, err := AtivaLancamento02(db2, id)
	if err != nil {
		t.Error(err)
	}

	if l != nil {
		idObtido := l.ID
		if idObtido != id {
			t.Errorf("ativação de lançamento retornou um lançamento com ID diferente do esperado. Esperado: %d, obtido: %d", id, idObtido)
		}

		estadoObtido := l.Estado
		estadoEsperado := true
		if estadoObtido != estadoEsperado {
			t.Errorf("ativação de lançamento retornou um lançamento com estado diferente do esperado. Esperado: %t, obtido: %t", estadoEsperado, estadoObtido)
		}
	} else {
		t.Errorf("func AtivaLancamento02(db2, %d) retornou um lançamento nulo(nil)", id)
	}
}

func TestCarregaLancamentosAtivo02(t *testing.T) {
	lancamentos, err := CarregaLancamentosAtivo02(db2)
	if err != nil {
		t.Error(err)
	}

	quantObtida := len(lancamentos)
	quantEsperada := 1
	if quantEsperada != quantObtida {
		t.Errorf("consulta de lançamentos ativos retornou uma quantidade de registros incorreta. Esperado: %d, obtido: %d", quantEsperada, quantObtida)
	}
}

func TestCarregaLancamentosAtivoPorCPF02(t *testing.T) {
	cpf := testLancamento01.CpfPessoa
	lancamentos, err := CarregaLancamentosAtivoPorCPF02(db2, cpf)
	if err != nil {
		t.Error(err)
	}

	quantEsperada := 1
	quantObtida := lancamentos.Len()
	if quantEsperada != quantObtida {
		t.Errorf("consulta de lançamentos ativos por CPF '%s' retornou uma quantidade de registros incorreta. Esperado: %d, obtido: %d", cpf, quantEsperada, quantObtida)
	}

	for _, lanc := range lancamentos {
		cpfObtido := lanc.CpfPessoa
		if lanc.CpfPessoa != cpf {
			t.Errorf("consulta de lançamentos ativos por CPF '%s' retornou um cpf diferente do esperado. Esperado '%[1]s', obtido: '%s'", cpf, cpfObtido)
		}
	}
}

func TestRemoveLancamento02(t *testing.T) {
	err := RemoveLancamento02(db2, testLancamento01.ID)
	if err != nil {
		t.Error(err)
	}

	err = RemovePessoa02(db2, testPessoaAdmin02.Cpf)
	if err != nil {
		t.Error(err)
	}
}
