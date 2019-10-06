package helper

import "fmt"

// MonetarioFormatado retorna uma string baseado em um número decimal(float32) no formato de 3 casas decimais
func MonetarioFormatado(valor float64) string {
	return fmt.Sprintf("%.3f", valor)
}
