package config

import (
	"testing"
)

func TestLenRotas(t *testing.T) {
	tamanhoEsperado := len(Rotas)
	tamanhoMetodoLen := Rotas.Len()
	if tamanhoEsperado != tamanhoMetodoLen {
		t.Errorf("método Len de map rotas retornou um tamanho inválido, esperado: %d, retorno: %d", tamanhoEsperado, tamanhoMetodoLen)
	}
}

func TestDefineDocumentacao(t *testing.T) {
	nomeRota := "Index"
	r := Rotas[nomeRota]
	documentacao := r.Documentacao
	if documentacao != "" {
		t.Errorf("documentação de rota '%s' preenchida antes da adição de documentação. Esperado: '%s', retorno: '%s'", nomeRota, "", documentacao)
	}

	documentacaoTeste := "Documentação teste"
	Rotas.DefineDocumentacao(nomeRota, documentacaoTeste)
	r = Rotas[nomeRota]
	documentacao = r.Documentacao
	if documentacao != documentacaoTeste {
		t.Errorf("documentacao de rota '%s' diferente da atribuida pelo método DefineDocumentacao. Esperado: '%s', retorno: '%s'", nomeRota, documentacaoTeste, documentacao)
	}
}
