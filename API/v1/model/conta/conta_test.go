package conta

import (
	"testing"
)

func TestIConta(t *testing.T) {
	c1 := &Conta{}
	c2 := new(Conta)

	var ic1 IConta = c1
	c3, ok := ic1.(*Conta)
	if !ok {
		t.Error(c3)
	}

	var ic2 IConta = c2
	c4, ok := ic2.(*Conta)
	if !ok {
		t.Error(c4)
	}
}

func TestConverteParaConta(t *testing.T) {
	c1 := &Conta{}
	c2 := new(Conta)

	c3, err := converteParaConta(c1)
	if err != nil {
		t.Error(c3)
	}

	c4, err := converteParaConta(c2)
	if err != nil {
		t.Error(c4)
	}
}
