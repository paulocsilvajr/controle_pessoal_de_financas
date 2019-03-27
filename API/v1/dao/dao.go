package dao

import (
	"database/sql"
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
