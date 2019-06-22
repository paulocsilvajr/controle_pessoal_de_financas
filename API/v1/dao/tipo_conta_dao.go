package dao

import (
	"controle_pessoal_de_financas/API/v1/model/tipo_conta"
	"database/sql"
)

var (
	tipoContaDB = map[string]string{
		"tabela":           "tipo_conta",
		"nome":             "nome",
		"descricaoDebito":  "descricao_debito",
		"descricaoCredito": "descricao_credito",
		"dataCriacao":      "data_criacao",
		"dataModificacao":  "data_modificacao",
		"estado":           "estado"}
)

func CarregaTiposConta(db *sql.DB) (tiposContas tipo_conta.TiposConta, err error) {
	sql := `
SELECT
	{{.nome}}, {{.descricaoDebito}}, {{.descricaoCredito}}, {{.dataCriacao}}, {{.dataModificacao}}
FROM
	{{.tabela}}
WHERE
	{{.estado}} = true
`
	query := getTemplateQuery("CarregaTipoConta", tipoContaDB, sql)

	return carregaTipoConta(db, query)
}

func CarregaTiposContaInativa(db *sql.DB) (tiposContas tipo_conta.TiposConta, err error) {
	sql := `
SELECT
	{{.nome}}, {{.descricaoDebito}}, {{.descricaoCredito}}, {{.dataCriacao}}, {{.dataModificacao}}
FROM
	{{.tabela}}
WHERE
	{{.estado}} = false
`
	query := getTemplateQuery("CarregaTipoContaInativa", tipoContaDB, sql)

	return carregaTipoConta(db, query)
}

func AdicionaTipoConta(db *sql.DB, novoTipoConta *tipo_conta.TipoConta) (tc *tipo_conta.TipoConta, err error) {
	tc, err = tipo_conta.NewTipoConta(novoTipoConta.Nome, novoTipoConta.DescricaoDebito, novoTipoConta.DescricaoCredito)
	if err != nil {
		return
	}

	sql := `
INSERT INTO {{.tabela}}(
	{{.nome}}, {{.descricaoDebito}}, {{.descricaoCredito}}, {{.dataCriacao}}, {{.dataModificacao}}, {{.estado}})
VALUES($1, $2, $3, $4, $5, $6)
`
	query := getTemplateQuery("AdicionaTipoConta", tipoContaDB, sql)

	return adicionaTipoConta(db, tc, query)
}

func adicionaTipoConta(db *sql.DB, novoTipoConta *tipo_conta.TipoConta, query string) (tc *tipo_conta.TipoConta, err error) {
	resultado, err := adiciona(db, novoTipoConta, query, setValoresTipoConta01)
	tipoContaTemp, ok := resultado.(*tipo_conta.TipoConta)
	if ok {
		tc = tipoContaTemp
	}
	return
}

func setValoresTipoConta01(stmt *sql.Stmt, novoRegistro interface{}) (r sql.Result, err error) {

	novoTipoConta, ok := novoRegistro.(*tipo_conta.TipoConta)

	if ok {
		r, err = stmt.Exec(
			novoTipoConta.Nome,
			novoTipoConta.DescricaoDebito,
			novoTipoConta.DescricaoCredito,
			novoTipoConta.DataCriacao,
			novoTipoConta.DataModificacao,
			novoTipoConta.Estado)
	}
	return
}

func carregaTipoConta(db *sql.DB, query string, args ...interface{}) (tiposConta tipo_conta.TiposConta, err error) {
	registros, err := carrega(db, query, registrosTipoConta01, args...)

	tiposConta = converteEmTiposConta(registros)

	return
}

func registrosTipoConta01(rows *sql.Rows, registros []interface{}) (novosRegistros []interface{}, err error) {
	tipoContaAtual := new(tipo_conta.TipoConta)
	err = scanTipoConta01(rows, tipoContaAtual)
	if err != nil {
		return
	}
	novosRegistros = append(registros, tipoContaAtual)

	return
}

func scanTipoConta01(rows *sql.Rows, tipoContaAtual *tipo_conta.TipoConta) error {
	return rows.Scan(
		&tipoContaAtual.Nome,
		&tipoContaAtual.DescricaoDebito,
		&tipoContaAtual.DescricaoCredito,
		&tipoContaAtual.DataCriacao,
		&tipoContaAtual.DataModificacao)
}

func converteEmTiposConta(registros []interface{}) (tiposConta tipo_conta.TiposConta) {
	for _, r := range registros {
		tc, ok := r.(*tipo_conta.TipoConta)
		if ok {
			tiposConta = append(tiposConta, tc)
		}
	}

	return
}
