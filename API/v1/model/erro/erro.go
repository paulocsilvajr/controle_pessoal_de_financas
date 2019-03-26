package erro

import (
	"errors"
	"fmt"
)

func ErroTamanho(msg string, tamanho int) error {
	msg += "[%d]"
	return errors.New(fmt.Sprintf(msg, tamanho))
}

func ErroDetalhe(msg, detalhe string) error {
	msg += "[%s]"
	return errors.New(fmt.Sprintf(msg, detalhe))
}
