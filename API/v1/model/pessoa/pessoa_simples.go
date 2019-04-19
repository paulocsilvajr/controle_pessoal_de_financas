package pessoa

import (
	"errors"
	"fmt"
	"time"
)

type PessoaSimples struct {
	Usuario         string    `json:"usuario"`
	Email           string    `json:"email"`
	DataCriacao     time.Time `json:"data_criacao"`
	DataModificacao time.Time `json:"data_modificacao"`
}

func (p PessoaSimples) GetEmail() string {
	return p.Email
}

type PessoasSimples []*PessoaSimples

func (ps PessoasSimples) Len() int {
	return len(ps)
}

func (ps PessoasSimples) ProcuraPessoaPorUsuario(usuario string) (p PessoaI, err error) {
	for _, pessoaLista := range ps {
		if pessoaLista.Usuario == usuario {
			p = pessoaLista
			return
		}
	}

	err = errors.New(fmt.Sprintf(
		"Pessoa com usuário %s informado não existe na listagem", usuario))

	return
}
