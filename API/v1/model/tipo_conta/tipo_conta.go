package tipo_conta

import (
	"controle_pessoal_de_financas/API/v1/helper"
	"controle_pessoal_de_financas/API/v1/model/erro"
	"fmt"
	"time"
)

// ITipoConta é uma interface que exige a implementação dos métodos obrigatórios em TipoConta
type ITipoConta interface {
	String() string
	Repr() string
	VerificaAtributos() error
	Altera(string, string, string) error
	AlteraCampos(map[string]string) error
	Ativa()
	Inativa()
}

// TipoConta é uma struct que representa um tipo de conta. Possui notações JSON para cada campo e tem a composição da interface ITipoConta
type TipoConta struct {
	Nome             string    `json:"nome"`
	DescricaoDebito  string    `json:"descricao_debito"`
	DescricaoCredito string    `json:"descricao_credito"`
	DataCriacao      time.Time `json:"data_criacao"`
	DataModificacao  time.Time `json:"data_modificacao"`
	Estado           bool      `json:"estado"`
	ITipoConta
}

// MaxNome: tamanho máximo para o nome do tipo de conta
// MaxDescricao: tamanho máximo para as descrições de débito e crédito de tipo de conta
// MsgErroNome01: mensagem erro padrão 01 para nome
// MsgErroDescricao01: mensagem erro padrão 01 para descrição(débito)
// MsgErroDescricao02: mensagem erro padrão 02 para descrição(crédito)
const (
	MaxNome            = 50
	MaxDescricao       = 20
	MsgErroNome01      = "Nome com tamanho inválido"
	MsgErroDescricao01 = "Descrição do débito com tamanho inválido"
	MsgErroDescricao02 = "Descrição do crédito com tamanho inválido"
)

// ITiposConta é uma interface que exige a implementação dos métodos ProcuraTipoConta e Len para representar um conjunto/lista(slice) de Tipos de Contas genéricas
type ITiposConta interface {
	Len() int
	ProcuraTipoConta(string) *TipoConta
}

// TiposConta representa um conjunto/lista(slice) de Tipos de Contas(*TipoConta)
type TiposConta []*TipoConta

// New retorna uma novo Tipo de Conta(*TipoConta) através dos parâmetros informados(nome, descDebito e descCredito). Função equivalente a criação de um TipoConta via literal &TipoConta{Nome: ..., ...}. Data de criação e modificação são definidos como o horário atual e o estado é definido como ativo
func New(nome, descDebito, descCredito string) *TipoConta {
	return &TipoConta{
		Nome:             nome,
		DescricaoDebito:  descDebito,
		DescricaoCredito: descCredito,
		DataCriacao:      time.Now().Local(),
		DataModificacao:  time.Now().Local(),
		Estado:           true}
}

// NewTipoConta cria uma novo TipoConta semelhante a função New(), mas faz a validação dos campos informados nos parâmetros nome, descDebito e descCredito
func NewTipoConta(nome, descDebito, descCredito string) (tipoConta *TipoConta, err error) {
	tipoConta = New(nome, descDebito, descCredito)

	if err = tipoConta.VerificaAtributos(); err != nil {
		tipoConta = nil
	}

	return
}

// GetTipoContaTest retorna um TipoConta teste usado para testes em geral
func GetTipoContaTest() (tipoConta *TipoConta) {
	tipoConta = New("banco teste 01", "saque", "depósito")
	tipoConta.DataCriacao = time.Date(2000, 2, 1, 12, 30, 0, 0, new(time.Location))
	tipoConta.DataModificacao = time.Date(2000, 2, 1, 12, 30, 0, 0, new(time.Location))
	tipoConta.Estado = true

	return
}

// String retorna uma representação textual de um TipoConta. Datas são formatadas usando a função helper.DataFormatada() e campo estado é formatado usando a função helper.GetEstado()
func (t *TipoConta) String() string {
	estado := helper.GetEstado(t.Estado)
	dataCriacao := helper.DataFormatada(t.DataCriacao)
	dataModificacao := helper.DataFormatada(t.DataModificacao)

	return fmt.Sprintf("%s\t%s\t%s\t%s\t%s\t%s", t.Nome, t.DescricaoDebito, t.DescricaoCredito, dataCriacao, dataModificacao, estado)
}

// Repr é um método que retorna uma string da representação de uma TipoConta, sem formatações especiais
func (t *TipoConta) Repr() string {
	return fmt.Sprintf("%s\t%s\t%s\t%s\t%s\t%v", t.Nome, t.DescricaoDebito, t.DescricaoCredito, t.DataCriacao, t.DataModificacao, t.Estado)
}

// VerificaAtributos é um método de tipoConta que verifica os campos Nome, DescricaoDebito e DescricaoCredito, retornando um erro != nil caso ocorra um problema
func (t *TipoConta) VerificaAtributos() error {
	return verifica(t.Nome, t.DescricaoDebito, t.DescricaoCredito)
}

// Altera é um método que modifica os dados do TipoConta a partir dos parâmetros ifnormados depois da verificação de cada parâmetro e atualiza a data de modificação dele. Retorna um erro != nil, caso algum parâmetro seja inválido
func (t *TipoConta) Altera(nome string, descDebito string, descCredito string) (err error) {
	if err = verifica(nome, descDebito, descCredito); err != nil {
		return
	}

	t.Nome = nome
	t.DescricaoDebito = descDebito
	t.DescricaoCredito = descCredito
	t.DataModificacao = time.Now().Local()

	return
}

// AlteraCampos é um método para alterar os campos de um TipoConta a partir de hashMap informado no parâmetro campos. Somente as chaves ifnormadas com um valor correto serão atualizadas. É atualizado a data de modificação do TipoConta. Cado ocorra um problema na validação dos campos, retorna um erro != nil. Campos permitidos: nome, descricaoDebito, descricaoCredito
func (t *TipoConta) AlteraCampos(campos map[string]string) (err error) {
	for chave, valor := range campos {
		switch chave {
		case "nome":
			if err = verificaNome(valor); err != nil {
				return
			}
			t.Nome = valor
		case "descricaoDebito":
			if err = verificaDescricao(valor, false); err != nil {
				return
			}
			t.DescricaoDebito = valor
		case "descricaoCredito":
			if err = verificaDescricao(valor, true); err != nil {
				return
			}
			t.DescricaoCredito = valor
		}
	}
	t.DataModificacao = time.Now().Local()

	return
}

// Ativa é um método que define o TipoConta como ativo e atualiza a sua data de modificação
func (t *TipoConta) Ativa() {
	t.alteraEstado(true)
}

// Inativa é um método que define o TipoConta como inativo e atualiza a sua data de modificação
func (t *TipoConta) Inativa() {
	t.alteraEstado(false)
}

// ProcuraTipoConta é um método que retorna um TipoConta a partir da busca em uma listagem de TiposConta. Caso não seja encontrado o TipoConta, retorna um erro != nil. A interface ITiposConta exige a implementação desse método
func (ts TiposConta) ProcuraTipoConta(tipoConta string) (t *TipoConta, err error) {
	for _, tipoContaLista := range ts {
		if tipoContaLista.Nome == tipoConta {
			t = tipoContaLista
			return
		}
	}

	err = fmt.Errorf("Tipo de conta %s informada não existe na listagem", tipoConta)

	return
}

// Len é um método de TiposConta que retorna a quantidade de elementos contidos dentro do slice de TipoConta. A interface ITiposConta exibe a implementação desse método
func (ts TiposConta) Len() int {
	return len(ts)
}

func (t *TipoConta) alteraEstado(estado bool) {
	t.DataModificacao = time.Now().Local()
	t.Estado = estado
}

func verifica(nome, descDebito, descCredito string) (err error) {
	if err = verificaNome(nome); err != nil {
		return
	} else if err = verificaDescricao(descDebito, false); err != nil {
		return
	} else if err = verificaDescricao(descCredito, true); err != nil {
		return
	}

	return
}

func verificaNome(nome string) (err error) {
	if len(nome) == 0 || len(nome) > MaxNome {
		err = erro.ErroTamanho(MsgErroNome01, len(nome))
	}

	return
}

func verificaDescricao(descricao string, tipoCredito bool) (err error) {
	if len(descricao) == 0 || len(descricao) > MaxDescricao {
		if tipoCredito {
			err = erro.ErroTamanho(MsgErroDescricao02, len(descricao))
		} else {
			err = erro.ErroTamanho(MsgErroDescricao01, len(descricao))
		}
	}

	return
}
