package dao

import (
	"testing"

	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/model/detalhe_lancamento"
)

type detLanc struct {
	idLancamento int
	nomeConta    string
}

var (
	detLanc01, detLanc02, detLanc03, detLanc04 detLanc
)

func TestAdicionaDetalheLancamento(t *testing.T) {
	TestAdicionaLancamento(t)
	TestAdicionaConta(t)

	dl1 := detalhe_lancamento.GetDetalheLancamentoCTest()
	dl1.IDLancamento = numLanc01
	dl1.NomeConta = "Ativos Teste 01"

	dl2 := detalhe_lancamento.GetDetalheLancamentoDTest()
	dl2.IDLancamento = numLanc01
	dl2.NomeConta = "Ativos Teste 02"

	dl3, err := AdicionaDetalheLancamento(db, dl1)
	if err != nil {
		t.Error(err, dl3)
	} else {
		detLanc01.idLancamento = dl3.IDLancamento
		detLanc01.nomeConta = dl3.NomeConta
	}

	dl4, err := AdicionaDetalheLancamento(db, dl2)
	if err != nil {
		t.Error(err, dl4)
	} else {
		detLanc02.idLancamento = dl4.IDLancamento
		detLanc02.nomeConta = dl4.NomeConta
	}

	dl5 := detalhe_lancamento.GetDetalheLancamentoDTest()
	dl5.IDLancamento = numLanc02
	dl5.NomeConta = "Ativos Teste 03"

	dl6 := detalhe_lancamento.GetDetalheLancamentoCTest()
	dl6.IDLancamento = numLanc02
	dl6.NomeConta = "Ativos Teste 02"

	dl7, err := AdicionaDetalheLancamento(db, dl5)
	if err != nil {
		t.Error(err, dl7)
	} else {
		detLanc03.idLancamento = dl7.IDLancamento
		detLanc03.nomeConta = dl7.NomeConta
	}

	dl8, err := AdicionaDetalheLancamento(db, dl6)
	if err != nil {
		t.Error(err, dl8)
	} else {
		detLanc04.idLancamento = dl8.IDLancamento
		detLanc04.nomeConta = dl8.NomeConta
	}
}

func TestCarregaDetalheLancamento(t *testing.T) {
	listaDetalhesLancamento, err := CarregaDetalheLancamentos(db)
	if err != nil {
		t.Error(err, listaDetalhesLancamento)
	}

	if len(listaDetalhesLancamento) == 0 {
		t.Error(listaDetalhesLancamento)
	}

	if len(listaDetalhesLancamento) < 4 {
		t.Errorf("%d elementos\n%v", len(listaDetalhesLancamento), listaDetalhesLancamento)
	}

	listaDetalhesLancamento, err = CarregaDetalheLancamentosPorIDLancamento(db, numLanc01)
	if err != nil {
		t.Error(err, listaDetalhesLancamento)
	}

	if len(listaDetalhesLancamento) == 0 {
		t.Error(listaDetalhesLancamento)
	}

	if len(listaDetalhesLancamento) != 2 {
		t.Error(listaDetalhesLancamento)
	}

	listaDetalhesLancamento, err = CarregaDetalheLancamentosPorNomeConta(db, "Ativos Teste 02")
	if err != nil {
		t.Error(err, listaDetalhesLancamento)
	}

	if len(listaDetalhesLancamento) == 0 {
		t.Error(listaDetalhesLancamento)
	}

	if len(listaDetalhesLancamento) != 2 {
		t.Error(listaDetalhesLancamento)
	}
}

func TestProcuraDetalheLancamento(t *testing.T) {
	dl1, err := ProcuraDetalheLancamento(db, detLanc01.idLancamento, detLanc01.nomeConta)
	if err != nil {
		t.Error(err, dl1)
	}

	if dl1.IDLancamento != detLanc01.idLancamento || dl1.NomeConta != detLanc01.nomeConta {
		t.Error("Detalhe lançamento localizado pela função ProduraDetalheLancamento retornou com valores inválidos", detLanc01, dl1)
	}

	dl2, err := ProcuraDetalheLancamento(db, detLanc02.idLancamento, detLanc02.nomeConta)
	if err != nil {
		t.Error(err, dl2)
	}

	dl3, err := ProcuraDetalheLancamento(db, detLanc03.idLancamento, detLanc03.nomeConta)
	if err != nil {
		t.Error(err, dl3)
	}

	dl4, err := ProcuraDetalheLancamento(db, detLanc04.idLancamento, detLanc04.nomeConta)
	if err != nil {
		t.Error(err, dl4)
	}
}

func TestAlteraDetalheLancamento(t *testing.T) {
	novoDetalheLancamento := detalhe_lancamento.GetDetalheLancamentoCTest()
	novoDetalheLancamento.IDLancamento = numLanc02
	novoDetalheLancamento.NomeConta = "Ativos Teste 01"
	novoDetalheLancamento.Credito = 632.25

	dl1, err := AlteraDetalheLancamento(db, detLanc04.idLancamento, detLanc04.nomeConta, novoDetalheLancamento)
	if err != nil {
		t.Error(err, dl1)
	} else {
		if dl1.IDLancamento != novoDetalheLancamento.IDLancamento ||
			dl1.NomeConta != novoDetalheLancamento.NomeConta ||
			dl1.Credito != novoDetalheLancamento.Credito ||
			dl1.Debito != novoDetalheLancamento.Debito {
			t.Error("Erro na alteração de DetalheLancamento(IDLancamento ou NomeConta ou Credito ou Debito)", dl1, novoDetalheLancamento)
		}
	}

}

func TestRemoveDetalheLancamento(t *testing.T) {
	err := RemoveDetalheLancamento(db, detLanc01.idLancamento, detLanc01.nomeConta)
	if err != nil {
		t.Error(err)
	}

	err = RemoveDetalheLancamento(db, detLanc02.idLancamento, detLanc02.nomeConta)
	if err != nil {
		t.Error(err)
	}

	err = RemoveDetalheLancamento(db, detLanc03.idLancamento, detLanc03.nomeConta)
	if err != nil {
		t.Error(err)
	}

	// detLanc4 deve ser removido automaticamente ao excluir os lancamentos, por causa da CONSTRAINT lancamento_detalhe_lancamento_fk em ON DELETE CASCADE
	TestRemoveLancamento(t)
	TestRemoveConta(t)
}
