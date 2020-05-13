package dao

import (
	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/helper"
	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/model/pessoa"
	"testing"
)

var (
	cpf     = "38674832680" // cpf inválido somente para teste(com Dígito Verificador[DV] incorreto)
	usuario = "teste_inclusao"
)

func TestAdicionaPessoa(t *testing.T) {
	p, _ := pessoa.GetPessoaTest()
	p.Cpf = cpf
	p.Usuario = usuario
	p.Email = "testei@gmail.com"
	p.Senha = "987321"
	p, err := AdicionaPessoa(db, p)

	strErroChavePrimariaDuplicada := "pq: duplicate key value violates unique constraint \"pessoa_pk\""
	if err != nil && err.Error() != strErroChavePrimariaDuplicada {
		t.Error(err, p)
	}
}

func TestCarregaPessoas(t *testing.T) {
	listaPessoas, err := CarregaPessoas(db)

	if err != nil {
		t.Error(err, listaPessoas)
	}

	if len(listaPessoas) == 0 {
		t.Error(listaPessoas)
	}
}

func TestProcuraPessoa(t *testing.T) {
	p, err := ProcuraPessoa(db, cpf)
	if err != nil {
		t.Error(err, p)
	}
}

func TestProcuraPessoaPorUsuario(t *testing.T) {
	p, err := ProcuraPessoaPorUsuario(db, usuario)
	if err != nil {
		t.Error(err, p)
	}
}

func TestAlteraPessoa(t *testing.T) {
	p1, _ := pessoa.GetPessoaTest()
	p1.NomeCompleto = "Teste Alteração"
	p1.Usuario = "teste_alteracao"
	p1.Senha = "123457"

	p2, err := AlteraPessoa(db, cpf, p1)
	if err != nil {
		t.Error(err, p2)
	}

	if p2.NomeCompleto != p1.NomeCompleto ||
		p2.Usuario != p1.Usuario ||
		p2.Senha != helper.GetSenhaSha256(p1.Senha) {
		t.Error("Erro na alteração de pessoa(NomeCompleto ou Usuario ou Senha)", p2)
	}
}

func TestInativaPessoa(t *testing.T) {
	p, err := InativaPessoa(db, cpf)
	if err != nil {
		t.Error(err, p)
	}

	if p.Estado != false {
		t.Error("Estado de pessoa inválido, deveria ser false", p)
	}
}

func TestAtivaPessoa(t *testing.T) {
	p, err := AtivaPessoa(db, cpf)
	if err != nil {
		t.Error(err, p)
	}

	if p.Estado != true {
		t.Error("Estado de pessoa inválido, deveria ser true", p)
	}
}

func TestRemovePessoa(t *testing.T) {
	err := RemovePessoa(db, cpf)
	if err != nil {
		t.Error(err, cpf)
	}
}
