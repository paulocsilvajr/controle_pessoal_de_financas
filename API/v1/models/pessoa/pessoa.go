package pessoa

import (
	"controle_pessoal_de_financas/API/v1/models/erros"
	"fmt"
	"regexp"
	"time"
)

type Pessoa struct {
	Cpf             string    `json:"cpf"`
	Nome            string    `json:"nome"`
	Usuario         string    `json:"usuario"`
	Senha           string    `json:"senha"`
	Email           string    `json:"email"`
	DataCriacao     time.Time `json:"data_criacao"`
	DataModificacao time.Time `json:"data_modificacao"`
	Estado          bool      `json:"estado"`
}

const (
	LenCpf           = 11
	MaxNome          = 100
	MaxUsuario       = 20
	MaxSenha         = 64
	MaxEmail         = 45
	MsgErroCpf01     = "CPF inválido, tamanho incorreto"
	MsgErroCpf02     = "CPF deve ser formado somente por números"
	MsgErroNome01    = "Nome com tamanho inválido"
	MsgErroUsuario01 = "Usuário com tamanho inválido"
	MsgErroSenha01   = "Senha com tamanho inválido"
	MsgErroEmail01   = "Email com tamanho inválido"
)

type Pessoas []Pessoa

func NewPessoa(cpf, nome, usuario, senha, email string) (pessoa *Pessoa, erro error) {
	if erro = verificaCPF(cpf); erro != nil {
		return
	} else if erro = verificaNome(nome); erro != nil {
		return
	} else if erro = verificaUsuario(usuario); erro != nil {
		return
	} else if erro = verificaSenha(senha); erro != nil {
		return
	} else if erro = verificaEmail(email); erro != nil {
		return
	}

	pessoa = &Pessoa{Cpf: cpf, Nome: nome, Usuario: usuario, Senha: senha, Email: email, DataCriacao: time.Now()}

	return
}

func (p *Pessoa) VerificaAtributos() (err error) {
	if err = verificaCPF(p.Cpf); err != nil {
		return
	} else if err = verificaNome(p.Nome); err != nil {
		return
	} else if err = verificaUsuario(p.Usuario); err != nil {
		return
	} else if err = verificaSenha(p.Senha); err != nil {
		return
	} else if err = verificaEmail(p.Email); err != nil {
		return
	}

	return
}

func verificaCPF(cpf string) (err error) {
	padrao, _ := regexp.Compile("[0-9]{11}")

	if len(cpf) != LenCpf {
		err = erros.ErroTamanho(MsgErroCpf01, len(cpf))
	} else if !padrao.MatchString(cpf) {
		err = erros.ErroDetalhe(MsgErroCpf02, cpf)
	}

	return
}

func verificaNome(nome string) (err error) {
	if len(nome) > MaxNome {
		err = erros.ErroTamanho(MsgErroNome01, len(nome))
	}

	return
}

func verificaUsuario(usuario string) (err error) {
	if len(usuario) > MaxUsuario {
		err = erros.ErroTamanho(MsgErroUsuario01, len(usuario))
	}

	return
}

func verificaSenha(senha string) (err error) {
	fmt.Println("verificando senha")
	if len(senha) > MaxSenha {
		err = erros.ErroTamanho(MsgErroSenha01, len(senha))
	}

	return
}

func verificaEmail(email string) (err error) {
	if len(email) > MaxEmail {
		err = erros.ErroTamanho(MsgErroEmail01, len(email))
	}

	return
}
