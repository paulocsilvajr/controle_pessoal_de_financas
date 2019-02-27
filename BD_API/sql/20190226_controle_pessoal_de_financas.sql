--
-- PostgreSQL database dump
--

-- Dumped from database version 9.6.11
-- Dumped by pg_dump version 9.6.11

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
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


--
-- Name: dm_cpf; Type: DOMAIN; Schema: public; Owner: pi
--

CREATE DOMAIN public.dm_cpf AS character(11) NOT NULL;


ALTER DOMAIN public.dm_cpf OWNER TO pi;

--
-- Name: dm_dinheiro; Type: DOMAIN; Schema: public; Owner: pi
--

CREATE DOMAIN public.dm_dinheiro AS numeric(19,3) NOT NULL;


ALTER DOMAIN public.dm_dinheiro OWNER TO pi;

--
-- Name: dm_nome_conta; Type: DOMAIN; Schema: public; Owner: pi
--

CREATE DOMAIN public.dm_nome_conta AS character varying(50);


ALTER DOMAIN public.dm_nome_conta OWNER TO pi;

--
-- Name: dm_nome_tipo_conta; Type: DOMAIN; Schema: public; Owner: pi
--

CREATE DOMAIN public.dm_nome_tipo_conta AS character varying(30) NOT NULL;


ALTER DOMAIN public.dm_nome_tipo_conta OWNER TO pi;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: conta; Type: TABLE; Schema: public; Owner: pi
--

CREATE TABLE public.conta (
    nome public.dm_nome_conta NOT NULL,
    tipo_conta_nome public.dm_nome_tipo_conta NOT NULL,
    codigo character varying(19),
    conta_pai public.dm_nome_conta,
    comentario character varying(150)
);


ALTER TABLE public.conta OWNER TO pi;

--
-- Name: detalhe_lancamento; Type: TABLE; Schema: public; Owner: pi
--

CREATE TABLE public.detalhe_lancamento (
    id integer NOT NULL,
    conta_nome public.dm_nome_conta NOT NULL,
    lancamento_id integer NOT NULL,
    debito public.dm_dinheiro NOT NULL,
    credito public.dm_dinheiro NOT NULL
);


ALTER TABLE public.detalhe_lancamento OWNER TO pi;

--
-- Name: detalhe_lancamento_id_seq; Type: SEQUENCE; Schema: public; Owner: pi
--

CREATE SEQUENCE public.detalhe_lancamento_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.detalhe_lancamento_id_seq OWNER TO pi;

--
-- Name: detalhe_lancamento_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: pi
--

ALTER SEQUENCE public.detalhe_lancamento_id_seq OWNED BY public.detalhe_lancamento.id;


--
-- Name: lancamento; Type: TABLE; Schema: public; Owner: pi
--

CREATE TABLE public.lancamento (
    id integer NOT NULL,
    pessoa_cpf public.dm_cpf NOT NULL,
    data timestamp without time zone DEFAULT now() NOT NULL,
    numero character varying(19),
    descricao character varying(100)
);


ALTER TABLE public.lancamento OWNER TO pi;

--
-- Name: lancamento_id_seq; Type: SEQUENCE; Schema: public; Owner: pi
--

CREATE SEQUENCE public.lancamento_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.lancamento_id_seq OWNER TO pi;

--
-- Name: lancamento_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: pi
--

ALTER SEQUENCE public.lancamento_id_seq OWNED BY public.lancamento.id;


--
-- Name: log; Type: TABLE; Schema: public; Owner: pi
--

CREATE TABLE public.log (
    id integer NOT NULL,
    pessoa_cpf public.dm_cpf NOT NULL,
    data timestamp without time zone DEFAULT now() NOT NULL,
    tabela character varying(20) NOT NULL,
    acao character varying(20) NOT NULL,
    dados text NOT NULL
);


ALTER TABLE public.log OWNER TO pi;

--
-- Name: log_id_seq; Type: SEQUENCE; Schema: public; Owner: pi
--

CREATE SEQUENCE public.log_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.log_id_seq OWNER TO pi;

--
-- Name: log_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: pi
--

ALTER SEQUENCE public.log_id_seq OWNED BY public.log.id;


--
-- Name: pessoa; Type: TABLE; Schema: public; Owner: pi
--

CREATE TABLE public.pessoa (
    cpf public.dm_cpf NOT NULL,
    nome character varying(100),
    usuario character varying(20) NOT NULL,
    senha character varying(64) NOT NULL,
    email character varying(45) NOT NULL
);


ALTER TABLE public.pessoa OWNER TO pi;

--
-- Name: tipo_conta; Type: TABLE; Schema: public; Owner: pi
--

CREATE TABLE public.tipo_conta (
    nome public.dm_nome_tipo_conta NOT NULL,
    descricao_debito character varying(20) NOT NULL,
    descricao_credito character varying(20) NOT NULL
);


ALTER TABLE public.tipo_conta OWNER TO pi;

--
-- Name: detalhe_lancamento id; Type: DEFAULT; Schema: public; Owner: pi
--

ALTER TABLE ONLY public.detalhe_lancamento ALTER COLUMN id SET DEFAULT nextval('public.detalhe_lancamento_id_seq'::regclass);


--
-- Name: lancamento id; Type: DEFAULT; Schema: public; Owner: pi
--

ALTER TABLE ONLY public.lancamento ALTER COLUMN id SET DEFAULT nextval('public.lancamento_id_seq'::regclass);


--
-- Name: log id; Type: DEFAULT; Schema: public; Owner: pi
--

ALTER TABLE ONLY public.log ALTER COLUMN id SET DEFAULT nextval('public.log_id_seq'::regclass);


--
-- Data for Name: conta; Type: TABLE DATA; Schema: public; Owner: pi
--

COPY public.conta (nome, tipo_conta_nome, codigo, conta_pai, comentario) FROM stdin;
ativos	ativo	1	\N	
despesas	despesa	2	\N	
líquidos	líquido	3	\N	
passivos	passivo	4	\N	
receitas	receita	5	\N	
conta corrente	banco	6	ativos	
conta poupança	banco	7	ativos	
dinheiro em carteira	carteira	8	ativos	
nu conta	banco	9	conta corrente	
bradesco	banco	10	conta poupança	
caixa econômica	banco	11	conta poupança	
internet	despesa	12	despesas	
telefone	despesa	13	despesas	
serviços	despesa	14	despesas	
eletricidade	despesa	15	serviços	
refeições fora	despesa	16	despesas	
computador	despesa	17	despesas	
cartão de crédito	passivo	18	passivos	
nubank	passivo	18	cartão de crédito	
submarino	passivo	19	cartão de crédito	
salário	receita	20	receitas	
juros recebidos	receita	21	receitas	
\.


--
-- Data for Name: detalhe_lancamento; Type: TABLE DATA; Schema: public; Owner: pi
--

COPY public.detalhe_lancamento (id, conta_nome, lancamento_id, debito, credito) FROM stdin;
\.


--
-- Name: detalhe_lancamento_id_seq; Type: SEQUENCE SET; Schema: public; Owner: pi
--

SELECT pg_catalog.setval('public.detalhe_lancamento_id_seq', 1, false);


--
-- Data for Name: lancamento; Type: TABLE DATA; Schema: public; Owner: pi
--

COPY public.lancamento (id, pessoa_cpf, data, numero, descricao) FROM stdin;
\.


--
-- Name: lancamento_id_seq; Type: SEQUENCE SET; Schema: public; Owner: pi
--

SELECT pg_catalog.setval('public.lancamento_id_seq', 1, false);


--
-- Data for Name: log; Type: TABLE DATA; Schema: public; Owner: pi
--

COPY public.log (id, pessoa_cpf, data, tabela, acao, dados) FROM stdin;
\.


--
-- Name: log_id_seq; Type: SEQUENCE SET; Schema: public; Owner: pi
--

SELECT pg_catalog.setval('public.log_id_seq', 1, false);


--
-- Data for Name: pessoa; Type: TABLE DATA; Schema: public; Owner: pi
--

COPY public.pessoa (cpf, nome, usuario, senha, email) FROM stdin;
\.


--
-- Data for Name: tipo_conta; Type: TABLE DATA; Schema: public; Owner: pi
--

COPY public.tipo_conta (nome, descricao_debito, descricao_credito) FROM stdin;
banco	saque	depósito
carteira	gastar	receber
despesa	desconto	despesa
cartão de crédito	cobrar	pagamento
ativo	descrescer	aumentar
líquido	aumentar	descrescer
passivo	aumentar	descrescer
receita	receita	cobrar
\.


--
-- Name: conta conta_pkey; Type: CONSTRAINT; Schema: public; Owner: pi
--

ALTER TABLE ONLY public.conta
    ADD CONSTRAINT conta_pkey PRIMARY KEY (nome);


--
-- Name: detalhe_lancamento detalhe_lancamento_pkey; Type: CONSTRAINT; Schema: public; Owner: pi
--

ALTER TABLE ONLY public.detalhe_lancamento
    ADD CONSTRAINT detalhe_lancamento_pkey PRIMARY KEY (id);


--
-- Name: lancamento lancamento_pkey; Type: CONSTRAINT; Schema: public; Owner: pi
--

ALTER TABLE ONLY public.lancamento
    ADD CONSTRAINT lancamento_pkey PRIMARY KEY (id);


--
-- Name: log log_pkey; Type: CONSTRAINT; Schema: public; Owner: pi
--

ALTER TABLE ONLY public.log
    ADD CONSTRAINT log_pkey PRIMARY KEY (id);


--
-- Name: pessoa pessoa_pkey; Type: CONSTRAINT; Schema: public; Owner: pi
--

ALTER TABLE ONLY public.pessoa
    ADD CONSTRAINT pessoa_pkey PRIMARY KEY (cpf);


--
-- Name: tipo_conta tipo_conta_pkey; Type: CONSTRAINT; Schema: public; Owner: pi
--

ALTER TABLE ONLY public.tipo_conta
    ADD CONSTRAINT tipo_conta_pkey PRIMARY KEY (nome);


--
-- Name: conta fk_conta_tipo_conta_nome; Type: FK CONSTRAINT; Schema: public; Owner: pi
--

ALTER TABLE ONLY public.conta
    ADD CONSTRAINT fk_conta_tipo_conta_nome FOREIGN KEY (tipo_conta_nome) REFERENCES public.tipo_conta(nome) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: detalhe_lancamento fk_detalhe_lancamento_conta_nome; Type: FK CONSTRAINT; Schema: public; Owner: pi
--

ALTER TABLE ONLY public.detalhe_lancamento
    ADD CONSTRAINT fk_detalhe_lancamento_conta_nome FOREIGN KEY (conta_nome) REFERENCES public.conta(nome) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: detalhe_lancamento fk_detalhe_lancamento_lancamento_id; Type: FK CONSTRAINT; Schema: public; Owner: pi
--

ALTER TABLE ONLY public.detalhe_lancamento
    ADD CONSTRAINT fk_detalhe_lancamento_lancamento_id FOREIGN KEY (lancamento_id) REFERENCES public.lancamento(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: lancamento fk_lancamento_pessoa_cpf; Type: FK CONSTRAINT; Schema: public; Owner: pi
--

ALTER TABLE ONLY public.lancamento
    ADD CONSTRAINT fk_lancamento_pessoa_cpf FOREIGN KEY (pessoa_cpf) REFERENCES public.pessoa(cpf) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: log fk_log_pessoa_cpf; Type: FK CONSTRAINT; Schema: public; Owner: pi
--

ALTER TABLE ONLY public.log
    ADD CONSTRAINT fk_log_pessoa_cpf FOREIGN KEY (pessoa_cpf) REFERENCES public.pessoa(cpf) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

