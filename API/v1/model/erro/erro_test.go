package erro

import "testing"

func TestErros(t *testing.T) {
	erroTamanho := ErroTamanho("mensagem", 42)
	if erroTamanho.Error() != "mensagem[42]" {
		t.Error("Retorno incorreto, função erros.ErroTamanho()")
	}

	erroDetalhe := ErroDetalhe("mensagem", "Resposta para a vida, o universo e tudo mais")
	if erroDetalhe.Error() != "mensagem[Resposta para a vida, o universo e tudo mais]" {
		t.Error("Retorno incorreto, função erros.ErroDetalhe()")
	}

}
