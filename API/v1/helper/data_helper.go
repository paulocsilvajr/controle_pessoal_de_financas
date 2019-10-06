package helper

import "time"

// DataFormatada retorna uma string no formato dd/mm/yyyy hh:MM:ss baseado na data passada como parâmetro
func DataFormatada(data time.Time) string {
	return data.Format("02/01/2006 15:04:05") // formatação de data aos moldes de Golang
}
