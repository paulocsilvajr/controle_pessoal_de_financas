// pessoa_test.go
package pessoa

import (
	"testing"
)

func TestCpf(t *testing.T) {
	p1, err := New("12345678910", "Teste 01", "teste01", "123456", "teste01@email.com")
	if err != nil {
		t.Error(err, p1)
	}

	p1, err = New("111.222.333-44", "Pedro Gonzaga", "pedro", "brasil", "email@gmail.com")
	if err != nil {
		t.Error(err, p1)
	}

	p1, err = New("1112223334A", "Pedro Gonzaga", "pedro", "brasil", "email@gmail.com")
	if err != nil {
		t.Error(err, p1)
	}
}

func TestNome(t *testing.T) {
	p1, err := New("11122233344", "Pedro de Alcântara João Carlos Leopoldo Salvador Bibiano Francisco Xavier de Paula Leocádio Miguel Gabriel Rafael Gonzaga", "pedro", "brasil", "pedro@gmail.com")
	if err != nil {
		t.Error(err, p1)
	}
}
