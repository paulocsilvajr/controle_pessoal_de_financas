package dao

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/model/tipo_conta"
	"gorm.io/gorm"
)

var (
	tipoContaDB = map[string]string{
		"tabela":           tipo_conta.GetNomeTabelaTipoConta(),
		"nome":             "nome",
		"descricaoDebito":  "descricao_debito",
		"descricaoCredito": "descricao_credito",
		"dataCriacao":      "data_criacao",
		"dataModificacao":  "data_modificacao",
		"estado":           "estado"}
)

// CarregaTiposConta02 retorna uma listagem de todos os tipos de conta(tipo_conta.TiposConta) e erro = nil do BD caso a consulta ocorra corretamente. erro != nil caso ocorra um problema. Deve ser informado uma conexão ao BD(*gorm.DB) como parâmetro obrigatório
func CarregaTiposConta02(db *gorm.DB) (tipo_conta.TiposConta, error) {
	var tTiposConta tipo_conta.TTiposConta
	resultado := db.Find(&tTiposConta)

	return ConverteTTiposContaParaTiposConta(resultado, &tTiposConta)
}

// CarregaTiposContaAtiva02 retorna uma listagem de tipos de conta ativos(tipo_conta.TiposConta) e erro = nil do BD caso a consulta ocorra corretamente. erro != nil caso ocorra um problema. Deve ser informado uma conexão ao BD(*gorm.DB) como parâmetro obrigatório
func CarregaTiposContaAtiva02(db *gorm.DB) (tiposContas tipo_conta.TiposConta, err error) {
	return carregaTiposContaEstado02(db, true)
}

// CarregaTiposContaInativa02 retorna uma listagem de tipos de conta(tipo_conta.TiposConta) no estado inativo e erro = nil do BD caso a consulta ocorra corretamente. erro != nil caso ocorra um problema. Deve ser informado uma conexão ao BD(*gorm.DB) como parâmetro obrigatório
func CarregaTiposContaInativa02(db *gorm.DB) (tiposContas tipo_conta.TiposConta, err error) {
	return carregaTiposContaEstado02(db, false)
}

func carregaTiposContaEstado02(db *gorm.DB, estado bool) (tiposContas tipo_conta.TiposConta, err error) {
	var tTiposConta tipo_conta.TTiposConta
	sql := fmt.Sprintf("%s = ?", tipoContaDB["estado"])
	resultado := db.Where(sql, estado).Find(&tTiposConta)

	return ConverteTTiposContaParaTiposConta(resultado, &tTiposConta)
}

// CarregaTiposConta retorna uma listagem de todos os tipos de conta(tipo_conta.TiposConta) e erro = nil do BD caso a consulta ocorra corretamente. erro != nil caso ocorra um problema. Deve ser informado uma conexão ao BD como parâmetro obrigatório
func CarregaTiposConta(db *sql.DB) (tiposContas tipo_conta.TiposConta, err error) {
	sql := `
SELECT
	{{.nome}}, {{.descricaoDebito}}, {{.descricaoCredito}}, {{.dataCriacao}}, {{.dataModificacao}}, {{.estado}}
FROM
	{{.tabela}}
`
	query := getTemplateQuery("CarregaTiposConta", tipoContaDB, sql)

	return carregaTiposConta(db, query)
}

// CarregaTiposContaAtiva retorna uma listagem de tipos de conta ativos(tipo_conta.TiposConta) e erro = nil do BD caso a consulta ocorra corretamente. erro != nil caso ocorra um problema. Deve ser informado uma conexão ao BD como parâmetro obrigatório
func CarregaTiposContaAtiva(db *sql.DB) (tiposContas tipo_conta.TiposConta, err error) {
	sql := `
SELECT
	{{.nome}}, {{.descricaoDebito}}, {{.descricaoCredito}}, {{.dataCriacao}}, {{.dataModificacao}}, {{.estado}}
FROM
	{{.tabela}}
WHERE
	{{.estado}} = true
`
	query := getTemplateQuery("CarregaTiposContaAtiva", tipoContaDB, sql)

	return carregaTiposConta(db, query)
}

// CarregaTiposContaInativa retorna uma listagem de tipos de conta(tipo_conta.TiposConta) no estado inativo e erro = nil do BD caso a consulta ocorra corretamente. erro != nil caso ocorra um problema. Deve ser informado uma conexão ao BD como parâmetro obrigatório
func CarregaTiposContaInativa(db *sql.DB) (tiposContas tipo_conta.TiposConta, err error) {
	sql := `
SELECT
	{{.nome}}, {{.descricaoDebito}}, {{.descricaoCredito}}, {{.dataCriacao}}, {{.dataModificacao}}, {{.estado}}
FROM
	{{.tabela}}
WHERE
	{{.estado}} = false
`
	query := getTemplateQuery("CarregaTiposContaInativa", tipoContaDB, sql)

	return carregaTiposConta(db, query)
}

// AdicionaTipoConta adiciona um tipo conta ao BD e retorna o tipo conta incluída(*TipoConta) com os dados de acordo como ficou no BD. erro != nil caso ocorra um problema no processo de inclusão. Deve ser informado uma conexão ao BD como parâmetro obrigatório e um tipo conta(*TipoConta)
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

// AdicionaTipoConta02 adiciona um tipo conta ao BD e retorna o tipo conta incluída(*TipoConta) com os dados de acordo como ficou no BD. erro != nil caso ocorra um problema no processo de inclusão. Deve ser informado uma conexão ao BD(*gorm.DB) como parâmetro obrigatório e um tipo conta(*TipoConta)
func AdicionaTipoConta02(db *gorm.DB, novoTipoConta *tipo_conta.TipoConta) (*tipo_conta.TipoConta, error) {
	tc, err := tipo_conta.NewTipoConta(novoTipoConta.Nome, novoTipoConta.DescricaoDebito, novoTipoConta.DescricaoCredito)
	if err != nil {
		return nil, err
	}

	tTipoConta := ConverteTipoContaParaTTipoConta(tc)
	err = db.Create(&tTipoConta).Error
	if err != nil {
		return nil, err
	}
	tipoConta := ConverteTTipoContaParaTipoConta(tTipoConta)

	return tipoConta, nil
}

// ProcuraTipoConta localiza um tipo conta no BD e retorna o tipo conta procurado(*TipoConta). erro != nil caso ocorra um problema no processo de procura. Deve ser informado uma conexão ao BD como parâmetro obrigatório e um NOME do tipo conta desejado
func ProcuraTipoConta(db *sql.DB, nome string) (tc *tipo_conta.TipoConta, err error) {
	sql := `
SELECT
	{{.nome}}, {{.descricaoDebito}}, {{.descricaoCredito}}, {{.dataCriacao}}, {{.dataModificacao}}, {{.estado}}
FROM
	{{.tabela}}
WHERE {{.nome}} = $1
`
	query := getTemplateQuery("ProcuraTipoConta", tipoContaDB, sql)

	tiposConta, err := carregaTiposConta(db, query, nome)
	if len(tiposConta) == 1 {
		tc = tiposConta[0]
	} else {
		err = errors.New("Não foi encontrado um registro com o nome " + nome)
	}

	return
}

// AtivaTipoConta ativa um tipo conta no BD e retorna o tipo conta(*TipoConta) com os dados atualizados. erro != nil caso ocorra um problema no processo de procura. Deve ser informado uma conexão ao BD como parâmetro obrigatório e um NOME do TipoConta desejado
func AtivaTipoConta(db *sql.DB, nome string) (tc *tipo_conta.TipoConta, err error) {
	tipoContaBanco, err := ProcuraTipoConta(db, nome)
	if err != nil {
		return
	}

	tipoContaBanco.Ativa()

	sql := `
UPDATE {{.tabela}}
SET {{.estado}} = $1, {{.dataModificacao}} = $2
WHERE {{.nome}} = $3
`

	query := getTemplateQuery("AtivaTipoConta", tipoContaDB, sql)

	return estadoTipoConta(db, tipoContaBanco, query, nome)
}

// InativaTipoConta inativa um tipo conta no BD e retorna o tipo conta(*TipoConta) com os dados atualizados. erro != nil caso ocorra um problema no processo de procura. Deve ser informado uma conexão ao BD como parâmetro obrigatório e um NOME do TipoConta desejado
func InativaTipoConta(db *sql.DB, nome string) (tc *tipo_conta.TipoConta, err error) {
	tipoContaBanco, err := ProcuraTipoConta(db, nome)
	if err != nil {
		return
	}

	tipoContaBanco.Inativa()

	sql := `
UPDATE {{.tabela}}
SET {{.estado}} = $1, {{.dataModificacao}} = $2
WHERE {{.nome}} = $3
`

	query := getTemplateQuery("InativaTipoConta", tipoContaDB, sql)

	return estadoTipoConta(db, tipoContaBanco, query, nome)
}

// AlteraTipoConta altera um tipo conta com o nome(string) informado a partir dos dados do *TipoConta informado no parâmetro tipoContaAlteracao. O campo Estado não é alterado, enquanto que o campo Nome sim. Use a função específica para essa tarefa(estado). Retorna um *TipoConta alterado no BD e um error. error != nil caso ocorra um problema.
func AlteraTipoConta(db *sql.DB, nome string, tipoContaAlteracao *tipo_conta.TipoConta) (tc *tipo_conta.TipoConta, err error) {
	tipoContaBanco, err := ProcuraTipoConta(db, nome)
	if err != nil {
		return
	}

	err = tipoContaBanco.Altera(tipoContaAlteracao.Nome, tipoContaAlteracao.DescricaoDebito, tipoContaAlteracao.DescricaoCredito)
	if err != nil {
		return
	}

	sql := `
UPDATE {{.tabela}}
SET {{.nome}} = $1, {{.descricaoDebito}} = $2, {{.descricaoCredito}} = $3, {{.dataModificacao}} = $4
WHERE {{.nome}} = $5
`

	query := getTemplateQuery("AlteraTipoConta", tipoContaDB, sql)

	return alteraTipoConta(db, tipoContaBanco, query, nome)
}

// RemoveTipoConta remove um tipo conta do BD e retorna erro != nil caso ocorra um problema no processo de remoção. Deve ser informado uma conexão ao BD como parâmetro obrigatório e uma string contendo o NOME do tipo conta desejado
func RemoveTipoConta(db *sql.DB, nome string) (err error) {
	sql := `
DELETE FROM
	{{.tabela}}
WHERE {{.nome}} = $1
`
	query := getTemplateQuery("RemoveTipoConta", tipoContaDB, sql)

	tc, err := ProcuraTipoConta(db, nome)
	if tc != nil {
		err = remove(db, tc.Nome, query)
	}

	return
}

func carregaTiposConta(db *sql.DB, query string, args ...interface{}) (tiposConta tipo_conta.TiposConta, err error) {
	registros, err := carrega(db, query, registrosTipoConta01, args...)

	tiposConta = converteEmTiposConta(registros)

	return
}

func adicionaTipoConta(db *sql.DB, novoTipoConta *tipo_conta.TipoConta, query string) (tc *tipo_conta.TipoConta, err error) {
	resultado, err := adiciona(db, novoTipoConta, query, setValoresTipoConta01)
	tipoContaTemp, ok := resultado.(*tipo_conta.TipoConta)
	if ok {
		tc = tipoContaTemp
	}
	return
}

func alteraTipoConta(db *sql.DB, tipoContaBanco *tipo_conta.TipoConta, query, chave string) (tc *tipo_conta.TipoConta, err error) {
	resultado, err := altera(db, tipoContaBanco, query, setValoresTipoConta03, chave)
	tipoContaTemp, ok := resultado.(*tipo_conta.TipoConta)
	if ok {
		tc = tipoContaTemp
		tc.DataCriacao = tc.DataCriacao.Local()
	}
	return
}

func estadoTipoConta(db *sql.DB, tipoContaBanco *tipo_conta.TipoConta, query, chave string) (tc *tipo_conta.TipoConta, err error) {
	resultado, err := altera(db, tipoContaBanco, query, setValoresTipoConta02, chave)
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

func setValoresTipoConta02(stmt *sql.Stmt, novoRegistro interface{}, chave string) (r sql.Result, err error) {
	novoTipoConta, ok := novoRegistro.(*tipo_conta.TipoConta)

	if ok {
		r, err = stmt.Exec(
			novoTipoConta.Estado,
			novoTipoConta.DataModificacao,
			chave)
	}
	return
}

func setValoresTipoConta03(stmt *sql.Stmt, novoRegistro interface{}, chave string) (r sql.Result, err error) {

	novoTipoConta, ok := novoRegistro.(*tipo_conta.TipoConta)

	if ok {
		r, err = stmt.Exec(
			novoTipoConta.Nome,
			novoTipoConta.DescricaoDebito,
			novoTipoConta.DescricaoCredito,
			novoTipoConta.DataModificacao,
			chave)
	}
	return
}

func registrosTipoConta01(rows *sql.Rows, registros []interface{}) (novosRegistros []interface{}, err error) {
	tipoContaAtual := new(tipo_conta.TipoConta)
	err = scanTipoConta01(rows, tipoContaAtual)
	if err != nil {
		return
	}
	tipoContaAtual.CorrigeData()

	novosRegistros = append(registros, tipoContaAtual)

	return
}

func scanTipoConta01(rows *sql.Rows, tipoContaAtual *tipo_conta.TipoConta) error {
	return rows.Scan(
		&tipoContaAtual.Nome,
		&tipoContaAtual.DescricaoDebito,
		&tipoContaAtual.DescricaoCredito,
		&tipoContaAtual.DataCriacao,
		&tipoContaAtual.DataModificacao,
		&tipoContaAtual.Estado)
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
