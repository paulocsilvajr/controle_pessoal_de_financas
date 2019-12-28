package dao

import (
	"controle_pessoal_de_financas/API/v1/model/lancamento"
	"testing"
	"time"
)

var (
	numLanc01, numLanc02, numLanc03 int
)

func TestAdicionaLancamento(t *testing.T) {
	TestAdicionaPessoa(t)

	l1 := lancamento.GetLancamentoTest()
	l1.CpfPessoa = cpf

	l2 := lancamento.GetLancamentoTest()
	l2.Numero = "5678A"
	l2.Descricao = "Lançamento teste 02 - parcela 1"
	l2.CpfPessoa = cpf

	l3 := lancamento.GetLancamentoTest()
	l3.Numero = "5678B"
	l3.Descricao = "Lançamento teste 02 - parcela 2"
	l3.CpfPessoa = cpf

	l5, err := AdicionaLancamento(db, l1)
	numLanc01 = l5.ID
	if err != nil {
		t.Error(err, l5)
	}

	l6, err := AdicionaLancamento(db, l2)
	numLanc02 = l6.ID
	if err != nil {
		t.Error(err, l6)
	}

	l7, err := AdicionaLancamento(db, l3)
	numLanc03 = l7.ID
	if err != nil {
		t.Error(err, l7)
	}

	l4 := lancamento.New(0, "", time.Date(2001, 1, 12, 12, 31, 0, 0, new(time.Location)), "", "")

	l8, err := AdicionaLancamento(db, l4)
	if err.Error() != "Tamanho de campo CPF inválido[0 caracter(es)]" {
		t.Error(err, l8)
	}

	l4.CpfPessoa = cpf
	l8, err = AdicionaLancamento(db, l4)
	if err.Error() != "Tamanho de campo Descrição inválido[0 caracter(es)]" {
		t.Error(err, l8)
	}

	l4.Descricao = "Lançamento teste 03"
	l8, err = AdicionaLancamento(db, l4)
	if err != nil {
		t.Error(err, l8)
	}
}

func TestInativaLancamentoECarregaLancamentosInativos(t *testing.T) {
	l1, err := InativaLancamento(db, numLanc01)
	if err != nil {
		t.Error(err, l1)
	}

	l2, err := InativaLancamento(db, numLanc02)
	if err != nil {
		t.Error(err, l2)
	}

	l3, err := InativaLancamento(db, numLanc03)
	if err != nil {
		t.Error(err, l3)
	}

	l4, err := InativaLancamento(db, 0)
	if err.Error() != "Não foi encontrado um registro com o ID 0" {
		t.Error(err, l4)
	}

	if l1.Estado != false {
		t.Error("Estado do lancamento inválido, deveria ser false", l1)
	}

	if l2.Estado != false {
		t.Error("Estado do lancamento inválido, deveria ser false", l2)
	}

	if l3.Estado != false {
		t.Error("Estado do lancamento inválido, deveria ser false", l3)
	}

	lancamentos, err := CarregaLancamentosInativo(db)
	if err != nil {
		t.Error(err, lancamentos)
	}

	if len(lancamentos) == 0 {
		t.Error(lancamentos)
	}

	if len(lancamentos) < 3 {
		t.Error(lancamentos)
	}
}

func TestAtivaLancamentoECarregaLancamentosAtivos(t *testing.T) {
	l1, err := AtivaLancamento(db, numLanc01)
	if err != nil {
		t.Error(err, l1)
	}

	l2, err := AtivaLancamento(db, numLanc02)
	if err != nil {
		t.Error(err, l2)
	}

	l3, err := AtivaLancamento(db, numLanc03)
	if err != nil {
		t.Error(err, l3)
	}

	l4, err := AtivaLancamento(db, 0)
	if err.Error() != "Não foi encontrado um registro com o ID 0" {
		t.Error(err, l4)
	}

	if l1.Estado != true {
		t.Error("Estado do lancamento inválido, deveria ser false", l1)
	}

	if l2.Estado != true {
		t.Error("Estado do lancamento inválido, deveria ser false", l2)
	}

	if l3.Estado != true {
		t.Error("Estado do lancamento inválido, deveria ser false", l3)
	}

	lancamentos, err := CarregaLancamentosAtivo(db)
	if err != nil {
		t.Error(err, lancamentos)
	}

	if len(lancamentos) == 0 {
		t.Error(lancamentos)
	}

	if len(lancamentos) < 4 {
		t.Error(lancamentos)
	}
}

func TestCarregaLancamentos(t *testing.T) {
	listaLancamentos, err := CarregaLancamentos(db)

	if err != nil {
		t.Error(err, listaLancamentos)
	}

	if len(listaLancamentos) == 0 {
		t.Error(listaLancamentos)
	}

	if len(listaLancamentos) < 4 {
		t.Error(listaLancamentos)
	}
}

func TestProcuraLancamento(t *testing.T) {
	l1, err := ProcuraLancamento(db, numLanc01)
	if err != nil {
		t.Error(err, l1)
	}

	l2, err := ProcuraLancamento(db, numLanc02)
	if err != nil {
		t.Error(err, l2)
	}

	l3, err := ProcuraLancamento(db, numLanc03)
	if err != nil {
		t.Error(err, l3)
	}
}

func TestAlteraLancamento(t *testing.T) {
	novoLancamento := lancamento.GetLancamentoTest()
	novoLancamento.CpfPessoa = cpf

	l1, err := AlteraLancamento(db, numLanc01, novoLancamento)
	if err != nil {
		t.Error(err, l1)
	}

	if l1.CpfPessoa != novoLancamento.CpfPessoa ||
		l1.Data.Unix() != novoLancamento.Data.Unix() ||
		l1.Descricao != novoLancamento.Descricao ||
		l1.Numero != novoLancamento.Numero {
		t.Error("Erro na alteração de lancamento(ID ou CpfPessoa ou Data ou Numero ou Descricao)", l1, novoLancamento)
	}
}

func TestRemoveLancamento(t *testing.T) {
	err := RemoveLancamento(db, numLanc01)
	if err != nil {
		t.Error(err, numLanc01)
	}

	err = RemoveLancamento(db, numLanc02)
	if err != nil {
		t.Error(err, numLanc02)
	}

	err = RemoveLancamento(db, numLanc03)
	if err != nil {
		t.Error(err, numLanc03)
	}

	// Lancamento com a descrição "Lançamento teste 03"(var l4 de TestAdicionaLancamento(...)) é excluida automaticamente ao excluir a pessoa com cpf da variável cpf, por causa da CONSTRAINT pessoa_lancamento_fk com ON DELETE CASCADE na tabela lancamento do DB
	TestRemovePessoa(t)
}
