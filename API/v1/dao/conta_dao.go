package dao

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/model/conta"
	"gorm.io/gorm"
)

var (
	contaDB = map[string]string{
		"tabela":          conta.GetNomeTabelaConta(),
		"nome":            "nome",
		"nomeTipoConta":   "nome_tipo_conta",
		"codigo":          "codigo",
		"contaPai":        "conta_pai",
		"comentario":      "comentario",
		"dataCriacao":     "data_criacao",
		"dataModificacao": "data_modificacao",
		"estado":          "estado",
		"tabelaTipoConta": tipoContaDB["tabela"],
		"fkTipoConta":     tipoContaDB["nome"],
	}
)

func GetPrimaryKeyConta() string {
	return contaDB["nome"]
}

// AdicionaConta02 adiciona uma conta ao BD e retorna a conta incluída(*Conta) com os dados de acordo como ficou no BD. erro != nil caso ocorra um problema no processo de inclusão. Deve ser informado uma conexão ao BD(*gorm.DB) como parâmetro obrigatório e uma conta(*Conta)
func AdicionaConta02(db *gorm.DB, novaConta *conta.Conta) (*conta.Conta, error) {
	c, err := conta.NewConta(novaConta.Nome, novaConta.NomeTipoConta, novaConta.Codigo, novaConta.ContaPai, novaConta.Comentario)
	if err != nil {
		return nil, err
	}

	tc := ConverteContaParaTConta(c)
	err = db.Create(&tc).Error
	if err != nil {
		return nil, err
	}

	return ConverteTContaParaConta(tc), nil
}

// RemoveConta02 remove uma conta do BD e retorna erro != nil caso ocorra um problema no processo de remoção. Deve ser informado uma conexão ao BD(*gorm.DB) como parâmetro obrigatório e uma string contendo o NOME da conta desejado
func RemoveConta02(db *gorm.DB, nome string) error {
	c := &conta.TConta{Nome: nome}

	tx := db.Delete(c)
	if err := tx.Error; err != nil {
		return err
	}

	linhaAfetadas := tx.RowsAffected
	var esperado int64 = 1
	if linhaAfetadas != esperado {
		return fmt.Errorf("remoção de conta com nome '%s' retornou uma quantidade de registros afetados incorreto. Esperado: %d, retorno: %d", nome, esperado, linhaAfetadas)
	}

	return nil
}

// ProcuraConta02 localiza uma conta no BD e retorna a conta procurada(*Conta). erro != nil caso ocorra um problema no processo de procura. Deve ser informado uma conexão ao BD(*gorm.DB) como parâmetro obrigatório e o NOME da conta desejada
func ProcuraConta02(db *gorm.DB, nome string) (*conta.Conta, error) {
	tc := new(conta.TConta)

	sql := getTemplateSQL(
		"ProcuraConta02",
		"{{.nome}} = ?",
		contaDB,
	)
	tx := db.Where(sql, nome).First(&tc)
	if err := tx.Error; err != nil {
		return nil, err
	}

	return ConverteTContaParaConta(tc), nil
}

// AlteraConta02 altera uma conta com o nome(string) informado a partir dos dados da *Conta informada no parâmetro contaAlteracao. O campo Estado não é alterado, enquanto que o campo Nome sim. Use a função específica para essa tarefa(estado). Retorna uma *Conta alterada no BD(*gorm.DB) e um error. error != nil caso ocorra um problema.
func AlteraConta02(db *gorm.DB, nome string, contaAlteracao *conta.Conta) (*conta.Conta, error) {
	contaBanco, err := ProcuraConta02(db, nome)
	if err != nil {
		return nil, err
	}

	err = contaBanco.Altera(contaAlteracao.Nome, contaAlteracao.NomeTipoConta, contaAlteracao.Codigo, contaAlteracao.ContaPai, contaAlteracao.Comentario)
	if err != nil {
		return nil, err
	}

	tc := ConverteContaParaTConta(contaBanco)
	var tx *gorm.DB
	if nome != tc.Nome {
		sql := getTemplateSQL(
			"AlteraConta02",
			"{{.nome}}",
			contaDB,
		)
		novoNome := tc.Nome
		tc.Nome = nome
		tx = db.Model(&tc).Update(sql, novoNome)
	} else {
		tx = db.Save(&tc)
	}

	linhaAfetadas := tx.RowsAffected
	var esperado int64 = 1
	if linhaAfetadas != esperado {
		return nil, fmt.Errorf("alteração de conta com nome '%s' retornou uma quantidade de registros afetados incorreto. Esperado: %d, obtido: %d", nome, esperado, linhaAfetadas)
	}

	if err := tx.Error; err != nil {
		return nil, err
	}

	return ConverteTContaParaConta(tc), nil
}

// AtivaConta02 ativa uma conta no BD e retorna a conta(*Conta) com os dados atualizados. erro != nil caso ocorra um problema no processo de procura. Deve ser informado uma conexão ao BD(*gorm.DB) como parâmetro obrigatório e um NOME(string) da Conta desejada
func AtivaConta02(db *gorm.DB, nome string) (*conta.Conta, error) {
	contaBanco, err := ProcuraConta02(db, nome)
	if err != nil {
		return nil, err
	}

	contaBanco.Ativa()

	tc := ConverteContaParaTConta(contaBanco)
	tx := db.Save(&tc)

	linhaAfetadas := tx.RowsAffected
	var esperado int64 = 1
	if linhaAfetadas != esperado {
		return nil, fmt.Errorf("ativação de conta com nome '%s' retornou uma quantidade de registros afetados incorreto. Esperado: %d, retorno: %d", nome, esperado, linhaAfetadas)
	}

	if err := tx.Error; err != nil {
		return nil, err
	}

	return ConverteTContaParaConta(tc), nil
}

// InativaConta02 inativa uma conta no BD e retorna a conta(*Conta) com os dados atualizados. erro != nil caso ocorra um problema no processo de procura. Deve ser informado uma conexão ao BD(*gorm.DB) como parâmetro obrigatório e um NOME(string) da Conta desejada
func InativaConta02(db *gorm.DB, nome string) (*conta.Conta, error) {
	contaBanco, err := ProcuraConta02(db, nome)
	if err != nil {
		return nil, err
	}

	contaBanco.Inativa()

	tc := ConverteContaParaTConta(contaBanco)
	tx := db.Save(&tc)

	linhaAfetadas := tx.RowsAffected
	var esperado int64 = 1
	if linhaAfetadas != esperado {
		return nil, fmt.Errorf("inativação de conta com nome '%s' retornou uma quantidade de registros afetados incorreto. Esperado: %d, retorno: %d", nome, esperado, linhaAfetadas)
	}

	if err := tx.Error; err != nil {
		return nil, err
	}

	return ConverteTContaParaConta(tc), nil
}

// CarregaContas02 retorna uma listagem de todos as contas(conta.Conta) e erro = nil do BD caso a consulta ocorra corretamente. erro != nil caso ocorra um problema. Deve ser informado uma conexão ao BD(*gorm.DB) como parâmetro obrigatório
func CarregaContas02(db *gorm.DB) (conta.Contas, error) {
	var tContas conta.TContas
	resultado := db.Find(&tContas)

	return ConverteTContasParaContas(resultado, &tContas)
}

// CarregaContasAtiva02 retorna uma listagem de contas ativas(conta.Conta) e erro = nil do BD caso a consulta ocorra corretamente. erro != nil caso ocorra um problema. Deve ser informado uma conexão ao BD(*gorm.DB) como parâmetro obrigatório
func CarregaContasAtiva02(db *gorm.DB) (conta.Contas, error) {
	var tContas conta.TContas
	sql := getTemplateSQL(
		"CarregaContasAtiva02",
		"{{.estado}} = true",
		contaDB,
	)
	resultado := db.Where(sql).Find(&tContas)

	return ConverteTContasParaContas(resultado, &tContas)
}

// CarregaContasInativa02 retorna uma listagem de contas inativas(conta.Conta) e erro = nil do BD caso a consulta ocorra corretamente. erro != nil caso ocorra um problema. Deve ser informado uma conexão ao BD(*gorm.DB) como parâmetro obrigatório
func CarregaContasInativa02(db *gorm.DB) (conta.Contas, error) {
	var tContas conta.TContas
	sql := getTemplateSQL(
		"CarregaContasAtiva02",
		"{{.estado}} = false",
		contaDB,
	)
	resultado := db.Where(sql).Find(&tContas)

	return ConverteTContasParaContas(resultado, &tContas)
}

// CarregaContas retorna uma listagem de todos as contas(conta.Conta) e erro = nil do BD caso a consulta ocorra corretamente. erro != nil caso ocorra um problema. Deve ser informado uma conexão ao BD como parâmetro obrigatório
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

// CarregaContasAtiva retorna uma listagem de contas ativas(conta.Conta) e erro = nil do BD caso a consulta ocorra corretamente. erro != nil caso ocorra um problema. Deve ser informado uma conexão ao BD como parâmetro obrigatório
func CarregaContasAtiva(db *sql.DB) (contas conta.Contas, err error) {
	sql := `
SELECT
	{{.nome}}, {{.nomeTipoConta}}, {{.codigo}}, {{.contaPai}}, {{.comentario}}, {{.dataCriacao}}, {{.dataModificacao}}, {{.estado}}
FROM
	{{.tabela}}
WHERE
	{{.estado}} = true
`

	query := getTemplateQuery("CarregaContasAtiva", contaDB, sql)

	return carregaContas(db, query)
}

// CarregaContasInativa retorna uma listagem de contas inativas(conta.Conta) e erro = nil do BD caso a consulta ocorra corretamente. erro != nil caso ocorra um problema. Deve ser informado uma conexão ao BD como parâmetro obrigatório
func CarregaContasInativa(db *sql.DB) (contas conta.Contas, err error) {
	sql := `
SELECT
	{{.nome}}, {{.nomeTipoConta}}, {{.codigo}}, {{.contaPai}}, {{.comentario}}, {{.dataCriacao}}, {{.dataModificacao}}, {{.estado}}
FROM
	{{.tabela}}
WHERE
	{{.estado}} = false
`

	query := getTemplateQuery("CarregaContasInativa", contaDB, sql)

	return carregaContas(db, query)
}

// AdicionaConta adiciona uma conta ao BD e retorna a conta incluída(*Conta) com os dados de acordo como ficou no BD. erro != nil caso ocorra um problema no processo de inclusão. Deve ser informado uma conexão ao BD como parâmetro obrigatório e uma conta(*Conta)
func AdicionaConta(db *sql.DB, novaConta *conta.Conta) (c *conta.Conta, err error) {
	c, err = conta.NewConta(novaConta.Nome, novaConta.NomeTipoConta, novaConta.Codigo, novaConta.ContaPai, novaConta.Comentario)
	if err != nil {
		return
	}

	sql := `
INSERT INTO {{.tabela}}(
	{{.nome}}, {{.nomeTipoConta}}, {{.codigo}}, {{.contaPai}}, {{.comentario}}, {{.dataCriacao}}, {{.dataModificacao}}, {{.estado}})
VALUES($1, $2, $3, $4, $5, $6, $7, $8)
`
	query := getTemplateQuery("AdicionaConta", contaDB, sql)

	return adicionaConta(db, c, query)
}

// ProcuraConta localiza uma conta no BD e retorna a conta procurada(*Conta). erro != nil caso ocorra um problema no processo de procura. Deve ser informado uma conexão ao BD como parâmetro obrigatório e o NOME da conta desejada
func ProcuraConta(db *sql.DB, nome string) (c *conta.Conta, err error) {
	sql := `
SELECT
	{{.nome}}, {{.nomeTipoConta}}, {{.codigo}}, {{.contaPai}}, {{.comentario}}, {{.dataCriacao}}, {{.dataModificacao}}, {{.estado}}
FROM
	{{.tabela}}
WHERE {{.nome}} = $1
`
	query := getTemplateQuery("ProcuraConta", contaDB, sql)

	contas, err := carregaContas(db, query, nome)
	if len(contas) == 1 {
		c = contas[0]
	} else {
		err = errors.New("Não foi encontrado um registro com o nome " + nome)
	}

	return
}

// AtivaConta ativa uma conta no BD e retorna a conta(*Conta) com os dados atualizados. erro != nil caso ocorra um problema no processo de procura. Deve ser informado uma conexão ao BD como parâmetro obrigatório e um NOME da Conta desejada
func AtivaConta(db *sql.DB, nome string) (c *conta.Conta, err error) {
	contaBanco, err := ProcuraConta(db, nome)
	if err != nil {
		return
	}

	contaBanco.Ativa()

	sql := `
UPDATE {{.tabela}}
SET {{.estado}} = $1, {{.dataModificacao}} = $2
WHERE {{.nome}} = $3
`

	query := getTemplateQuery("AtivaConta", contaDB, sql)

	return estadoConta(db, contaBanco, query, nome)
}

// InativaConta ativa uma conta no BD e retorna a conta(*Conta) com os dados atualizados. erro != nil caso ocorra um problema no processo de procura. Deve ser informado uma conexão ao BD como parâmetro obrigatório e um NOME da Conta desejada
func InativaConta(db *sql.DB, nome string) (c *conta.Conta, err error) {
	contaBanco, err := ProcuraConta(db, nome)
	if err != nil {
		return
	}

	contaBanco.Inativa()

	sql := `
UPDATE {{.tabela}}
SET {{.estado}} = $1, {{.dataModificacao}} = $2
WHERE {{.nome}} = $3
`

	query := getTemplateQuery("InativaConta", contaDB, sql)

	return estadoConta(db, contaBanco, query, nome)
}

// AlteraConta altera uma conta com o nome(string) informado a partir dos dados da *Conta informada no parâmetro contaAlteracao. O campo Estado não é alterado, enquanto que o campo Nome sim. Use a função específica para essa tarefa(estado). Retorna uma *Conta alterada no BD e um error. error != nil caso ocorra um problema.
func AlteraConta(db *sql.DB, nome string, contaAlteracao *conta.Conta) (c *conta.Conta, err error) {
	contaBanco, err := ProcuraConta(db, nome)
	if err != nil {
		return
	}

	err = contaBanco.Altera(contaAlteracao.Nome, contaAlteracao.NomeTipoConta, contaAlteracao.Codigo, contaAlteracao.ContaPai, contaAlteracao.Comentario)
	if err != nil {
		return
	}

	sql := `
UPDATE {{.tabela}}
SET {{.nome}} = $1, {{.nomeTipoConta}} = $2, {{.codigo}} = $3, {{.contaPai}} = $4, {{.comentario}} = $5, {{.dataModificacao}} = $6
WHERE {{.nome}} = $7
`

	query := getTemplateQuery("AlteraConta", contaDB, sql)

	return alteraConta(db, contaBanco, query, nome)
}

// RemoveConta remove uma conta do BD e retorna erro != nil caso ocorra um problema no processo de remoção. Deve ser informado uma conexão ao BD como parâmetro obrigatório e uma string contendo o NOME da conta desejado
func RemoveConta(db *sql.DB, nome string) (err error) {
	sql := `
DELETE FROM
	{{.tabela}}
WHERE {{.nome}} = $1
`

	query := getTemplateQuery("RemoveConta", contaDB, sql)

	c, err := ProcuraConta(db, nome)
	if c != nil {
		err = remove(db, c.Nome, query)
	}

	return
}

func alteraConta(db *sql.DB, contaBanco *conta.Conta, query, chave string) (c *conta.Conta, err error) {
	resultado, err := altera(db, contaBanco, query, setValoresConta03, chave)
	contaTemp, ok := resultado.(*conta.Conta)
	if ok {
		c = contaTemp
		c.DataCriacao = c.DataCriacao.Local()
	}
	return
}

func setValoresConta03(stmt *sql.Stmt, novoRegistro interface{}, chave string) (r sql.Result, err error) {
	novaConta, ok := novoRegistro.(*conta.Conta)

	codigo, contaPai, comentario := getCamposConvertidosConta(novaConta)

	if ok {
		r, err = stmt.Exec(
			novaConta.Nome,
			novaConta.NomeTipoConta,
			codigo,
			contaPai,
			comentario,
			novaConta.DataModificacao,
			chave)
	}
	return
}

func getCamposConvertidosConta(novaConta *conta.Conta) (codigo, contaPai, comentario sql.NullString) {
	temCodigo := len(novaConta.Codigo) > 0
	codigo = sql.NullString{String: novaConta.Codigo, Valid: temCodigo}

	temContaPai := len(novaConta.ContaPai) > 0
	contaPai = sql.NullString{String: novaConta.ContaPai, Valid: temContaPai}

	temComentario := len(novaConta.Comentario) > 0
	comentario = sql.NullString{String: novaConta.Comentario, Valid: temComentario}

	return
}

func estadoConta(db *sql.DB, contaBanco *conta.Conta, query, chave string) (c *conta.Conta, err error) {
	resultado, err := altera(db, contaBanco, query, setValoresConta02, chave)
	contaTemp, ok := resultado.(*conta.Conta)
	if ok {
		c = contaTemp
	}
	return
}

func setValoresConta02(stmt *sql.Stmt, novoRegistro interface{}, chave string) (r sql.Result, err error) {
	novaConta, ok := novoRegistro.(*conta.Conta)

	if ok {
		r, err = stmt.Exec(
			novaConta.Estado,
			novaConta.DataModificacao,
			chave)
	}
	return
}

func adicionaConta(db *sql.DB, novaConta *conta.Conta, query string) (c *conta.Conta, err error) {
	resultado, err := adiciona(db, novaConta, query, setValoresConta01)
	contaTemp, ok := resultado.(*conta.Conta)
	if ok {
		c = contaTemp
	}
	return
}

func setValoresConta01(stmt *sql.Stmt, novoRegistro interface{}) (r sql.Result, err error) {
	novaConta, ok := novoRegistro.(*conta.Conta)

	codigo, contaPai, comentario := getCamposConvertidosConta(novaConta)

	if ok {
		r, err = stmt.Exec(
			novaConta.Nome,
			novaConta.NomeTipoConta,
			codigo,
			contaPai,
			comentario,
			novaConta.DataCriacao,
			novaConta.DataModificacao,
			novaConta.Estado)
	}
	return
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
	contaAtual.CorrigeData()

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
