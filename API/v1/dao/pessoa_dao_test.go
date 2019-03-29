package dao

import (
	"controle_pessoal_de_financas/API/v1/model/pessoa"
	"testing"
)

var (
	db  = GetDB()
	cpf = "38674832680"
)

func TestDaoAdicionaPessoa(t *testing.T) {
	p, _ := pessoa.GetPessoaTest()
	p.Cpf = cpf
	p.Usuario = "teste_inclusao"
	p.Email = "testei@gmail.com"
	p, err := DaoAdicionaPessoa(db, p)

	strErroChavePrimariaDuplicada := "pq: duplicate key value violates unique constraint \"pessoa_pk\""
	if err != nil && err.Error() != strErroChavePrimariaDuplicada {
		t.Error(err, p)
	}
}

func TestDaoCarregaPessoas(t *testing.T) {
	listaPessoas, err := DaoCarregaPessoas(db)

	if err != nil {
		t.Error(err, listaPessoas)
	}

	if len(listaPessoas) == 0 {
		t.Error(listaPessoas)
	}
}

func TestDaoProcuraPessoa(t *testing.T) {
	p, err := DaoProcuraPessoa(db, cpf)
	if err != nil {
		t.Error(err, p)
	}
}

func TestDaoAlteraPessoa(t *testing.T) {
	p1, _ := pessoa.GetPessoaTest()
	p1.NomeCompleto = "Teste Alteração"
	p1.Usuario = "teste_alteracao"
	p1.Senha = "123457"

	p2, err := DaoAlteraPessoa(db, cpf, p1)
	if err != nil {
		t.Error(err, p2)
	}
}

func TestDaoRemovePessoa(t *testing.T) {
	err := DaoRemovePessoa(db, cpf)
	if err != nil {
		t.Error(err, cpf)
	}
}
