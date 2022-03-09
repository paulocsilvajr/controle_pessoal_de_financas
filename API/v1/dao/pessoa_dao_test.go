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

func TestProcuraPessoa02(t *testing.T) {
	cpfProcurado := testPessoa01.Cpf
	p, err := ProcuraPessoa02(db2, cpfProcurado)

	if err != nil {
		t.Error(err)
	}

	cpfEncontrado := p.Cpf
	if cpfEncontrado != cpfProcurado {
		t.Errorf("CPF procurado diferente de CPF encontrado. Esperado: '%s', encontrado: '%s'", cpfProcurado, cpfEncontrado)
	}
}

func TestProcuraPessoaPorUsuario02(t *testing.T) {
	usuarioProcurado := testPessoaAdmin01.Usuario
	p, err := ProcuraPessoaPorUsuario02(db2, usuarioProcurado)

	if err != nil {
		t.Error(err)
	}

	usuarioEncontrado := p.Usuario
	if usuarioEncontrado != usuarioProcurado {
		t.Errorf("Usuário procurado diferente de usuário encontrado. Esperado: '%s', encontrado: '%s'", usuarioProcurado, usuarioEncontrado)
	}
}

func TestAlteraPessoa02(t *testing.T) {
	cpf := testPessoa01.Cpf
	novoNomeCompleto := "teste alteração"
	novoUsuario := "testeAlt01"
	novoEmail := "novoemail@gmail.com"

	testPessoa01.NomeCompleto = novoNomeCompleto
	testPessoa01.Usuario = novoUsuario
	testPessoa01.Email = novoEmail

	p, err := AlteraPessoa02(db2, cpf, testPessoa01)
	if err != nil {
		t.Error(err)
	}

	if p == nil {
		// feito RETURN para evitar PANIC por ponteiro vazio nas próximas verificações
		t.Errorf("Alteração retornou um ponteiro vazio[%v]", p)
		return
	}

	nomeCompleto := p.NomeCompleto
	if nomeCompleto != novoNomeCompleto {
		t.Errorf("Alteração de pessoa com CPF %s retornou um 'Nome Completo' diferente do esperado. Esperado: '%s', retorno: '%s'", cpf, novoNomeCompleto, nomeCompleto)
	}

	usuario := p.Usuario
	if usuario != novoUsuario {
		t.Errorf("Alteração de pessoa com CPF %s retornou um 'Usuário' diferente do esperado. Esperado: '%s', retorno: '%s'", cpf, novoUsuario, usuario)
	}

	email := p.Email
	if email != novoEmail {
		t.Errorf("Alteração de pessoa com CPF %s retornou um 'Email' diferente do esperado. Esperado: '%s', retorno: '%s'", cpf, novoEmail, email)
	}
}

func TestAlteraPessoaPorUsuario02(t *testing.T) {
	usuario := testPessoaAdmin01.Usuario
	novoNomeCompleto := "teste alteração admin"
	novoEmail := "novoemailadmin@gmail.com"

	testPessoaAdmin01.NomeCompleto = novoNomeCompleto
	testPessoaAdmin01.Email = novoEmail

	p, err := AlteraPessoaPorUsuario02(db2, usuario, testPessoaAdmin01)
	if err != nil {
		t.Error(err)
	}

	if p == nil {
		// feito RETURN para evitar PANIC por ponteiro vazio nas próximas verificações
		t.Errorf("Alteração retornou um ponteiro vazio[%v]", p)
		return
	}

	nomeCompleto := p.NomeCompleto
	if nomeCompleto != novoNomeCompleto {
		t.Errorf("Alteração de pessoa com usuário '%s' retornou um 'Nome Completo' diferente do esperado. Esperado: '%s', retorno: '%s'", usuario, novoNomeCompleto, nomeCompleto)
	}

	email := p.Email
	if email != novoEmail {
		t.Errorf("Alteração de pessoa com usuário '%s' retornou um 'Email' diferente do esperado. Esperado: '%s', retorno: '%s'", usuario, novoEmail, email)
	}
}

func TestInativaPessoa02(t *testing.T) {
	cpf := testPessoa01.Cpf

	p, err := InativaPessoa02(db2, cpf)
	if err != nil {
		t.Error(err)
	}

	if p == nil {
		// feito RETURN para evitar PANIC por ponteiro vazio nas próximas verificações
		t.Errorf("Inativação retornou um ponteiro vazio[%v]", p)
		return
	}

	estadoObtido := p.Estado
	estadoEsperado := false
	if estadoObtido != estadoEsperado {
		t.Errorf("Inativação de pessoa com CPF '%s' retornou um 'estado' diferente do esperado. Esperado: '%t', obtido: '%t'", cpf, estadoEsperado, estadoObtido)
	}
}

func TestAtivaPessoa02(t *testing.T) {
	cpf := testPessoa01.Cpf

	p, err := AtivaPessoa02(db2, cpf)
	if err != nil {
		t.Error(err)
	}

	if p == nil {
		// feito RETURN para evitar PANIC por ponteiro vazio nas próximas verificações
		t.Errorf("Ativação retornou um ponteiro vazio[%v]", p)
		return
	}

	estadoObtido := p.Estado
	estadoEsperado := true
	if estadoObtido != estadoEsperado {
		t.Errorf("Ativação de pessoa com CPF '%s' retornou um 'estado' diferente do esperado. Esperado: '%t', obtido: '%t'", cpf, estadoEsperado, estadoObtido)
	}
}

func TestSetAdministrador02(t *testing.T) {
	cpf := testPessoa01.Cpf
	administrador := true
	comum := false

	p, err := SetAdministrador02(db2, cpf, administrador)
	if err != nil {
		t.Error(err)
	}

	if p == nil {
		// feito RETURN para evitar PANIC por ponteiro vazio nas próximas verificações
		t.Errorf("Alterar administrador retornou um ponteiro vazio[%v]", p)
		return
	}

	valorAdministrador := p.Administrador
	if valorAdministrador != administrador {
		t.Errorf("Alteração de usuário comum para administrador de pessoa com CPF '%s' retornou um valor para chave 'administrador' diferente do esperado. Esperado: '%t', obtido: '%t'", cpf, administrador, valorAdministrador)
	}

	p, err = SetAdministrador02(db2, cpf, comum)
	if err != nil {
		t.Error(err)
	}

	if p == nil {
		// feito RETURN para evitar PANIC por ponteiro vazio nas próximas verificações
		t.Errorf("Alterar chave administrador retornou um ponteiro vazio[%v]", p)
		return
	}

	valorAdministrador = p.Administrador
	if valorAdministrador != comum {
		t.Errorf("Alteração de administrador para usuário comum de pessoa com CPF '%s' retornou um valor para chave 'administrador' diferente do esperado. Esperado: '%t', obtido: '%t'", cpf, comum, valorAdministrador)
	}
}

func TestCarregaPessoas02(t *testing.T) {
	pessoas, err := CarregaPessoas02(db2)
	if err != nil {
		t.Error(err)
	}

	quantPessoas := len(pessoas)
	quantEsperada := 2
	if quantPessoas != quantEsperada {
		t.Errorf("consulta de Pessoas retornou uma quantidade de registros diferente do esperado. Esperado: %d, obtido: %d", quantEsperada, quantPessoas)
	}

	cpf1 := testPessoa01.Cpf
	cpf2 := testPessoaAdmin01.Cpf
	for _, p := range pessoas {
		if p.Cpf == cpf1 || p.Cpf == cpf2 {
			continue
		}
		t.Errorf("registro de pessoa(CPF) encontrado diferente do esperado. Esperado: '%s' ou '%s', obtido: '%s'", cpf1, cpf2, p.Cpf)
	}
}

func TestCarregaPessoasSimples02(t *testing.T) {
	pessoasSimples, err := CarregaPessoasSimples02(db2)
	if err != nil {
		t.Error(err)
	}

	quantPessoas := len(pessoasSimples)
	quantEsperada := 2
	if quantPessoas != quantEsperada {
		t.Errorf("consulta de Pessoas Simples retornou uma quantidade de registros diferente do esperado. Esperado: %d, obtido: %d", quantEsperada, quantPessoas)
	}

	usuario1 := testPessoa01.Usuario
	usuario2 := testPessoaAdmin01.Usuario
	for _, p := range pessoasSimples {
		if p.Usuario == usuario1 || p.Usuario == usuario2 {
			continue
		}
		t.Errorf("registro de pessoa(Usuário) encontrado diferente do esperado. Esperado: '%s' ou '%s', obtido: '%s'", usuario1, usuario2, p.Usuario)
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
