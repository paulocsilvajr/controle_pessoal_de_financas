package dao

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/model/detalhe_lancamento"
	"gorm.io/gorm"
)

var (
	detalheLancamentoDB = map[string]string{
		"tabela":           detalhe_lancamento.GetNomeTabelaDetalheLancamento(),
		"idLancamento":     "id_lancamento",
		"nomeConta":        "nome_conta",
		"debito":           "debito",
		"credito":          "credito",
		"tabelaConta":      contaDB["tabela"],
		"fkConta":          contaDB["nome"],
		"tabelaLancamento": lancamentoDB["tabela"],
		"fkLancamento":     lancamentoDB["id"],
	}
)

// AdicionaDetalheLancamento02 adiciona um detalhe lancamento ao BD e retorna o detalhe lancamento incluída(*DetalheLancamento) com os dados de acordo como ficou no BD. erro != nil caso ocorra um problema no processo de inclusão. Deve ser informado uma conexão ao BD(*gorm.DB) como parâmetro obrigatório e um detalhe lancamento(*DetalheLancamento)
func AdicionaDetalheLancamento02(db *gorm.DB, novoDetalheLancamento *detalhe_lancamento.DetalheLancamento) (*detalhe_lancamento.DetalheLancamento, error) {
	dl, err := detalhe_lancamento.NewDetalheLancamento(novoDetalheLancamento.IDLancamento, novoDetalheLancamento.NomeConta, novoDetalheLancamento.Debito, novoDetalheLancamento.Credito)
	if err != nil {
		return nil, err
	}

	tDetLancamento := ConverteDetalheLancamentoParaTDetalheLancamento(dl)
	err = db.Create(&tDetLancamento).Error
	if err != nil {
		return nil, err
	}

	return ConverteTDetalheLancamentoParaDetalheLancamento(tDetLancamento), nil
}

// RemoveDetalheLancamento02 remove um detalhe lancamento do BD e retorna erro != nil caso ocorra um problema no processo de remoção. Deve ser informado uma conexão ao BD(*gorm.DB) como parâmetro obrigatório e um int contendo o IDLancamento e um string contendo o NomeConta desejado
func RemoveDetalheLancamento02(db *gorm.DB, idLancamento int, nomeConta string) (err error) {
	dl := &detalhe_lancamento.TDetalheLancamento{
		IDLancamento: idLancamento,
		NomeConta:    nomeConta,
	}

	tx := db.Delete(dl)
	if err := tx.Error; err != nil {
		return err
	}

	linhaAfetadas := tx.RowsAffected
	var esperado int64 = 1
	if linhaAfetadas != esperado {
		return fmt.Errorf("remoção de detalhe lançamento com ID %d e nome conta '%s' retornou uma quantidade de registros afetados incorreto. Esperado: %d, obtido: %d", idLancamento, nomeConta, esperado, linhaAfetadas)
	}

	return nil
}

// CarregaDetalheLancamentos retorna uma listagem de todos os detalhe lancamentos(detalhe_lancamento.detalheLancamentos) e erro = nil do BD caso a consulta ocorra corretamente. erro != nil caso ocorra um problema. Deve ser informado uma conexão ao BD como parâmetro obrigatório
func CarregaDetalheLancamentos(db *sql.DB) (detalheLancamentos detalhe_lancamento.DetalheLancamentos, err error) {
	sql := `
SELECT
	{{.idLancamento}}, {{.nomeConta}}, {{.debito}}, {{.credito}}
FROM
	{{.tabela}}
`

	query := getTemplateQuery("CarregaDetalheLancamentos", detalheLancamentoDB, sql)

	return carregaDetalheLancamentos(db, query)
}

// CarregaDetalheLancamentosPorIDLancamento retorna uma listagem de todos os detalhe lancamentos(detalhe_lancamento.detalheLancamentos) ref ao id de lancamento informado e erro = nil do BD caso a consulta ocorra corretamente. erro != nil caso ocorra um problema. Deve ser informado uma conexão ao BD e um idLancamento(int) como parâmetro obrigatório
func CarregaDetalheLancamentosPorIDLancamento(db *sql.DB, idLancamento int) (detalheLancamentos detalhe_lancamento.DetalheLancamentos, err error) {
	sql := `
SELECT
	{{.idLancamento}}, {{.nomeConta}}, {{.debito}}, {{.credito}}
FROM
	{{.tabela}}
WHERE {{.idLancamento}} = $1
`

	query := getTemplateQuery("CarregaDetalheLancamentosPorIdLancamento", detalheLancamentoDB, sql)

	return carregaDetalheLancamentos(db, query, idLancamento)
}

// CarregaDetalheLancamentosPorNomeConta retorna uma listagem de todos os detalhe lancamentos(detalhe_lancamento.detalheLancamentos) ref ao nome de conta informado e erro = nil do BD caso a consulta ocorra corretamente. erro != nil caso ocorra um problema. Deve ser informado uma conexão ao BD e um nomeConta(string) como parâmetro obrigatório
func CarregaDetalheLancamentosPorNomeConta(db *sql.DB, nomeConta string) (detalheLancamentos detalhe_lancamento.DetalheLancamentos, err error) {
	sql := `
SELECT
	{{.idLancamento}}, {{.nomeConta}}, {{.debito}}, {{.credito}}
FROM
	{{.tabela}}
WHERE {{.nomeConta}} = $1
`

	query := getTemplateQuery("CarregaDetalheLancamentosPorNomeConta", detalheLancamentoDB, sql)

	return carregaDetalheLancamentos(db, query, nomeConta)
}

// AdicionaDetalheLancamento adiciona um detalhe lancamento ao BD e retorna o detalhe lancamento incluída(*DetalheLancamento) com os dados de acordo como ficou no BD. erro != nil caso ocorra um problema no processo de inclusão. Deve ser informado uma conexão ao BD como parâmetro obrigatório e um detalhe lancamento(*DetalheLancamento)
func AdicionaDetalheLancamento(db *sql.DB, novoDetalheLancamento *detalhe_lancamento.DetalheLancamento) (dl *detalhe_lancamento.DetalheLancamento, err error) {
	dl, err = detalhe_lancamento.NewDetalheLancamento(novoDetalheLancamento.IDLancamento, novoDetalheLancamento.NomeConta, novoDetalheLancamento.Debito, novoDetalheLancamento.Credito)
	if err != nil {
		return
	}

	sql := `
INSERT INTO {{.tabela}}(
	{{.idLancamento}}, {{.nomeConta}}, {{.debito}}, {{.credito}})
VALUES($1, $2, $3, $4)
`
	query := getTemplateQuery("AdicionaDetalheLancamento", detalheLancamentoDB, sql)

	return adicionaDetalheLancamento(db, dl, query)
}

// ProcuraDetalheLancamento localiza um detalhe lancamento no BD e retorna o detalhe lancamento procurado(*DetalheLancamento). erro != nil caso ocorra um problema no processo de procura. Deve ser informado uma conexão ao BD como parâmetro obrigatório e o ID e NomeConta do detalhe lancamento desejado
func ProcuraDetalheLancamento(db *sql.DB, idLancamento int, nomeConta string) (dl *detalhe_lancamento.DetalheLancamento, err error) {
	sql := `
SELECT
	{{.idLancamento}}, {{.nomeConta}}, {{.debito}}, {{.credito}}
FROM
	{{.tabela}}
WHERE {{.idLancamento}} = $1 AND {{.nomeConta}} = $2
`
	query := getTemplateQuery("ProcuraDetalheLancamento", detalheLancamentoDB, sql)

	detalheLancamentos, err := carregaDetalheLancamentos(db, query, idLancamento, nomeConta)
	if len(detalheLancamentos) == 1 {
		dl = detalheLancamentos[0]
	} else {
		err = fmt.Errorf("não foi encontrado um registro com o ID %d e o NomeConta %s", idLancamento, nomeConta)
	}

	return
}

// AlteraDetalheLancamento altera um detalhe lancamento com o IDLancamento(int) e NomeConta(string) informado a partir dos dados do *DetalheLancamento informado no parâmetro detalheLancamentoAlteracao. O IDLancamento não é alterado. Retorna um *DetalheLancamento alterado no BD e um error. error != nil caso ocorra um problema.
func AlteraDetalheLancamento(db *sql.DB, idLancamento int, nomeConta string, detalheLancamentoAlteracao *detalhe_lancamento.DetalheLancamento) (dl *detalhe_lancamento.DetalheLancamento, err error) {
	detalheLancamentoBanco, err := ProcuraDetalheLancamento(db, idLancamento, nomeConta)

	if err != nil {
		return
	}

	err = detalheLancamentoBanco.Altera(detalheLancamentoAlteracao.NomeConta, detalheLancamentoAlteracao.Debito, detalheLancamentoAlteracao.Credito)
	if err != nil {
		return
	}

	sql := `
UPDATE {{.tabela}}
SET {{.nomeConta}} = $1, {{.debito}} = $2, {{.credito}} = $3
WHERE {{.idLancamento}} = $4 AND {{.nomeConta}} = $5
`

	query := getTemplateQuery("AlteraLancamento", detalheLancamentoDB, sql)

	return alteraDetalheLancamento(db, detalheLancamentoBanco, query, idLancamento, nomeConta)
}

// AlteraDetalheLancamento2 altera um detalhe lancamento com o IDLancamento(int) e NomeConta(string) informado a partir dos dados do *DetalheLancamento informado no parâmetro detalheLancamentoAlteracao. O IDLancamento não é alterado. Retorna um *DetalheLancamento alterado no BD e um error. error != nil caso ocorra um problema.
func AlteraDetalheLancamento2(db *sql.DB, transacao *sql.Tx, idLancamento int, nomeConta string, detalheLancamentoAlteracao *detalhe_lancamento.DetalheLancamento) (dl *detalhe_lancamento.DetalheLancamento, err error) {
	detalheLancamentoBanco, err := ProcuraDetalheLancamento(db, idLancamento, nomeConta)

	if err != nil {
		return
	}

	err = detalheLancamentoBanco.Altera(detalheLancamentoAlteracao.NomeConta, detalheLancamentoAlteracao.Debito, detalheLancamentoAlteracao.Credito)
	if err != nil {
		return
	}

	sql := `
UPDATE {{.tabela}}
SET {{.nomeConta}} = $1, {{.debito}} = $2, {{.credito}} = $3
WHERE {{.idLancamento}} = $4 AND {{.nomeConta}} = $5
`

	query := getTemplateQuery("AlteraLancamento", detalheLancamentoDB, sql)

	return alteraDetalheLancamento2(transacao, detalheLancamentoBanco, query, idLancamento, nomeConta)
}

// RemoveDetalheLancamento remove um detalhe lancamento do BD e retorna erro != nil caso ocorra um problema no processo de remoção. Deve ser informado uma conexão ao BD como parâmetro obrigatório e um int contendo o IDLancamento e um string contendo o NomeConta desejado
func RemoveDetalheLancamento(db *sql.DB, idLancamento int, nomeConta string) (err error) {
	sql := `
DELETE FROM
	{{.tabela}}
WHERE {{.idLancamento}} = $1 AND {{.nomeConta}} = $2
`

	query := getTemplateQuery("RemoveDetalheLancamento", detalheLancamentoDB, sql)

	dl, err := ProcuraDetalheLancamento(db, idLancamento, nomeConta)
	if dl != nil {
		err = remove2(db, query, dl.IDLancamento, dl.NomeConta)
	}

	return
}

func alteraDetalheLancamento(db *sql.DB, detalheLancamentoBanco *detalhe_lancamento.DetalheLancamento, query string, chave1 int, chave2 string) (dl *detalhe_lancamento.DetalheLancamento, err error) {
	chave1String := strconv.Itoa(chave1)

	resultado, err := altera2(db, detalheLancamentoBanco, query, setValoresDetalheLancamento03, chave1String, chave2)
	detalheLancamentoTemp, ok := resultado.(*detalhe_lancamento.DetalheLancamento)
	if ok {
		dl = detalheLancamentoTemp
	}
	return
}

func alteraDetalheLancamento2(transacao *sql.Tx, detalheLancamentoBanco *detalhe_lancamento.DetalheLancamento, query string, chave1 int, chave2 string) (dl *detalhe_lancamento.DetalheLancamento, err error) {
	chave1String := strconv.Itoa(chave1)

	resultado, err := altera2T(transacao, detalheLancamentoBanco, query, setValoresDetalheLancamento03, chave1String, chave2)
	detalheLancamentoTemp, ok := resultado.(*detalhe_lancamento.DetalheLancamento)
	if ok {
		dl = detalheLancamentoTemp
	}
	return
}

func setValoresDetalheLancamento03(stmt *sql.Stmt, novoRegistro interface{}, chave ...interface{}) (r sql.Result, err error) {
	novoDetalheLancamento, ok := novoRegistro.(*detalhe_lancamento.DetalheLancamento)

	debito, credito := getCamposConvertidosDetalheLancamento(novoDetalheLancamento)

	if ok {
		r, err = stmt.Exec(
			novoDetalheLancamento.NomeConta,
			debito,
			credito,
			chave[0],
			chave[1])
	}
	return
}

func adicionaDetalheLancamento(db *sql.DB, novoDetalheLancamento *detalhe_lancamento.DetalheLancamento, query string) (dl *detalhe_lancamento.DetalheLancamento, err error) {
	resultado, err := adiciona(db, novoDetalheLancamento, query, setValoresDetalheLancamento01)
	detalheLancamentoTemp, ok := resultado.(*detalhe_lancamento.DetalheLancamento)
	if ok {
		dl = detalheLancamentoTemp
	}
	return
}

func setValoresDetalheLancamento01(stmt *sql.Stmt, novoRegistro interface{}) (r sql.Result, err error) {
	novoDetalheLancamento, ok := novoRegistro.(*detalhe_lancamento.DetalheLancamento)

	debito, credito := getCamposConvertidosDetalheLancamento(novoDetalheLancamento)

	if ok {
		r, err = stmt.Exec(
			novoDetalheLancamento.IDLancamento,
			novoDetalheLancamento.NomeConta,
			debito,
			credito)
	}

	return
}

func getCamposConvertidosDetalheLancamento(novoDetalheLancamento *detalhe_lancamento.DetalheLancamento) (debito, credito sql.NullFloat64) {
	temDebito := novoDetalheLancamento.Debito != 0
	debito = sql.NullFloat64{Float64: novoDetalheLancamento.Debito, Valid: temDebito}

	temCredito := novoDetalheLancamento.Credito != 0
	credito = sql.NullFloat64{Float64: novoDetalheLancamento.Credito, Valid: temCredito}

	return
}

func carregaDetalheLancamentos(db *sql.DB, query string, args ...interface{}) (detalheLancamentos detalhe_lancamento.DetalheLancamentos, err error) {
	registros, err := carrega(db, query, registrosDetalheLancamento01, args...)

	detalheLancamentos = converterEmDetalheLancamento(registros)

	return
}

func registrosDetalheLancamento01(rows *sql.Rows, registros []interface{}) (novosRegistros []interface{}, err error) {
	detalheLancamentoAtual := new(detalhe_lancamento.DetalheLancamento)
	err = scanDetalheLancamento01(rows, detalheLancamentoAtual)
	if err != nil {
		return
	}
	novosRegistros = append(registros, detalheLancamentoAtual)

	return
}

func scanDetalheLancamento01(rows *sql.Rows, detalheLancamentoAtual *detalhe_lancamento.DetalheLancamento) error {
	var debito, credito sql.NullFloat64 // debito e crédito são campos na tabela detalhe_lancamento do DB que pode ser nulo

	err := rows.Scan(
		&detalheLancamentoAtual.IDLancamento,
		&detalheLancamentoAtual.NomeConta,
		&debito,
		&credito)
	detalheLancamentoAtual.Debito = debito.Float64
	detalheLancamentoAtual.Credito = credito.Float64

	return err
}

func converterEmDetalheLancamento(registros []interface{}) (detalheLancamentos detalhe_lancamento.DetalheLancamentos) {
	for _, r := range registros {
		dl, ok := r.(*detalhe_lancamento.DetalheLancamento)
		if ok {
			detalheLancamentos = append(detalheLancamentos, dl)
		}
	}

	return
}
