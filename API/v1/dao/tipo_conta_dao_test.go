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

// TESTES ANTIGOS
// import (
// 	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/model/tipo_conta"
// 	"testing"
// )

// func TestAdicionaTipoConta(t *testing.T) {
// tc1 := tipo_conta.GetTipoContaTest()

// 	tc2 := tipo_conta.GetTipoContaTest()
// 	tc2.Nome = "banco teste 02"

// 	tc5 := tipo_conta.GetTipoContaTest()
// 	tc5.Nome = "banco teste 03"

// 	tc7 := tipo_conta.New("EVA01", "", "")

// 	tc3, err := AdicionaTipoConta(db, tc1)
// 	strErroChavePrimariaDuplicada := "pq: duplicate key value violates unique constraint \"tipo_conta_pk\""
// 	if err != nil && err.Error() != strErroChavePrimariaDuplicada {
// 		t.Error(err, tc3)
// 	}

// 	tc4, err := AdicionaTipoConta(db, tc2)
// 	if err != nil && err.Error() != strErroChavePrimariaDuplicada {
// 		t.Error(err, tc4)
// 	}

// 	tc6, err := AdicionaTipoConta(db, tc5)
// 	if err != nil && err.Error() != strErroChavePrimariaDuplicada {
// 		t.Error(err, tc6)
// 	}

// 	tc8, err := AdicionaTipoConta(db, tc7)
// 	if err.Error() != "Descrição do débito com tamanho inválido[0]" {
// 		t.Error(err, tc8)
// 	}
// }

// func TestInativaTipoConta(t *testing.T) {
// 	nome01 := "banco teste 01"
// 	nome02 := "banco teste 02"
// 	nome03 := "banco teste 03"
// 	nome04 := "banco teste 04"

// 	tc01, err := InativaTipoConta(db, nome01)
// 	if err != nil {
// 		t.Error(err, tc01)
// 	}

// 	tc02, err := InativaTipoConta(db, nome02)
// 	if err != nil {
// 		t.Error(err, tc02)
// 	}

// 	tc03, err := InativaTipoConta(db, nome03)
// 	if err != nil {
// 		t.Error(err, tc03)
// 	}

// 	tc04, err := InativaTipoConta(db, nome04)
// 	if err.Error() != "Não foi encontrado um registro com o nome banco teste 04" {
// 		t.Error(err, tc04)
// 	}

// 	if tc01.Estado != false {
// 		t.Error("Estado de tipo conta inválido, deveria ser false", tc01)
// 	}

// 	if tc02.Estado != false {
// 		t.Error("Estado de tipo conta inválido, deveria ser false", tc02)
// 	}

// 	if tc03.Estado != false {
// 		t.Error("Estado de tipo conta inválido, deveria ser false", tc03)
// 	}
// }

// func TestAtivaTipoConta(t *testing.T) {
// 	nome01 := "banco teste 01"
// 	nome02 := "banco teste 02"
// 	nome04 := "banco teste 04"

// 	tc01, err := AtivaTipoConta(db, nome01)
// 	if err != nil {
// 		t.Error(err, tc01)
// 	}

// 	tc02, err := AtivaTipoConta(db, nome02)
// 	if err != nil {
// 		t.Error(err, tc02)
// 	}

// 	tc04, err := AtivaTipoConta(db, nome04)
// 	if err.Error() != "Não foi encontrado um registro com o nome banco teste 04" {
// 		t.Error(err, tc04)
// 	}

// 	if tc01.Estado != true {
// 		t.Error("Estado de tipo conta inválido, deveria ser true", tc01)
// 	}

// 	if tc02.Estado != true {
// 		t.Error("Estado de tipo conta inválido, deveria ser true", tc02)
// 	}
// }

// func TestCarregaTiposConta(t *testing.T) {
// 	listaTiposConta, err := CarregaTiposConta(db)

// 	if err != nil {
// 		t.Error(err, listaTiposConta)
// 	}

// 	if len(listaTiposConta) == 0 {
// 		t.Error(listaTiposConta)
// 	}

// 	if len(listaTiposConta) < 3 {
// 		t.Error(listaTiposConta)
// 	}
// }

// func TestCarregaTiposContaInativa(t *testing.T) {
// 	listaTiposConta, err := CarregaTiposContaInativa(db)

// 	if err != nil {
// 		t.Error(err, listaTiposConta)
// 	}

// 	if len(listaTiposConta) == 0 {
// 		t.Error(listaTiposConta)
// 	}

// 	if len(listaTiposConta) < 1 {
// 		t.Error(listaTiposConta)
// 	}
// }

// func TestProcuraTipoConta(t *testing.T) {
// 	nome01 := "banco teste 01"
// 	nome02 := "banco teste 02"
// 	nome03 := "banco teste 03"
// 	nome04 := "banco teste 04"

// 	tc01, err := ProcuraTipoConta(db, nome01)
// 	if err != nil {
// 		t.Error(err, tc01)
// 	}

// 	tc02, err := ProcuraTipoConta(db, nome02)
// 	if err != nil {
// 		t.Error(err, tc02)
// 	}

// 	tc03, err := ProcuraTipoConta(db, nome03)
// 	if err != nil {
// 		t.Error(err, tc03)
// 	}

// 	tc04, err := ProcuraTipoConta(db, nome04)
// 	if err.Error() != "Não foi encontrado um registro com o nome banco teste 04" {
// 		t.Error(err, tc04)
// 	}
// }

// func TestAlteraTipoConta(t *testing.T) {
// 	tc01 := tipo_conta.GetTipoContaTest()
// 	nome := tc01.Nome
// 	tc01.Nome = "Teste alteração 01"
// 	tc01.DescricaoDebito = "Débito"
// 	tc01.DescricaoCredito = "Crédito"

// 	tc02, err := AlteraTipoConta(db, nome, tc01)
// 	if err != nil {
// 		t.Error(err, tc02)
// 	}

// 	if tc02.Nome != tc01.Nome || tc02.DescricaoDebito != tc01.DescricaoDebito || tc02.DescricaoCredito != tc01.DescricaoCredito {
// 		t.Error("Erro na alteração de tipo conta(Nome ou DescricaoDebito ou DescricaoCredito)", tc02)
// 	}

// 	nomeAlterado := tc02.Nome
// 	tc02.Nome = nome
// 	tc02, err = AlteraTipoConta(db, nomeAlterado, tc02)
// 	if err != nil {
// 		t.Error(err, tc02)
// 	}

// 	nome = "banco teste 04"
// 	_, err = AlteraTipoConta(db, nome, tc01)
// 	if err.Error() != "Não foi encontrado um registro com o nome banco teste 04" {
// 		t.Error(err)
// 	}

// 	tc03 := tipo_conta.GetTipoContaTest()
// 	tc03.DescricaoDebito = ""
// 	_, err = AlteraTipoConta(db, tc02.Nome, tc03)
// 	if err.Error() != "Descrição do débito com tamanho inválido[0]" {
// 		t.Error(err)
// 	}
// }

// func TestRemoveTipoConta(t *testing.T) {
// 	nome01 := "banco teste 01"
// 	nome02 := "banco teste 02"
// 	nome03 := "banco teste 03"
// 	nome04 := "banco teste 04"

// 	err := RemoveTipoConta(db, nome01)
// 	if err != nil {
// 		t.Error(err, nome01)
// 	}

// 	err = RemoveTipoConta(db, nome02)
// 	if err != nil {
// 		t.Error(err, nome02)
// 	}

// 	err = RemoveTipoConta(db, nome03)
// 	if err != nil {
// 		t.Error(err, nome03)
// 	}

// 	err = RemoveTipoConta(db, nome04)
// 	if err.Error() != "Não foi encontrado um registro com o nome banco teste 04" {
// 		t.Error(err, nome04)
// 	}
// }
