package lancamento

import (
	"controle_pessoal_de_financas/API/v1/helper"
	"controle_pessoal_de_financas/API/v1/model/conta"
	"controle_pessoal_de_financas/API/v1/model/pessoa"
	"errors"
	"fmt"
	"time"
)

// ILancamento é uma interface que exige a implementação dos métodos obrigatórios em Lancamento
type ILancamento interface {
	String() string
	Repr() string
	VerificaAtributos() error
	Altera(string, time.Time, string, string) error
	AlteraCampos(map[string]interface{}) error
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
	ProcuraLancamentoID(int) *Lancamento
	ProcuraLancamentoCPF(string) Lancamentos
}

// Lancamentos representa um conjunto/lista(slice) de Lancamentos(*Lancamento)
type Lancamentos []*Lancamento

// converterParaLancamento converte um intergface ILancamento em um tipo Lancamento, se possível. Caso contrário retorna nil para lancamento e um erro
func converterParaLancamento(lb ILancamento) (*Lancamento, error) {
	l, ok := lb.(*Lancamento)
	if ok {
		return l, nil
	}
	return nil, errors.New("Erro ao converter para Lancamento")
}

// New retorna um novo Lancamento(*Lancamento) através dos parâmetros informados(id, cpfPessoa, data, numero, descricao). Função equivalente a criação de um Lancamento via literal &Lancamemto{ID: ..., ...}. Data de criação e modificação são definidos com o horário atual e o estado é definido como ativo
func New(id int, cpfPessoa string, data time.Time, numero, descricao string) *Lancamento {
	return &Lancamento{
		ID:              id,
		CpfPessoa:       cpfPessoa,
		Data:            data,
		Numero:          numero,
		Descricao:       descricao,
		DataCriacao:     time.Now().Local(),
		DataModificacao: time.Now().Local(),
		Estado:          true}
}

// NewLancamento cria uma novo Lancamento semelhante a função New(), mas faz a validação dos campos informados nos parâmetros id, cpfPeddoa, data, numero e descricao
func NewLancamento(id int, cpfPessoa string, data time.Time, numero, descricao string) (lancamento *Lancamento, err error) {
	lancamento = New(id, cpfPessoa, data, numero, descricao)

	if err = lancamento.VerificaAtributos(); err != nil {
		lancamento = nil
	}

	return
}

// GetLancamentoTest retorna um Lancamento teste usado para os testes em geral
func GetLancamentoTest() (lancamento *Lancamento) {
	data := time.Date(2000, 2, 1, 12, 30, 0, 0, new(time.Location))
	lancamento = New(9999, "12345678910", data, "1234A", "Pgto conta energia")
	lancamento.DataCriacao = data
	lancamento.DataModificacao = data
	lancamento.Estado = true

	return
}

// String retorna uma representação textual de um Lancamento. Datas são formatadas usando a função helper.DataFormatada() e campo estado é formatado usando a função helper.GetEstado()
func (l *Lancamento) String() string {
	estado := helper.GetEstado(l.Estado)
	data := helper.DataFormatada(l.Data)
	dataCriacao := helper.DataFormatada(l.DataCriacao)
	dataModificacao := helper.DataFormatada(l.DataModificacao)

	return fmt.Sprintf("%d\t%s\t%s\t%s\t%s\t%s\t%s\t%s", l.ID, l.CpfPessoa, data, l.Numero, l.Descricao, dataCriacao, dataModificacao, estado)
}

// Repr é um método que retorna uma string de representação de um Lancamento, sem formatações especiais
func (l *Lancamento) Repr() string {
	return fmt.Sprintf("%d\t%s\t%s\t%s\t%s\t%s\t%s\t%v", l.ID, l.CpfPessoa, l.Data, l.Numero, l.Descricao, l.DataCriacao, l.DataModificacao, l.Estado)
}

// VerificaAtributos é um método de Lancamento que verifica os campos CpfPessoa, Numero e Descricao, retornando um erro != nil caso ocorra um problema
func (l *Lancamento) VerificaAtributos() error {
	return verifica(l.CpfPessoa, l.Numero, l.Descricao)
}

// Altera é um método que modifica os dados de Lancamento a partir dos parâmetros informados depois da verificação de cada parâmetro e atualiza a data de modificação dela. Retorna um erro != nil, caso algum parâmetro seja inválido
func (l *Lancamento) Altera(cpf string, data time.Time, numero, descricao string) (err error) {
	if err = verifica(cpf, numero, descricao); err != nil {
		return
	}

	l.CpfPessoa = cpf
	l.Data = data
	l.Numero = numero
	l.Descricao = descricao
	l.DataModificacao = time.Now().Local()

	return
}

// AlteraCampos é um método para alterar os campos de um Lancamento a partir de hashMap informado no parâmetro campos. Somente as chaves informadas com um valor correto serão atualizadas. É atualizado a data de modificação do Lancamento. Caso ocorra um problema na validação dos campos, retorna um erro != nil. Campos permitidos: cpf, data, numero, descricao
func (l *Lancamento) AlteraCampos(campos map[string]interface{}) (err error) {
	for chave, valor := range campos {
		switch chave {
		case "cpf":
			if cpf, ok := valor.(string); ok {
				if err = helper.VerificaCampoTexto("CPF", cpf, MaxCPFPessoa); err != nil {
					return
				}

				l.CpfPessoa = cpf
			}
		case "data":
			if data, ok := valor.(time.Time); ok {
				l.Data = data
			} else {
				return fmt.Errorf("Data inválida %v[%T]", valor, valor)
			}
		case "numero":
			if numero, ok := valor.(string); ok {
				if err = helper.VerificaCampoTextoOpcional("Número", numero, MaxNumero); err != nil {
					return
				}

				l.Numero = numero
			}
		case "descricao":
			if descricao, ok := valor.(string); ok {
				if err = helper.VerificaCampoTexto("Descrição", descricao, MaxDescricao); err != nil {
					return
				}

				l.Numero = descricao
			}
		}
	}
	l.DataModificacao = time.Now().Local()

	return
}

// Ativa é um método que define o Lancamento como ativo e atualiza a sua data de modificação
func (l *Lancamento) Ativa() {
	l.alteraEstado(true)
}

// Inativa é um método que define o Lancamento como inativo e atualiza a sua data de modificação
func (l *Lancamento) Inativa() {
	l.alteraEstado(false)
}

// ProcuraLancamentoID é um método que retorna um Lancamento a partir da busca em uma listagem de Lancamentos por um id informado. Caso não seja encontrado o Lancamento, retorna um erro != nil. A interface ILancamentos exige a implementação desse método
func (ls Lancamentos) ProcuraLancamentoID(id int) (l *Lancamento, err error) {
	for _, lancamentoLista := range ls {
		if lancamentoLista.ID == id {
			l = lancamentoLista
			return
		}
	}

	err = fmt.Errorf("Lançamento com ID:%d informado não existe na listagem", id)

	return
}

// ProcuraLancamentoCPF é um método que retorna um slice de Lancamento a partir da busca em uma listagem de Lancamentos por um cpf informado. Caso não seja encontrado o Lancamento, retorna um erro != nil. A interface ILancamentos exige a implementação desse método
func (ls Lancamentos) ProcuraLancamentoCPF(cpf string) (l Lancamentos, err error) {
	for _, lancamentoLista := range ls {
		if lancamentoLista.CpfPessoa == cpf {
			l = append(l, lancamentoLista)
		}
	}

	if l.Len() == 0 {
		err = fmt.Errorf("Lançamento(s) com CPF:%s informado não existe na listagem", cpf)
	}

	return
}

// Len é um método de Lancamentos que retorna a quantidade de elementos contidos dentro do slice de Lancamentos. A interface ILancamentos exige a implementação desse método
func (ls Lancamentos) Len() int {
	return len(ls)
}

func verifica(cpf, numero, descricao string) (err error) {
	if err = helper.VerificaCampoTexto("CPF", cpf, MaxCPFPessoa); err != nil {
		return
	} else if err = helper.VerificaCampoTextoOpcional("Número", numero, MaxNumero); err != nil {
		return
	} else if err = helper.VerificaCampoTexto("Descrição", descricao, MaxDescricao); err != nil {
		return
	}

	return
}

func (l *Lancamento) alteraEstado(estado bool) {
	l.DataModificacao = time.Now().Local()
	l.Estado = estado
}
