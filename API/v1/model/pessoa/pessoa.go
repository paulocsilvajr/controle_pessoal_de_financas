package pessoa

import (
	"controle_pessoal_de_financas/API/v1/helper"
	"controle_pessoal_de_financas/API/v1/model/erro"
	"fmt"
	"regexp"
	"time"
)

type Pessoa struct {
	Cpf             string    `json:"cpf"`
	NomeCompleto    string    `json:"nome_completo"`
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

type Pessoas []*Pessoa

func NewPessoa(cpf, nome, usuario, senha, email string) (pessoa *Pessoa, err error) {
	pessoa = &Pessoa{
		Cpf:             cpf,
		NomeCompleto:    nome,
		Usuario:         usuario,
		Senha:           senha,
		Email:           email,
		DataCriacao:     time.Now().Local(),
		DataModificacao: time.Now().Local(),
		Estado:          true}

	if err = pessoa.VerificaAtributos(); err != nil {
		pessoa = nil
	}

	return
}

func (p *Pessoa) Altera(cpf, nome, usuario, senha, email string) (err error) {
	if err = verifica(cpf, nome, usuario, senha, email); err != nil {
		return
	}

	p.Cpf = cpf
	p.NomeCompleto = nome
	p.Usuario = usuario
	p.Senha = senha
	p.Email = email
	p.DataModificacao = time.Now().Local()

	return
}

func (p *Pessoa) alteraEstado(estado bool) {
	p.Estado = estado
}

func (p *Pessoa) Ativa() {
	p.alteraEstado(true)
}

func (p *Pessoa) Inativa() {
	p.alteraEstado(false)
}

func (p *Pessoa) String() string {
	estado := "ativo"
	if !p.Estado {
		estado = "inativo"
	}

	dataCriacao := helper.DataFormatada(p.DataCriacao)
	dataModificacao := helper.DataFormatada(p.DataModificacao)

	return fmt.Sprintf("%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s", p.Cpf, p.NomeCompleto, p.Usuario, p.Senha, p.Email, dataCriacao, dataModificacao, estado)
}

func (p *Pessoa) Repr() string {
	return fmt.Sprintf("%s\t%s\t%s\t%s\t%s\t%s\t%s\t%v", p.Cpf, p.NomeCompleto, p.Usuario, p.Senha, p.Email, p.DataCriacao, p.DataModificacao, p.Estado)
}

func (p *Pessoa) VerificaAtributos() (err error) {
	return verifica(p.Cpf, p.NomeCompleto, p.Usuario, p.Senha, p.Email)
}

func verifica(cpf, nome, usuario, senha, email string) (err error) {
	if err = verificaCPF(cpf); err != nil {
		return
	} else if err = verificaNome(nome); err != nil {
		return
	} else if err = verificaUsuario(usuario); err != nil {
		return
	} else if err = verificaSenha(senha); err != nil {
		return
	} else if err = verificaEmail(email); err != nil {
		return
	}

	return
}

func verificaCPF(cpf string) (err error) {
	padrao, _ := regexp.Compile("[0-9]{11}")

	if len(cpf) != LenCpf {
		err = erro.ErroTamanho(MsgErroCpf01, len(cpf))
	} else if len(cpf) == 0 {
		err = erro.ErroTamanho(MsgErroCpf01, len(cpf))
	} else if !padrao.MatchString(cpf) {
		err = erro.ErroDetalhe(MsgErroCpf02, cpf)
	}

	return
}

func verificaNome(nome string) (err error) {
	if len(nome) > MaxNome {
		err = erro.ErroTamanho(MsgErroNome01, len(nome))
	}

	return
}

func verificaUsuario(usuario string) (err error) {
	if len(usuario) > MaxUsuario {
		err = erro.ErroTamanho(MsgErroUsuario01, len(usuario))
	} else if len(usuario) == 0 {
		err = erro.ErroTamanho(MsgErroUsuario01, len(usuario))
	}

	return
}

func verificaSenha(senha string) (err error) {
	if len(senha) > MaxSenha {
		err = erro.ErroTamanho(MsgErroSenha01, len(senha))
	} else if len(senha) == 0 {
		err = erro.ErroTamanho(MsgErroSenha01, len(senha))
	}

	return
}

func verificaEmail(email string) (err error) {
	if len(email) > MaxEmail {
		err = erro.ErroTamanho(MsgErroEmail01, len(email))
	} else if len(email) == 0 {
		err = erro.ErroTamanho(MsgErroEmail01, len(email))
	}

	return
}
