package dao

import (
	"controle_pessoal_de_financas/API/v1/model/pessoa"
	"database/sql"
)

var (
	pessoaDB = map[string]string{
		"tabela":          "pessoa",
		"cpf":             "cpf",
		"nomeCompleto":    "nome_completo",
		"usuario":         "usuario",
		"senha":           "senha",
		"email":           "email",
		"dataCriacao":     "data_criacao",
		"dataModificacao": "data_modificacao",
		"estado":          "estado"}
)

func DaoCarregaPessoas(db *sql.DB) (pessoas pessoa.Pessoas, err error) {
	sql := `
SELECT
	{{.cpf}}, {{.nomeCompleto}}, {{.usuario}}, {{.senha}}, {{.email}}, {{.dataCriacao}}, {{.dataModificacao}}, {{.estado}}
FROM
	{{.tabela}}
`
	// fmt.Println(fmt.Sprintf(`sql
	// nova linha{{.atributoParaTemplate}}: %d`, 123)) // exemplo para adicionar numero concatenado ao sql template

	query := getTemplateQuery("CarregaPessoas", pessoaDB, sql)

	return carregaPessoas(db, query)
}

func carregaPessoas(db *sql.DB, query string) (pessoas pessoa.Pessoas, err error) {
	registros, err := carrega(db, query, registrosPessoas01)

	pessoas = converteEmPessoas(registros)

	return
}

func converteEmPessoas(registros []interface{}) (pessoas pessoa.Pessoas) {
	for _, r := range registros {
		// fmt.Printf(">>> %T\n", r)
		p, ok := r.(*pessoa.Pessoa)
		if ok {
			pessoas = append(pessoas, p)
		}
	}

	return
}

func registrosPessoas01(rows *sql.Rows, registros []interface{}) (novosRegistros []interface{}, err error) {
	pessoaAtual := new(pessoa.Pessoa)
	err = scanPessoas01(rows, pessoaAtual)
	if err != nil {
		return
	}
	novosRegistros = append(registros, pessoaAtual)

	return
}

func scanPessoas01(rows *sql.Rows, pessoaAtual *pessoa.Pessoa) error {
	return rows.Scan(
		&pessoaAtual.Cpf,
		&pessoaAtual.NomeCompleto,
		&pessoaAtual.Usuario,
		&pessoaAtual.Senha,
		&pessoaAtual.Email,
		&pessoaAtual.DataCriacao,
		&pessoaAtual.DataModificacao,
		&pessoaAtual.Estado)
}

// func OLDcarregaPessoas(db *sql.DB, query string) (pessoas pessoa.Pessoas, err error) {
// 	queryStmt, err := db.Prepare(query)
// 	if err != nil {
// 		return
// 	}

// 	rows, err := queryStmt.Query()
// 	defer queryStmt.Close()
// 	if err != nil {
// 		return
// 	}

// 	for rows.Next() {
// 		pessoaAtual := new(pessoa.Pessoa)
// 		err = scanPessoas01(rows, pessoaAtual)
// 		if err != nil {
// 			return
// 		}
// 		pessoas = append(pessoas, pessoaAtual)
// 	}
// 	err = rows.Err()
// 	if err != nil {
// 		pessoas = nil
// 		return
// 	}

// 	return
// }
