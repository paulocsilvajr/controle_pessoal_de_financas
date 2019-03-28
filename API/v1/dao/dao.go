package dao

import (
	"bytes"
	"database/sql"
	"html/template"
	"log"

	_ "github.com/lib/pq"
)

func GetDB() *sql.DB {
	connStr := "host=localhost port=15432 user=postgres password=Postgres2019! dbname=controle_pessoal_financas sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func getTemplateQuery(nome string, campos map[string]string, sql string) string {
	t := template.Must(template.New(nome).Parse(sql))
	query := new(bytes.Buffer)
	t.Execute(query, campos)

	return query.String()
}

func carrega(db *sql.DB, query string, appendRegistros func(rows *sql.Rows, registros []interface{}) (novosRegistros []interface{}, err error)) (registros []interface{}, err error) {
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
		registros, err = appendRegistros(rows, registros)
		if err != nil {
			return
		}

	}
	err = rows.Err()
	if err != nil {
		registros = nil
		return
	}

	return
}

func adiciona(db *sql.DB, novoRegistro interface{}, query string, setValores func(*sql.Stmt, interface{}) (sql.Result, error)) (r interface{}, err error) {

	transacao, err := db.Begin()
	if err != nil {
		return
	}

	stmt, err := transacao.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = setValores(stmt, novoRegistro)
	if err != nil {
		return
	}

	err = transacao.Commit()
	if err != nil {
		return
	}

	r = novoRegistro

	return
}
