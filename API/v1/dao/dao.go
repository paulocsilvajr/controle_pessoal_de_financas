package dao

import (
	"bytes"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"strings"
	"time"

	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/config"
	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	loggerGORM "gorm.io/gorm/logger"

	_ "github.com/lib/pq"
)

// funcSetValores é um tipo representando a função para setar os valores para a alteração de um dado no BD. ...interface{} representa todos os campos chave(primária). Veja o exemplo de aplicação em detalhe_lancamento_dao.go na função setValoresDetalheLancamento03. Tipo usado pela função altera2
type funcSetValores func(*sql.Stmt, interface{}, ...interface{}) (sql.Result, error)

// GetDB2 retorna uma conexão com o banco de dados(*gorm.DB) de acordo com as informações obtida de configurações
func GetDB02() *gorm.DB {
	config := config.AbrirConfiguracoes()
	return getDB02(config)
}

func GetDB02ParaTestes() *gorm.DB {
	config := config.AbrirConfiguracoesParaTestes()
	return getDB02(config)
}

func CreateDB() error {
	config := config.AbrirConfiguracoes()
	return createDB(config)
}

func CriarTabelas() error {
	err := criarTabelas()
	return err
}

func criarTabelas() error {
	db2 := GetDB02()

	err := CriarTabelaPessoa(db2)
	if err != nil {
		return verificaErroChaveJaExiste(err)
	}

	err = CriarTabelaTipoConta(db2)
	if err != nil {
		return verificaErroChaveJaExiste(err)
	}

	err = CriarTabelaConta(db2)
	if err != nil {
		return verificaErroChaveJaExiste(err)
	}

	err = CriarTabelaLancamento(db2)
	if err != nil {
		return verificaErroChaveJaExiste(err)
	}

	err = CriarTabelaDetalheLancamento(db2)
	if err != nil {
		return verificaErroChaveJaExiste(err)
	}

	return nil
}

func verificaErroChaveJaExiste(err error) error {
	jaExiste := "already exists"
	if strings.Contains(err.Error(), jaExiste) {
		return nil
	}
	return err
}

func CreateDBParaTestes() error {
	config := config.AbrirConfiguracoesParaTestes()
	return createDB(config)
}

func CloseDB(db *gorm.DB) error {
	db2, _ := db.DB()
	return db2.Close()
}

func PingDB(db *gorm.DB) error {
	db2, _ := db.DB()

	return db2.Ping()
}

func createDB(config config.Configuracoes) error {
	// Fonte: https://stackoverflow.com/questions/55555836/is-it-possible-to-create-postgresql-databases-with-dynamic-names-with-the-help-o
	conninfo := getStringConexao2(config)

	db, err := sql.Open("postgres", conninfo)
	if err != nil {
		log.Fatal(err)
	}

	dbName := config["DBnome"]
	_, err = db.Exec("CREATE DATABASE " + dbName)
	if err != nil {
		bancoDadosJaExiste := fmt.Sprintf("database \"%s\" already exists", dbName)
		if strings.Contains(err.Error(), bancoDadosJaExiste) {
			logger.GeraLogFS(
				fmt.Sprintf("Banco de Dados \"%s\" já existe", dbName),
				time.Now(),
			)

			return nil
		}

		return err
	}

	return nil
}

func getDB02(config config.Configuracoes) *gorm.DB {
	connStr := getStringConexao(config)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  connStr,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		Logger: loggerGORM.Default.LogMode(loggerGORM.Error),
	})

	if err != nil {
		logger.GeraLogFS(
			fmt.Sprintf("Erro ao conectar em servidor do Banco de dados[%s]", err),
			time.Now(),
		)
		log.Fatal(err)
	}

	if err := PingDB(db); err != nil {
		logger.GeraLogFS(
			fmt.Sprintf("Erro em PING em servidor de Banco de Dados[%s]", err),
			time.Now(),
		)
		log.Fatal(err)
	}

	return db
}

// getStringConexao retorna uma string contendo os dados para se conectar ao banco de dados de acordo com configurações(config.Configuracoes) informadas como parâmetro
func getStringConexao(config config.Configuracoes) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config["DBhost"], config["DBporta"], config["DBusuario"], config["DBsenha"], config["DBnome"])
}

// getStringConexao2 retorna uma string contendo os dados para se conectar ao banco de dados de acordo com configurações(config.Configuracoes) informadas como parâmetro, mas SEM conter o nome do banco(dbname)
func getStringConexao2(config config.Configuracoes) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable", config["DBhost"], config["DBporta"], config["DBusuario"], config["DBsenha"])
}

func setNullString(value string) sql.NullString {
	if len(value) > 0 {
		return sql.NullString{
			String: value,
			Valid:  true,
		}
	}
	return sql.NullString{}
}

func setNullFloat64(value float64) sql.NullFloat64 {
	if value != 0 {
		return sql.NullFloat64{
			Float64: value,
			Valid:   true,
		}
	}
	return sql.NullFloat64{}
}

func getTemplateQuery(nome string, campos map[string]string, sql string) string {
	t := template.Must(template.New(nome).Parse(sql))
	query := new(bytes.Buffer)
	t.Execute(query, campos)

	return query.String()
}

func getTemplateSQL(nome string, sql string, campos map[string]string) string {
	return getTemplateQuery(nome, campos, sql)
}

func carrega(db *sql.DB, query string, appendRegistros func(rows *sql.Rows, registros []interface{}) (novosRegistros []interface{}, err error), args ...interface{}) (registros []interface{}, err error) {
	queryStmt, err := db.Prepare(query)
	if err != nil {
		return
	}

	rows, err := queryStmt.Query(args...)
	defer queryStmt.Close()
	if err != nil {
		return
	}

	for rows.Next() {
		registros, err = appendRegistros(rows, registros)
		if err != nil {
			return
		}

	}
	err = rows.Err()
	if err != nil {
		registros = nil
		return
	}

	return
}

func adiciona(db *sql.DB, novoRegistro interface{}, query string, setValores func(*sql.Stmt, interface{}) (sql.Result, error)) (r interface{}, err error) {

	transacao, err := db.Begin()
	if err != nil {
		return
	}

	stmt, err := transacao.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = setValores(stmt, novoRegistro)
	if err != nil {
		return
	}

	err = transacao.Commit()
	if err != nil {
		return
	}

	r = novoRegistro

	return
}

func remove(db *sql.DB, chavePrimaria interface{}, query string) (err error) {
	transacao, err := db.Begin()
	if err != nil {
		return
	}

	stmt, err := transacao.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(chavePrimaria)
	if err != nil {
		return
	}

	err = transacao.Commit()
	if err != nil {
		return
	}

	return
}

func remove2(db *sql.DB, query string, chaves ...interface{}) (err error) {
	transacao, err := db.Begin()
	if err != nil {
		return
	}

	stmt, err := transacao.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(chaves...)
	if err != nil {
		return
	}

	err = transacao.Commit()
	if err != nil {
		return
	}

	return
}

func altera(db *sql.DB, novoRegistro interface{}, query string, setValores func(*sql.Stmt, interface{}, string) (sql.Result, error), chave string) (r interface{}, err error) {

	transacao, err := db.Begin()
	if err != nil {
		return
	}

	stmt, err := transacao.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = setValores(stmt, novoRegistro, chave)
	if err != nil {
		return
	}

	err = transacao.Commit()
	if err != nil {
		return
	}

	r = novoRegistro

	return
}

func altera2(db *sql.DB, novoRegistro interface{}, query string, setValores funcSetValores, chave ...interface{}) (r interface{}, err error) {

	transacao, err := db.Begin()
	if err != nil {
		return
	}

	stmt, err := transacao.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = setValores(stmt, novoRegistro, chave...)
	if err != nil {
		return
	}

	err = transacao.Commit()
	if err != nil {
		return
	}

	r = novoRegistro

	return
}

func altera2T(transacao *sql.Tx, novoRegistro interface{}, query string, setValores funcSetValores, chave ...interface{}) (r interface{}, err error) {
	stmt, err := transacao.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = setValores(stmt, novoRegistro, chave...)
	if err != nil {
		return
	}

	r = novoRegistro

	return
}

// GetDB retorna uma conexão com o banco de dados(*sql.DB) de acordo com as informações obtida de configurações.
func GetDB() *sql.DB {
	config := config.AbrirConfiguracoes()
	connStr := getStringConexao(config)
	db, err := sql.Open(config["DB"], connStr)

	if err != nil {
		logger.GeraLogFS(
			"Erro ao conectar em servidor do Banco de dados",
			time.Now(),
		)
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		logger.GeraLogFS(
			fmt.Sprintf("Erro em PING em servidor de Banco de Dados[%s]", err),
			time.Now(),
		)
	}

	return db
}
