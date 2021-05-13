package dao

import (
	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/model/pessoa"
	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/model/tipo_conta"
	"gorm.io/gorm"
)

// Definição de campos GORM de Pessoa e Método TableName() em model/pessoa.dao na struct Pessoa

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

func CriarTabelaPessoa(db *gorm.DB) error {
	return db.AutoMigrate(&pessoa.TPessoa{})
}

func CriarTabelaTipoConta(db *gorm.DB) error {
	return db.AutoMigrate(&tipo_conta.TTipoConta{})
}
