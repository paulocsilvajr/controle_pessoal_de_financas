package conta

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/helper"
)

// IConta é uma interface que exige a implementação dos métodos obrigatórios em Conta
type IConta interface {
	String() string
	Repr() string
	VerificaAtributos() error
	Altera(string, string, string, string, string) error
	AlteraCampos(map[string]string) error
	Ativa()
	Inativa()
	CorrigeData()
}

// Conta é uma struct que representa uma conta. Possui notação para JSON para cada campo
type Conta struct {
	Nome            string    `json:"nome"`
	NomeTipoConta   string    `json:"nome_tipo_conta"`
	Codigo          string    `json:"codigo"`
	ContaPai        string    `json:"conta_pai"`
	Comentario      string    `json:"comentario"`
	DataCriacao     time.Time `json:"data_criacao"`
	DataModificacao time.Time `json:"data_modificacao"`
	Estado          bool      `json:"estado"`
}

type TConta struct {
	Nome            string         `gorm:"primaryKey;size:50;not null"`
	NomeTipoConta   string         `gorm:"size:50;not null"`
	Codigo          sql.NullString `gorm:"size:19;unique"`
	ContaPai        sql.NullString `gorm:"size:50"`
	Comentario      sql.NullString `gorm:"size:150"`
	DataCriacao     time.Time      `gorm:"not null;autoCreateTime"`
	DataModificacao time.Time      `gorm:"not null;autoUpdateTime"`
	Estado          bool           `gorm:"not null;default:true"`
}

func (TConta) TableName() string {
	return "conta"
}

func GetNomeTabelaConta() string {
	return new(TConta).TableName()
}

// MaxConta: tamanho máximo para os campos de texto(string) Nome, NomeTipoConta e ContaPai
// MaxCodigo: tamanho máximo para o Código
// MaxComentario: tamanho máximo para o Comentário
const (
	MaxNome       = 50
	MaxCodigo     = 19
	MaxComentario = 150
)

// IContas é uma interface que exige a implementação obrigatórios em conjunto/lista(slice) de Contas
type IContas interface {
	Len() int
	ProcuraConta(string) *Conta
}

// Contas representa um conjunto/lista(slice) de Contas(*Conta)
type Contas []*Conta

// converteParaConta converte uma interface IConta em um tipo Conta, se possível. Caso contrário, retorna nil para conta e um erro
func converteParaConta(cb IConta) (*Conta, error) {
	c, ok := cb.(*Conta)
	if ok {
		return c, nil
	}
	return nil, errors.New("Erro ao converter para Conta")
}

// New retorna uma nova Conta(*Conta) através dos parâmetros informados(nome, nomeTipoConta, codigo, contaPai, comentario). Função equivalente a criação de uma Conta via literal &Conta{Nome: ..., ...}. Data de criação e modificação são definidos com o horário atual e o estado é definido como ativo
func New(nome, nomeTipoConta, codigo, contaPai, comentario string) *Conta {
	return &Conta{
		Nome:            nome,
		NomeTipoConta:   nomeTipoConta,
		Codigo:          codigo,
		ContaPai:        contaPai,
		Comentario:      comentario,
		DataCriacao:     time.Now().Local(),
		DataModificacao: time.Now().Local(),
		Estado:          true}
}

// NewConta cria uma nova Conta semelhante a função New(), mas faz a validação dos campos informados nos parâmetros nome, nomeTipoConta, codigo, contaPai, comentario
func NewConta(nome, nomeTipoConta, codigo, contaPai, comentario string) (conta *Conta, err error) {
	conta = New(nome, nomeTipoConta, codigo, contaPai, comentario)

	if err = conta.VerificaAtributos(); err != nil {
		conta = nil
	}

	return
}

// GetContaTest retorna uma Conta teste usado para os testes em geral
func GetContaTest() (conta *Conta) {
	conta = New("Ativos Teste 01", "banco teste 01", "001", "", "teste de Conta")
	conta.DataCriacao = time.Date(2000, 2, 1, 12, 30, 0, 0, new(time.Location))
	conta.DataModificacao = time.Date(2000, 2, 1, 12, 30, 0, 0, new(time.Location))
	conta.Estado = true

	return
}

// String retorna uma representação textual de uma Conta. Datas são formatadas usando a função helper.DataFormatada() e campo estado é formatado usando a função helper.GetEstado()
func (c *Conta) String() string {
	estado := helper.GetEstado(c.Estado)
	dataCriacao := helper.DataFormatada(c.DataCriacao)
	dataModificacao := helper.DataFormatada(c.DataModificacao)

	return fmt.Sprintf("%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s", c.Nome, c.NomeTipoConta, c.Codigo, c.ContaPai, c.Comentario, dataCriacao, dataModificacao, estado)
}

// Repr é um método que retorna uma string de representação de uma Conta, sem formatações especiais
func (c *Conta) Repr() string {
	return fmt.Sprintf("%s\t%s\t%s\t%s\t%s\t%s\t%s\t%v", c.Nome, c.NomeTipoConta, c.Codigo, c.ContaPai, c.Comentario, c.DataCriacao, c.DataModificacao, c.Estado)
}

// VerificaAtributos é um método de Conta que verifica os campos Nome, NomeTipoConta, Codigo, ContaPai, Comentario, retornando um erro != nil caso ocorra um problema
func (c *Conta) VerificaAtributos() error {
	return verifica(c.Nome, c.NomeTipoConta, c.Codigo, c.ContaPai, c.Comentario)
}

// Altera é um método que modifica os dados da Conta a partir dos parâmetros informados depois da verificação de cada parâmetro e atualiza a data de modificação dela. Retorna um erro != nil, caso algum parâmetro seja inválido
func (c *Conta) Altera(nome, nomeTipoConta, codigo, contaPai, comentario string) (err error) {
	if err = verifica(nome, nomeTipoConta, codigo, contaPai, comentario); err != nil {
		return
	}

	c.Nome = nome
	c.NomeTipoConta = nomeTipoConta
	c.Codigo = codigo
	c.ContaPai = contaPai
	c.Comentario = comentario
	c.DataModificacao = time.Now().Local()

	return
}

// AlteraCampos é um método para alterar os campos de uma Conta a partir de hashMap informado no parâmetro campos. Somente as chaves informadas com um valor correto serão atualizadas. É atualizado a data de modificação da Conta. Caso ocorra um problema na validação dos campos, retorna um erro != nil. Campos permitidos: nome, nomeTipoConta, codigo, contaPai, comentario
func (c *Conta) AlteraCampos(campos map[string]string) (err error) {
	for chave, valor := range campos {
		switch chave {
		case "nome":
			if err = helper.VerificaCampoTexto("Nome", valor, MaxNome); err != nil {
				return
			}
			c.Nome = valor
		case "nomeTipoConta":
			if err = helper.VerificaCampoTexto("Nome do Tipo de Conta", valor, MaxNome); err != nil {
				return
			}
			c.NomeTipoConta = valor
		case "codigo":
			if err = helper.VerificaCampoTextoOpcional("Código", valor, MaxCodigo); err != nil {
				return
			}
			c.Codigo = valor
		case "contaPai":
			if err = helper.VerificaCampoTextoOpcional("Nome da Conta Pai", valor, MaxNome); err != nil {
				return
			}
			c.ContaPai = valor
		case "comentario":
			if err = helper.VerificaCampoTexto("Comentário", valor, MaxComentario); err != nil {
				return
			}
			c.ContaPai = valor
		}
	}
	c.DataModificacao = time.Now().Local()

	return
}

// Ativa é um método que define a Conta como ativa e atualiza a sua data de modificação
func (c *Conta) Ativa() {
	c.alteraEstado(true)
}

// Inativa é um método que define a Conta como inativa e atualiza a sua data de modificação
func (c *Conta) Inativa() {
	c.alteraEstado(false)
}

// CorrigeData é um método que converte a data(Time) no formato do timezone local
func (c *Conta) CorrigeData() {
	c.DataCriacao = c.DataCriacao.Local()
	c.DataModificacao = c.DataModificacao.Local()
}

// ProcuraConta é um método que retorna uma Conta a partir da busca em uma listagem de Contas. Caso não seja encontrado a Conta, retorna um erro != nil. A interface IContas exige a implementação desse método
func (cs Contas) ProcuraConta(nomeConta string) (c *Conta, err error) {
	for _, contaLista := range cs {
		if contaLista.Nome == nomeConta {
			c = contaLista
			return
		}
	}

	err = fmt.Errorf("Conta %s informada não existe na listagem", nomeConta)

	return
}

// Len é um método de Contas que retorna a quantidade de elementos contidos dentro do slice de Conta. A interface IContas exige a implementação desse método
func (cs Contas) Len() int {
	return len(cs)
}

func verifica(nome, nomeTipoConta, codigo, contaPai, comentario string) (err error) {
	if err = helper.VerificaCampoTexto("Nome", nome, MaxNome); err != nil {
		return
	} else if err = helper.VerificaCampoTexto("Nome do Tipo da Conta", nomeTipoConta, MaxNome); err != nil {
		return
	} else if err = helper.VerificaCampoTextoOpcional("Código", codigo, MaxCodigo); err != nil {
		return
	} else if err = helper.VerificaCampoTextoOpcional("Nome da Conta pai", contaPai, MaxNome); err != nil {
		return
	} else if err = helper.VerificaCampoTextoOpcional("Comentário", comentario, MaxComentario); err != nil {
		return
	}

	return
}

func (c *Conta) alteraEstado(estado bool) {
	c.DataModificacao = time.Now().Local()
	c.Estado = estado
}

// movido funções comentadas para helper.texto_helper.go
// func verificaCampoTexto(nomeCampo, campo string, tamanho int) error {
// 	campoValido := len(campo) > 0 && len(campo) <= tamanho
// 	if campoValido {
// 		return nil
// 	}
// 	return fmt.Errorf("Tamanho de campo %s inválido[%d caracter(es)]", nomeCampo, len(campo))
// }

// func verificaCampoTextoOpcional(nomeCampo, campo string, tamanho int) error {
// 	campoValido := len(campo) >= 0 && len(campo) <= tamanho
// 	if campoValido {
// 		return nil
// 	}
// 	return fmt.Errorf("Tamanho de campo %s inválido[%d caracter(es)]", nomeCampo, len(campo))
// }
