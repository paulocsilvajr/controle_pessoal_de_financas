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

func converteParaConta(cb IConta) (*Conta, error) {
	c, ok := cb.(*Conta)
	if ok {
		return c, nil
	}
	return nil, errors.New("Erro ao converter para Conta")
}

// String retorna uma representação textual de uma Conta. Datas são formatadas usando a função helper.DataFormatada() e campo estado é formatado usando a função helper.GetEstado()
func (c *Conta) String() string {
	estado := helper.GetEstado(c.Estado)
	dataCriacao := helper.DataFormatada(c.DataCriacao)
	dataModificacao := helper.DataFormatada(c.DataModificacao)

	return fmt.Sprintf("%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s", c.Nome, c.NomeTipoConta, c.Codigo, c.ContaPai, c.Comentario, dataCriacao, dataModificacao, estado)
}

func (c *Conta) Repr() string {
	return ""
}

func (c *Conta) VerificaAtributos() error {
	return nil
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
