package dao

import (
	"controle_pessoal_de_financas/API/v1/model/pessoa"
	"database/sql"
	"errors"
)

var (
	pessoaDB = map[string]string{
		"tabela":          "pessoa",
		"cpf":             "cpf",
		"nomeCompleto":    "nome_completo",
		"usuario":         "usuario",
		"senha":           "senha",
		"email":           "email",
		"dataCriacao":     "data_criacao",
		"dataModificacao": "data_modificacao",
		"estado":          "estado"}
)

func DaoCarregaPessoas(db *sql.DB) (pessoas pessoa.Pessoas, err error) {
	sql := `
SELECT
	{{.cpf}}, {{.nomeCompleto}}, {{.usuario}}, {{.senha}}, {{.email}}, {{.dataCriacao}}, {{.dataModificacao}}, {{.estado}}
FROM
	{{.tabela}}
`
	// fmt.Println(fmt.Sprintf(`sql
	// nova linha{{.atributoParaTemplate}}: %d`, 123)) // exemplo para adicionar numero concatenado ao sql template

	query := getTemplateQuery("CarregaPessoas", pessoaDB, sql)

	return carregaPessoas(db, query)
}

func DaoAdicionaPessoa(db *sql.DB, novaPessoa *pessoa.Pessoa) (p *pessoa.Pessoa, err error) {
	p, err = pessoa.NewPessoa(novaPessoa.Cpf, novaPessoa.NomeCompleto, novaPessoa.Usuario, novaPessoa.Senha, novaPessoa.Email)
	if err != nil {
		return
	}

	sql := `
INSERT INTO {{.tabela}}(
	{{.cpf}}, {{.nomeCompleto}}, {{.usuario}}, {{.senha}}, {{.email}}, {{.dataCriacao}}, {{.dataModificacao}}, {{.estado}})
VALUES($1, $2, $3, $4, $5, $6, $7, $8)
`
	query := getTemplateQuery("AdicionaPessoa", pessoaDB, sql)

	return adicionaPessoa(db, p, query)
}

func DaoRemovePessoa(db *sql.DB, cpf string) (err error) {
	sql := `
DELETE FROM
	{{.tabela}}
WHERE {{.cpf}} = $1
`
	query := getTemplateQuery("RemovePessoa", pessoaDB, sql)

	p, err := DaoProcuraPessoa(db, cpf)
	if p != nil {
		err = remove(db, p.Cpf, query)
	}

	return
}

func DaoProcuraPessoa(db *sql.DB, cpf string) (p *pessoa.Pessoa, err error) {
	sql := `
SELECT
	{{.cpf}}, {{.nomeCompleto}}, {{.usuario}}, {{.senha}}, {{.email}}, {{.dataCriacao}}, {{.dataModificacao}}, {{.estado}}
FROM
	{{.tabela}}
WHERE {{.cpf}} = $1
`
	query := getTemplateQuery("ProcuraPessoa", pessoaDB, sql)

	pessoas, err := carregaPessoas(db, query, cpf)
	if len(pessoas) == 1 {
		p = pessoas[0]
	} else {
		err = errors.New("Não foi encontrado um registro com o cpf " + cpf)
	}

	return
}

// DaoAlteraPessoa altera uma pessoa com o cpf(string) informado a partir dos dados da *Pessoa informada no parâmetro pessoaAlteracao.
// Retorna uma *Pessoa alterada no BD e um error. error != nil caso ocorra um problema.
func DaoAlteraPessoa(db *sql.DB, cpf string, pessoaAlteracao *pessoa.Pessoa) (p *pessoa.Pessoa, err error) {
	pessoaBanco, err := DaoProcuraPessoa(db, cpf)
	if err != nil {
		return
	}

	err = pessoaBanco.Altera(pessoaAlteracao.Cpf, pessoaAlteracao.NomeCompleto, pessoaAlteracao.Usuario, pessoaAlteracao.Senha, pessoaAlteracao.Email)
	if err != nil {
		return
	}

	sql := `
UPDATE {{.tabela}}
SET {{.nomeCompleto}} = $1, {{.usuario}} = $2, {{.senha}} = $3, {{.email}} = $4, {{.dataModificacao}} = $5
WHERE {{.cpf}} = $6
`

	query := getTemplateQuery("AlteraPessoa", pessoaDB, sql)

	return alteraPessoa(db, pessoaBanco, query, cpf)
}

func alteraPessoa(db *sql.DB, pessoaBanco *pessoa.Pessoa, query, chave string) (p *pessoa.Pessoa, err error) {
	resultado, err := altera(db, pessoaBanco, query, setValores02, chave)
	pessoaTemp, ok := resultado.(*pessoa.Pessoa)
	if ok {
		p = pessoaTemp
	}
	return
}

func setValores02(stmt *sql.Stmt, novoRegistro interface{}, chave string) (r sql.Result, err error) {

	novaPessoa, ok := novoRegistro.(*pessoa.Pessoa)

	if ok {
		r, err = stmt.Exec(
			novaPessoa.NomeCompleto,
			novaPessoa.Usuario,
			novaPessoa.Senha,
			novaPessoa.Email,
			novaPessoa.DataModificacao,
			chave)
	}
	return
}

func adicionaPessoa(db *sql.DB, novaPessoa *pessoa.Pessoa, query string) (p *pessoa.Pessoa, err error) {
	resultado, err := adiciona(db, novaPessoa, query, setValores01)
	pessoaTemp, ok := resultado.(*pessoa.Pessoa)
	if ok {
		p = pessoaTemp
	}
	return
}

func setValores01(stmt *sql.Stmt, novoRegistro interface{}) (r sql.Result, err error) {

	novaPessoa, ok := novoRegistro.(*pessoa.Pessoa)

	if ok {
		r, err = stmt.Exec(
			novaPessoa.Cpf,
			novaPessoa.NomeCompleto,
			novaPessoa.Usuario,
			novaPessoa.Senha,
			novaPessoa.Email,
			novaPessoa.DataCriacao,
			novaPessoa.DataModificacao,
			novaPessoa.Estado)
	}
	return
}

func carregaPessoas(db *sql.DB, query string, args ...interface{}) (pessoas pessoa.Pessoas, err error) {
	registros, err := carrega(db, query, registrosPessoas01, args...)

	pessoas = converteEmPessoas(registros)

	return
}

func converteEmPessoas(registros []interface{}) (pessoas pessoa.Pessoas) {
	for _, r := range registros {
		// fmt.Printf(">>> %T\n", r)
		p, ok := r.(*pessoa.Pessoa)
		if ok {
			pessoas = append(pessoas, p)
		}
	}

	return
}

func registrosPessoas01(rows *sql.Rows, registros []interface{}) (novosRegistros []interface{}, err error) {
	pessoaAtual := new(pessoa.Pessoa)
	err = scanPessoas01(rows, pessoaAtual)
	if err != nil {
		return
	}
	novosRegistros = append(registros, pessoaAtual)

	return
}

func scanPessoas01(rows *sql.Rows, pessoaAtual *pessoa.Pessoa) error {
	return rows.Scan(
		&pessoaAtual.Cpf,
		&pessoaAtual.NomeCompleto,
		&pessoaAtual.Usuario,
		&pessoaAtual.Senha,
		&pessoaAtual.Email,
		&pessoaAtual.DataCriacao,
		&pessoaAtual.DataModificacao,
		&pessoaAtual.Estado)
}

// func OLDcarregaPessoas(db *sql.DB, query string) (pessoas pessoa.Pessoas, err error) {
// 	queryStmt, err := db.Prepare(query)
// 	if err != nil {
// 		return
// 	}

// 	rows, err := queryStmt.Query()
// 	defer queryStmt.Close()
// 	if err != nil {
// 		return
// 	}

// 	for rows.Next() {
// 		pessoaAtual := new(pessoa.Pessoa)
// 		err = scanPessoas01(rows, pessoaAtual)
// 		if err != nil {
// 			return
// 		}
// 		pessoas = append(pessoas, pessoaAtual)
// 	}
// 	err = rows.Err()
// 	if err != nil {
// 		pessoas = nil
// 		return
// 	}

// 	return
// }

// func OLDadicionaPessoa(db *sql.DB, novaPessoa *pessoa.Pessoa, query string) (p *pessoa.Pessoa, err error) {
// 	transacao, err := db.Begin()
// 	if err != nil {
// 		return
// 	}

// 	stmt, err := transacao.Prepare(query)
// 	if err != nil {
// 		return
// 	}
// 	defer stmt.Close()

// 	_, err = stmt.Exec(
// 		novaPessoa.Cpf,
// 		novaPessoa.NomeCompleto,
// 		novaPessoa.Usuario,
// 		novaPessoa.Senha,
// 		novaPessoa.Email,
// 		novaPessoa.DataCriacao,
// 		novaPessoa.DataModificacao,
// 		novaPessoa.Estado)
// 	if err != nil {
// 		return
// 	}

// 	err = transacao.Commit()
// 	if err != nil {
// 		return
// 	}

// 	p = novaPessoa
// 	return
// }

// func DaoRemovePessoa(db *sql.DB, cpf string) (err error) {
// 	transacao, err := db.Begin()
// 	if err != nil {
// 		return
// 	}

// 	stmt, err := transacao.Prepare(`
// DELETE FROM pessoa WHERE cpf = $1
// `)
// 	if err != nil {
// 		return
// 	}
// 	defer stmt.Close()

// 	_, err = stmt.Exec(cpf)
// 	if err != nil {
// 		return
// 	}

// 	err = transacao.Commit()
// 	if err != nil {
// 		return
// 	}

// 	return
// }

// func DaoAlteraPessoa(db *sql.DB, cpf string, pessoaAlteracao *pessoa.Pessoa) (p *pessoa.Pessoa, err error) {
// 	pessoaBanco, err := DaoProcuraPessoa(db, cpf)
// 	if err != nil {
// 		return
// 	}

// 	err = pessoaBanco.Altera(pessoaAlteracao.Cpf, pessoaAlteracao.NomeCompleto, pessoaAlteracao.Usuario, pessoaAlteracao.Senha, pessoaAlteracao.Email)
// 	if err != nil {
// 		return
// 	}

// 	transacao, err := db.Begin()
// 	if err != nil {
// 		return
// 	}

// 	stmt, err := transacao.Prepare(`
// UPDATE pessoa
// SET nome_completo = $1, usuario = $2, senha = $3, email = $4, data_modificacao = $5
// WHERE cpf = $6
// `)
// 	if err != nil {
// 		return
// 	}
// 	defer stmt.Close()

// 	_, err = stmt.Exec(
// 		pessoaBanco.NomeCompleto,
// 		pessoaBanco.Usuario,
// 		pessoaBanco.Senha,
// 		pessoaBanco.Email,
// 		pessoaBanco.DataModificacao,
// 		cpf)
// 	if err != nil {
// 		return
// 	}

// 	err = transacao.Commit()
// 	if err != nil {
// 		return
// 	}

// 	p = pessoaBanco

// 	return
// }
