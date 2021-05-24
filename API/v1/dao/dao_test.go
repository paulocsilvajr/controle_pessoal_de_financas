package dao

import "testing"

func TestCreateDB(t *testing.T) {
	err := CreateDBParaTestes()
	if err != nil {
		t.Error(err)
	}
}

func TestGetDB02(t *testing.T) {
	db := GetDB02ParaTestes()

	err := PingDB(db)
	if err != nil {
		t.Error("Não foi possível estabelecer conexão com o Banco de Dados", db)
	}
}

// TESTES ANTIGOS

// var (
// 	db = GetDB()
// )

// func TestGetDB(t *testing.T) {
// 	db := GetDB()
// 	err := db.Ping()

// 	if err != nil {
// 		t.Error("Não foi possível estabelecer conexão com o Banco de Dados", db)
// 	}
// }

// func TestGetTemplateQuery(t *testing.T) {
// 	sql := `
// SELECT
// 	{{.cpf}}, {{.nomeCompleto}}, {{.usuario}}, {{.senha}}, {{.email}}, {{.dataCriacao}}, {{.dataModificacao}}, {{.estado}}
// FROM
// 	{{.tabela}}
// `

// 	query := getTemplateQuery("carrega pessoas", pessoaDB, sql)
// 	test := `
// SELECT
// 	cpf, nome_completo, usuario, senha, email, data_criacao, data_modificacao, estado
// FROM
// 	pessoa
// `
// 	if query != test {
// 		t.Error(query)
// 	}
// }
