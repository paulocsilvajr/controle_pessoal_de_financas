package conta

import (
	"controle_pessoal_de_financas/API/v1/helper"
	"errors"
	"fmt"
	"time"
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

// MaxConta: tamanho máximo para os campos de texto(string) Nome, NomeTipoConta e ContaPai
// MaxCodigo: tamanho máximo para o Código
// MaxComentario: tamanho máximo para o Comentário
const (
	MaxConta      = 50
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

// New retorna um nova Conta(*Conta) através dos parâmetros informados(nome, nomeTipoConta, codigo, contaPai, comentario). Função equivalente a criação de uma Conta via literal &Conta{Nome: ..., ...}. Data de criação e modificação são definidos com o horário atual e o estado é definido como ativo
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

// VerificaAtributos é um método de Conta que verifica os compos Nome, NomeTipoConta, Codigo, ContaPai, Comentario, retorna um erro != nil caso ocorra um problema
func (c *Conta) VerificaAtributos() error {
	return verifica(nome, nomeTipoConta, codigo, contaPai, comentario)
}

func (c *Conta) Altera(string, string, string, string, string) error {
	return nil
}

func (c *Conta) AlteraCampos(map[string]string) error {
	return nil
}

func (c *Conta) Ativa() {

}

func (c *Conta) Inativa() {

}
