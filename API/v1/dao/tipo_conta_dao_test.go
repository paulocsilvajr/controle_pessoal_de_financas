package dao

import (
	"testing"

	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/model/tipo_conta"
)

var (
	testTipoConta01, testTipoConta02, testTipoConta03 *tipo_conta.TipoConta
)

func TestAdicionaTipoConta02(t *testing.T) {
	nome01 := "banco teste 01"
	testTipoConta01 = ConverteTTipoContaParaTipoConta(getTTipoConta1())
	testTipoConta01.Nome = nome01

	nome02 := "banco teste 02"
	testTipoConta02 = ConverteTTipoContaParaTipoConta(getTTipoConta1())
	testTipoConta02.Nome = nome02

	nome03 := "banco teste 03"
	testTipoConta03 = ConverteTTipoContaParaTipoConta(getTTipoConta1())
	testTipoConta03.Nome = nome03

	testTipoConta04 := tipo_conta.New("EVA01", "", "")

	testTipoConta01, err := AdicionaTipoConta02(db2, testTipoConta01)
	if err != nil {
		t.Error(err, testTipoConta01)
	}

	testTipoConta02, err := AdicionaTipoConta02(db2, testTipoConta02)
	if err != nil {
		t.Error(err, testTipoConta02)
	}

	testTipoConta03, err := AdicionaTipoConta02(db2, testTipoConta03)
	if err != nil {
		t.Error(err, testTipoConta03)
	}

	testTipoConta04, err = AdicionaTipoConta02(db2, testTipoConta04)
	erroDescricao := "Descrição do débito com tamanho inválido[0]"
	if err.Error() != erroDescricao {
		t.Error(err, testTipoConta04)
	}
}

func TestProcuraTipoConta02(t *testing.T) {
	tiposConta := tipo_conta.TiposConta{
		testTipoConta01,
		testTipoConta02,
		testTipoConta03,
	}

	for _, tcs := range tiposConta {
		nomeTipoContaProcurado := tcs.Nome
		tc, err := ProcuraTipoConta02(db2, nomeTipoContaProcurado)
		if err != nil {
			t.Error(err, tc)
		}

		nomeTipoContaEncontrado := tc.Nome
		if nomeTipoContaEncontrado != nomeTipoContaProcurado {
			t.Errorf("Nome de tipo de conta procurado diferente do encontrado. Esperado: '%s', obtido: '%s'", nomeTipoContaProcurado, nomeTipoContaEncontrado)
		}
	}

	nomeTipoContaProcurado := "Tipo Conta inexistente"
	tc, err := ProcuraTipoConta02(db2, nomeTipoContaProcurado)
	if err == nil {
		t.Error(err, tc)
	}

	if tc != nil {
		t.Error("Retornou um 'TipoConta' para um nome de tipo de conta inexistente", tc)
	}
}

func TestInativaTipoConta02(t *testing.T) {
	nome01 := testTipoConta01.Nome
	tc, err := InativaTipoConta02(db2, nome01)
	if err != nil {
		t.Error(err)
	}

	if tc != nil {
		estadoObtido := tc.Estado
		estafoEsperado := false

		if estadoObtido != estafoEsperado {
			t.Errorf("Inativação de tipo conta com nome '%s' retornou um estado diferente do esperado. Esperado: '%t', obtido: '%t'", nome01, estafoEsperado, estadoObtido)
		}
	} else {
		t.Errorf("Inativação retornou um ponteiro vazio[%v]", tc)
	}
}

func TestCarregaTiposContaInativa02(t *testing.T) {
	tiposConta, err := CarregaTiposContaInativa02(db2)
	if err != nil {
		t.Error(err)
	}

	quantTiposConta := len(tiposConta)
	quantEsperada := 1
	if quantTiposConta != quantEsperada {
		t.Errorf("consulta de tipos conta retornou uma quantidade de registros diferente do esperado. Esperado: %d, obtido: %d", quantEsperada, quantTiposConta)
	}
}

func TestAtivaTipoConta02(t *testing.T) {
	nome01 := testTipoConta01.Nome
	tc, err := AtivaTipoConta02(db2, nome01)
	if err != nil {
		t.Error(err)
	}

	if tc != nil {
		estadoObtido := tc.Estado
		estafoEsperado := true

		if estadoObtido != estafoEsperado {
			t.Errorf("Ativação de tipo conta(inativa) com nome '%s' retornou um estado diferente do esperado. Esperado: '%t', obtido: '%t'", nome01, estafoEsperado, estadoObtido)
		}
	} else {
		t.Errorf("Ativação retornou um ponteiro vazio[%v]", tc)
	}
}

func TestCarregaTiposContaAtiva02(t *testing.T) {
	tiposConta, err := CarregaTiposContaAtiva02(db2)
	if err != nil {
		t.Error(err)
	}

	quantTiposConta := len(tiposConta)
	quantEsperada := 3
	if quantTiposConta != quantEsperada {
		t.Errorf("consulta de tipos conta(ativa) retornou uma quantidade de registros diferente do esperado. Esperado: %d, obtido: %d", quantEsperada, quantTiposConta)
	}
}

func TestAlteraTipoConta02(t *testing.T) {
	nome := testTipoConta01.Nome
	novaDescricaoDebito := "débito alterado"
	novaDescricaoCredito := "crédito alterado"

	testTipoConta01.DescricaoDebito = novaDescricaoDebito
	testTipoConta01.DescricaoCredito = novaDescricaoCredito

	tc, err := AlteraTipoConta02(db2, nome, testTipoConta01)
	if err != nil {
		t.Error(err)
	}

	if tc != nil {
		descricaoDebito := tc.DescricaoDebito
		if descricaoDebito != novaDescricaoDebito {
			t.Errorf("Alteração de tipo conta com nome '%s' retornou um descrição de débito diferente do esperado. Esperado: '%s', obtido: '%s'", nome, novaDescricaoDebito, descricaoDebito)
		}

		descricaoCredito := tc.DescricaoCredito
		if descricaoCredito != novaDescricaoCredito {
			t.Errorf("Alteração de tipo conta com nome '%s' retornou um descrição de crédito diferente do esperado. Esperado: '%s', obtido: '%s'", nome, novaDescricaoCredito, descricaoCredito)
		}
	}
}

func TestCarregaTiposConta02(t *testing.T) {
	tiposConta, err := CarregaTiposConta02(db2)
	if err != nil {
		t.Error(err)
	}

	quantTiposConta := len(tiposConta)
	quantEsperada := 3
	if quantTiposConta != quantEsperada {
		t.Errorf("consulta de tipos conta retornou uma quantidade de registros diferente do esperado. Esperado: %d, obtido: %d", quantEsperada, quantTiposConta)
	}

	nome01 := testTipoConta01.Nome
	nome02 := testTipoConta02.Nome
	nome03 := testTipoConta03.Nome
	for _, p := range tiposConta {
		if p.Nome == nome01 || p.Nome == nome02 || p.Nome == nome03 {
			continue
		}
		t.Errorf("registro de tipo conta(nome) encontrado diferente do esperado. Esperado: '%s' ou '%s' ou '%s', obtido: '%s'", nome01, nome02, nome03, p.Nome)
	}
}

func TestRemoveTipoConta02(t *testing.T) {
	tiposConta := tipo_conta.TiposConta{
		testTipoConta01,
		testTipoConta02,
		testTipoConta03,
	}

	for _, tcs := range tiposConta {
		err := RemoveTipoConta02(db2, tcs.Nome)
		if err != nil {
			t.Error(err)
		}
	}
}
