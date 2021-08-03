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
		t.Errorf("Alteração de conta com nome %s retornou um 'codigo' diferente do esperado. Esperado: '%s', retorno: '%s'", nome, novoCodigo, codigo)
	}

	comentario := c.Comentario
	if comentario != novoComentario {
		t.Errorf("Alteração de conta com nome %s retornou um 'comentario' diferente do esperado. Esperado: '%s', retorno: '%s'", nome, novoComentario, comentario)
	}
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

// TESTES ANTIGOS
// import (
// 	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/model/conta"
// 	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/model/tipo_conta"
// 	"testing"
// )

// func TestAdicionaConta(t *testing.T) {
// 	strErroChaveEstrangeira := "pq: insert or update on table \"conta\" violates foreign key constraint \"tipo_conta_conta_fk\""
// 	strErroChaveUnica := "pq: duplicate key value violates unique constraint \"codigo_uq\""
// 	strErroChavePrimaria := "pq: duplicate key value violates unique constraint \"conta_pk\""

// 	c1 := conta.GetContaTest()

// 	c2 := conta.GetContaTest()
// 	c2.Nome = "Ativos Teste 02"

// 	c3 := conta.GetContaTest()
// 	c3.Nome = "Ativos Teste 03"

// 	c4 := conta.New("Teste conta C", "", "", "", "")

// 	c5, err := AdicionaConta(db, c1)
// 	temErro := err != nil
// 	erroChaveEstr := err.Error() == strErroChaveEstrangeira
// 	erroChavePrim := err.Error() == strErroChaveUnica
// 	erroNaoEConhecido := !(erroChaveEstr || erroChavePrim)
// 	if temErro && erroNaoEConhecido {
// 		t.Error(err, c5)
// 	}

// 	tc1, err := AdicionaTipoConta(db, tipo_conta.GetTipoContaTest())
// 	cadastrouNovoTipoConta := err == nil
// 	if cadastrouNovoTipoConta {
// 		c1.NomeTipoConta = tc1.Nome
// 	}
// 	c5, err = AdicionaConta(db, c1)
// 	temErro = err != nil
// 	erroChaveUnica := false
// 	if temErro {
// 		erroChaveUnica = err.Error() == strErroChaveUnica
// 	}
// 	if temErro && !erroChaveUnica {
// 		t.Error(err, c5)
// 	}

// 	c6, err := AdicionaConta(db, c2)
// 	if err != nil && err.Error() != strErroChaveUnica {
// 		t.Error(err, c6)
// 	}

// 	c2.Codigo = "002"
// 	c6, err = AdicionaConta(db, c2)
// 	temErro = err != nil
// 	erroChaveUnica = false
// 	if temErro {
// 		erroChaveUnica = err.Error() == strErroChaveUnica
// 	}
// 	if temErro && !erroChaveUnica {
// 		t.Error(err, c6)
// 	}

// 	c3.Codigo = "003"
// 	c7, err := AdicionaConta(db, c3)
// 	temErro = err != nil
// 	erroChaveUnica = false
// 	if temErro {
// 		erroChaveUnica = err.Error() == strErroChaveUnica
// 	}
// 	if temErro && !erroChaveUnica {
// 		t.Error(err, c7)
// 	}

// 	c8, err := AdicionaConta(db, c4)
// 	if err.Error() != "Tamanho de campo Nome do Tipo da Conta inválido[0 caracter(es)]" {
// 		t.Error(err, c8)
// 	}

// 	c3.Codigo = "004"
// 	c9, err := AdicionaConta(db, c3)
// 	if err.Error() != strErroChavePrimaria {
// 		t.Error(err, c9)
// 	}

// 	c10 := conta.New("Teste Conta 04", c1.NomeTipoConta, "", "", "")
// 	c11, err := AdicionaConta(db, c10)
// 	if err != nil {
// 		t.Error(err, c11)
// 	}
// }

// func TestInativaContaECarregaContasInativas(t *testing.T) {
// 	nome01 := "Ativos Teste 01"
// 	nome02 := "Ativos Teste 02"
// 	nome03 := "Ativos Teste 03"
// 	nome04 := "Ativos Teste 04"

// 	c01, err := InativaConta(db, nome01)
// 	if err != nil {
// 		t.Error(err, c01)
// 	}

// 	c02, err := InativaConta(db, nome02)
// 	if err != nil {
// 		t.Error(err, c02)
// 	}

// 	c03, err := InativaConta(db, nome03)
// 	if err != nil {
// 		t.Error(err, c03)
// 	}

// 	c04, err := InativaConta(db, nome04)
// 	if err.Error() != "Não foi encontrado um registro com o nome Ativos Teste 04" {
// 		t.Error(err, c04)
// 	}

// 	if c01.Estado != false {
// 		t.Error("Estado da conta inválido, deveria ser false", c01)
// 	}

// 	if c02.Estado != false {
// 		t.Error("Estado da conta inválido, deveria ser false", c02)
// 	}

// 	if c03.Estado != false {
// 		t.Error("Estado da conta inválido, deveria ser false", c03)
// 	}

// 	contas, err := CarregaContasInativa(db)
// 	if err != nil {
// 		t.Error(err, contas)
// 	}

// 	if len(contas) == 0 {
// 		t.Error(contas)
// 	}

// 	if len(contas) < 3 {
// 		t.Error(contas)
// 	}
// }

// func TestAtivaContaECarregaContasAtivas(t *testing.T) {
// 	nome01 := "Ativos Teste 01"
// 	nome02 := "Ativos Teste 02"
// 	nome03 := "Ativos Teste 03"
// 	nome04 := "Ativos Teste 04"

// 	c01, err := AtivaConta(db, nome01)
// 	if err != nil {
// 		t.Error(err, c01)
// 	}

// 	c02, err := AtivaConta(db, nome02)
// 	if err != nil {
// 		t.Error(err, c02)
// 	}

// 	c03, err := AtivaConta(db, nome03)
// 	if err != nil {
// 		t.Error(err, c03)
// 	}

// 	c04, err := AtivaConta(db, nome04)
// 	if err.Error() != "Não foi encontrado um registro com o nome Ativos Teste 04" {
// 		t.Error(err, c04)
// 	}

// 	if c01.Estado != true {
// 		t.Error("Estado da conta inválido, deveria ser true", c01)
// 	}

// 	if c02.Estado != true {
// 		t.Error("Estado da conta inválido, deveria ser true", c02)
// 	}

// 	if c03.Estado != true {
// 		t.Error("Estado da conta inválido, deveria ser true", c03)
// 	}

// 	contas, err := CarregaContasAtiva(db)
// 	if err != nil {
// 		t.Error(err, contas)
// 	}

// 	if len(contas) == 0 {
// 		t.Error(contas)
// 	}

// 	if len(contas) < 3 {
// 		t.Error(contas)
// 	}
// }

// func TestCarregaContas(t *testing.T) {
// 	listaContas, err := CarregaContas(db)

// 	if err != nil {
// 		t.Error(err, listaContas)
// 	}

// 	if len(listaContas) == 0 {
// 		t.Error(listaContas)
// 	}

// 	if len(listaContas) < 3 {
// 		t.Error(listaContas)
// 	}
// }

// func TestProcuraConta(t *testing.T) {
// 	nome01 := "Ativos Teste 01"
// 	nome02 := "Ativos Teste 02"
// 	nome03 := "Ativos Teste 03"

// 	c1, err := ProcuraConta(db, nome01)
// 	if err != nil {
// 		t.Error(err, c1)
// 	}

// 	c2, err := ProcuraConta(db, nome02)
// 	if err != nil {
// 		t.Error(err, c2)
// 	}

// 	c3, err := ProcuraConta(db, nome03)
// 	if err != nil {
// 		t.Error(err, c3)
// 	}
// }

// func TestAlteraConta(t *testing.T) {
// 	nome01 := "Ativos Teste 01"
// 	novaConta := conta.GetContaTest()
// 	novaConta.Nome = nome01

// 	c1, err := AlteraConta(db, nome01, novaConta)
// 	if err != nil {
// 		t.Error(err, c1)
// 	}

// 	if c1.Codigo != novaConta.Codigo ||
// 		c1.Comentario != novaConta.Comentario ||
// 		c1.ContaPai != novaConta.ContaPai ||
// 		c1.Estado != novaConta.Estado {
// 		t.Error("Erro na alteração de conta(Codigo ou Comentario ou ContaPai ou Estado)", c1, novaConta)
// 	}
// }

// func TestRemoveConta(t *testing.T) {
// 	nome01 := "Ativos Teste 01"
// 	nome02 := "Ativos Teste 02"
// 	nome03 := "Ativos Teste 03"
// 	nome04 := "Teste Conta 04"
// 	nomeTipoConta01 := "banco teste 01"

// 	err := RemoveConta(db, nome01)
// 	if err != nil {
// 		t.Error(err, nome01)
// 	}

// 	err = RemoveConta(db, nome02)
// 	if err != nil {
// 		t.Error(err, nome02)
// 	}

// 	err = RemoveConta(db, nome03)
// 	if err != nil {
// 		t.Error(err, nome03)
// 	}

// 	err = RemoveConta(db, nome04)
// 	if err != nil {
// 		t.Error(err, nome04)
// 	}

// 	err = RemoveTipoConta(db, nomeTipoConta01)
// 	if err != nil {
// 		t.Error(err, nomeTipoConta01)
// 	}

// }
