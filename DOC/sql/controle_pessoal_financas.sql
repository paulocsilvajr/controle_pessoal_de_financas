CREATE DOMAIN public.codigo_numero
    AS INTEGER
    NOT NULL;
	
CREATE DOMAIN public.codigo_texto
    AS VARCHAR(19);
	
CREATE DOMAIN public.cpf
    AS VARCHAR(11)
    NOT NULL;
	
CREATE DOMAIN public.data_completa
    AS TIMESTAMP WITH TIME ZONE;
	
CREATE DOMAIN public.descricao_curta
    AS VARCHAR(20)
    NOT NULL;

CREATE DOMAIN public.dinheiro
    AS NUMERIC(19,3);

CREATE DOMAIN public.nome_conta
    AS VARCHAR(50);

CREATE TABLE public.tipo_conta (
                nome nome_conta NOT NULL,
                descricao_debito descricao_curta NOT NULL,
                descricao_credito descricao_curta NOT NULL,
                data_criacao data_completa NOT NULL,
                data_modificacao data_completa DEFAULT now() NOT NULL,
                estado BOOLEAN DEFAULT true NOT NULL,
                CONSTRAINT tipo_conta_pk PRIMARY KEY (nome)
);

CREATE TABLE public.conta (
                nome nome_conta NOT NULL,
                nome_tipo_conta nome_conta NOT NULL,
                codigo codigo_texto,
                conta_pai nome_conta,
                comentario VARCHAR(150),
                data_criacao data_completa NOT NULL,
                data_modificacao data_completa DEFAULT now() NOT NULL,
                estado BOOLEAN DEFAULT true NOT NULL,
                CONSTRAINT conta_pk PRIMARY KEY (nome)
);

CREATE TABLE public.pessoa (
                cpf cpf NOT NULL,
                nome_completo VARCHAR(100),
                usuario VARCHAR(20) NOT NULL,
                senha VARCHAR(64) NOT NULL,
                email VARCHAR(45) NOT NULL,
                data_criacao data_completa NOT NULL,
                data_modificacao data_completa DEFAULT now() NOT NULL,
                estado BOOLEAN DEFAULT true NOT NULL,
                CONSTRAINT pessoa_pk PRIMARY KEY (cpf)
);

CREATE TABLE public.log (
                id codigo_numero NOT NULL,
                cpf_pessoa cpf NOT NULL,
                data data_completa NOT NULL,
                sql TEXT NOT NULL,
                CONSTRAINT log_pk PRIMARY KEY (id)
);

CREATE TABLE public.lancamento (
                id codigo_numero NOT NULL,
                cpf_pessoa cpf NOT NULL,
                data data_completa DEFAULT now() NOT NULL,
                numero codigo_numero,
                descricao VARCHAR(100),
                data_criacao data_completa NOT NULL,
                data_modificacao data_completa DEFAULT now() NOT NULL,
                estado BOOLEAN DEFAULT true NOT NULL,
                CONSTRAINT lancamento_pk PRIMARY KEY (id)
);

CREATE TABLE public.detalhe_lancamento (
                id_lancamento codigo_numero NOT NULL,
                nome_conta nome_conta NOT NULL,
                debito dinheiro,
                credito dinheiro,
                CONSTRAINT detalhe_lancamento_pk PRIMARY KEY (id_lancamento, nome_conta)
);

ALTER TABLE public.conta ADD CONSTRAINT tipo_conta_conta_fk
FOREIGN KEY (nome_tipo_conta)
REFERENCES public.tipo_conta (nome)
ON DELETE RESTRICT
ON UPDATE CASCADE;

ALTER TABLE public.detalhe_lancamento ADD CONSTRAINT conta_detalhe_lancamento_fk
FOREIGN KEY (nome_conta)
REFERENCES public.conta (nome)
ON DELETE RESTRICT
ON UPDATE CASCADE;

ALTER TABLE public.lancamento ADD CONSTRAINT pessoa_lancamento_fk
FOREIGN KEY (cpf_pessoa)
REFERENCES public.pessoa (cpf)
ON DELETE CASCADE
ON UPDATE CASCADE;

ALTER TABLE public.log ADD CONSTRAINT pessoa_log_fk
FOREIGN KEY (cpf_pessoa)
REFERENCES public.pessoa (cpf)
ON DELETE CASCADE
ON UPDATE CASCADE;

ALTER TABLE public.detalhe_lancamento ADD CONSTRAINT lancamento_detalhe_lancamento_fk
FOREIGN KEY (id_lancamento)
REFERENCES public.lancamento (id)
ON DELETE CASCADE
ON UPDATE CASCADE;


 