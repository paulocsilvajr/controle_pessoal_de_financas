package dao

import (
	"controle_pessoal_de_financas/API/v1/model/conta"
	"controle_pessoal_de_financas/API/v1/model/tipo_conta"
	"testing"
)

func TestAdicionaConta(t *testing.T) {
	c1 := conta.GetContaTest()

	c2 := conta.GetContaTest()
	c2.Nome = "Ativos Teste 02"

	c3 := conta.GetContaTest()
	c3.Nome = "Ativos Teste 03"

	c4 := conta.New("Teste conta C", "", "", "", "")

	c5, err := AdicionaConta(db, c1)
	strErroChaveEstrangeira := "pq: insert or update on table \"conta\" violates foreign key constraint \"tipo_conta_conta_fk\""
	if err != nil && err.Error() != strErroChaveEstrangeira {
		t.Error(err, c5)
	}

	tc1, _ := AdicionaTipoConta(db, tipo_conta.GetTipoContaTest())
	c1.NomeTipoConta = tc1.Nome
	c5, err = AdicionaConta(db, c1)
	if err != nil {
		t.Error(err, c5)
	}

	c6, err := AdicionaConta(db, c2)
	strErroChaveUnica := "pq: duplicate key value violates unique constraint \"codigo_uq\""
	if err != nil && err.Error() != strErroChaveUnica {
		t.Error(err, c6)
	}

	c2.Codigo = "002"
	c6, err = AdicionaConta(db, c2)
	if err != nil {
		t.Error(err, c6)
	}

	c3.Codigo = "003"
	c7, err := AdicionaConta(db, c3)
	if err != nil {
		t.Error(err, c7)
	}

	c8, err := AdicionaConta(db, c4)
	if err.Error() != "Tamanho de campo Nome do Tipo da Conta inválido[0 caracter(es)]" {
		t.Error(err, c8)
	}
}

func TestCarregaContas(t *testing.T) {
	listaContas, err := CarregaContas(db)

	if err != nil {
		t.Error(err, listaContas)
	}

	if len(listaContas) == 0 {
		t.Error(listaContas)
	}

	if len(listaContas) < 3 {
		t.Error(listaContas)
	}
}

func TestRemoveConta(t *testing.T) {
	nome01 := "Ativos Teste 01"
	nome02 := "Ativos Teste 02"
	nome03 := "Ativos Teste 03"
	nomeTipoConta01 := "banco teste 01"

	err := RemoveConta(db, nome01)
	if err != nil {
		t.Error(err, nome01)
	}

	err = RemoveConta(db, nome02)
	if err != nil {
		t.Error(err, nome02)
	}

	err = RemoveConta(db, nome03)
	if err != nil {
		t.Error(err, nome03)
	}

	err = RemoveTipoConta(db, nomeTipoConta01)
	if err != nil {
		t.Error(err, nomeTipoConta01)
	}
}
