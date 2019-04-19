package dao

import (
	"controle_pessoal_de_financas/API/v1/model/pessoa"
	"database/sql"
)

func CarregaPessoasSimples(db *sql.DB) (pessoas pessoa.PessoasSimples, err error) {
	sql := `
SELECT
	{{.usuario}}, {{.email}}, {{.dataCriacao}}, {{.dataModificacao}}
FROM
	{{.tabela}}
WHERE
	{{.estado}} = true
`
	query := getTemplateQuery("CarregaPessoas", pessoaDB, sql)

	return carregaPessoasSimples(db, query)
}

func carregaPessoasSimples(db *sql.DB, query string, args ...interface{}) (pessoas pessoa.PessoasSimples, err error) {
	registros, err := carrega(db, query, registrosPessoas02, args...)

	pessoas = converteEmPessoasSimples(registros)

	return
}

func converteEmPessoasSimples(registros []interface{}) (pessoas pessoa.PessoasSimples) {
	for _, r := range registros {
		// fmt.Printf(">>> %T\n", r)
		p, ok := r.(*pessoa.PessoaSimples)
		if ok {
			pessoas = append(pessoas, p)
		}
	}

	return
}
