package lancamento

import (
	"controle_pessoal_de_financas/API/v1/model/conta"
	"controle_pessoal_de_financas/API/v1/model/pessoa"
	"time"
)

// ILancamento é uma interface que exige a implementação dos métodos obrigatórios em Lancamento
type ILancamento interface {
	String() string
	Repr() string
	VerificaAtributos() error
	Altera(string, time.Time, string, string) error
	AlteraCampos(map[string]string) error
	Ativa()
	Inativa()
}

// Lancamento é uma struct que representa um lancamento. Possui notação JSON para cada campo
type Lancamento struct {
	ID              int       `json:"id"`
	CpfPessoa       string    `json:"cpf_pessoa"`
	Data            time.Time `json:"data"`
	Numero          string    `json:"numero"`
	Descricao       string    `json:"descricao"`
	DataCriacao     time.Time `json:"data_criacao"`
	DataModificacao time.Time `json:"data_modificacao"`
	Estado          bool      `json:"estado"`
}

// MaxCPFPessoa: tamanho máximo para o campo CpfPessoa, baseado em modelo em model.pessoa
// MaxNumero: tamamho máximo para o campo Numero, baseado em modelo em model.conta
// MaxDescricao: tamanho máximo para o campo Descricao
const (
	MaxCPFPessoa = pessoa.LenCpf
	MaxNumero    = conta.MaxCodigo
	MaxDescricao = 100
)

// ILancamentos é uma interface que exige a implementação obrigatória em conjunta/lista(slice) de Lancamentos
type ILancamentos interface {
	Len() int
	ProcuraLancamento(int) *Lancamento
}

// Lancamentos representa um conjunto/lista(slice) de Lancamentos(*Lancamento)
type Lancamentos []*Lancamento
