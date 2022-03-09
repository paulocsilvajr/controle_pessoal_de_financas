--
-- PostgreSQL database dump
--

-- Dumped from database version 11.2 (Debian 11.2-1.pgdg90+1)
-- Dumped by pg_dump version 11.2

-- Started on 2019-03-25 13:46:23 UTC

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
-- TOC entry 637 (class 1247 OID 16709)
-- Name: codigo_numero; Type: DOMAIN; Schema: public; Owner: postgres
--

--
-- TOC entry 596 (class 1247 OID 16711)
-- Name: codigo_texto; Type: DOMAIN; Schema: public; Owner: postgres
--

CREATE DOMAIN public.codigo_texto AS character varying(19);


ALTER DOMAIN public.codigo_texto OWNER TO postgres;

--
-- TOC entry 599 (class 1247 OID 16713)
-- Name: cpf; Type: DOMAIN; Schema: public; Owner: postgres
--

CREATE DOMAIN public.cpf AS character varying(11) NOT NULL;


ALTER DOMAIN public.cpf OWNER TO postgres;

--
-- TOC entry 602 (class 1247 OID 16715)
-- Name: data_completa; Type: DOMAIN; Schema: public; Owner: postgres
--

CREATE DOMAIN public.data_completa AS timestamp with time zone;


ALTER DOMAIN public.data_completa OWNER TO postgres;

--
-- TOC entry 605 (class 1247 OID 16717)
-- Name: descricao_curta; Type: DOMAIN; Schema: public; Owner: postgres
--

CREATE DOMAIN public.descricao_curta AS character varying(20) NOT NULL;


ALTER DOMAIN public.descricao_curta OWNER TO postgres;

--
-- TOC entry 608 (class 1247 OID 16719)
-- Name: dinheiro; Type: DOMAIN; Schema: public; Owner: postgres
--

CREATE DOMAIN public.dinheiro AS numeric(19,3);


ALTER DOMAIN public.dinheiro OWNER TO postgres;

--
-- TOC entry 611 (class 1247 OID 16721)
-- Name: nome_conta; Type: DOMAIN; Schema: public; Owner: postgres
--

CREATE DOMAIN public.nome_conta AS character varying(50);


ALTER DOMAIN public.nome_conta OWNER TO postgres;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- TOC entry 197 (class 1259 OID 16732)
-- Name: conta; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.conta (
    nome public.nome_conta NOT NULL,
    nome_tipo_conta public.nome_conta NOT NULL,
    codigo public.codigo_texto,
    conta_pai public.nome_conta,
    comentario character varying(150),
    data_criacao public.data_completa NOT NULL,
    data_modificacao public.data_completa DEFAULT now() NOT NULL,
    estado boolean DEFAULT true NOT NULL
);


ALTER TABLE public.conta OWNER TO postgres;

--
-- TOC entry 201 (class 1259 OID 16771)
-- Name: detalhe_lancamento; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.detalhe_lancamento (
    id_lancamento serial NOT NULL,
    nome_conta public.nome_conta NOT NULL,
    debito public.dinheiro,
    credito public.dinheiro
);


ALTER TABLE public.detalhe_lancamento OWNER TO postgres;

--
-- TOC entry 200 (class 1259 OID 16760)
-- Name: lancamento; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.lancamento (
    id serial NOT NULL,
    cpf_pessoa public.cpf NOT NULL,
    data public.data_completa DEFAULT now() NOT NULL,
    numero public.codigo_texto,
    descricao character varying(100) NOT NULL,
    data_criacao public.data_completa NOT NULL,
    data_modificacao public.data_completa DEFAULT now() NOT NULL,
    estado boolean DEFAULT true NOT NULL
);


ALTER TABLE public.lancamento OWNER TO postgres;

--
-- TOC entry 199 (class 1259 OID 16752)
-- Name: log; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.log (
    id serial NOT NULL,
    cpf_pessoa public.cpf NOT NULL,
    data public.data_completa NOT NULL,
    sql text NOT NULL
);


ALTER TABLE public.log OWNER TO postgres;

--
-- TOC entry 198 (class 1259 OID 16742)
-- Name: pessoa; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.pessoa (
    cpf public.cpf NOT NULL,
    nome_completo character varying(100),
    usuario character varying(20) NOT NULL,
    senha character varying(64) NOT NULL,
    email character varying(45) NOT NULL,
    data_criacao public.data_completa NOT NULL,
    data_modificacao public.data_completa DEFAULT now() NOT NULL,
    estado boolean DEFAULT true NOT NULL
);


ALTER TABLE public.pessoa OWNER TO postgres;

--
-- TOC entry 196 (class 1259 OID 16722)
-- Name: tipo_conta; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tipo_conta (
    nome public.nome_conta NOT NULL,
    descricao_debito public.descricao_curta NOT NULL,
    descricao_credito public.descricao_curta NOT NULL,
    data_criacao public.data_completa NOT NULL,
    data_modificacao public.data_completa DEFAULT now() NOT NULL,
    estado boolean DEFAULT true NOT NULL
);


ALTER TABLE public.tipo_conta OWNER TO postgres;

--
-- TOC entry 2935 (class 0 OID 16732)
-- Dependencies: 197
-- Data for Name: conta; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.conta (nome, nome_tipo_conta, codigo, conta_pai, comentario, data_criacao, data_modificacao, estado) FROM stdin;
ativos	ativo	1	\N	\N	2019-03-25 12:51:04.064096+00	2019-03-25 12:51:04.064096+00	t
despesas	despesa	2	\N	\N	2019-03-25 12:51:04.064096+00	2019-03-25 12:51:04.064096+00	t
líquidos	líquido	3	\N	\N	2019-03-25 12:51:04.064096+00	2019-03-25 12:51:04.064096+00	t
passivos	passivo	4	\N	\N	2019-03-25 12:51:04.064096+00	2019-03-25 12:51:04.064096+00	t
receitas	receita	5	\N	\N	2019-03-25 12:51:04.064096+00	2019-03-25 12:51:04.064096+00	t
conta corrente	banco	6	ativos	\N	2019-03-25 12:51:04.064096+00	2019-03-25 12:51:04.064096+00	t
conta poupança	banco	7	ativos	\N	2019-03-25 12:51:04.064096+00	2019-03-25 12:51:04.064096+00	t
dinheiro em carteira	carteira	8	ativos	\N	2019-03-25 12:51:04.064096+00	2019-03-25 12:51:04.064096+00	t
nu conta	banco	9	conta corrente	\N	2019-03-25 12:51:04.064096+00	2019-03-25 12:51:04.064096+00	t
bradesco	banco	10	conta poupança	\N	2019-03-25 12:51:04.064096+00	2019-03-25 12:51:04.064096+00	t
caixa econômica	banco	11	conta poupança	\N	2019-03-25 12:51:04.064096+00	2019-03-25 12:51:04.064096+00	t
internet	despesa	12	despesas	\N	2019-03-25 12:51:04.064096+00	2019-03-25 12:51:04.064096+00	t
telefone	despesa	13	despesas	\N	2019-03-25 12:51:04.064096+00	2019-03-25 12:51:04.064096+00	t
serviços	despesa	14	despesas	\N	2019-03-25 12:51:04.064096+00	2019-03-25 12:51:04.064096+00	t
eletricidade	despesa	15	serviços	\N	2019-03-25 12:51:04.064096+00	2019-03-25 12:51:04.064096+00	t
refeições fora	despesa	16	despesas	\N	2019-03-25 12:51:04.064096+00	2019-03-25 12:51:04.064096+00	t
computador	despesa	17	despesas	\N	2019-03-25 12:51:04.064096+00	2019-03-25 12:51:04.064096+00	t
\.


--
-- TOC entry 2939 (class 0 OID 16771)
-- Dependencies: 201
-- Data for Name: detalhe_lancamento; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.detalhe_lancamento (id_lancamento, nome_conta, debito, credito) FROM stdin;
\.


--
-- TOC entry 2938 (class 0 OID 16760)
-- Dependencies: 200
-- Data for Name: lancamento; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.lancamento (id, cpf_pessoa, data, numero, descricao, data_criacao, data_modificacao, estado) FROM stdin;
\.


--
-- TOC entry 2937 (class 0 OID 16752)
-- Dependencies: 199
-- Data for Name: log; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.log (id, cpf_pessoa, data, sql) FROM stdin;
\.


--
-- TOC entry 2936 (class 0 OID 16742)
-- Dependencies: 198
-- Data for Name: pessoa; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.pessoa (cpf, nome_completo, usuario, senha, email, data_criacao, data_modificacao, estado) FROM stdin;
12345678910	Paulo C Silva Jr	paulo	123456	pauluscave@gmail.com	2019-03-25 12:17:11.18173+00	2019-03-25 12:17:11.18173+00	t
11111111111	teste 01	teste01	123456	teste01@gmail.com	2019-03-25 12:21:35.918739+00	2019-03-25 12:21:35.918739+00	t
22222222222	Teste 02	teste02	654321	teste02@gmail.com	2019-03-25 12:22:28.045346+00	2019-03-25 12:22:28.045346+00	t
33333333333	João Alcântara	joao02	121212	joaoa@gmail.com	2019-03-25 12:24:10.260758+00	2019-03-25 12:24:10.260758+00	t
\.


--
-- TOC entry 2934 (class 0 OID 16722)
-- Dependencies: 196
-- Data for Name: tipo_conta; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.tipo_conta (nome, descricao_debito, descricao_credito, data_criacao, data_modificacao, estado) FROM stdin;
banco	saque	depósito	2019-03-25 12:40:58.187736+00	2019-03-25 12:40:58.187736+00	t
carteira	gastar	receber	2019-03-25 12:40:58.187736+00	2019-03-25 12:40:58.187736+00	t
despesa	desconto	despesa	2019-03-25 12:40:58.187736+00	2019-03-25 12:40:58.187736+00	t
cartão de crédito	cobrar	pagamento	2019-03-25 12:40:58.187736+00	2019-03-25 12:40:58.187736+00	t
receita	receita	cobrar	2019-03-25 12:40:58.187736+00	2019-03-25 12:40:58.187736+00	t
ativo	descrescer	aumentar	2019-03-25 12:40:58.187736+00	2019-03-25 12:40:58.187736+00	t
passivo	aumentar	descrescer	2019-03-25 12:40:58.187736+00	2019-03-25 12:40:58.187736+00	t
líquido	aumentar	descrescer	2019-03-25 12:40:58.187736+00	2019-03-25 12:40:58.187736+00	t
\.


--
-- TOC entry 2797 (class 2606 OID 16812)
-- Name: conta codigo_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.conta
    ADD CONSTRAINT codigo_uq UNIQUE (codigo);


--
-- TOC entry 2799 (class 2606 OID 16741)
-- Name: conta conta_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.conta
    ADD CONSTRAINT conta_pk PRIMARY KEY (nome);


--
-- TOC entry 2807 (class 2606 OID 16778)
-- Name: detalhe_lancamento detalhe_lancamento_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.detalhe_lancamento
    ADD CONSTRAINT detalhe_lancamento_pk PRIMARY KEY (id_lancamento, nome_conta);


--
-- TOC entry 2805 (class 2606 OID 16770)
-- Name: lancamento lancamento_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.lancamento
    ADD CONSTRAINT lancamento_pk PRIMARY KEY (id);


--
-- TOC entry 2803 (class 2606 OID 16759)
-- Name: log log_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.log
    ADD CONSTRAINT log_pk PRIMARY KEY (id);


--
-- TOC entry 2801 (class 2606 OID 16751)
-- Name: pessoa pessoa_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pessoa
    ADD CONSTRAINT pessoa_pk PRIMARY KEY (cpf);


--
-- TOC entry 2795 (class 2606 OID 16731)
-- Name: tipo_conta tipo_conta_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tipo_conta
    ADD CONSTRAINT tipo_conta_pk PRIMARY KEY (nome);


--
-- TOC entry 2811 (class 2606 OID 16784)
-- Name: detalhe_lancamento conta_detalhe_lancamento_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.detalhe_lancamento
    ADD CONSTRAINT conta_detalhe_lancamento_fk FOREIGN KEY (nome_conta) REFERENCES public.conta(nome) ON UPDATE CASCADE ON DELETE RESTRICT;


--
-- TOC entry 2812 (class 2606 OID 16799)
-- Name: detalhe_lancamento lancamento_detalhe_lancamento_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.detalhe_lancamento
    ADD CONSTRAINT lancamento_detalhe_lancamento_fk FOREIGN KEY (id_lancamento) REFERENCES public.lancamento(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- TOC entry 2810 (class 2606 OID 16789)
-- Name: lancamento pessoa_lancamento_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.lancamento
    ADD CONSTRAINT pessoa_lancamento_fk FOREIGN KEY (cpf_pessoa) REFERENCES public.pessoa(cpf) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- TOC entry 2809 (class 2606 OID 16794)
-- Name: log pessoa_log_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.log
    ADD CONSTRAINT pessoa_log_fk FOREIGN KEY (cpf_pessoa) REFERENCES public.pessoa(cpf) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- TOC entry 2808 (class 2606 OID 16779)
-- Name: conta tipo_conta_conta_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.conta
    ADD CONSTRAINT tipo_conta_conta_fk FOREIGN KEY (nome_tipo_conta) REFERENCES public.tipo_conta(nome) ON UPDATE CASCADE ON DELETE RESTRICT;


-- Completed on 2019-03-25 13:46:23 UTC

--
-- PostgreSQL database dump complete
--

