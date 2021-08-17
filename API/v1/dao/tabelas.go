package dao

import (
	"fmt"
	"strings"

	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/model/conta"
	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/model/detalhe_lancamento"
	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/model/lancamento"
	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/model/pessoa"
	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/model/tipo_conta"
	"gorm.io/gorm"
)

// Definição de campos GORM de cada modelo e Método TableName() em cada arquivo de entidade da pasta model. Foi procurado aproveitar as structs já existentes adicionando as anotações do GORM quando possível, ou foi criado structs iniciadas com 'T' para representar o modelo GORM

// SQL Gerado via GORM:
//
// CREATE TABLE public.pessoa
// (
//     cpf character varying(11) COLLATE pg_catalog."default" NOT NULL,
//     nome_completo character varying(100) COLLATE pg_catalog."default",
//     usuario character varying(20) COLLATE pg_catalog."default" NOT NULL,
//     senha character varying(64) COLLATE pg_catalog."default" NOT NULL,
//     email character varying(45) COLLATE pg_catalog."default" NOT NULL,
//     data_criacao timestamp with time zone NOT NULL,
//     data_modificacao timestamp with time zone NOT NULL,
//     estado boolean NOT NULL DEFAULT true,
//     administrador boolean NOT NULL DEFAULT false,
//     CONSTRAINT pessoa_pkey PRIMARY KEY (cpf),
//     CONSTRAINT pessoa_email_key UNIQUE (email),
//     CONSTRAINT pessoa_usuario_key UNIQUE (usuario)
// )
//
// CREATE TABLE public.tipo_conta
// (
//     nome character varying(50) COLLATE pg_catalog."default" NOT NULL,
//     descricao_debito character varying(20) COLLATE pg_catalog."default" NOT NULL,
//     descricao_credito character varying(20) COLLATE pg_catalog."default" NOT NULL,
//     data_criacao timestamp with time zone NOT NULL,
//     data_modificacao timestamp with time zone NOT NULL,
//     estado boolean NOT NULL DEFAULT true,
//     CONSTRAINT tipo_conta_pkey PRIMARY KEY (nome)
// )
//
// CREATE TABLE public.conta
// (
//     nome character varying(50) COLLATE pg_catalog."default" NOT NULL,
//     nome_tipo_conta character varying(50) COLLATE pg_catalog."default" NOT NULL,
//     codigo character varying(19) COLLATE pg_catalog."default",
//     conta_pai character varying(50) COLLATE pg_catalog."default",
//     comentario character varying(150) COLLATE pg_catalog."default",
//     data_criacao timestamp with time zone NOT NULL,
//     data_modificacao timestamp with time zone NOT NULL,
//     estado boolean NOT NULL DEFAULT true,
//     CONSTRAINT conta_pkey PRIMARY KEY (nome),
//     CONSTRAINT conta_codigo_key UNIQUE (codigo),
//     CONSTRAINT conta_fk FOREIGN KEY (conta_pai)
//         REFERENCES public.conta (nome) MATCH SIMPLE
//         ON UPDATE CASCADE
//         ON DELETE CASCADE,
//     CONSTRAINT tipo_conta_fk FOREIGN KEY (nome_tipo_conta)
//         REFERENCES public.tipo_conta (nome) MATCH SIMPLE
//         ON UPDATE CASCADE
//         ON DELETE RESTRICT
// )
//
// CREATE TABLE public.lancamento
// (
//     id bigint NOT NULL DEFAULT nextval('lancamento_id_seq'::regclass),
//     cpf_pessoa character varying(11) COLLATE pg_catalog."default" NOT NULL,
//     data timestamp with time zone NOT NULL,
//     numero character varying(19) COLLATE pg_catalog."default",
//     descricao character varying(100) COLLATE pg_catalog."default" NOT NULL,
//     data_criacao timestamp with time zone NOT NULL,
//     data_modificacao timestamp with time zone NOT NULL,
//     estado boolean NOT NULL DEFAULT true,
//     CONSTRAINT lancamento_pkey PRIMARY KEY (id),
//     CONSTRAINT pessoa_lancamento_fk FOREIGN KEY (cpf_pessoa)
//         REFERENCES public.pessoa (cpf) MATCH SIMPLE
//         ON UPDATE CASCADE
//         ON DELETE CASCADE
// )
//
// CREATE TABLE public.detalhe_lancamento
// (
//     id_lancamento bigint NOT NULL,
//     nome_conta text COLLATE pg_catalog."default" NOT NULL,
//     debito numeric(19,3),
//     credito numeric(19,3),
//     CONSTRAINT detalhe_lancamento_pkey PRIMARY KEY (id_lancamento, nome_conta),
//     CONSTRAINT conta_detalhe_lancamento_fk FOREIGN KEY (nome_conta)
//         REFERENCES public.conta (nome) MATCH SIMPLE
//         ON UPDATE CASCADE
//         ON DELETE RESTRICT,
//     CONSTRAINT lancamento_detalhe_lancamento_fk FOREIGN KEY (id_lancamento)
//         REFERENCES public.lancamento (id) MATCH SIMPLE
//         ON UPDATE CASCADE
//         ON DELETE CASCADE
// )

// CriarTabelaPessoa cria a tabela "pessoa" em banco de dados via AutoMigrate de GORM. Deve ser informado obrigatoriamente com parâmetro um *gorm.DB. Retorna um erro != nil caso ocorra um problema
func CriarTabelaPessoa(db *gorm.DB) error {
	return db.AutoMigrate(&pessoa.TPessoa{})
}

// CriarTabelaTipoConta cria a tabela "tipo_conta" em banco de dados via AutoMigrate de GORM. Deve ser informado obrigatoriamente com parâmetro um *gorm.DB. Retorna um erro != nil caso ocorra um problema
func CriarTabelaTipoConta(db *gorm.DB) error {
	return db.AutoMigrate(&tipo_conta.TTipoConta{})
}

// CriarTabelaConta cria a tabela "conta" em banco de dados via AutoMigrate de GORM e altera a tabela para criar as suas chaves estrangeiras. Deve ser informado obrigatoriamente com parâmetro um *gorm.DB. Retorna um erro != nil caso ocorra um problema
func CriarTabelaConta(db *gorm.DB) error {
	err := db.AutoMigrate(&conta.TConta{})
	if err != nil {
		return err
	}

	err = criarFKTabelaConta(db)
	if err != nil {
		return err
	}

	return nil
}

// CriarTabelaLancamento cria a tabela "lancamento" em banco de dados via AutoMigrate de GORM e altera a tabela para criar as suas chaves estrangeiras. Deve ser informado obrigatoriamente com parâmetro um *gorm.DB. Retorna um erro != nil caso ocorra um problema
func CriarTabelaLancamento(db *gorm.DB) error {
	err := db.AutoMigrate(&lancamento.TLancamento{})
	if err != nil {
		return err
	}

	err = criarFKTabelaLancamento(db)
	if err != nil {
		return err
	}

	return nil
}

// CriarTabelaDetalheLancamento cria a tabela "detalhe_lancamento" em banco de dados via AutoMigrate de GORM e altera a tabela para criar as suas chaves estrangeiras. Deve ser informado obrigatoriamente com parâmetro um *gorm.DB. Retorna um erro != nil caso ocorra um problema
func CriarTabelaDetalheLancamento(db *gorm.DB) error {
	err := db.AutoMigrate(&detalhe_lancamento.TDetalheLancamento{})
	if err != nil {
		return err
	}

	err = criarFKTabelaDetalheLancamento(db)
	if err != nil {
		return err
	}

	return nil
}

func criarFKTabelaConta(db *gorm.DB) error {
	sql1 := `
ALTER TABLE {{.tabela}}
ADD CONSTRAINT tipo_conta_fk FOREIGN KEY ({{.nomeTipoConta}}) REFERENCES {{.tabelaTipoConta}}({{.fkTipoConta}})
ON UPDATE CASCADE
ON DELETE RESTRICT;
`
	sql1 = getTemplateSQL("tipoContaFK", sql1, contaDB)

	sql2 := `
ALTER TABLE {{.tabela}}
ADD CONSTRAINT conta_fk FOREIGN KEY ({{.contaPai}}) REFERENCES {{.tabela}}({{.nome}})
ON UPDATE CASCADE
ON DELETE CASCADE;
`
	sql2 = getTemplateSQL("contaFK", sql2, contaDB)

	return db.Exec(strings.Join([]string{sql1, sql2}, "")).Error
}

func criarFKTabelaLancamento(db *gorm.DB) error {
	sql := `
ALTER TABLE {{.tabela}}
ADD	CONSTRAINT pessoa_lancamento_fk FOREIGN KEY ({{.cpfPessoa}})
REFERENCES {{.tabelaPessoa}}({{.fkPessoa}})
ON UPDATE CASCADE
ON DELETE CASCADE;
`
	sql = getTemplateSQL("cpfPessoaFK", sql, lancamentoDB)

	return db.Exec(sql).Error
}

func criarFKTabelaDetalheLancamento(db *gorm.DB) error {
	sql := `
ALTER TABLE {{.tabela}}
ADD	CONSTRAINT conta_detalhe_lancamento_fk FOREIGN KEY ({{.nomeConta}})
REFERENCES {{.tabelaConta}}({{.fkConta}})
ON UPDATE CASCADE
ON DELETE RESTRICT;

ALTER TABLE {{.tabela}}
ADD	CONSTRAINT lancamento_detalhe_lancamento_fk FOREIGN KEY ({{.idLancamento}})
REFERENCES {{.tabelaLancamento}}({{.fkLancamento}})
ON UPDATE CASCADE
ON DELETE CASCADE;
`
	sql = getTemplateSQL("detalheLancamentoFK", sql, detalheLancamentoDB)

	return db.Exec(sql).Error
}

// ConvertePessoaParaTPessoa recebe um ponteiro do tipo da struct Pessoa como parâmetro e retorna um ponteiro do tipo TPessoa
func ConvertePessoaParaTPessoa(p *pessoa.Pessoa) *pessoa.TPessoa {
	return &pessoa.TPessoa{
		Cpf:             p.Cpf,
		NomeCompleto:    p.NomeCompleto,
		Usuario:         p.Usuario,
		Senha:           p.Senha,
		Email:           p.Email,
		DataCriacao:     p.DataCriacao,
		DataModificacao: p.DataModificacao,
		Estado:          p.Estado,
		Administrador:   p.Administrador,
	}
}

// ConverteTPessoaParaPessoa recebe um ponteiro do tipo da struct TPessoa como parâmetro e retorna um ponteiro do tipo Pessoa
func ConverteTPessoaParaPessoa(p *pessoa.TPessoa) *pessoa.Pessoa {
	return &pessoa.Pessoa{
		Cpf:             p.Cpf,
		NomeCompleto:    p.NomeCompleto,
		Usuario:         p.Usuario,
		Senha:           p.Senha,
		Email:           p.Email,
		DataCriacao:     p.DataCriacao,
		DataModificacao: p.DataModificacao,
		Estado:          p.Estado,
		Administrador:   p.Administrador,
	}
}

// ConverteTPessoaParaPessoaSimples recebe um ponteiro do tipo da struct TPessoa como parâmetro e retorna um ponteiro do tipo PessoaSimples
func ConverteTPessoaParaPessoaSimples(p *pessoa.TPessoa) *pessoa.PessoaSimples {
	return &pessoa.PessoaSimples{
		Usuario:         p.Usuario,
		Email:           p.Email,
		DataCriacao:     p.DataCriacao,
		DataModificacao: p.DataModificacao,
	}
}

// ConverteTipoContaParaTTipoConta recebe uma variável do tipo da struct TipoConta como parâmetro e retorna uma variável do tipo TTipoConta
func ConverteTipoContaParaTTipoConta(tc *tipo_conta.TipoConta) *tipo_conta.TTipoConta {
	return &tipo_conta.TTipoConta{
		Nome:             tc.Nome,
		DescricaoDebito:  tc.DescricaoDebito,
		DescricaoCredito: tc.DescricaoCredito,
		DataCriacao:      tc.DataCriacao,
		DataModificacao:  tc.DataModificacao,
		Estado:           tc.Estado,
	}
}

// ConverteTTipoContaParaTipoConta recebe um ponteiro do tipo da struct TTipoConta como parâmetro e retorna um ponteiro do tipo TipoConta
func ConverteTTipoContaParaTipoConta(tc *tipo_conta.TTipoConta) *tipo_conta.TipoConta {
	return &tipo_conta.TipoConta{
		Nome:             tc.Nome,
		DescricaoDebito:  tc.DescricaoDebito,
		DescricaoCredito: tc.DescricaoCredito,
		DataCriacao:      tc.DataCriacao,
		DataModificacao:  tc.DataModificacao,
		Estado:           tc.Estado,
	}
}

// ConverteContaParaTConta recebe um ponteiro do tipo da struct Conta como parâmetro e retorna um ponteiro do tipo TConta
func ConverteContaParaTConta(c *conta.Conta) *conta.TConta {
	return &conta.TConta{
		Nome:            c.Nome,
		NomeTipoConta:   c.NomeTipoConta,
		Codigo:          setNullString(c.Codigo),
		ContaPai:        setNullString(c.ContaPai),
		Comentario:      setNullString(c.Comentario),
		DataCriacao:     c.DataCriacao,
		DataModificacao: c.DataModificacao,
		Estado:          c.Estado,
	}
}

// ConverteTContaParaConta recebe um ponteiro do tipo da struct TConta como parâmetro e retorna um ponteiro do tipo Conta
func ConverteTContaParaConta(c *conta.TConta) *conta.Conta {
	return &conta.Conta{
		Nome:            c.Nome,
		NomeTipoConta:   c.NomeTipoConta,
		Codigo:          c.Codigo.String,
		ContaPai:        c.ContaPai.String,
		Comentario:      c.Comentario.String,
		DataCriacao:     c.DataCriacao,
		DataModificacao: c.DataModificacao,
		Estado:          c.Estado,
	}
}

// ConverteTPessoasParaPessoas recebe um resultado do tipo *gorm.DB e um slice de TPessoas e retorna um slice de Pessoas e um erro != nil se ocorrer algum problema
func ConverteTContasParaContas(resultado *gorm.DB, tContas *conta.TContas) (conta.Contas, error) {
	err := resultado.Error
	if err != nil {
		return nil, err
	}

	contas := conta.Contas{}
	encontrouRegistros := resultado.RowsAffected > 0
	if encontrouRegistros {
		for _, tc := range *tContas {
			c := ConverteTContaParaConta(tc)
			contas = append(contas, c)
		}
	}

	return contas, nil
}

// ConverteLancamentoParaTLancamento recebe um ponteiro do tipo da struct Lancamento como parâmetro e retorna um ponteiro do tipo TLancamento
func ConverteLancamentoParaTLancamento(l *lancamento.Lancamento) *lancamento.TLancamento {
	return &lancamento.TLancamento{
		ID:              l.ID,
		CpfPessoa:       l.CpfPessoa,
		Data:            l.Data,
		Numero:          setNullString(l.Numero),
		Descricao:       l.Descricao,
		DataCriacao:     l.DataCriacao,
		DataModificacao: l.DataModificacao,
		Estado:          l.Estado,
	}
}

// ConverteTLancamentoParaLancamento recebe um ponteiro do tipo da struct TLancamento como parâmetro e retorna um ponteiro do tipo Lancamento
func ConverteTLancamentoParaLancamento(l *lancamento.TLancamento) *lancamento.Lancamento {
	return &lancamento.Lancamento{
		ID:              l.ID,
		CpfPessoa:       l.CpfPessoa,
		Data:            l.Data,
		Numero:          l.Numero.String,
		Descricao:       l.Descricao,
		DataCriacao:     l.DataCriacao,
		DataModificacao: l.DataModificacao,
		Estado:          l.Estado,
	}
}

// ConverteDetalheLancamentoParaTDetalheLancamento recebe um ponteiro do tipo da struct DetalheLancamento como parâmetro e retorna um ponteiro do tipo TDetalheLancamento
func ConverteDetalheLancamentoParaTDetalheLancamento(dl *detalhe_lancamento.DetalheLancamento) *detalhe_lancamento.TDetalheLancamento {
	return &detalhe_lancamento.TDetalheLancamento{
		IDLancamento: dl.IDLancamento,
		NomeConta:    dl.NomeConta,
		Debito:       setNullFloat64(dl.Debito),
		Credito:      setNullFloat64(dl.Credito),
	}
}

// ConverteTDetalheLancamentoParaDetalheLancamento recebe um ponteiro do tipo da struct TDetalheLancamento como parâmetro e retorna um ponteiro do tipo DetalheLancamento
func ConverteTDetalheLancamentoParaDetalheLancamento(dl *detalhe_lancamento.TDetalheLancamento) *detalhe_lancamento.DetalheLancamento {
	return &detalhe_lancamento.DetalheLancamento{
		IDLancamento: dl.IDLancamento,
		NomeConta:    dl.NomeConta,
		Debito:       dl.Debito.Float64,
		Credito:      dl.Credito.Float64,
	}
}

// ConverteTPessoasParaPessoas recebe um resultado do tipo *gorm.DB e um slice de TPessoas e retorna um slice de Pessoas e um erro != nil se ocorrer algum problema
func ConverteTPessoasParaPessoas(resultado *gorm.DB, tpessoas *pessoa.TPessoas) (pessoa.Pessoas, error) {
	err := resultado.Error
	if err != nil {
		return nil, err
	}

	pessoas := pessoa.Pessoas{}
	encontrouRegistros := resultado.RowsAffected > 0
	if encontrouRegistros {
		for _, tp := range *tpessoas {
			p := ConverteTPessoaParaPessoa(tp)
			pessoas = append(pessoas, p)
		}
	}

	return pessoas, nil
}

// ConverteTPessoasParaPessoasSimples recebe um resultado do tipo *gorm.DB e um slice de TPessoas e retorna um slice de PessoasSimples e um erro != nil se ocorrer algum problema
func ConverteTPessoasParaPessoasSimples(resultado *gorm.DB, tpessoas *pessoa.TPessoas) (pessoa.PessoasSimples, error) {
	err := resultado.Error
	if err != nil {
		return nil, err
	}

	pessoasSimples := pessoa.PessoasSimples{}
	encontrouRegistros := resultado.RowsAffected > 0
	if encontrouRegistros {
		for _, tp := range *tpessoas {
			p := ConverteTPessoaParaPessoaSimples(tp)
			pessoasSimples = append(pessoasSimples, p)
		}
	}

	return pessoasSimples, nil
}

// ConverteTTiposContaParaTiposConta recebe um resultado do tipo *gorm.DB e um slice de TTiposConta e retorna um slice de TiposConta e um erro != nil se ocorrer algum problema
func ConverteTTiposContaParaTiposConta(resultado *gorm.DB, ttiposConta *tipo_conta.TTiposConta) (tipo_conta.TiposConta, error) {
	err := resultado.Error
	if err != nil {
		return nil, err
	}

	tiposConta := tipo_conta.TiposConta{}
	encontrouRegistros := resultado.RowsAffected > 0
	if encontrouRegistros {
		for _, ttc := range *ttiposConta {
			tc := ConverteTTipoContaParaTipoConta(ttc)
			tiposConta = append(tiposConta, tc)
		}
	}

	return tiposConta, nil
}

// ConverteTLancamentosParaLancamentos recebe um resultado do tipo *gorm.DB e um slice de TLancamentos e retorna um slice de Lancamentos e um erro != nil se ocorrer algum problema
func ConverteTLancamentosParaLancamentos(resultado *gorm.DB, tlancamentos *lancamento.TLancamentos) (lancamento.Lancamentos, error) {
	err := resultado.Error
	if err != nil {
		return nil, err
	}

	lancamentos := lancamento.Lancamentos{}
	encontrouRegistros := resultado.RowsAffected > 0
	if encontrouRegistros {
		for _, tp := range *tlancamentos {
			l := ConverteTLancamentoParaLancamento(tp)
			lancamentos = append(lancamentos, l)
		}
	}

	return lancamentos, nil
}

// VerificaQuantidadeLinhasAfetadas retorna um erro != nil se a quantidade esperada(quantiadeEsperada int64) for diferente da quantidade realmente afetada obtida a partir do parâmetro tx(*gorm.DB)
func VerificaQuantidadeRegistrosAfetados(tx *gorm.DB, quantidadeEsperada int64) error {
	quantidadeAfetada := tx.RowsAffected
	if quantidadeAfetada != quantidadeEsperada {
		return fmt.Errorf("quantidade de registros afetados errada. Esperado %d, afetado %d", quantidadeEsperada, quantidadeAfetada)
	}
	return nil
}
