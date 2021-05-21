package dao

import (
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

func CriarTabelaPessoa(db *gorm.DB) error {
	return db.AutoMigrate(&pessoa.TPessoa{})
}

func CriarTabelaTipoConta(db *gorm.DB) error {
	return db.AutoMigrate(&tipo_conta.TTipoConta{})
}

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
ADD	CONSTRAINT conta_detalhe_lancamento_fk FOREIGN KEY ({{.fkConta}})
REFERENCES {{.tabelaConta}}({{.fkConta}})
ON UPDATE CASCADE
ON DELETE CASCADE;

ALTER TABLE {{.tabela}}
ADD	CONSTRAINT lancamento_detalhe_lancamento_fk FOREIGN KEY ({{.fkLancamento}})
REFERENCES {{.tabelaLancamento}}({{.fkLancamento}})
ON UPDATE CASCADE
ON DELETE CASCADE;
`
	sql = getTemplateSQL("detalheLancamentoFK", sql, detalheLancamentoDB)

	return db.Exec(sql).Error
}
