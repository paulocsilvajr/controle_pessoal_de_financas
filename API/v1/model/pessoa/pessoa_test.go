// pessoa_test.go
package pessoa

import (
	"controle_pessoal_de_financas/API/v1/model/erro"
	"testing"
	"time"
)

func TestMakePessoa(t *testing.T) {
	p1, err := GetPessoaTest()
	if err != nil {
		t.Error(err, p1)
	}

	if p1.String() != "12345678910	Teste 01	teste01	8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92	teste01@email.com	01/02/2000 12:30:00	01/02/2000 12:30:00	ativo	Comum" {
		t.Error("Erro na método String de Pessoa", p1)
	}

	if p1.Repr() != "12345678910	Teste 01	teste01	8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92	teste01@email.com	2000-02-01 12:30:00 +0000 UTC	2000-02-01 12:30:00 +0000 UTC	true	false" {
		t.Error("Erro no método Repr de Pessoa", p1)
	}

	p1, err = NewPessoa("11122233345", "Pedro de Alcântara João Gonzaga", "pedro", "654321", "pedroajoao@gmail.com")
	if err != nil {
		t.Error(err, p1)
	}

	if p1.VerificaAtributos() != nil {
		t.Error(err, p1)
	}
}

func TestMakePessoaAdmin(t *testing.T) {
	p1, err := NewPessoaAdmin("12345678910", "Teste 01", "teste01", "123456", "teste01@email.com")
	p1.DataCriacao = time.Date(2000, 2, 1, 12, 30, 0, 0, new(time.Location))
	p1.DataModificacao = time.Date(2000, 2, 1, 12, 30, 0, 0, new(time.Location))

	if !p1.Administrador {
		t.Error("Função NewPessoaAdmin não está criando uma Pessoa como Administrador", p1, err)
	}

	p2, err := GetPessoaTest()
	p2.SetAdmin(true)
	if !p2.Administrador {
		t.Error("Função SetAdmin(true) não está definindo como Administrador uma Pessoa", p1, err)
	}
}

func TestVerificaAtributosPessoa(t *testing.T) {
	p1, err := GetPessoaTest()

	cpf := "123456789101"
	p1.Cpf = cpf
	msgTamanhoIncorreto := erro.ErroTamanho(MsgErroCpf01, len(cpf)).Error()
	if p1.VerificaAtributos().Error() != msgTamanhoIncorreto {
		t.Error(err, p1)
	}

	cpf = "123.456.789-10"
	p1.Cpf = cpf
	msgTamanhoIncorreto = erro.ErroTamanho(MsgErroCpf01, len(cpf)).Error()
	if p1.VerificaAtributos().Error() != msgTamanhoIncorreto {
		t.Error(err, p1)
	}

	cpf = "123456789xx"
	p1.Cpf = cpf
	msgFormatoIncorreto := erro.ErroDetalhe(MsgErroCpf02, cpf).Error()
	if p1.VerificaAtributos().Error() != msgFormatoIncorreto {
		t.Error(err, p1)
	}

	p1, err = GetPessoaTest()

	nome := "Pedro de Alcântara João Carlos Leopoldo Salvador Bibiano Francisco Xavier de Paula Leocádio Miguel Gabriel Rafael Gonzaga"
	p1.NomeCompleto = nome
	msgTamanhoIncorreto = erro.ErroTamanho(MsgErroNome01, len(nome)).Error()
	if p1.VerificaAtributos().Error() != msgTamanhoIncorreto {
		t.Error(err, p1)
	}

	p1, err = GetPessoaTest()

	senha := `e_-'zWK6$y6Z7!cB>]2d!c7p;]d\bKFLku"ejf*+]g\sE=yjNuKD5Z.~p)#"A=C@[(V,^$`
	p1.Senha = senha
	msgTamanhoIncorreto = erro.ErroTamanho(MsgErroSenha01, len(senha)).Error()
	if p1.VerificaAtributos().Error() != msgTamanhoIncorreto {
		t.Error(err, p1)
	}

	p1, err = GetPessoaTest()

	email := "pedro_alcantara_joao_carlos_leopoldo_salvador_rafael_gonzaga@gmail.com"
	p1.Email = email
	msgTamanhoIncorreto = erro.ErroTamanho(MsgErroEmail01, len(email)).Error()
	if p1.VerificaAtributos().Error() != msgTamanhoIncorreto {
		t.Error(err, p1)
	}

	p1, err = GetPessoaTest()

	usuario := "pedro_alcantara_joao_rafael_gonzaga"
	p1.Usuario = usuario
	msgTamanhoIncorreto = erro.ErroTamanho(MsgErroUsuario01, len(usuario)).Error()
	if p1.VerificaAtributos().Error() != msgTamanhoIncorreto {
		t.Error(err, p1)
	}
}

func TestFuncoesInternasPessoa(t *testing.T) {
	cpf := "123.456.789-10"
	msgTamanhoIncorreto := erro.ErroTamanho(MsgErroCpf01, len(cpf)).Error()
	err := verificaCPF(cpf)
	if err.Error() != msgTamanhoIncorreto {
		t.Error(err, cpf)
	}

	cpf = "123456789xx"
	msgFormatoIncorreto := erro.ErroDetalhe(MsgErroCpf02, cpf).Error()
	err = verificaCPF(cpf)
	if err.Error() != msgFormatoIncorreto {
		t.Error(err, cpf)
	}

	nome := "Pedro de Alcântara João Carlos Leopoldo Salvador Bibiano Francisco Xavier de Paula Leocádio Miguel Gabriel Rafael Gonzaga"
	msgTamanhoIncorreto = erro.ErroTamanho(MsgErroNome01, len(nome)).Error()
	err = verificaNome(nome)
	if verificaNome(nome).Error() != msgTamanhoIncorreto {
		t.Error(err, nome)
	}

	senha := `e_-'zWK6$y6Z7!cB>]2d!c7p;]d\bKFLku"ejf*+]g\sE=yjNuKD5Z.~p)#"A=C@[(V,^$`
	msgTamanhoIncorreto = erro.ErroTamanho(MsgErroSenha01, len(senha)).Error()
	err = verificaSenha(senha)
	if err.Error() != msgTamanhoIncorreto {
		t.Error(err, senha)
	}

	email := "pedro_alcantara_joao_carlos_leopoldo_salvador_rafael_gonzaga@gmail.com"
	msgTamanhoIncorreto = erro.ErroTamanho(MsgErroEmail01, len(email)).Error()
	err = verificaEmail(email)
	if err.Error() != msgTamanhoIncorreto {
		t.Error(err, email)
	}

	usuario := "pedro_alcantara_joao_rafael_gonzaga"
	msgTamanhoIncorreto = erro.ErroTamanho(MsgErroUsuario01, len(usuario)).Error()
	err = verificaUsuario(usuario)
	if err.Error() != msgTamanhoIncorreto {
		t.Error(err, usuario)
	}

	usuario = "pedro_@ 2"
	msgFormatoIncorreto = erro.ErroDetalhe(MsgErroUsuario02, usuario).Error()
	err = verificaUsuario(usuario)
	if err.Error() != msgFormatoIncorreto {
		t.Error(err, usuario)
	}
}

func TestAlteraPessoa(t *testing.T) {
	p1, _ := GetPessoaTest()

	err := p1.Altera("12365487910", "Teste alterado", "usuario_2", "S3nh4N0v4", "teste@teste.com.br")
	if err != nil {
		t.Error(err)
	}

	err = p1.AlteraCampos(map[string]string{"senha": "123", "nome": "Teste alterado 2"})
	if err != nil {
		t.Error(err)
	}
}

func TestAlteraEstadoPessoa(t *testing.T) {
	p1, _ := GetPessoaTest()

	p1.Inativa()
	if p1.Estado != false {
		t.Error("Estado inválido, esperando false, obtido ", p1.Estado)
	}

	p1.Ativa()
	if p1.Estado != true {
		t.Error("Estado inválido, esperando true, obtido ", p1.Estado)
	}
}

func TestProcuraPessoaPorUsuario(t *testing.T) {
	p1, err := GetPessoaTest()
	p2, err := NewPessoa("12378945610", "Usuário Teste 02", "teste02", "147852", "teste02@email.com")

	pessoas := Pessoas{p1, p2}

	p3i, err := pessoas.ProcuraPessoaPorUsuario("teste02")
	if err != nil {
		t.Error(err, p3i)
	}

	p3 := p3i.(*Pessoa)
	if p3.Email != "teste02@email.com" {
		t.Error("Recuperado pelo método Pessoas.ProcuraPessoaPorUsuario() uma Pessoa com email diferente do informado na sua criação", p3)
	}
}
