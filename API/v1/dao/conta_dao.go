package dao

import (
	"controle_pessoal_de_financas/API/v1/model/conta"
	"database/sql"
)

var (
	contaDB = map[string]string{
		"tabela":          "conta",
		"nome":            "nome",
		"nomeTipoConta":   "nome_tipo_conta",
		"codigo":          "codigo",
		"contaPai":        "conta_pai",
		"comentario":      "comentario",
		"dataCriacao":     "data_criacao",
		"dataModificacao": "data_modificacao",
		"estado":          "estado"}
)

func CarregaContas(db *sql.DB) (contas conta.Contas, err error) {
	sql := `
SELECT
	{{.nome}}, {{.nomeTipoConta}}, {{.codigo}}, {{.contaPai}}, {{.comentario}}, {{.dataCriacao}}, {{.dataModificacao}}, {{.estado}}
FROM
	{{.tabela}}
`

	query := getTemplateQuery("CarregaContas", contaDB, sql)

	return carregaContas(db, query)
}

func carregaContas(db *sql.DB, query string, args ...interface{}) (contas conta.Contas, err error) {
	registros, err := carrega(db, query, registrosConta01, args...)

	contas = converteEmConta(registros)

	return
}

func registrosConta01(rows *sql.Rows, registros []interface{}) (novosRegistros []interface{}, err error) {
	contaAtual := new(conta.Conta)
	err = scanConta01(rows, contaAtual)
	if err != nil {
		return
	}
	novosRegistros = append(registros, contaAtual)

	return
}

func scanConta01(rows *sql.Rows, contaAtual *conta.Conta) error {
	var codigo sql.NullString
	var contaPai sql.NullString
	var comentario sql.NullString

	err := rows.Scan(
		&contaAtual.Nome,
		&contaAtual.NomeTipoConta,
		&codigo,
		&contaPai,
		&comentario,
		&contaAtual.DataCriacao,
		&contaAtual.DataModificacao,
		&contaAtual.Estado)

	contaAtual.Codigo = codigo.String
	contaAtual.ContaPai = contaPai.String
	contaAtual.Comentario = comentario.String

	return err
}

func converteEmConta(registros []interface{}) (contas conta.Contas) {
	for _, r := range registros {
		c, ok := r.(*conta.Conta)
		if ok {
			contas = append(contas, c)
		}
	}

	return
}
