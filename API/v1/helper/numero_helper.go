package helper

import "fmt"

// MonetarioFormatado retorna uma string baseado em um número decimal(float64) no formato de 3 casas decimais
func MonetarioFormatado(valor float64) string {
	return fmt.Sprintf("%.3f", valor)
}

// VerificaValor retorna um erro caso o parâmetro valor informado seja < 0. O primeiro parâmetro é o nome do campo testado, para exibir este nome na mensagem de erro, caso ocorra
func VerificaValor(campo string, valor float64) (err error) {
	if valor < 0 {
		err = fmt.Errorf("O campo %s deve ser >= 0", campo)
	}

	return
}
