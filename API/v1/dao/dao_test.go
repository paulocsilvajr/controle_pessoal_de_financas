package dao

import "testing"

func TestGetDB(t *testing.T) {
	db := GetDB()
	err := db.Ping()

	if err != nil {
		t.Error("Não foi possível estabelecer conexão com o Banco de Dados", db)
	}
}

func TestGetTemplateQuery(t *testing.T) {
	sql := `
SELECT
	{{.cpf}}, {{.nomeCompleto}}, {{.usuario}}, {{.senha}}, {{.email}}, {{.dataCriacao}}, {{.dataModificacao}}, {{.estado}}
FROM
	{{.tabela}}
`

	query := getTemplateQuery("carrega pessoas", pessoaDB, sql)
	test := `
SELECT
	cpf, nome_completo, usuario, senha, email, data_criacao, data_modificacao, estado
FROM
	pessoa
`
	if query != test {
		t.Error(query)
	}
}
