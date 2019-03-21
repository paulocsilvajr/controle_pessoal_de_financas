
CREATE TABLE public.tipo_conta (
                nome VARCHAR(50) NOT NULL,
                descricao_debito VARCHAR(20) NOT NULL,
                descricao_credito VARCHAR(20) NOT NULL,
                data_criacao TIMESTAMP NOT NULL,
                data_modificacao TIMESTAMP DEFAULT now() NOT NULL,
                estado BOOLEAN DEFAULT true NOT NULL,
                CONSTRAINT tipo_conta_pk PRIMARY KEY (nome)
);


CREATE TABLE public.conta (
                nome VARCHAR(50) NOT NULL,
                nome_tipo_conta VARCHAR(50) NOT NULL,
                codigo VARCHAR(19),
                conta_pai VARCHAR(50),
                comentario VARCHAR(150),
                data_criacao TIMESTAMP NOT NULL,
                data_modificacao TIMESTAMP DEFAULT now() NOT NULL,
                estado BOOLEAN DEFAULT true NOT NULL,
                CONSTRAINT conta_pk PRIMARY KEY (nome)
);


CREATE TABLE public.pessoa (
                cpf VARCHAR(11) NOT NULL,
                nome_completo VARCHAR(100),
                usuario VARCHAR(20) NOT NULL,
                senha VARCHAR(64) NOT NULL,
                email VARCHAR(45) NOT NULL,
                data_criacao TIMESTAMP NOT NULL,
                data_modificacao TIMESTAMP DEFAULT now() NOT NULL,
                estado BOOLEAN DEFAULT true NOT NULL,
                CONSTRAINT pessoa_pk PRIMARY KEY (cpf)
);


CREATE TABLE public.log (
                id INTEGER NOT NULL,
                cpf_pessoa VARCHAR(11) NOT NULL,
                data TIMESTAMP NOT NULL,
                sql VARCHAR(500) NOT NULL,
                CONSTRAINT log_pk PRIMARY KEY (id)
);


CREATE TABLE public.lancamento (
                id INTEGER NOT NULL,
                cpf_pessoa VARCHAR(11) NOT NULL,
                data TIMESTAMP DEFAULT now() NOT NULL,
                numero VARCHAR(19),
                descricao VARCHAR(100),
                data_criacao TIMESTAMP NOT NULL,
                data_modificacao TIMESTAMP DEFAULT now() NOT NULL,
                estado BOOLEAN DEFAULT true NOT NULL,
                CONSTRAINT lancamento_pk PRIMARY KEY (id)
);


CREATE TABLE public.detalhe_lancamento (
                id_lancamento INTEGER NOT NULL,
                nome_conta VARCHAR(50) NOT NULL,
                debito NUMERIC(19,3),
                credito NUMERIC(19,3),
                CONSTRAINT detalhe_lancamento_pk PRIMARY KEY (id_lancamento, nome_conta)
);


ALTER TABLE public.conta ADD CONSTRAINT tipo_conta_conta_fk
FOREIGN KEY (nome_tipo_conta)
REFERENCES public.tipo_conta (nome)
ON DELETE RESTRICT
ON UPDATE CASCADE
NOT DEFERRABLE;

ALTER TABLE public.detalhe_lancamento ADD CONSTRAINT conta_detalhe_lancamento_fk
FOREIGN KEY (nome_conta)
REFERENCES public.conta (nome)
ON DELETE RESTRICT
ON UPDATE CASCADE
NOT DEFERRABLE;

ALTER TABLE public.lancamento ADD CONSTRAINT pessoa_lancamento_fk
FOREIGN KEY (cpf_pessoa)
REFERENCES public.pessoa (cpf)
ON DELETE CASCADE
ON UPDATE CASCADE
NOT DEFERRABLE;

ALTER TABLE public.log ADD CONSTRAINT pessoa_log_fk
FOREIGN KEY (cpf_pessoa)
REFERENCES public.pessoa (cpf)
ON DELETE CASCADE
ON UPDATE CASCADE
NOT DEFERRABLE;

ALTER TABLE public.detalhe_lancamento ADD CONSTRAINT lancamento_detalhe_lancamento_fk
FOREIGN KEY (id_lancamento)
REFERENCES public.lancamento (id)
ON DELETE CASCADE
ON UPDATE CASCADE
NOT DEFERRABLE;
