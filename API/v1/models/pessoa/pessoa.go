package pessoa

import (
	"cpf_api_go/v1/models/erros"
	"regexp"
)

type Pessoa struct {
	Cpf     string `json:"cpf"`
	Nome    string `json:"nome"`
	Usuario string `json:"usuario"`
	Senha   string `json:"senha"`
	Email   string `json:"email"`
}

const (
	LenCpf     = 11
	MaxNome    = 100
	MaxUsuario = 20
	MaxSenha   = 64
	MaxEmail   = 45
)

type Pessoas []Pessoa

func New(cpf, nome, usuario, senha, email string) (pessoa Pessoa, erro error) {
	if erro = verificaCPF(cpf); erro == nil {
		if erro = verificaNome(nome); erro == nil {
			if erro = verificaUsuario(usuario); erro == nil {
				if erro = verificaSenha(nome); erro == nil {
					if erro = verificaEmail(nome); erro == nil {
						pessoa = Pessoa{cpf, nome, usuario, senha, email}
					}
				}
			}
		}
	}

	return
}

func verificaCPF(cpf string) (erro error) {
	padrao, _ := regexp.Compile("[0-9]{11}")

	if len(cpf) != LenCpf {
		erro = erros.ErroTamanho("CPF inválido, tamanho incorreto", len(cpf))
	} else if !padrao.MatchString(cpf) {
		erro = erros.ErroDetalhe("CPF deve ser formado somente por números", cpf)
	}

	return
}

func verificaNome(nome string) (erro error) {
	if len(nome) > MaxNome {
		erro = erros.ErroTamanho("Nome com tamanho inválido", len(nome))
	}

	return
}

func verificaUsuario(usuario string) (erro error) {
	if len(usuario) > MaxUsuario {
		erro = erros.ErroTamanho("Usuário com tamanho inválido", len(usuario))
	}

	return
}

func verificaSenha(senha string) (erro error) {
	if len(senha) > MaxSenha {
		erro = erros.ErroTamanho("Senha com tamanho inválido", len(senha))
	}

	return
}

func verificaEmail(email string) (erro error) {
	if len(email) > MaxEmail {
		erro = erros.ErroTamanho("Email com tamanho inválido", len(email))
	}

	return
}
