package lancamento

import "time"

type ILancamento interface {
	String() string
	Repr() string
	VerificaAtributos() error
	Altera(string, time.Time, string, string) error
	AlteraCampos(map[string]string) error
	Ativa()
	Inativa()
}

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

const (
	MaxCPFPessoa = 11
	MaxNumero    = 19
	MaxDescricao = 100
)

type ILancamentos interface {
	Len() int
	ProcuraLancamento(int) *Lancamento
}

type Lancamentos []*Lancamento
