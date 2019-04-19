package pessoa

import (
	"controle_pessoal_de_financas/API/v1/helper"
	"controle_pessoal_de_financas/API/v1/model/erro"
	"errors"
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
	Administrador   bool      `json:administrador`
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
	MsgErroUsuario02 = "Usuário deve ser formado somente por letras, números ou _(underline)"
	MsgErroSenha01   = "Senha com tamanho inválido"
	MsgErroEmail01   = "Email com tamanho inválido"
)

type Pessoas []*Pessoa

type PessoaI interface {
	GetEmail() string
}

type PessoasI interface {
	ProcuraPessoaPorUsuario(string) (PessoaI, error)
	Len() int
}

func NewPessoa(cpf, nome, usuario, senha, email string) (*Pessoa, error) {
	return newPessoa(cpf, nome, usuario, senha, email, false)
}

func NewPessoaAdmin(cpf, nome, usuario, senha, email string) (*Pessoa, error) {
	return newPessoa(cpf, nome, usuario, senha, email, true)
}

func (p *Pessoa) GetEmail() string {
	return p.Email
}

func (p *Pessoa) Altera(cpf, nome, usuario, senha, email string) (err error) {
	if err = verifica(cpf, nome, usuario, senha, email); err != nil {
		return
	}

	p.Cpf = cpf
	p.NomeCompleto = nome
	p.Usuario = usuario
	if p.Senha != senha {
		p.Senha = helper.GetSenhaSha256(senha)
	}
	p.Email = email
	p.DataModificacao = time.Now().Local()

	return
}

func (p *Pessoa) Ativa() {
	p.alteraEstado(true)
}

func (p *Pessoa) Inativa() {
	p.alteraEstado(false)
}

func (p *Pessoa) SetAdmin(admin bool) {
	p.DataModificacao = time.Now().Local()
	p.Administrador = admin
}

func (p *Pessoa) AlteraCampos(campos map[string]string) (err error) {
	for chave, valor := range campos {
		switch chave {
		case "cpf":
			if err = verificaCPF(valor); err != nil {
				return
			}
			p.Cpf = valor
		case "nome":
			if err = verificaNome(valor); err != nil {
				return
			}
			p.NomeCompleto = valor
		case "usuario":
			if err = verificaUsuario(valor); err != nil {
				return
			}
			p.Usuario = valor
		case "senha":
			if err = verificaSenha(valor); err != nil {
				return
			}
			p.Senha = helper.GetSenhaSha256(valor)
		case "email":
			if err = verificaEmail(valor); err != nil {
				return
			}
			p.Email = valor
		}
	}
	p.DataModificacao = time.Now().Local()

	return
}

func (p *Pessoa) String() string {
	estado := "ativo"
	tipo := "Administrador"
	if !p.Estado {
		estado = "inativo"
	}

	if !p.Administrador {
		tipo = "Comum"
	}

	dataCriacao := helper.DataFormatada(p.DataCriacao)
	dataModificacao := helper.DataFormatada(p.DataModificacao)

	return fmt.Sprintf("%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s", p.Cpf, p.NomeCompleto, p.Usuario, p.Senha, p.Email, dataCriacao, dataModificacao, estado, tipo)
}

func (p *Pessoa) Repr() string {
	return fmt.Sprintf("%s\t%s\t%s\t%s\t%s\t%s\t%s\t%v\t%v", p.Cpf, p.NomeCompleto, p.Usuario, p.Senha, p.Email, p.DataCriacao, p.DataModificacao, p.Estado, p.Administrador)
}

func (p *Pessoa) VerificaAtributos() (err error) {
	return verifica(p.Cpf, p.NomeCompleto, p.Usuario, p.Senha, p.Email)
}

// func (ps Pessoas) ProcuraPessoaPorUsuario(usuario string) (p *Pessoa, err error) {
func (ps Pessoas) ProcuraPessoaPorUsuario(usuario string) (p PessoaI, err error) {
	for _, pessoaLista := range ps {
		if pessoaLista.Usuario == usuario {
			p = pessoaLista
			return
		}
	}

	err = errors.New(fmt.Sprintf(
		"Pessoa com usuário %s informado não existe na listagem", usuario))

	return
}

func (ps Pessoas) Len() int {
	return len(ps)
}

func GetPessoaTest() (pessoa *Pessoa, err error) {
	pessoa, err = NewPessoa("12345678910", "Teste 01", "teste01", "123456", "teste01@email.com")
	pessoa.DataCriacao = time.Date(2000, 2, 1, 12, 30, 0, 0, new(time.Location))
	pessoa.DataModificacao = time.Date(2000, 2, 1, 12, 30, 0, 0, new(time.Location))

	return
}

func newPessoa(cpf, nome, usuario, senha, email string, admin bool) (pessoa *Pessoa, err error) {
	pessoa = &Pessoa{
		Cpf:             cpf,
		NomeCompleto:    nome,
		Usuario:         usuario,
		Senha:           helper.GetSenhaSha256(senha),
		Email:           email,
		DataCriacao:     time.Now().Local(),
		DataModificacao: time.Now().Local(),
		Estado:          true,
		Administrador:   admin}

	if err = pessoa.VerificaAtributos(); err != nil {
		pessoa = nil
	}

	return
}

func (p *Pessoa) alteraEstado(estado bool) {
	p.DataModificacao = time.Now().Local()
	p.Estado = estado
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
	padrao, _ := regexp.Compile("^[a-zA-Z0-9_]*$")
	if len(usuario) > MaxUsuario {
		err = erro.ErroTamanho(MsgErroUsuario01, len(usuario))
	} else if len(usuario) == 0 {
		err = erro.ErroTamanho(MsgErroUsuario01, len(usuario))
	} else if !padrao.MatchString(usuario) {
		err = erro.ErroDetalhe(MsgErroUsuario02, usuario)
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
