package dao

import (
	"database/sql"
	"errors"

	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/model/pessoa"
	"gorm.io/gorm"
)

var (
	pessoaDB = map[string]string{
		"tabela":          pessoa.GetNomeTabelaPessoa(),
		"cpf":             "cpf",
		"nomeCompleto":    "nome_completo",
		"usuario":         "usuario",
		"senha":           "senha",
		"email":           "email",
		"dataCriacao":     "data_criacao",
		"dataModificacao": "data_modificacao",
		"estado":          "estado",
		"administrador":   "administrador"}
)

// CarregaPessoas retorna uma listagem de pessoas(pessoa.Pessoas) e erro = nil do BD caso a consulta ocorra corretamente. erro != nil caso ocorra um problema. Deve ser informado uma conexão ao BD como parâmetro obrigatório
func CarregaPessoas(db *sql.DB) (pessoas pessoa.Pessoas, err error) {
	sql := `
SELECT
	{{.cpf}}, {{.nomeCompleto}}, {{.usuario}}, {{.senha}}, {{.email}}, {{.dataCriacao}}, {{.dataModificacao}}, {{.estado}}, {{.administrador}}
FROM
	{{.tabela}}
`
	// fmt.Println(fmt.Sprintf(`sql
	// nova linha{{.atributoParaTemplate}}: %d`, 123)) // exemplo para adicionar numero concatenado ao sql template

	query := getTemplateQuery("CarregaPessoas", pessoaDB, sql)

	return carregaPessoas(db, query)
}

func CarregaPessoas02(db *gorm.DB) (pessoas pessoa.Pessoas, err error) {
	var tpessoas pessoa.TPessoas
	resultado := db.Find(&tpessoas)

	return ConverteTPessoasParaPessoas(resultado, &tpessoas)
}

// AdicionaPessoa adiciona uma pessoa comum ao BD e retorna a pessoa incluída(*Pessoa) com os dados de acordo como ficou no BD. erro != nil caso ocorra um problema no processo de inclusão. Deve ser informado uma conexão ao BD como parâmetro obrigatório e uma pessoa(*Pessoa)
func AdicionaPessoa(db *sql.DB, novaPessoa *pessoa.Pessoa) (p *pessoa.Pessoa, err error) {
	return adicionaPessoaBase(db, novaPessoa, pessoa.NewPessoa)
}

// AdicionaPessoaAdmin adiciona uma pessoa administradora ao BD e retorna a pessoa incluída(*Pessoa) com os dados de acordo como ficou no BD. erro != nil caso ocorra um problema no processo de inclusão. Deve ser informado uma conexão ao BD como parâmetro obrigatório e uma pessoa(*Pessoa)
func AdicionaPessoaAdmin(db *sql.DB, novaPessoa *pessoa.Pessoa) (p *pessoa.Pessoa, err error) {
	return adicionaPessoaBase(db, novaPessoa, pessoa.NewPessoaAdmin)
}

// AdicionaPessoa02 adiciona uma pessoa comum ao BD e retorna a pessoa incluída(*Pessoa) com os dados de acordo como ficou no BD. erro != nil caso ocorra um problema no processo de inclusão. Deve ser informado uma conexão ao BD(*gorm.DB) e uma pessoa(*Pessoa) como parâmetros obrigatórios
func AdicionaPessoa02(db *gorm.DB, novaPessoa *pessoa.Pessoa) (*pessoa.Pessoa, error) {
	return adicionaPessoaBase02(db, novaPessoa, pessoa.NewPessoa)
}

// AdicionaPessoaAdmin02 adiciona uma pessoa administradora ao BD e retorna a pessoa incluída(*Pessoa) com os dados de acordo como ficou no BD. erro != nil caso ocorra um problema no processo de inclusão. Deve ser informado uma conexão ao BD(*gorm.DB) e uma pessoa(*Pessoa) como parâmetros obrigatórios
func AdicionaPessoaAdmin02(db *gorm.DB, novaPessoa *pessoa.Pessoa) (*pessoa.Pessoa, error) {
	return adicionaPessoaBase02(db, novaPessoa, pessoa.NewPessoaAdmin)
}

// RemovePessoa remove uma pessoa do BD e retorna erro != nil caso ocorra um problema no processo de remoção. Deve ser informado uma conexão ao BD como parâmetro obrigatório e uma string contendo o CPF da pessoa desejada
func RemovePessoa(db *sql.DB, cpf string) (err error) {
	sql := `
DELETE FROM
	{{.tabela}}
WHERE {{.cpf}} = $1
`
	query := getTemplateQuery("RemovePessoa", pessoaDB, sql)

	p, err := ProcuraPessoa(db, cpf)
	if p != nil {
		err = remove(db, p.Cpf, query)
	}

	return
}

// RemovePessoaPorUsuario remove uma pessoa do BD e retorna erro != nil caso ocorra um problema no processo de remoção. Deve ser informado uma conexão ao BD como parâmetro obrigatório e uma string contendo o USUÁRIO da pessoa desejada
func RemovePessoaPorUsuario(db *sql.DB, usuario string) (err error) {
	sql := `
DELETE FROM
	{{.tabela}}
WHERE {{.usuario}} = $1
`
	query := getTemplateQuery("RemovePessoaPorUsuario", pessoaDB, sql)

	p, err := ProcuraPessoaPorUsuario(db, usuario)
	if p != nil {
		err = remove(db, p.Usuario, query)
	}

	return
}

// ProcuraPessoa localiza uma pessoa no BD e retorna a pessoa procurada(*Pessoa). erro != nil caso ocorra um problema no processo de procura. Deve ser informado uma conexão ao BD como parâmetro obrigatório e um CPF da pessoa desejada
func ProcuraPessoa(db *sql.DB, cpf string) (p *pessoa.Pessoa, err error) {
	sql := `
SELECT
	{{.cpf}}, {{.nomeCompleto}}, {{.usuario}}, {{.senha}}, {{.email}}, {{.dataCriacao}}, {{.dataModificacao}}, {{.estado}}, {{.administrador}}
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

// ProcuraPessoaPorUsuario localiza uma pessoa no BD e retorna a pessoa procurada(*Pessoa). erro != nil caso ocorra um problema no processo de procura. Deve ser informado uma conexão ao BD como parâmetro obrigatório e uma string contendo o USUÁRIO da pessoa desejada
func ProcuraPessoaPorUsuario(db *sql.DB, usuario string) (p *pessoa.Pessoa, err error) {
	sql := `
SELECT
	{{.cpf}}, {{.nomeCompleto}}, {{.usuario}}, {{.senha}}, {{.email}}, {{.dataCriacao}}, {{.dataModificacao}}, {{.estado}}, {{.administrador}}
FROM
	{{.tabela}}
WHERE {{.usuario}} = $1
`
	query := getTemplateQuery("ProcuraPessoaPorUsuario", pessoaDB, sql)

	pessoas, err := carregaPessoas(db, query, usuario)
	if len(pessoas) == 1 {
		p = pessoas[0]
	} else {
		err = errors.New("Não foi encontrado um registro com o usuário " + usuario)
	}

	return
}

// AlteraPessoa altera uma pessoa com o cpf(string) informado a partir dos dados da *Pessoa informada no parâmetro pessoaAlteracao. Os campos Cpf(PK) e Estado não são alterados. Use a função específica para essa tarefa. Retorna uma *Pessoa alterada no BD e um error. error != nil caso ocorra um problema.
func AlteraPessoa(db *sql.DB, cpf string, pessoaAlteracao *pessoa.Pessoa) (p *pessoa.Pessoa, err error) {
	pessoaBanco, err := ProcuraPessoa(db, cpf)
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

// AlteraPessoaPorUsuario altera uma pessoa com o usuário(string) informado a partir dos dados da *Pessoa informada no parâmetro pessoaAlteracao. Os campos Cpf(PK) e Estado não são alterados. Use a função específica para essa tarefa. Retorna uma *Pessoa alterada no BD e um error. error != nil caso ocorra um problema.
func AlteraPessoaPorUsuario(db *sql.DB, usuario string, pessoaAlteracao *pessoa.Pessoa) (p *pessoa.Pessoa, err error) {
	pessoaBanco, err := ProcuraPessoaPorUsuario(db, usuario)
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

	return alteraPessoa(db, pessoaBanco, query, pessoaBanco.Cpf)
}

// AtivaPessoa ativa uma pessoa no BD e retorna a pessoa(*Pessoa) com os dados atualizados. erro != nil caso ocorra um problema no processo de procura. Deve ser informado uma conexão ao BD como parâmetro obrigatório e um CPF da pessoa desejada
func AtivaPessoa(db *sql.DB, cpf string) (p *pessoa.Pessoa, err error) {
	pessoaBanco, err := ProcuraPessoa(db, cpf)
	if err != nil {
		return
	}

	pessoaBanco.Ativa()

	sql := `
UPDATE {{.tabela}}
SET {{.estado}} = $1, {{.dataModificacao}} = $2
WHERE {{.cpf}} = $3
`

	query := getTemplateQuery("AtivaPessoa", pessoaDB, sql)

	return estadoPessoa(db, pessoaBanco, query, cpf)
}

// InativaPessoa inativa uma pessoa no BD e retorna a pessoa(*Pessoa) com os dados atualizados. erro != nil caso ocorra um problema no processo de procura. Deve ser informado uma conexão ao BD como parâmetro obrigatório e um CPF da pessoa desejada
func InativaPessoa(db *sql.DB, cpf string) (p *pessoa.Pessoa, err error) {
	pessoaBanco, err := ProcuraPessoa(db, cpf)
	if err != nil {
		return
	}

	pessoaBanco.Inativa()

	sql := `
UPDATE {{.tabela}}
SET {{.estado}} = $1, {{.dataModificacao}} = $2
WHERE {{.cpf}} = $3
`

	query := getTemplateQuery("InativaPessoa", pessoaDB, sql)

	return estadoPessoa(db, pessoaBanco, query, cpf)
}

// SetAdministrador define com administrador de acordo com parâmetro boleado admin informado e retorna a pessoa com os dados atualizados. erro != nil caso ocorra um problema no processo de procura. Deve ser informado uma conexão ao BD como parâmetro obrigatório, um CPF da pessoa desejada e o valor boleano no parâmetro admin
func SetAdministrador(db *sql.DB, cpf string, admin bool) (p *pessoa.Pessoa, err error) {
	pessoaBanco, err := ProcuraPessoa(db, cpf)
	if err != nil {
		return
	}

	pessoaBanco.SetAdmin(admin)

	sql := `
UPDATE {{.tabela}}
SET {{.administrador}} = $1, {{.dataModificacao}} = $2
WHERE {{.cpf}} = $3
`

	query := getTemplateQuery("SetAdministrador", pessoaDB, sql)

	return setAdminPessoa(db, pessoaBanco, query, cpf)
}

func adicionaPessoaBase(db *sql.DB, novaPessoa *pessoa.Pessoa, newPessoa func(string, string, string, string, string) (*pessoa.Pessoa, error)) (p *pessoa.Pessoa, err error) {
	p, err = newPessoa(novaPessoa.Cpf, novaPessoa.NomeCompleto, novaPessoa.Usuario, novaPessoa.Senha, novaPessoa.Email)
	if err != nil {
		return
	}

	sql := `
INSERT INTO {{.tabela}}(
	{{.cpf}}, {{.nomeCompleto}}, {{.usuario}}, {{.senha}}, {{.email}}, {{.dataCriacao}}, {{.dataModificacao}}, {{.estado}}, {{.administrador}})
VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)
`
	query := getTemplateQuery("AdicionaPessoa", pessoaDB, sql)

	return adicionaPessoa(db, p, query)
}

func setAdminPessoa(db *sql.DB, pessoaBanco *pessoa.Pessoa, query, chave string) (p *pessoa.Pessoa, err error) {
	resultado, err := altera(db, pessoaBanco, query, setValoresPessoa04, chave)
	pessoaTemp, ok := resultado.(*pessoa.Pessoa)
	if ok {
		p = pessoaTemp
	}
	return
}

func estadoPessoa(db *sql.DB, pessoaBanco *pessoa.Pessoa, query, chave string) (p *pessoa.Pessoa, err error) {
	resultado, err := altera(db, pessoaBanco, query, setValoresPessoa03, chave)
	pessoaTemp, ok := resultado.(*pessoa.Pessoa)
	if ok {
		p = pessoaTemp
	}
	return
}

func setValoresPessoa03(stmt *sql.Stmt, novoRegistro interface{}, chave string) (r sql.Result, err error) {
	novaPessoa, ok := novoRegistro.(*pessoa.Pessoa)

	if ok {
		r, err = stmt.Exec(
			novaPessoa.Estado,
			novaPessoa.DataModificacao,
			chave)
	}
	return
}

func setValoresPessoa04(stmt *sql.Stmt, novoRegistro interface{}, chave string) (r sql.Result, err error) {
	novaPessoa, ok := novoRegistro.(*pessoa.Pessoa)

	if ok {
		r, err = stmt.Exec(
			novaPessoa.Administrador,
			novaPessoa.DataModificacao,
			chave)
	}
	return
}

func alteraPessoa(db *sql.DB, pessoaBanco *pessoa.Pessoa, query, chave string) (p *pessoa.Pessoa, err error) {
	resultado, err := altera(db, pessoaBanco, query, setValoresPessoa02, chave)
	pessoaTemp, ok := resultado.(*pessoa.Pessoa)
	if ok {
		p = pessoaTemp
		p.DataCriacao = p.DataCriacao.Local()
	}
	return
}

func setValoresPessoa02(stmt *sql.Stmt, novoRegistro interface{}, chave string) (r sql.Result, err error) {

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
	resultado, err := adiciona(db, novaPessoa, query, setValoresPessoa01)
	pessoaTemp, ok := resultado.(*pessoa.Pessoa)
	if ok {
		p = pessoaTemp
	}
	return
}

func setValoresPessoa01(stmt *sql.Stmt, novoRegistro interface{}) (r sql.Result, err error) {

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
			novaPessoa.Estado,
			novaPessoa.Administrador)
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
	pessoaAtual.CorrigeData()

	novosRegistros = append(registros, pessoaAtual)

	return
}

func registrosPessoas02(rows *sql.Rows, registros []interface{}) (novosRegistros []interface{}, err error) {
	pessoaAtual := new(pessoa.PessoaSimples)
	err = scanPessoas02(rows, pessoaAtual)
	if err != nil {
		return
	}
	pessoaAtual.CorrigeData()

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
		&pessoaAtual.Estado,
		&pessoaAtual.Administrador)
}

func scanPessoas02(rows *sql.Rows, pessoaAtual *pessoa.PessoaSimples) error {
	return rows.Scan(
		&pessoaAtual.Usuario,
		&pessoaAtual.Email,
		&pessoaAtual.DataCriacao,
		&pessoaAtual.DataModificacao)
}

func adicionaPessoaBase02(db *gorm.DB, novaPessoa *pessoa.Pessoa, newPessoa pessoa.FuncaoNewPessoa) (p *pessoa.Pessoa, err error) {
	p, err = newPessoa(novaPessoa.Cpf, novaPessoa.NomeCompleto, novaPessoa.Usuario, novaPessoa.Senha, novaPessoa.Email)
	if err != nil {
		return
	}

	tpessoa := ConvertePessoaParaTPessoa(*p)
	err = db.Create(&tpessoa).Error
	if err != nil {
		return
	}
	pessoa := ConverteTPessoaParaPessoa(tpessoa)

	return &pessoa, nil
}
