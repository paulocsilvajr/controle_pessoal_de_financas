package dao

import (
	"controle_pessoal_de_financas/API/v1/model/lancamento"
	"database/sql"
	"fmt"
	"strconv"
)

var (
	lancamentoDB = map[string]string{
		"tabela":          "lancamento",
		"id":              "id",
		"cpfPessoa":       "cpf_pessoa",
		"data":            "data",
		"numero":          "numero",
		"descricao":       "descricao",
		"dataCriacao":     "data_criacao",
		"dataModificacao": "data_modificacao",
		"estado":          "estado"}
)

// CarregaLancamentos retorna uma listagem de todos os lancamentos(lancamento.Lancamentos) e erro = nil do BD caso a consulta ocorra corretamente. erro != nil caso ocorra um problema. Deve ser informado uma conexão ao BD como parâmetro obrigatório
func CarregaLancamentos(db *sql.DB) (lancamentos lancamento.Lancamentos, err error) {
	sql := `
SELECT
	{{.id}}, {{.cpfPessoa}}, {{.data}}, {{.numero}}, {{.descricao}}, {{.dataCriacao}}, {{.dataModificacao}}, {{.estado}}
FROM
	{{.tabela}}
`

	query := getTemplateQuery("CarregaLancamentos", lancamentoDB, sql)

	return carregaLancamentos(db, query)
}

// CarregaLancamentosAtivo retorna uma listagem de lancamentos ativos(lancamento.Lancamento) e erro = nil do BD caso a consulta ocorra corretamente. erro != nil caso ocorra um problema. Deve ser informado uma conexão ao BD como parâmetro obrigatório
func CarregaLancamentosAtivo(db *sql.DB) (lancamentos lancamento.Lancamentos, err error) {
	sql := `
SELECT
	{{.id}}, {{.cpfPessoa}}, {{.data}}, {{.numero}}, {{.descricao}}, {{.dataCriacao}}, {{.dataModificacao}}, {{.estado}}
FROM
	{{.tabela}}
WHERE
	{{.estado}} = true
`

	query := getTemplateQuery("CarregaLancamentosAtivo", lancamentoDB, sql)

	return carregaLancamentos(db, query)
}

// CarregaLancamentosInativo retorna uma listagem de lancamentos inativos(lancamento.Lancamento) e erro = nil do BD caso a consulta ocorra corretamente. erro != nil caso ocorra um problema. Deve ser informado uma conexão ao BD como parâmetro obrigatório
func CarregaLancamentosInativo(db *sql.DB) (lancamentos lancamento.Lancamentos, err error) {
	sql := `
SELECT
	{{.id}}, {{.cpfPessoa}}, {{.data}}, {{.numero}}, {{.descricao}}, {{.dataCriacao}}, {{.dataModificacao}}, {{.estado}}
FROM
	{{.tabela}}
WHERE
	{{.estado}} = false
`

	query := getTemplateQuery("CarregaLancamentosAtivo", lancamentoDB, sql)

	return carregaLancamentos(db, query)
}

// AdicionaLancamento adiciona um lancamento ao BD e retorna o lancamento incluída(*Lancamento) com os dados de acordo como ficou no BD. erro != nil caso ocorra um problema no processo de inclusão. Deve ser informado uma conexão ao BD como parâmetro obrigatório e um lancamento(*Lancamento)
func AdicionaLancamento(db *sql.DB, novoLancamento *lancamento.Lancamento) (l *lancamento.Lancamento, err error) {
	l, err = lancamento.NewLancamento(novoLancamento.ID, novoLancamento.CpfPessoa, novoLancamento.Data, novoLancamento.Numero, novoLancamento.Descricao)
	if err != nil {
		return
	}

	// foi necessário usar a instrução RETURNING em sql para obter a chave gerada ao inserir o Lancamento. O driver "github.com/lib/pq" não tem o método LastInsertId funcional. Detalhes e fonte em função setValoresLancamento01
	sql := `
INSERT INTO {{.tabela}}(
	{{.cpfPessoa}}, {{.data}}, {{.numero}}, {{.descricao}}, {{.dataCriacao}}, {{.dataModificacao}}, {{.estado}})
VALUES($1, $2, $3, $4, $5, $6, $7)
RETURNING {{.id}}
`
	query := getTemplateQuery("AdicionaLancamento", lancamentoDB, sql)

	return adicionaLancamento(db, l, query)
}

// ProcuraLancamento localiza um lancamento no BD e retorna o lancamento procurado(*Lancamento). erro != nil caso ocorra um problema no processo de procura. Deve ser informado uma conexão ao BD como parâmetro obrigatório e o ID do lancamento desejado
func ProcuraLancamento(db *sql.DB, id int) (l *lancamento.Lancamento, err error) {
	sql := `
SELECT
	{{.id}}, {{.cpfPessoa}}, {{.data}}, {{.numero}}, {{.descricao}}, {{.dataCriacao}}, {{.dataModificacao}}, {{.estado}}
FROM
	{{.tabela}}
WHERE {{.id}} = $1
`
	query := getTemplateQuery("ProcuraLancamento", lancamentoDB, sql)

	lancamentos, err := carregaLancamentos(db, query, id)
	if len(lancamentos) == 1 {
		l = lancamentos[0]
	} else {
		err = fmt.Errorf("Não foi encontrado um registro com o ID %d", id)
	}

	return
}

// AtivaLancamento ativa um lancamento no BD e retorna o lancamento(*Lancamento) com os dados atualizados. erro != nil caso ocorra um problema no processo de procura. Deve ser informado uma conexão ao BD como parâmetro obrigatório e um ID do Lancamento desejado
func AtivaLancamento(db *sql.DB, id int) (l *lancamento.Lancamento, err error) {
	lancamentoBanco, err := ProcuraLancamento(db, id)
	if err != nil {
		return
	}

	lancamentoBanco.Ativa()

	sql := `
UPDATE {{.tabela}}
SET {{.estado}} = $1, {{.dataModificacao}} = $2
WHERE {{.id}} = $3
`

	query := getTemplateQuery("AtivaLancamento", lancamentoDB, sql)

	return estadoLancamento(db, lancamentoBanco, query, id)
}

// InativaLancamento inativa um lancamento no BD e retorna o lancamento(*Lancamento) com os dados atualizados. erro != nil caso ocorra um problema no processo de procura. Deve ser informado uma conexão ao BD como parâmetro obrigatório e um ID do Lancamento desejado
func InativaLancamento(db *sql.DB, id int) (l *lancamento.Lancamento, err error) {
	lancamentoBanco, err := ProcuraLancamento(db, id)
	if err != nil {
		return
	}

	lancamentoBanco.Inativa()

	sql := `
UPDATE {{.tabela}}
SET {{.estado}} = $1, {{.dataModificacao}} = $2
WHERE {{.id}} = $3
`

	query := getTemplateQuery("InativaLancamento", lancamentoDB, sql)

	return estadoLancamento(db, lancamentoBanco, query, id)
}

// AlteraLancamento altera um lancamento com o id(int) informado a partir dos dados do *Lancamento informada no parâmetro lancamentoAlteracao. O campo Estado não é alterado. Use a função específica para essa tarefa(estado). Retorna um *Lancamento alterado no BD e um error. error != nil caso ocorra um problema.
func AlteraLancamento(db *sql.DB, id int, lancamentoAlteracao *lancamento.Lancamento) (l *lancamento.Lancamento, err error) {
	lancamentoBanco, err := ProcuraLancamento(db, id)

	if err != nil {
		return
	}

	err = lancamentoBanco.Altera(lancamentoAlteracao.CpfPessoa, lancamentoAlteracao.Data, lancamentoAlteracao.Numero, lancamentoAlteracao.Descricao)
	if err != nil {
		return
	}

	sql := `
UPDATE {{.tabela}}
SET {{.cpfPessoa}} = $1, {{.data}} = $2, {{.numero}} = $3, {{.descricao}} = $4, {{.dataModificacao}} = $5
WHERE {{.id}} = $6
`

	query := getTemplateQuery("AlteraLancamento", lancamentoDB, sql)

	return alteraLancamento(db, lancamentoBanco, query, id)
}

// RemoveLancamento remove um lancamento do BD e retorna erro != nil caso ocorra um problema no processo de remoção. Deve ser informado uma conexão ao BD como parâmetro obrigatório e um int contendo o ID do lancamento desejado
func RemoveLancamento(db *sql.DB, id int) (err error) {
	sql := `
DELETE FROM
	{{.tabela}}
WHERE {{.id}} = $1
`

	query := getTemplateQuery("RemoveLancamento", lancamentoDB, sql)

	l, err := ProcuraLancamento(db, id)
	if l != nil {
		err = remove(db, l.ID, query)
	}

	return
}

func alteraLancamento(db *sql.DB, lancamentoBanco *lancamento.Lancamento, query string, chave int) (l *lancamento.Lancamento, err error) {
	chaveString := strconv.Itoa(chave)
	resultado, err := altera(db, lancamentoBanco, query, setValoresLancamento03, chaveString)
	lancamentoTemp, ok := resultado.(*lancamento.Lancamento)
	if ok {
		l = lancamentoTemp
	}
	return
}

func setValoresLancamento03(stmt *sql.Stmt, novoRegistro interface{}, chave string) (r sql.Result, err error) {
	novoLancamento, ok := novoRegistro.(*lancamento.Lancamento)

	numero := getCamposConvertidosLancamento(novoLancamento)

	if ok {
		r, err = stmt.Exec(
			novoLancamento.CpfPessoa,
			novoLancamento.Data,
			numero,
			novoLancamento.Descricao,
			novoLancamento.DataModificacao,
			chave)
	}
	return
}

func estadoLancamento(db *sql.DB, lancamentoBanco *lancamento.Lancamento, query string, chave int) (l *lancamento.Lancamento, err error) {
	chaveString := strconv.Itoa(chave) // convertido chave em string para satisfazer a interface da função setValoresLancamento02. Ao fazer update em banco, é reconhecido o campo como int(5) ou string('5') em where
	resultado, err := altera(db, lancamentoBanco, query, setValoresLancamento02, chaveString)
	lancamentoTemp, ok := resultado.(*lancamento.Lancamento)
	if ok {
		l = lancamentoTemp
	}
	return
}

func setValoresLancamento02(stmt *sql.Stmt, novoRegistro interface{}, chave string) (r sql.Result, err error) {
	novoLancamento, ok := novoRegistro.(*lancamento.Lancamento)

	if ok {
		r, err = stmt.Exec(
			novoLancamento.Estado,
			novoLancamento.DataModificacao,
			chave)
	}
	return
}

func adicionaLancamento(db *sql.DB, novoLancamento *lancamento.Lancamento, query string) (l *lancamento.Lancamento, err error) {
	resultado, err := adiciona(db, novoLancamento, query, setValoresLancamento01)
	lancamentoTemp, ok := resultado.(*lancamento.Lancamento)
	if ok {
		l = lancamentoTemp
	}
	return
}

func setValoresLancamento01(stmt *sql.Stmt, novoRegistro interface{}) (r sql.Result, err error) {
	novoLancamento, ok := novoRegistro.(*lancamento.Lancamento)

	numero := getCamposConvertidosLancamento(novoLancamento)
	lastInsertID := 0

	if ok {
		err = stmt.QueryRow(
			novoLancamento.CpfPessoa,
			novoLancamento.Data,
			numero,
			novoLancamento.Descricao,
			novoLancamento.DataCriacao,
			novoLancamento.DataModificacao,
			novoLancamento.Estado).Scan(&lastInsertID)

		novoLancamento.ID = lastInsertID
		// o método de sql.Result LastInsertId não é suportado pelo driver "github.com/lib/pq"
		// foi usado um método alternativo(stmt.QueryRow) de acordo com link https://stackoverflow.com/questions/33382981/go-how-to-get-last-insert-id-on-postgresql-with-namedexec
	}

	return
}

func getCamposConvertidosLancamento(novoLancamento *lancamento.Lancamento) (numero sql.NullString) {
	temNumero := len(novoLancamento.Numero) > 0
	numero = sql.NullString{String: novoLancamento.Numero, Valid: temNumero}

	return
}

func carregaLancamentos(db *sql.DB, query string, args ...interface{}) (lancamentos lancamento.Lancamentos, err error) {
	registros, err := carrega(db, query, registrosLancamento01, args...)

	lancamentos = converterEmLancamento(registros)

	return
}

func registrosLancamento01(rows *sql.Rows, registros []interface{}) (novosRegistros []interface{}, err error) {
	lancamentoAtual := new(lancamento.Lancamento)
	err = scanLancamento01(rows, lancamentoAtual)
	if err != nil {
		return
	}
	novosRegistros = append(registros, lancamentoAtual)

	return
}

func scanLancamento01(rows *sql.Rows, lancamentoAtual *lancamento.Lancamento) error {
	var numero sql.NullString // numero é um campo na tabela lancamento do BD que pode ser nulo

	err := rows.Scan(
		&lancamentoAtual.ID,
		&lancamentoAtual.CpfPessoa,
		&lancamentoAtual.Data,
		&numero,
		&lancamentoAtual.Descricao,
		&lancamentoAtual.DataCriacao,
		&lancamentoAtual.DataModificacao,
		&lancamentoAtual.Estado)
	lancamentoAtual.Numero = numero.String

	return err
}

func converterEmLancamento(registros []interface{}) (lancamentos lancamento.Lancamentos) {
	for _, r := range registros {
		l, ok := r.(*lancamento.Lancamento)
		if ok {
			lancamentos = append(lancamentos, l)
		}
	}

	return
}
