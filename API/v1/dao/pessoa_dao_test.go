package dao

import "testing"

var db = GetDB()

func TestDaoCarregaPessoas(t *testing.T) {
	listaPessoas, err := DaoCarregaPessoas(db)

	if err != nil {
		t.Error(err, listaPessoas)
	}

	// if len(listaPessoas) > 0 {
	// 	t.Error(listaPessoas)
	// }
}
