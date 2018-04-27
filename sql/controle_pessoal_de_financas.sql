--
-- PostgreSQL database dump
--

-- Dumped from database version 9.6.7
-- Dumped by pg_dump version 9.6.7

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

--
-- Name: dm_cpf; Type: DOMAIN; Schema: public; Owner: pi
--

CREATE DOMAIN dm_cpf AS character(11) NOT NULL;


ALTER DOMAIN dm_cpf OWNER TO pi;

--
-- Name: dm_dinheiro; Type: DOMAIN; Schema: public; Owner: pi
--

CREATE DOMAIN dm_dinheiro AS numeric(19,3) NOT NULL;


ALTER DOMAIN dm_dinheiro OWNER TO pi;

--
-- Name: dm_nome_conta; Type: DOMAIN; Schema: public; Owner: pi
--

CREATE DOMAIN dm_nome_conta AS character varying(50) NOT NULL;


ALTER DOMAIN dm_nome_conta OWNER TO pi;

--
-- Name: dm_nome_tipo_conta; Type: DOMAIN; Schema: public; Owner: pi
--

CREATE DOMAIN dm_nome_tipo_conta AS character varying(30) NOT NULL;


ALTER DOMAIN dm_nome_tipo_conta OWNER TO pi;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: conta; Type: TABLE; Schema: public; Owner: pi
--

CREATE TABLE conta (
    nome dm_nome_conta NOT NULL,
    tipo_conta_nome dm_nome_tipo_conta NOT NULL,
    codigo character varying(19),
    conta_pai dm_nome_conta,
    comentario character varying(150)
);


ALTER TABLE conta OWNER TO pi;

--
-- Name: detalhe_lancamento; Type: TABLE; Schema: public; Owner: pi
--

CREATE TABLE detalhe_lancamento (
    id integer NOT NULL,
    conta_nome dm_nome_conta NOT NULL,
    lancamento_id integer NOT NULL,
    debito dm_dinheiro NOT NULL,
    credito dm_dinheiro NOT NULL
);


ALTER TABLE detalhe_lancamento OWNER TO pi;

--
-- Name: detalhe_lancamento_id_seq; Type: SEQUENCE; Schema: public; Owner: pi
--

CREATE SEQUENCE detalhe_lancamento_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE detalhe_lancamento_id_seq OWNER TO pi;

--
-- Name: detalhe_lancamento_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: pi
--

ALTER SEQUENCE detalhe_lancamento_id_seq OWNED BY detalhe_lancamento.id;


--
-- Name: lancamento; Type: TABLE; Schema: public; Owner: pi
--

CREATE TABLE lancamento (
    id integer NOT NULL,
    pessoa_cpf dm_cpf NOT NULL,
    data timestamp without time zone DEFAULT now() NOT NULL,
    numero character varying(19),
    descricao character varying(100)
);


ALTER TABLE lancamento OWNER TO pi;

--
-- Name: lancamento_id_seq; Type: SEQUENCE; Schema: public; Owner: pi
--

CREATE SEQUENCE lancamento_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE lancamento_id_seq OWNER TO pi;

--
-- Name: lancamento_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: pi
--

ALTER SEQUENCE lancamento_id_seq OWNED BY lancamento.id;


--
-- Name: log; Type: TABLE; Schema: public; Owner: pi
--

CREATE TABLE log (
    id integer NOT NULL,
    pessoa_cpf dm_cpf NOT NULL,
    data timestamp without time zone DEFAULT now() NOT NULL,
    tabela character varying(20) NOT NULL,
    acao character varying(20) NOT NULL,
    dados text NOT NULL
);


ALTER TABLE log OWNER TO pi;

--
-- Name: log_id_seq; Type: SEQUENCE; Schema: public; Owner: pi
--

CREATE SEQUENCE log_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE log_id_seq OWNER TO pi;

--
-- Name: log_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: pi
--

ALTER SEQUENCE log_id_seq OWNED BY log.id;


--
-- Name: pessoa; Type: TABLE; Schema: public; Owner: pi
--

CREATE TABLE pessoa (
    cpf dm_cpf NOT NULL,
    nome character varying(100),
    usuario character varying(20) NOT NULL,
    senha character varying(64) NOT NULL,
    email character varying(45) NOT NULL
);


ALTER TABLE pessoa OWNER TO pi;

--
-- Name: tipo_conta; Type: TABLE; Schema: public; Owner: pi
--

CREATE TABLE tipo_conta (
    nome dm_nome_tipo_conta NOT NULL,
    descricao_debito character varying(20) NOT NULL,
    descricao_credito character varying(20) NOT NULL
);


ALTER TABLE tipo_conta OWNER TO pi;

--
-- Name: detalhe_lancamento id; Type: DEFAULT; Schema: public; Owner: pi
--

ALTER TABLE ONLY detalhe_lancamento ALTER COLUMN id SET DEFAULT nextval('detalhe_lancamento_id_seq'::regclass);


--
-- Name: lancamento id; Type: DEFAULT; Schema: public; Owner: pi
--

ALTER TABLE ONLY lancamento ALTER COLUMN id SET DEFAULT nextval('lancamento_id_seq'::regclass);


--
-- Name: log id; Type: DEFAULT; Schema: public; Owner: pi
--

ALTER TABLE ONLY log ALTER COLUMN id SET DEFAULT nextval('log_id_seq'::regclass);


--
-- Data for Name: conta; Type: TABLE DATA; Schema: public; Owner: pi
--

COPY conta (nome, tipo_conta_nome, codigo, conta_pai, comentario) FROM stdin;
\.


--
-- Data for Name: detalhe_lancamento; Type: TABLE DATA; Schema: public; Owner: pi
--

COPY detalhe_lancamento (id, conta_nome, lancamento_id, debito, credito) FROM stdin;
\.


--
-- Name: detalhe_lancamento_id_seq; Type: SEQUENCE SET; Schema: public; Owner: pi
--

SELECT pg_catalog.setval('detalhe_lancamento_id_seq', 1, false);


--
-- Data for Name: lancamento; Type: TABLE DATA; Schema: public; Owner: pi
--

COPY lancamento (id, pessoa_cpf, data, numero, descricao) FROM stdin;
\.


--
-- Name: lancamento_id_seq; Type: SEQUENCE SET; Schema: public; Owner: pi
--

SELECT pg_catalog.setval('lancamento_id_seq', 1, false);


--
-- Data for Name: log; Type: TABLE DATA; Schema: public; Owner: pi
--

COPY log (id, pessoa_cpf, data, tabela, acao, dados) FROM stdin;
\.


--
-- Name: log_id_seq; Type: SEQUENCE SET; Schema: public; Owner: pi
--

SELECT pg_catalog.setval('log_id_seq', 1, false);


--
-- Data for Name: pessoa; Type: TABLE DATA; Schema: public; Owner: pi
--

COPY pessoa (cpf, nome, usuario, senha, email) FROM stdin;
\.


--
-- Data for Name: tipo_conta; Type: TABLE DATA; Schema: public; Owner: pi
--

COPY tipo_conta (nome, descricao_debito, descricao_credito) FROM stdin;
\.


--
-- Name: conta conta_pkey; Type: CONSTRAINT; Schema: public; Owner: pi
--

ALTER TABLE ONLY conta
    ADD CONSTRAINT conta_pkey PRIMARY KEY (nome);


--
-- Name: detalhe_lancamento detalhe_lancamento_pkey; Type: CONSTRAINT; Schema: public; Owner: pi
--

ALTER TABLE ONLY detalhe_lancamento
    ADD CONSTRAINT detalhe_lancamento_pkey PRIMARY KEY (id);


--
-- Name: lancamento lancamento_pkey; Type: CONSTRAINT; Schema: public; Owner: pi
--

ALTER TABLE ONLY lancamento
    ADD CONSTRAINT lancamento_pkey PRIMARY KEY (id);


--
-- Name: log log_pkey; Type: CONSTRAINT; Schema: public; Owner: pi
--

ALTER TABLE ONLY log
    ADD CONSTRAINT log_pkey PRIMARY KEY (id);


--
-- Name: pessoa pessoa_pkey; Type: CONSTRAINT; Schema: public; Owner: pi
--

ALTER TABLE ONLY pessoa
    ADD CONSTRAINT pessoa_pkey PRIMARY KEY (cpf);


--
-- Name: tipo_conta tipo_conta_pkey; Type: CONSTRAINT; Schema: public; Owner: pi
--

ALTER TABLE ONLY tipo_conta
    ADD CONSTRAINT tipo_conta_pkey PRIMARY KEY (nome);


--
-- Name: conta fk_conta_tipo_conta_nome; Type: FK CONSTRAINT; Schema: public; Owner: pi
--

ALTER TABLE ONLY conta
    ADD CONSTRAINT fk_conta_tipo_conta_nome FOREIGN KEY (tipo_conta_nome) REFERENCES tipo_conta(nome) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: detalhe_lancamento fk_detalhe_lancamento_conta_nome; Type: FK CONSTRAINT; Schema: public; Owner: pi
--

ALTER TABLE ONLY detalhe_lancamento
    ADD CONSTRAINT fk_detalhe_lancamento_conta_nome FOREIGN KEY (conta_nome) REFERENCES conta(nome) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: detalhe_lancamento fk_detalhe_lancamento_lancamento_id; Type: FK CONSTRAINT; Schema: public; Owner: pi
--

ALTER TABLE ONLY detalhe_lancamento
    ADD CONSTRAINT fk_detalhe_lancamento_lancamento_id FOREIGN KEY (lancamento_id) REFERENCES lancamento(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: lancamento fk_lancamento_pessoa_cpf; Type: FK CONSTRAINT; Schema: public; Owner: pi
--

ALTER TABLE ONLY lancamento
    ADD CONSTRAINT fk_lancamento_pessoa_cpf FOREIGN KEY (pessoa_cpf) REFERENCES pessoa(cpf) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: log fk_log_pessoa_cpf; Type: FK CONSTRAINT; Schema: public; Owner: pi
--

ALTER TABLE ONLY log
    ADD CONSTRAINT fk_log_pessoa_cpf FOREIGN KEY (pessoa_cpf) REFERENCES pessoa(cpf) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

