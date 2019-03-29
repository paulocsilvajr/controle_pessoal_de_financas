package dao

import (
	"controle_pessoal_de_financas/API/v1/model/pessoa"
	"testing"
)

var (
	db  = GetDB()
	cpf = "38674832680"
)

func TestDaoCarregaPessoas(t *testing.T) {
	listaPessoas, err := DaoCarregaPessoas(db)

	if err != nil {
		t.Error(err, listaPessoas)
	}

	// if len(listaPessoas) > 0 {
	// 	t.Error(listaPessoas)
	// }
}

func TestDaoAdicionaPessoa(t *testing.T) {
	p, _ := pessoa.GetPessoaTest()
	p.Cpf = cpf
	p.Usuario = "teste_inclusao"
	p, err := DaoAdicionaPessoa(db, p)

	strErroChavePrimariaDuplicada := "pq: duplicate key value violates unique constraint \"pessoa_pk\""
	if err != nil && err.Error() != strErroChavePrimariaDuplicada {
		t.Error(err, p)
	}
}

func TestDaoRemovePessoa(t *testing.T) {
	err := DaoRemovePessoa(db, cpf)
	if err != nil {
		t.Error(err, cpf)
	}
}

func TestDaoProcuraPessoa(t *testing.T) {
	TestDaoAdicionaPessoa(t)
	p, err := DaoProcuraPessoa(db, cpf)
	if err != nil {
		t.Error(err, p)
	}
}
