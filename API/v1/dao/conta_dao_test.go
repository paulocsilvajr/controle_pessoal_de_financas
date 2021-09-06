package dao

import (
	"testing"

	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/model/conta"
	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/model/tipo_conta"
)

var (
	testTipoContaB01         *tipo_conta.TipoConta
	testConta01, testConta02 *conta.Conta
)

func TestAdicionaConta02(t *testing.T) {
	ttc01 := getTTipoConta1()
	testTipoContaB01 = ConverteTTipoContaParaTipoConta(ttc01)
	testTipoContaB01, err := AdicionaTipoConta02(
		db2,
		testTipoContaB01,
	)
	if err != nil {
		t.Error("Erro em inclusão de tipo conta em TestAdicionaConta02", testTipoContaB01, err)
	}

	tc01 := getTConta1(ttc01)
	testConta01 = ConverteTContaParaConta(tc01)
	testConta01, err = AdicionaConta02(
		db2,
		testConta01,
	)
	if err != nil {
		t.Error("Erro em inclusão de conta em TestAdicionaConta02", testConta01, err)
	}

	tc02 := getTConta2(ttc01, tc01)
	testConta02 = ConverteTContaParaConta(tc02)
	testConta02, err = AdicionaConta02(
		db2,
		testConta02,
	)
	if err != nil {
		t.Error("Erro em inclusão de conta em TestAdicionaConta02", testConta02, err)
	}

	nomeEsperado := ttc01.Nome
	nomeObtido := testTipoContaB01.Nome
	if nomeEsperado != nomeObtido {
		t.Errorf("Nome de tipo de conta inserida em BD diferente do esperado. Esperado: '%s', obtido: '%s'", nomeEsperado, nomeObtido)
	}

	nomeEsperado = tc01.Nome
	nomeObtido = testConta01.Nome
	if nomeEsperado != nomeObtido {
		t.Errorf("Nome de conta inserida em BD diferente do esperado. Esperado: '%s', obtido: '%s'", nomeEsperado, nomeObtido)
	}

	nomeEsperado = tc02.Nome
	nomeObtido = testConta02.Nome
	if nomeEsperado != nomeObtido {
		t.Errorf("Nome de conta inserida em BD diferente do esperado. Esperado: '%s', obtido: '%s'", nomeEsperado, nomeObtido)
	}
}

func TestProcuraConta02(t *testing.T) {
	nomeProcurado := testConta02.Nome
	c, err := ProcuraConta02(db2, nomeProcurado)
	if err != nil {
		t.Error(err)
	}

	nomeEncontrado := c.Nome
	if nomeEncontrado != nomeProcurado {
		t.Errorf("Nome procurado diferente de nome encontrado. Esperado: '%s', encontrado: '%s'", nomeProcurado, nomeEncontrado)
	}
}

func TestAlteraConta02(t *testing.T) {
	nome := testConta01.Nome
	novoCodigo := "C0D1G0 N0V0"
	novoComentario := "Comentário NOVO"

	testConta01.Codigo = novoCodigo
	testConta01.Comentario = novoComentario

	c, err := AlteraConta02(db2, nome, testConta01)
	if err != nil {
		t.Error(err)
	}

	codigo := c.Codigo
	if codigo != novoCodigo {
		t.Errorf("Alteração de conta com nome %s retornou um 'codigo' diferente do esperado. Esperado: '%s', obtido: '%s'", nome, novoCodigo, codigo)
	}

	comentario := c.Comentario
	if comentario != novoComentario {
		t.Errorf("Alteração de conta com nome %s retornou um 'comentario' diferente do esperado. Esperado: '%s', obtido: '%s'", nome, novoComentario, comentario)
	}

	novoNome := "N0V0 N0m3 d3 c0nt4"
	novoComentario2 := "C0M3N3NT4R10..."
	testConta01.Nome = novoNome
	testConta01.Comentario = novoComentario2
	c2, err := AlteraConta02(db2, nome, testConta01)
	if err != nil {
		t.Error(err)
	}

	nomeAlterado := c2.Nome
	if novoNome != nomeAlterado {
		t.Errorf("Alteração de conta com nome %s retornou um 'nome' diferente do esperado. Esperado: '%s', obtido: '%s'", nome, novoNome, nomeAlterado)
	}

	comentario = c2.Comentario
	if comentario != novoComentario2 {
		t.Errorf("Alteração de conta com nome %s retornou um 'comentario' diferente do esperado. Esperado: '%s', obtido: '%s'", nome, novoComentario2, comentario)
	}
	// t.Error(c2)
}

func TestInativaConta02(t *testing.T) {
	nome := testConta01.Nome

	c, err := InativaConta02(db2, nome)
	if err != nil {
		t.Error(err)
	}

	estadoObtido := c.Estado
	estadoEsperado := false
	if estadoObtido != estadoEsperado {
		t.Errorf("Inativação de conta com nome '%s' retornou um 'estado' diferente do esperado. Esperado: '%t', obtido: '%t'", nome, estadoEsperado, estadoObtido)
	}
}

func TestCarregaContasInativa02(t *testing.T) {
	contas, err := CarregaContasInativa02(db2)
	if err != nil {
		t.Error(err)
	}

	quantObtida := len(contas)
	quantEsperada := 1
	if quantObtida != quantEsperada {
		t.Errorf("consulta de Contas Inativas retornou uma quantidade de registros diferente do esperado. Esperado: %d, obtido: %d", quantEsperada, quantObtida)
	}
}

func TestAtivaConta02(t *testing.T) {
	nome := testConta01.Nome

	c, err := AtivaConta02(db2, nome)
	if err != nil {
		t.Error(err)
	}

	estadoObtido := c.Estado
	estadoEsperado := true
	if estadoObtido != estadoEsperado {
		t.Errorf("Ativação de conta com nome '%s' retornou um 'estado' diferente do esperado. Esperado: '%t', obtido: '%t'", nome, estadoEsperado, estadoObtido)
	}
}

func TestCarregaContas02(t *testing.T) {
	contas, err := CarregaContas02(db2)
	if err != nil {
		t.Error(err)
	}

	quantObtida := len(contas)
	quantEsperada := 2
	if quantObtida != quantEsperada {
		t.Errorf("consulta de Contas retornou uma quantidade de registros diferente do esperado. Esperado: %d, obtido: %d", quantEsperada, quantObtida)
	}

	nome1 := testConta01.Nome
	nome2 := testConta02.Nome
	for _, c := range contas {
		if c.Nome == nome1 || c.Nome == nome2 {
			continue
		}
		t.Errorf("registro de conta(nome) encontrado diferente do esperado. Esperado: '%s' ou '%s', obtido: '%s'", nome1, nome2, c.Nome)
	}
}

func TestCarregaContasAtiva02(t *testing.T) {
	contas, err := CarregaContasAtiva02(db2)
	if err != nil {
		t.Error(err)
	}

	quantObtida := len(contas)
	quantEsperada := 2
	if quantObtida != quantEsperada {
		t.Errorf("consulta de Contas ativas retornou uma quantidade de registros diferente do esperado. Esperado: %d, obtido: %d", quantEsperada, quantObtida)
	}
}

func TestRemoveConta02(t *testing.T) {
	err := RemoveConta02(
		db2,
		testConta02.Nome,
	)
	if err != nil {
		t.Error(err)
	}

	err = RemoveConta02(
		db2,
		testConta01.Nome,
	)
	if err != nil {
		t.Error(err)
	}

	err = RemoveTipoConta02(
		db2,
		testTipoContaB01.Nome,
	)
	if err != nil {
		t.Error(err)
	}
}
