// pessoa_test.go
package pessoa

import (
	"controle_pessoal_de_financas/API/v1/models/erros"
	"testing"
)

func getPessoaTest() (*Pessoa, error) {
	return NewPessoa("12345678910", "Teste 01", "teste01", "123456", "teste01@email.com")
}

func TestMakePessoa(t *testing.T) {
	p1, err := getPessoaTest()
	if err != nil {
		t.Error(err, p1)
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
	msgTamanhoIncorreto := erros.ErroTamanho(MsgErroCpf01, len(cpf)).Error()
	if p1.VerificaAtributos().Error() != msgTamanhoIncorreto {
		t.Error(err, p1)
	}

	cpf = "123.456.789-10"
	p1.Cpf = cpf
	msgTamanhoIncorreto = erros.ErroTamanho(MsgErroCpf01, len(cpf)).Error()
	if p1.VerificaAtributos().Error() != msgTamanhoIncorreto {
		t.Error(err, p1)
	}

	cpf = "123456789xx"
	p1.Cpf = cpf
	msgFormatoIncorreto := erros.ErroDetalhe(MsgErroCpf02, cpf).Error()
	if p1.VerificaAtributos().Error() != msgFormatoIncorreto {
		t.Error(err, p1)
	}

	p1, err = getPessoaTest()

	nome := "Pedro de Alcântara João Carlos Leopoldo Salvador Bibiano Francisco Xavier de Paula Leocádio Miguel Gabriel Rafael Gonzaga"
	p1.Nome = nome
	msgTamanhoIncorreto = erros.ErroTamanho(MsgErroNome01, len(nome)).Error()
	if p1.VerificaAtributos().Error() != msgTamanhoIncorreto {
		t.Error(err, p1)
	}

	p1, err = getPessoaTest()

	senha := `e_-'zWK6$y6Z7!cB>]2d!c7p;]d\bKFLku"ejf*+]g\sE=yjNuKD5Z.~p)#"A=C@[(V,^$`
	p1.Senha = senha
	msgTamanhoIncorreto = erros.ErroTamanho(MsgErroSenha01, len(senha)).Error()
	if p1.VerificaAtributos().Error() != msgTamanhoIncorreto {
		t.Error(err, p1)
	}

	p1, err = getPessoaTest()

	email := "pedro_alcantara_joao_carlos_leopoldo_salvador_rafael_gonzaga"
	p1.Email = email
	msgTamanhoIncorreto = erros.ErroTamanho(MsgErroEmail01, len(email)).Error()
	if p1.VerificaAtributos().Error() != msgTamanhoIncorreto {
		t.Error(err, p1)
	}
}
