package helper

import "fmt"

// VerificaCampoTexto retorna um erro caso o tamanho do texto informado no parâmetro 'campo' não seja > 0 e <= ao tamanho informado em parâmetro 'tamanho'. O primeiro parâmetro é o nome do campo testado, para exibir este nome na mensagem de erro, caso ocorra
func VerificaCampoTexto(nomeCampo, campo string, tamanho int) error {
	campoValido := len(campo) > 0 && len(campo) <= tamanho
	if campoValido {
		return nil
	}
	return fmt.Errorf("Tamanho de campo %s inválido[%d caracter(es)]", nomeCampo, len(campo))
}

// VerificaCampoTextoOpcional retorna um erro caso o tamanho do texto informado no parâmetro 'campo' não seja >= 0 e <= ao tamanho informado em parâmetro 'tamanho'. O primeiro parâmetro é o nome do campo testado, para exibir este nome na mensagem de erro, caso ocorra
func VerificaCampoTextoOpcional(nomeCampo, campo string, tamanho int) error {
	campoValido := len(campo) >= 0 && len(campo) <= tamanho
	if campoValido {
		return nil
	}
	return fmt.Errorf("Tamanho de campo %s inválido[%d caracter(es)]", nomeCampo, len(campo))
}
