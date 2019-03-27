// pessoa_test.go
package pessoa

import (
	"controle_pessoal_de_financas/API/v1/model/erro"
	"testing"
	"time"
)

func getPessoaTest() (pessoa *Pessoa, err error) {
	// location, _ := time.LoadLocation("America/São Paulo")
	pessoa, err = NewPessoa("12345678910", "Teste 01", "teste01", "123456", "teste01@email.com")
	pessoa.DataCriacao = time.Date(2000, 2, 1, 12, 30, 0, 0, new(time.Location))
	pessoa.DataModificacao = time.Date(2000, 2, 1, 12, 30, 0, 0, new(time.Location))

	return
}

func TestMakePessoa(t *testing.T) {
	p1, err := getPessoaTest()
	if err != nil {
		t.Error(err, p1)
	}

	if p1.String() != "12345678910	Teste 01	teste01	123456	teste01@email.com	01/02/2000 12:30:00	01/02/2000 12:30:00	ativo" {
		t.Error("Erro na método String de Pessoa")
	}

	if p1.Repr() != "12345678910	Teste 01	teste01	123456	teste01@email.com	2000-02-01 12:30:00 +0000 UTC	2000-02-01 12:30:00 +0000 UTC	true" {
		t.Error("Erro no método Repr de Pessoa")
	}

	p1, err = NewPessoa("11122233345", "Pedro de Alcântara João Gonzaga", "pedro", "654321", "pedroajoao@gmail.com")
	if err != nil {
		t.Error(err, p1)
	}

	if p1.VerificaAtributos() != nil {
		t.Error(err, p1)
	}
}

func TestVerificaAtributosPessoa(t *testing.T) {
	p1, err := getPessoaTest()

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

	p1, err = getPessoaTest()

	nome := "Pedro de Alcântara João Carlos Leopoldo Salvador Bibiano Francisco Xavier de Paula Leocádio Miguel Gabriel Rafael Gonzaga"
	p1.NomeCompleto = nome
	msgTamanhoIncorreto = erro.ErroTamanho(MsgErroNome01, len(nome)).Error()
	if p1.VerificaAtributos().Error() != msgTamanhoIncorreto {
		t.Error(err, p1)
	}

	p1, err = getPessoaTest()

	senha := `e_-'zWK6$y6Z7!cB>]2d!c7p;]d\bKFLku"ejf*+]g\sE=yjNuKD5Z.~p)#"A=C@[(V,^$`
	p1.Senha = senha
	msgTamanhoIncorreto = erro.ErroTamanho(MsgErroSenha01, len(senha)).Error()
	if p1.VerificaAtributos().Error() != msgTamanhoIncorreto {
		t.Error(err, p1)
	}

	p1, err = getPessoaTest()

	email := "pedro_alcantara_joao_carlos_leopoldo_salvador_rafael_gonzaga@gmail.com"
	p1.Email = email
	msgTamanhoIncorreto = erro.ErroTamanho(MsgErroEmail01, len(email)).Error()
	if p1.VerificaAtributos().Error() != msgTamanhoIncorreto {
		t.Error(err, p1)
	}

	p1, err = getPessoaTest()

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
}

func TestAlteraEstadoPessoa(t *testing.T) {
	p1, _ := getPessoaTest()

	p1.Inativa()
	if p1.Estado != false {
		t.Error("Estado inválido, esperando false, obtido ", p1.Estado)
	}

	p1.Ativa()
	if p1.Estado != true {
		t.Error("Estado inválido, esperando true, obtido ", p1.Estado)
	}
}
