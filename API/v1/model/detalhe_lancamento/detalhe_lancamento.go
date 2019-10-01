package detalhe_lancamento

type IDetalheLancamento interface {
	String() string
	Repr() string
	VerificaAtributos() error
	Altera(string, int64, int64) error
	AlteraCampos(map[string]string) error
	Ativa()
	Inativa()
}

type DetalheLancamento struct {
	ID        int    `json:"id"`
	NomeConta string `json:"nome_conta"`
	Debito    int64  `json:"debito"`
	Credito   int64  `json:"credito"`
}

const (
	MaxNomeConta = 50
)

type IDetalheLancamentos interface {
	Len()
	ProcutaDetalheLancamento(int) *DetalheLancamento
}

type DetalheLancamentos []*DetalheLancamento
