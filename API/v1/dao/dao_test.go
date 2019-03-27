package dao

import "testing"

func TestGetDB(t *testing.T) {
	db := GetDB()
	err := db.Ping()

	if err != nil {
		t.Error("Não foi possível estabelecer conexão com o Banco de Dados", db)
	}
}
