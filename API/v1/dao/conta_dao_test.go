package dao

import "testing"

func TestCarregaContas(t *testing.T) {
	listaContas, err := CarregaContas(db)

	if err != nil {
		t.Error(err, listaContas)
	}

	if len(listaContas) == 0 {
		t.Error(listaContas)
	}

	if len(listaContas) < 3 {
		t.Error(listaContas)
	}
}
