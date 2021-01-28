package pessoa

import (
	"fmt"
	"time"
)

// PessoaSimples é uma estrutura que representa uma pessoa com menos informações, para gerar retorno em JSON mais simples
type PessoaSimples struct {
	Usuario         string    `json:"usuario"`
	Email           string    `json:"email"`
	DataCriacao     time.Time `json:"data_criacao"`
	DataModificacao time.Time `json:"data_modificacao"`
}

// GetEmail é um método de PessoaSimples que retorna uma string contendo o email da pessoa. Exigência da interface IPessoa
func (p *PessoaSimples) GetEmail() string {
	return p.Email
}

// CorrigeData é um método que converte a data(Time) no formato do timezone local
func (p *PessoaSimples) CorrigeData() {
	p.DataCriacao = p.DataCriacao.Local()
	p.DataModificacao = p.DataModificacao.Local()
}

// PessoasSimples é a representação de um conjunto/lista(slice) de pessoas simples(*PessoaSimples)
type PessoasSimples []*PessoaSimples

// Len é um método de PessoasSimples que retorna a quantidade de elementos contidos dentro do slice de pessoas simples. A interface IPessoas exibe a implementação desse método
func (ps PessoasSimples) Len() int {
	return len(ps)
}

// ProcuraPessoaPorUsuario é um método que retorna uma pessoa a partir da busca em uma listagem de pessoas(PessoasSimples). Caso não seja encontrado a pessoa, retorna um erro != nil. A interface IPessoas exige a implementação desse método
func (ps PessoasSimples) ProcuraPessoaPorUsuario(usuario string) (p IPessoa, err error) {
	for _, pessoaLista := range ps {
		if pessoaLista.Usuario == usuario {
			p = pessoaLista
			return
		}
	}

	err = fmt.Errorf(
		"Pessoa com usuário %s informado não existe na listagem", usuario)

	return
}
