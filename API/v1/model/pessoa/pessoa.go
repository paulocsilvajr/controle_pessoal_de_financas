package pessoa

import (
	"fmt"
	"regexp"
	"time"

	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/helper"
	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/model/erro"
)

// Pessoa é uma struct que representa uma pessoa. Possui notações para JSON para cada campo
type Pessoa struct {
	Cpf             string    `json:"cpf" gorm:"primaryKey;size:11;not null"`
	NomeCompleto    string    `json:"nome_completo" gorm:"size:100"`
	Usuario         string    `json:"usuario" gorm:"size:20;not null;unique"`
	Senha           string    `json:"senha" gorm:"size:64;not null"`
	Email           string    `json:"email" gorm:"size:45;not null;unique"`
	DataCriacao     time.Time `json:"data_criacao" gorm:"not null;autoCreateTime"`
	DataModificacao time.Time `json:"data_modificacao" gorm:"not null;autoUpdateTime"`
	Estado          bool      `json:"estado" gorm:"not null;default:true"`
	Administrador   bool      `json:"administrador" gorm:"not null;default:false"`
}

// TableName define o nome da tabela ao efetuar o AutoMigrate do GORM
func (Pessoa) TableName() string {
	return "pessoa"
}

// LenCpf: tamanho obrigatório do CPF;
// MaxNome: tamanho máximo do Nome
// MaxUsuario: tamanho máximo do Usuário
// MaxSenha: tamanho máximo do Senha
// MaxEmail: tamanho máximo do Email
// MsgErroCpf01: mensagem erro padrão 01 para CPF
// MsgErroCpf02: mensagem erro padrão 02 para CPF
// MsgErroNome01: mensagem erro padrão 01 para Nome
// MsgErroUsuario01: mensagem erro padrão 01 para Usuario
// MsgErroUsuario02: mensagem erro padrão 02 para Usuario
// MsgErroSenha01: mensagem erro padrão 01 para Senha
// MsgErroEmail01: mensagem erro padrão 01 para Email
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

// Pessoas representa um conjunto/lista(slice) de pessoas(*Pessoa)
type Pessoas []*Pessoa

// IPessoa é uma interface que exibe o método GetMail para representar uma pessoa genérica(Pessoa e PessoaSimples)
type IPessoa interface {
	GetEmail() string
	CorrigeData()
}

// IPessoas é uma interface que exige a implementação dos métodos ProcuraPessoaPorUsuario e Len para representar um conjunto/lista(slice) de pessoas genéricas(Pessoas e PessoasSimples)
type IPessoas interface {
	ProcuraPessoaPorUsuario(string) (IPessoa, error)
	Len() int
}

// New retorna uma nova Pessoa(*Pessoa) comum através dos parâmetros informados(cpf, nome, usuario, senha e email). Função equivalente a criação de pessoa via literal &Pessoa{Cpf: ..., NomeCompleto: ..., ...}. Data de criação e modificação são definidos como o horário atual e o estado é definido como ativo. OBS: senha NÃO é hasheada e parâmetros NÃO são validados
func New(cpf, nome, usuario, senha, email string) *Pessoa {
	return &Pessoa{
		Cpf:             cpf,
		NomeCompleto:    nome,
		Usuario:         usuario,
		Senha:           senha,
		Email:           email,
		DataCriacao:     time.Now().Local(),
		DataModificacao: time.Now().Local(),
		Estado:          true,
		Administrador:   false}
}

// NewPessoa retorna uma nova Pessoa(*Pessoa) comum através dos parâmetros informados(cpf, nome, usuario, senha e email). São verificados os parâmetros se são válidos e a senha é hasheada. Data de criação e modificação são definidos como o horário atual e o estado é definido como ativo
func NewPessoa(cpf, nome, usuario, senha, email string) (*Pessoa, error) {
	return newPessoaGeral(cpf, nome, usuario, senha, email, false)
}

// NewPessoaAdmin retorna uma nova Pessoa(*Pessoa) ADMINISTRADORA através dos parâmetros informados(cpf, nome, usuario, senha e email). São verificados os parâmetros se são válidos e a senha é hasheada.  Data de criação e modificação são definidos como o horário atual e o estado é definido como ativo
func NewPessoaAdmin(cpf, nome, usuario, senha, email string) (*Pessoa, error) {
	return newPessoaGeral(cpf, nome, usuario, senha, email, true)
}

// GetEmail é um método que retorna uma string com o email atribuído a pessoa. A interface IPessoa exige a implementação desse método
func (p *Pessoa) GetEmail() string {
	return p.Email
}

// CorrigeData é um método que converte a data(Time) no formato do timezone local
func (p *Pessoa) CorrigeData() {
	p.DataCriacao = p.DataCriacao.Local()
	p.DataModificacao = p.DataModificacao.Local()
}

// Altera é um método que modifica os dados da pessoa a partir dos parâmetros informados depois da verificação de cada parâmetro e atualiza a data de modificação dela. Retorna um erro != nil, caso algum parâmetro seja inválido
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

// Ativa é um método que define a pessoa como ativa e atualiza a data de modificação dela
func (p *Pessoa) Ativa() {
	p.alteraEstado(true)
}

// Inativa é um métido que define a pessoa como inativa  e atualiza a data de modificação dela
func (p *Pessoa) Inativa() {
	p.alteraEstado(false)
}

// SetAdmin é um método que define como administrador uma pessoa de acordo com o parâmetro boleano informado em admin. É atualizado a data de modificação.
func (p *Pessoa) SetAdmin(admin bool) {
	p.DataModificacao = time.Now().Local()
	p.Administrador = admin
}

// AlteraCampos é um método para alterar os campos de uma pesssoa a partir do hashMap informado no parâmetro campos. Somente as chaves informadas com um valor correto serão atualizados. É atualizado a data de modificação dessa pessoa. Caso ocorra um problema na validação dos campos, retorna um erro != nil. Campos permitidos: cpf, nome, usuario, senha, email
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

// String é um método de pessoa que retorna uma string representando uma pessoa. A data é formatada usando a função helper.DataFormatada, os campos estado e administrador são formatados para uma forma mais amigável/legível
func (p *Pessoa) String() string {
	estado := helper.GetEstado(p.Estado)

	tipo := "Administrador"
	if !p.Administrador {
		tipo = "Comum"
	}

	dataCriacao := helper.DataFormatada(p.DataCriacao)
	dataModificacao := helper.DataFormatada(p.DataModificacao)

	return fmt.Sprintf("%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s", p.Cpf, p.NomeCompleto, p.Usuario, p.Senha, p.Email, dataCriacao, dataModificacao, estado, tipo)
}

// Repr é um método que retorna uma string da representação de uma pessoa, sem formatações especiais
func (p *Pessoa) Repr() string {
	return fmt.Sprintf("%s\t%s\t%s\t%s\t%s\t%s\t%s\t%v\t%v", p.Cpf, p.NomeCompleto, p.Usuario, p.Senha, p.Email, p.DataCriacao, p.DataModificacao, p.Estado, p.Administrador)
}

// VerificaAtributos é um método de pessoa que verifica os campos Cpf, NomeCompleto, Usuario, Senha e Email, retornando um erro != nil caso ocorra um problema
func (p *Pessoa) VerificaAtributos() error {
	return verifica(p.Cpf, p.NomeCompleto, p.Usuario, p.Senha, p.Email)
}

// ProcuraPessoaPorUsuario é um método que retorna uma pessoa a partir da busca em uma listagem de pessoas(Pessoas). Caso não seja encontrado a pessoa, retorna um erro != nil. A interface IPessoas exige a implementação desse método
func (ps Pessoas) ProcuraPessoaPorUsuario(usuario string) (p IPessoa, err error) {
	for _, pessoaLista := range ps {
		if pessoaLista.Usuario == usuario {
			p = pessoaLista
			return
		}
	}

	err = fmt.Errorf(
		"Pessoa com usuário %s informado não existe na listagem", usuario)

	return
}

// Len é um método de Pessoas que retorna a quantidade de elementos contidos dentro do slice de pessoas. A interface IPessoas exibe a implementação desse método
func (ps Pessoas) Len() int {
	return len(ps)
}

// GetPessoaTest retorna uma pessoa teste para efetuar testes na API
func GetPessoaTest() (pessoa *Pessoa, err error) {
	pessoa, err = NewPessoa("12345678910", "Teste 01", "teste01", "123456", "teste01@email.com")
	pessoa.DataCriacao = time.Date(2000, 2, 1, 12, 30, 0, 0, new(time.Location))
	pessoa.DataModificacao = time.Date(2000, 2, 1, 12, 30, 0, 0, new(time.Location))

	return
}

func newPessoaGeral(cpf, nome, usuario, senha, email string, admin bool) (pessoa *Pessoa, err error) {
	pessoa = New(cpf, nome, usuario, senha, email)
	pessoa.Administrador = admin

	if err = pessoa.VerificaAtributos(); err != nil {
		pessoa = nil
	} else {
		pessoa.Senha = helper.GetSenhaSha256(senha)
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
	basePadrao := fmt.Sprintf("[0-9]{%d}", LenCpf)
	padrao, _ := regexp.Compile(basePadrao)

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
