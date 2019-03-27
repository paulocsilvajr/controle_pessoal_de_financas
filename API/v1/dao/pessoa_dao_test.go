package dao

import "testing"

var db = GetDB()

func TestGetTemplateQuery(t *testing.T) {
	sql := `
SELECT
	{{.cpf}}, {{.nomeCompleto}}, {{.usuario}}, {{.senha}}, {{.email}}, {{.dataCriacao}}, {{.dataModificacao}}, {{.estado}}
FROM
	{{.tabela}}
`

	query := getTemplateQuery("carrega pessoas", sql)
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

func TestDaoCarregaPessoas(t *testing.T) {
	listaPessoas, err := DaoCarregaPessoas(db)

	if err != nil {
		t.Error(err, listaPessoas)
	}
}
