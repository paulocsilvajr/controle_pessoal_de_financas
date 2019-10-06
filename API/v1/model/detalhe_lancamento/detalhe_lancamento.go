package detalhe_lancamento

import (
	"controle_pessoal_de_financas/API/v1/model/conta"
	"errors"
)

// IDetalheLancamento é uma interface que exige a implementação dos métodos origatórios em DetalheLancamento
type IDetalheLancamento interface {
	String() string
	Repr() string
	VerificaAtributos() error
	Altera(string, int64, int64) error
	AlteraCampos(map[string]string) error
	Ativa()
	Inativa()
	DebitoToStr() string
	CreditoToStr() string
}

// DetalheLancamento é uma struct que representa uma detalhe de lançamento. Possui notação para JSON para cada campo
type DetalheLancamento struct {
	ID        int     `json:"id"`
	NomeConta string  `json:"nome_conta"`
	Debito    float32 `json:"debito"`
	Credito   float32 `json:"credito"`
}

// MaxNomeConta: tamanho máximo para o campo NomeConta, baseado em valor informado em modelo.conta
const (
	MaxNomeConta = conta.MaxNome
)

// IDetalheLancamentos é uma interface que exige a implementação obrigatória em conjunto/lista(slice) de DetalheLancamentos
type IDetalheLancamentos interface {
	Len()
	ProcuraDetalheLancamento(int) *DetalheLancamento
}

// DetalheLancamentos representa um conjunto/lista(slice) de DetalheLancamentos(*DetalheLancamento)
type DetalheLancamentos []*DetalheLancamento

// converteParaDetalheLancamento converte uma interface IDetalheLancamento em um tipo DetalheLancamento, se possível. Caso contrário, retornal nil para DetalheLancamento e um erro
func converteParaDetalheLancamento(dlb IDetalheLancamento) (*DetalheLancamento, error) {
	dl, ok := dlb.(*DetalheLancamento)
	if ok {
		return dl, nil
	}
	return nil, errors.New("Erro ao converter para DetalheLancamento")
}

func New(id int, nomeConta string, debito, credito float32) *DetalheLancamento {
	return &DetalheLancamento{
		ID:        id,
		NomeConta: nomeConta,
		Debito:    debito,
		Credito:   credito}
}

func NewDetalheLancamento(id int, nomeConta string, debito, credito float32) (detalheLancamento *DetalheLancamento, err error) {
	detalheLancamento = New(id, nomeConta, debito, credito)

	if err = detalheLancamento.VerificaAtributos(); err != nil {
		detalheLancamento = nil
	}

	return
}

func GetDetalheLancamentoTest() (detalheLancamento *DetalheLancamento) {
	detalheLancamento = New(9999, "Detalhe de conta teste", 100, 0)

	return
}

func (dl *DetalheLancamento) String() string {
	return ""
}
