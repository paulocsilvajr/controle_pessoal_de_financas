package dao

import (
	"controle_pessoal_de_financas/API/v1/model/tipo_conta"
	"testing"
)

func TestAdicionaTipoConta(t *testing.T) {
	tc1 := tipo_conta.GetTipoContaTest()

	tc2 := tipo_conta.GetTipoContaTest()
	tc2.Nome = "banco teste 02"

	tc5 := tipo_conta.GetTipoContaTest()
	tc5.Nome = "banco teste 03"

	tc3, err := AdicionaTipoConta(db, tc1)
	strErroChavePrimariaDuplicada := "pq: duplicate key value violates unique constraint \"tipo_conta_pk\""
	if err != nil && err.Error() != strErroChavePrimariaDuplicada {
		t.Error(err, tc3)
	}

	tc4, err := AdicionaTipoConta(db, tc2)
	if err != nil && err.Error() != strErroChavePrimariaDuplicada {
		t.Error(err, tc4)
	}

	tc6, err := AdicionaTipoConta(db, tc5)
	if err != nil && err.Error() != strErroChavePrimariaDuplicada {
		t.Error(err, tc6)
	}
}

func TestCarregaTiposConta(t *testing.T) {
	listaTiposConta, err := CarregaTiposConta(db)

	if err != nil {
		t.Error(err, listaTiposConta)
	}

	if len(listaTiposConta) == 0 {
		t.Error(listaTiposConta)
	}

	if len(listaTiposConta) < 2 {
		t.Error(listaTiposConta)
	}
}

func TestCarregaTiposContaInativa(t *testing.T) {
	// listaTiposConta, err := CarregaTiposContaInativa(db)

	///////////// DESCOMENTAR QUANDO IMPLEMENTAR dao.InativaTipoConta ///////////////////////
	// if err != nil {
	// 	t.Error(err, listaTiposConta)
	// }

	// if len(listaTiposConta) == 0 {
	// 	t.Error(listaTiposConta)
	// }

	// if len(listaTiposConta) < 1 {
	// 	t.Error(listaTiposConta)
	// }
}
