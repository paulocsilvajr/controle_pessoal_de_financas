package detalhe_lancamento

import (
	"controle_pessoal_de_financas/API/v1/helper"
	"controle_pessoal_de_financas/API/v1/model/conta"
	"errors"
	"fmt"
	"strconv"
)

// IDetalheLancamento é uma interface que exige a implementação dos métodos origatórios em DetalheLancamento
type IDetalheLancamento interface {
	String() string
	Repr() string
	VerificaAtributos() error
	Altera(string, float64, float64) error
	AlteraCampos(map[string]string) error
	DebitoToStr() string
	CreditoToStr() string
}

// DetalheLancamento é uma struct que representa uma detalhe de lançamento.
type DetalheLancamento struct {
	IDLancamento int
	NomeConta    string
	Debito       float64
	Credito      float64
}

// MaxNomeConta: tamanho máximo para o campo NomeConta, baseado em valor informado em modelo.conta
const (
	MaxNomeConta = conta.MaxNome
)

// IDetalheLancamentos é uma interface que exige a implementação obrigatória em conjunto/lista(slice) de DetalheLancamentos
type IDetalheLancamentos interface {
	Len() int
	ProcuraDetalheLancamentos(string) (DetalheLancamentos, error)
	ProcuraDetalheLancamento(int) (*DetalheLancamento, error)
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

// New retorna um novo Detalhe Lançamento(*DetalheLancamento) através dos parâmetros informados(idLancamento, nomeConta, debito, credito). Função equivalente a criação de um DetalheLancamento via literal &DetalheLancamento{IDLancamento: ..., ...}
func New(idLancamento int, nomeConta string, debito, credito float64) *DetalheLancamento {
	return &DetalheLancamento{
		IDLancamento: idLancamento,
		NomeConta:    nomeConta,
		Debito:       debito,
		Credito:      credito}
}

// NewD retorna um novo *DetalheLancamento com o valor de debito informado e credito zerado. Preferencialmente, use esta função para criar DetalheLancamento com valor de debito
func NewD(idLancamento int, nomeConta string, debito float64) *DetalheLancamento {
	return New(idLancamento, nomeConta, debito, 0.0)
}

// NewC retorna um novo *DetalheLancamento com o valor de credito informado e debito zerado. Preferencialmente, use esta função para criar DetalheLancamento com valor de credito
func NewC(idLancamento int, nomeConta string, credito float64) *DetalheLancamento {
	return New(idLancamento, nomeConta, 0.0, credito)
}

// NewDetalheLancamento cria uma novo DetalheLancamento semelhante a função New(), mas faz a validação dos campos informados nos parâmetros nomeConta, debito e credito. Se debito for zerado, cria o novo DetalheLancamento com a função NewC, caso contrário, cria com a função NewD
func NewDetalheLancamento(idLancamento int, nomeConta string, debito, credito float64) (detalheLancamento *DetalheLancamento, err error) {
	if debito == 0.0 {
		detalheLancamento = NewC(idLancamento, nomeConta, credito)
	} else if credito == 0.0 {
		detalheLancamento = NewD(idLancamento, nomeConta, debito)
	} else {
		detalheLancamento = New(idLancamento, nomeConta, debito, credito)
	}

	if err = detalheLancamento.VerificaAtributos(); err != nil {
		detalheLancamento = nil
	}

	return
}

// GetDetalheLancamentoDTest retorna um DetalheLancamento teste usado para os testes em geral com valor de credito zerado
func GetDetalheLancamentoDTest() (detalheLancamento *DetalheLancamento) {
	detalheLancamento = NewD(9999, "Detalhe de conta teste A", 100)

	return
}

// GetDetalheLancamentoCTest retorna um DetalheLancamento teste usado para os testes em geral com valor de debito zerado
func GetDetalheLancamentoCTest() (detalheLancamento *DetalheLancamento) {
	detalheLancamento = NewC(9998, "Detalhe de conta teste B", 200)

	return
}

// String retorna uma representação textual de um DetalheLancamento. Valores de debito e credito são formacadas usando a função helper.MonetarioFormatado()
func (dl *DetalheLancamento) String() string {
	debito := helper.MonetarioFormatado(dl.Debito)
	credito := helper.MonetarioFormatado(dl.Credito)
	return fmt.Sprintf("%d\t%s\t%s\t%s", dl.IDLancamento, dl.NomeConta, debito, credito)
}

// Repr é um método que retorna uma string de representação de um DetalheLancamento, sem formatações especiais
func (dl *DetalheLancamento) Repr() string {
	return fmt.Sprintf("%d\t%s\t%f\t%f", dl.IDLancamento, dl.NomeConta, dl.Debito, dl.Credito)
}

// VerificaAtributos é um método de DetalheLancamento que verifica os campos NomeConta, Debito e Credito, retornando um erro != nil caso ocorra um problema
func (dl *DetalheLancamento) VerificaAtributos() error {
	return verifica(dl.NomeConta, dl.Debito, dl.Credito)
}

// Altera é um método que modifica os dados do DetalheLancamento a partir dos parâmetros informados depois da verificação de cada parâmetro. Retorna um erro != nil, caso algun parâmetro seja inválido
func (dl *DetalheLancamento) Altera(nomeConta string, debito, credito float64) (err error) {
	if err = verifica(nomeConta, debito, credito); err != nil {
		return
	}

	dl.NomeConta = nomeConta
	dl.Debito = debito
	dl.Credito = credito

	return
}

// AlteraCampos é um método para alterar os campos de um DetalheLancamento a partir de hashMap informado no parâmetro campos. Somente as chaves informadas com um valor correto serão atualizadas. Caso ocorra um problema na validação dos campos, retorna um erro != nil. Campos permitidos: nomeConta, debito, credito. Campos debito e credito devem ser informados como string, é feita uma verificação e conversão para atribuir em seus campos equivalente em float64
func (dl *DetalheLancamento) AlteraCampos(campos map[string]string) (err error) {
	for chave, valor := range campos {
		switch chave {
		case "nomeConta":
			if err = helper.VerificaCampoTexto("NomeConta", valor, MaxNomeConta); err != nil {
				return
			}
			dl.NomeConta = valor

		case "debito", "credito":
			decimal, err := strconv.ParseFloat(valor, 64)
			if err != nil {
				return errors.New("Erro ao converter string para float64")
			}

			err = helper.VerificaValor(chave, decimal)
			if err != nil {
				return err
			}

			if chave == "debito" {
				dl.Debito = decimal
			} else if chave == "credito" {
				dl.Credito = decimal
			}
		}

	}

	return
}

// CreditoToStr retorna uma string formatada referente ao atributo Credito no formato definido em helper.MonetarioFormatado
func (dl *DetalheLancamento) CreditoToStr() string {
	return helper.MonetarioFormatado(dl.Credito)
}

// DebitoToStr retorna uma string formatada referente ao atributo Credito no formato definido em helper.MonetarioFormatado
func (dl *DetalheLancamento) DebitoToStr() string {
	return helper.MonetarioFormatado(dl.Debito)
}

// ProcuraDetalheLancamentosPorNomeConta é um método que retorna um slice de *DetalheLancamento a partir da busca em uma listagem de DetalheLancamentos pelo nome de Conta informada. Caso não seja encontrado nenhum DetalheLancamento, retorna um erro != nil. A interface IDetalheLancamentos exige a implementação desse método
func (dls DetalheLancamentos) ProcuraDetalheLancamentosPorNomeConta(nomeConta string) (dlse DetalheLancamentos, err error) {
	for _, detalheLancamentoLista := range dls {
		if detalheLancamentoLista.NomeConta == nomeConta {
			dlse = append(dlse, detalheLancamentoLista)
		}
	}

	if len(dlse) == 0 {
		err = fmt.Errorf("Não foi encontrado nenhum DetalheLancamento com o NomeConta[%s] informado", nomeConta)
	}

	return
}

// ProcuraDetalheLancamentosPorIDLancamento é um método que retorna um slice de *DetalheLancamento a partir da busca em uma listagem de DetalheLancamentos pelo ID de Lancamento informada. Caso não seja encontrado nenhum DetalheLancamento, retorna um erro != nil. A interface IDetalheLancamentos exige a implementação desse método
func (dls DetalheLancamentos) ProcuraDetalheLancamentosPorIDLancamento(idLancamento int) (dlse DetalheLancamentos, err error) {
	for _, detalheLancamentoLista := range dls {
		if detalheLancamentoLista.IDLancamento == idLancamento {
			dlse = append(dlse, detalheLancamentoLista)
		}
	}

	if len(dlse) == 0 {
		err = fmt.Errorf("Não foi encontrado nenhum DetalheLancamento com o IDLancamento[%d] informado", idLancamento)
	}

	return
}

// ProcuraDetalheLancamento é um método que retorna um DetalheLancamento a partir da busca em uma listagem de DetalheLancamentos. Caso não seja encontrado um DetalheLancamento, retorna um erro != nil. A interface IDetalheLancamentos exige a implementação desse método
func (dls DetalheLancamentos) ProcuraDetalheLancamento(idLancamento int, nomeConta string) (dl *DetalheLancamento, err error) {
	for _, detalheLancamentoLista := range dls {
		temID := detalheLancamentoLista.IDLancamento == idLancamento
		temConta := detalheLancamentoLista.NomeConta == nomeConta

		if temID && temConta {
			dl = detalheLancamentoLista
			return
		}
	}

	err = fmt.Errorf("DetalheLancamento [%d:%s] informado não existe na listagem", idLancamento, nomeConta)

	return
}

// Len é um método de DetalheLancamentos que retorna a quantidade de elementos contidos dentro do slice de DetalheLancamento. A interface IDetalheLancamentos exige a implementação desse método
func (dls DetalheLancamentos) Len() int {
	return len(dls)
}

func verifica(nomeConta string, debito, credito float64) (err error) {
	if err = helper.VerificaCampoTexto("NomeConta", nomeConta, MaxNomeConta); err != nil {
		return
	} else if err = helper.VerificaValor("Debito", debito); err != nil {
		return
	} else if err = helper.VerificaValor("Credito", credito); err != nil {
		return
	} else if (debito + credito) == 0 {
		return errors.New("Campos débito e crédito não podem ter simultaneamente valor zero(0)")
	}

	return
}
