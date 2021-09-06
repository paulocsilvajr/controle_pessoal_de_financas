package dao

import (
	"testing"

	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/model/conta"
	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/model/detalhe_lancamento"
	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/model/lancamento"
	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/model/pessoa"
	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/model/tipo_conta"
)

var (
	testLancamentoB01        *lancamento.Lancamento
	testPessoaAdminB02       *pessoa.Pessoa
	testContaB01             *conta.Conta
	testTipoContaC01         *tipo_conta.TipoConta
	testDetalheLancamentoB01 *detalhe_lancamento.DetalheLancamento
)

func TestAdicionaDetalheLancamento02(t *testing.T) {
	var err error

	p := getPessoaAdmin2()
	testPessoaAdminB02, err = AdicionaPessoa02(db2, p)
	if err != nil {
		t.Error(err)
	}

	l := getLancamento2(testPessoaAdminB02)
	testLancamentoB01, err = AdicionaLancamento02(db2, l)
	if err != nil {
		t.Error(err)
	}

	tc := getTipoConta1()
	testTipoContaC01, err = AdicionaTipoConta02(db2, tc)
	if err != nil {
		t.Error(err)
	}

	c := getConta1(testTipoContaC01)
	testContaB01, err = AdicionaConta02(db2, c)
	if err != nil {
		t.Error(err)
	}

	dl := getDetalheLancamento(testLancamentoB01, testContaB01)
	testDetalheLancamentoB01, err = AdicionaDetalheLancamento02(db2, dl)
	if err != nil {
		t.Error(err)
	}

	idLancamentoEsperado := testLancamentoB01.ID
	nomeContaEsperado := testContaB01.Nome
	creditoEsperado := dl.Credito
	debitoEsperado := dl.Debito

	idLancamento := testDetalheLancamentoB01.IDLancamento
	nomeConta := testDetalheLancamentoB01.NomeConta
	credito := testDetalheLancamentoB01.Credito
	debito := testDetalheLancamentoB01.Debito

	if idLancamento != idLancamentoEsperado {
		t.Errorf("ID Lancamento do detalhe lançamento diferente do esperado. Esperado: %d, obtido: %d", idLancamentoEsperado, idLancamento)
	}

	if nomeConta != nomeContaEsperado {
		t.Errorf("Nome da conta do detalhe lançamento %d diferente do esperado. Esperado: '%s', obtido: '%s'", idLancamento, nomeContaEsperado, nomeConta)
	}

	if credito != creditoEsperado {
		t.Errorf("Crédito do detalhe lançamento %d diferente do esperado. Esperado: %f, obtido: %f", idLancamento, credito, creditoEsperado)
	}

	if debito != debitoEsperado {
		t.Errorf("Débito do detalhe lançamento %d diferente do esperado. Esperado: %f, obtido: %f", idLancamento, debito, debitoEsperado)
	}
}

func TestCarregaDetalheLancamento02(t *testing.T) {
	detLancamentos, err := CarregaDetalheLancamentos02(db2)
	if err != nil {
		t.Error(err)
	}

	quantDetLancamentos := len(detLancamentos)
	quantEsperada := 1
	if quantDetLancamentos != quantEsperada {
		t.Errorf("consulta de Detalhe Lancamentos retornou uma quantidade de registros diferente do esperado. Esperado: %d, obtido: %d", quantEsperada, quantDetLancamentos)
	}

	idLancamento := testDetalheLancamentoB01.IDLancamento
	nomeConta := testDetalheLancamentoB01.NomeConta
	for _, dl := range detLancamentos {
		if dl.IDLancamento != idLancamento {
			t.Errorf("registro de detalhe lancamento(IDLancamento) encontrado diferente do esperado. Esperado: %d, obtido: %d", idLancamento, dl.IDLancamento)
		}

		if dl.NomeConta != nomeConta {
			t.Errorf("registro de detalhe lançamento(NomeConta) encontrado diferente do esperado. Esperado: '%s', obtido: '%s'", nomeConta, dl.NomeConta)
		}
	}

}

func TestCarregaDetalheLancamentosPorIDLancamento02(t *testing.T) {
	idLanc := testLancamentoB01.ID
	detLancamentos, err := CarregaDetalheLancamentosPorIDLancamento02(db2, idLanc)
	if err != nil {
		t.Error(err)
	}

	quantDetLancamentos := len(detLancamentos)
	quantEsperada := 1
	if quantDetLancamentos != quantEsperada {
		t.Errorf("consulta de Detalhe Lançamentos por id lançamento retornou uma quantidade de registros diferente do esperado. Esperado: %d, obtido: %d", quantEsperada, quantDetLancamentos)
	}
}

func TestCarregaDetalheLancamentosPorNomeConta02(t *testing.T) {
	nomeConta := testContaB01.Nome
	detLancamentos, err := CarregaDetalheLancamentosPorNomeConta02(db2, nomeConta)
	if err != nil {
		t.Error(err)
	}

	quantDetLancamentos := len(detLancamentos)
	quantEsperada := 1
	if quantDetLancamentos != quantEsperada {
		t.Errorf("consulta de Detalhe Lançamentos por nome de conta retornou uma quantidade de registros diferente do esperado. Esperado: %d, obtido: %d", quantEsperada, quantDetLancamentos)
	}
}

func TestProcuraDetalheLancamento02(t *testing.T) {
	idLancamento := testLancamentoB01.ID
	nomeConta := testContaB01.Nome
	dl, err := ProcuraDetalheLancamento02(db2, idLancamento, nomeConta)
	if err != nil {
		t.Error(err)
	}

	idLancamentoEncontrado := dl.IDLancamento
	if idLancamentoEncontrado != idLancamento {
		t.Errorf("ID lançamento procurado diferente do encontrado. Esperado: %d, obtido: %d", idLancamento, idLancamentoEncontrado)
	}

	nomeContaEncontrada := dl.NomeConta
	if nomeContaEncontrada != nomeConta {
		t.Errorf("Nome de conta procurado diferente do encontrado. Esperado: '%s', obtido: '%s'", nomeConta, nomeContaEncontrada)
	}
}

func TestAlteraDetalheLancamento02(t *testing.T) {
	idLancamento := testDetalheLancamentoB01.IDLancamento
	nomeConta := testDetalheLancamentoB01.NomeConta
	novoValorCredito := 0.0
	novoValorDebito := 250.0

	testDetalheLancamentoB01.Credito = novoValorCredito
	testDetalheLancamentoB01.Debito = novoValorDebito

	transacao := db2.Begin()
	dl, err := AlteraDetalheLancamento02(db2, transacao, idLancamento, nomeConta, testDetalheLancamentoB01)
	transacao.Commit()
	if err != nil {
		t.Error(err)
	}

	credito := dl.Credito
	if credito != novoValorCredito {
		t.Errorf("alteração de detalhe lançamento com ID %d e nome conta '%s' retornou um 'Crédito' diferente do esperado. Esperado: %f, obtido: %f", idLancamento, nomeConta, novoValorCredito, credito)
	}

	debito := dl.Debito
	if debito != novoValorDebito {
		t.Errorf("alteração de detalhe lançamento com ID %d e nome conta '%s' retornou um 'Débito' diferente do esperado. Esperado: %f, obtido: %f", idLancamento, nomeConta, novoValorDebito, debito)
	}
}

func TestCarregaLancamentosPorCPFeConta02(t *testing.T) {
	cpf := testPessoaAdminB02.Cpf
	nomeConta := testContaB01.Nome
	l, err := CarregaLancamentosPorCPFeConta02(db2, cpf, nomeConta)
	if err != nil {
		t.Error(err)
	}

	quantEsperada := 1
	quantObtida := l.Len()
	if quantObtida != quantEsperada {
		t.Errorf("consulta de lançamentos por CPF '%s' e conta '%s' retornou uma quantidade de registros incorreta. Esperado: %d, obtido: %d", cpf, nomeConta, quantEsperada, quantObtida)
	}
}

func TestCarregaLancamentosInativosPorCPFeConta02(t *testing.T) {
	cpf := testPessoaAdminB02.Cpf
	nomeConta := testContaB01.Nome
	l, err := CarregaLancamentosInativosPorCPFeConta02(db2, cpf, nomeConta)
	if err != nil {
		t.Error(err)
	}

	quantEsperada := 0
	quantObtida := l.Len()
	if quantObtida != quantEsperada {
		t.Errorf("consulta de lançamentos por CPF '%s' e conta '%s' retornou uma quantidade de registros incorreta. Esperado: %d, obtido: %d", cpf, nomeConta, quantEsperada, quantObtida)
	}
}

func TestCarregaLancamentosAtivosPorCPFeConta02(t *testing.T) {
	cpf := testPessoaAdminB02.Cpf
	nomeConta := testContaB01.Nome
	l, err := CarregaLancamentosAtivosPorCPFeConta02(db2, cpf, nomeConta)
	if err != nil {
		t.Error(err)
	}

	quantEsperada := 1
	quantObtida := l.Len()
	if quantObtida != quantEsperada {
		t.Errorf("consulta de lançamentos por CPF '%s' e conta '%s' retornou uma quantidade de registros incorreta. Esperado: %d, obtido: %d", cpf, nomeConta, quantEsperada, quantObtida)
	}
}

func TestRemoveDetalheLancamento02(t *testing.T) {
	var err error
	err = RemoveDetalheLancamento02(db2,
		testDetalheLancamentoB01.IDLancamento,
		testDetalheLancamentoB01.NomeConta,
	)
	if err != nil {
		t.Error(err)
	}

	err = RemoveLancamento02(db2, testLancamentoB01.ID)
	if err != nil {
		t.Error(err)
	}

	err = RemoveConta02(db2, testContaB01.Nome)
	if err != nil {
		t.Error(err)
	}

	err = RemoveTipoConta02(db2, testTipoContaC01.Nome)
	if err != nil {
		t.Error(err)
	}

	err = RemovePessoa02(db2, testPessoaAdminB02.Cpf)
	if err != nil {
		t.Error(err)
	}
}
