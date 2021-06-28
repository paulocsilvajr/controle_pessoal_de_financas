package dao

import (
	"fmt"
	"strings"
	"testing"

	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/model/pessoa"
)

var (
	testPessoa01      *pessoa.Pessoa
	testPessoaAdmin01 *pessoa.Pessoa
)

func TestAdicionaPessoa02(t *testing.T) {
	tpessoa := getTPessoa1()
	testPessoa01 = ConverteTPessoaParaPessoa(tpessoa)

	p, err := AdicionaPessoa02(db2, testPessoa01)

	if err := verificaErroChaveDuplicada(err); err != nil {
		t.Error(err)
	}

	if err := verificaCamposPessoa(p, tpessoa); err != nil {
		t.Error(err)
	}

	if p.Administrador != false {
		t.Error("Pessoa Comum salva com flag Administrador true")
	}
}

func TestAdicionaPessoaAdmin02(t *testing.T) {
	tpessoa := getTPessoaAdmin1()
	testPessoaAdmin01 = ConverteTPessoaParaPessoa(tpessoa)

	p, err := AdicionaPessoaAdmin02(db2, testPessoaAdmin01)

	if err := verificaErroChaveDuplicada(err); err != nil {
		t.Error(err)
	}

	if err := verificaCamposPessoa(p, tpessoa); err != nil {
		t.Error(err)
	}

	if p.Administrador != true {
		t.Error("Pessoa Administrador salva com flag Adminstrador false")
	}
}

func TestRemovePessoa02(t *testing.T) {
	err := RemovePessoa02(db2, testPessoaAdmin01.Cpf)
	if err != nil {
		t.Error(err)
	}
}

func TestRemovePessoaPorUsuario02(t *testing.T) {
	err := RemovePessoaPorUsuario02(db2, testPessoa01.Usuario)
	if err != nil {
		t.Error(err)
	}
}

func verificaErroChaveDuplicada(err error) error {
	if err != nil {
		strErroChaveDuplicada := "duplicate key value violates unique constraint"
		erroNaoEhChaveDuplicada := !strings.Contains(err.Error(), strErroChaveDuplicada)

		if erroNaoEhChaveDuplicada {
			return err
		}
	}
	return nil
}

func verificaCamposPessoa(p *pessoa.Pessoa, tp *pessoa.TPessoa) error {
	if p.Cpf != tp.Cpf {
		return fmt.Errorf("Retornou CPF incorreto[%s != %s]", p.Cpf, tp.Cpf)
	}

	if p.NomeCompleto != tp.NomeCompleto {
		return fmt.Errorf("Retornou Nome Completo incorreto[%s != %s]", p.NomeCompleto, tp.NomeCompleto)
	}

	if p.Usuario != tp.Usuario {
		return fmt.Errorf("Retornou Usuário incorreto[%s != %s]", p.Usuario, tp.Usuario)
	}

	if p.Senha == tp.Senha {
		return fmt.Errorf("Retornou Senha sem HASH[%s != %s]", p.Senha, tp.Senha)
	}

	if p.Email != tp.Email {
		return fmt.Errorf("Retornou Email incorreto[%s != %s]", p.Email, tp.Email)
	}

	if p.Estado != true {
		return fmt.Errorf("Retornou Estado incorreto[%t != %t]", p.Estado, true)
	}

	return nil
}

// TESTES ANTIGOS
// import (
// 	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/helper"
// 	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/model/pessoa"
// 	"testing"
// )

// var (
// 	cpf     = "38674832680" // cpf inválido somente para teste(com Dígito Verificador[DV] incorreto)
// 	usuario = "teste_inclusao"
// )

// func TestAdicionaPessoa(t *testing.T) {
// 	p, _ := pessoa.GetPessoaTest()
// 	p.Cpf = cpf
// 	p.Usuario = usuario
// 	p.Email = "testei@gmail.com"
// 	p.Senha = "987321"
// 	p, err := AdicionaPessoa(db, p)

// 	strErroChavePrimariaDuplicada := "pq: duplicate key value violates unique constraint \"pessoa_pk\""
// 	if err != nil && err.Error() != strErroChavePrimariaDuplicada {
// 		t.Error(err, p)
// 	}
// }

// func TestCarregaPessoas(t *testing.T) {
// 	listaPessoas, err := CarregaPessoas(db)

// 	if err != nil {
// 		t.Error(err, listaPessoas)
// 	}

// 	if len(listaPessoas) == 0 {
// 		t.Error(listaPessoas)
// 	}
// }

// func TestProcuraPessoa(t *testing.T) {
// 	p, err := ProcuraPessoa(db, cpf)
// 	if err != nil {
// 		t.Error(err, p)
// 	}
// }

// func TestProcuraPessoaPorUsuario(t *testing.T) {
// 	p, err := ProcuraPessoaPorUsuario(db, usuario)
// 	if err != nil {
// 		t.Error(err, p)
// 	}
// }

// func TestAlteraPessoa(t *testing.T) {
// 	p1, _ := pessoa.GetPessoaTest()
// 	p1.NomeCompleto = "Teste Alteração"
// 	p1.Usuario = "teste_alteracao"
// 	p1.Senha = "123457"

// 	p2, err := AlteraPessoa(db, cpf, p1)
// 	if err != nil {
// 		t.Error(err, p2)
// 	}

// 	if p2.NomeCompleto != p1.NomeCompleto ||
// 		p2.Usuario != p1.Usuario ||
// 		p2.Senha != helper.GetSenhaSha256(p1.Senha) {
// 		t.Error("Erro na alteração de pessoa(NomeCompleto ou Usuario ou Senha)", p2)
// 	}
// }

// func TestInativaPessoa(t *testing.T) {
// 	p, err := InativaPessoa(db, cpf)
// 	if err != nil {
// 		t.Error(err, p)
// 	}

// 	if p.Estado != false {
// 		t.Error("Estado de pessoa inválido, deveria ser false", p)
// 	}
// }

// func TestAtivaPessoa(t *testing.T) {
// 	p, err := AtivaPessoa(db, cpf)
// 	if err != nil {
// 		t.Error(err, p)
// 	}

// 	if p.Estado != true {
// 		t.Error("Estado de pessoa inválido, deveria ser true", p)
// 	}
// }

// func TestRemovePessoa(t *testing.T) {
// 	err := RemovePessoa(db, cpf)
// 	if err != nil {
// 		t.Error(err, cpf)
// 	}
// }
