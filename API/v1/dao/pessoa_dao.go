package dao

import (
	"bytes"
	"controle_pessoal_de_financas/API/v1/model/pessoa"
	"database/sql"
	"html/template"
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

	query := getTemplateQuery("carrega pessoas", sql)

	return carregaPessoas(db, query)
}

func carregaPessoas(db *sql.DB, query string) (pessoas pessoa.Pessoas, err error) {
	queryStmt, err := db.Prepare(query)
	if err != nil {
		return
	}

	rows, err := queryStmt.Query()
	defer queryStmt.Close()
	if err != nil {
		return
	}

	for rows.Next() {
		pessoaAtual := new(pessoa.Pessoa)
		err = rows.Scan(
			&pessoaAtual.Cpf,
			&pessoaAtual.NomeCompleto,
			&pessoaAtual.Usuario,
			&pessoaAtual.Senha,
			&pessoaAtual.Email,
			&pessoaAtual.DataCriacao,
			&pessoaAtual.DataModificacao,
			&pessoaAtual.Estado)
		if err != nil {
			return
		}
		pessoas = append(pessoas, pessoaAtual)
	}
	err = rows.Err()
	if err != nil {
		pessoas = nil
		return
	}

	return
}

func getTemplateQuery(nome, sql string) string {
	t := template.Must(template.New(nome).Parse(sql))
	query := new(bytes.Buffer)
	t.Execute(query, pessoaDB)

	return query.String()
}
