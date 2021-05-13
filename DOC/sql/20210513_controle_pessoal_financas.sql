--
-- PostgreSQL database dump
--

-- Dumped from database version 11.9 (Debian 11.9-1.pgdg90+1)
-- Dumped by pg_dump version 12.0

-- Started on 2021-05-13 23:04:22 UTC

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 597 (class 1247 OID 16386)
-- Name: codigo_texto; Type: DOMAIN; Schema: public; Owner: postgres
--

CREATE DOMAIN public.codigo_texto AS character varying(19);


ALTER DOMAIN public.codigo_texto OWNER TO postgres;

--
-- TOC entry 600 (class 1247 OID 16388)
-- Name: cpf; Type: DOMAIN; Schema: public; Owner: postgres
--

CREATE DOMAIN public.cpf AS character varying(11) NOT NULL;


ALTER DOMAIN public.cpf OWNER TO postgres;

--
-- TOC entry 603 (class 1247 OID 16390)
-- Name: data_completa; Type: DOMAIN; Schema: public; Owner: postgres
--

CREATE DOMAIN public.data_completa AS timestamp with time zone;


ALTER DOMAIN public.data_completa OWNER TO postgres;

--
-- TOC entry 606 (class 1247 OID 16392)
-- Name: descricao_curta; Type: DOMAIN; Schema: public; Owner: postgres
--

CREATE DOMAIN public.descricao_curta AS character varying(20) NOT NULL;


ALTER DOMAIN public.descricao_curta OWNER TO postgres;

--
-- TOC entry 609 (class 1247 OID 16394)
-- Name: dinheiro; Type: DOMAIN; Schema: public; Owner: postgres
--

CREATE DOMAIN public.dinheiro AS numeric(19,3);


ALTER DOMAIN public.dinheiro OWNER TO postgres;

--
-- TOC entry 612 (class 1247 OID 16396)
-- Name: nome_conta; Type: DOMAIN; Schema: public; Owner: postgres
--

CREATE DOMAIN public.nome_conta AS character varying(50);


ALTER DOMAIN public.nome_conta OWNER TO postgres;

SET default_tablespace = '';

--
-- TOC entry 196 (class 1259 OID 16397)
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
-- TOC entry 198 (class 1259 OID 16407)
-- Name: detalhe_lancamento; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.detalhe_lancamento (
    id_lancamento integer NOT NULL,
    nome_conta public.nome_conta NOT NULL,
    debito public.dinheiro,
    credito public.dinheiro
);


ALTER TABLE public.detalhe_lancamento OWNER TO postgres;

--
-- TOC entry 197 (class 1259 OID 16405)
-- Name: detalhe_lancamento_id_lancamento_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.detalhe_lancamento_id_lancamento_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.detalhe_lancamento_id_lancamento_seq OWNER TO postgres;

--
-- TOC entry 2961 (class 0 OID 0)
-- Dependencies: 197
-- Name: detalhe_lancamento_id_lancamento_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.detalhe_lancamento_id_lancamento_seq OWNED BY public.detalhe_lancamento.id_lancamento;


--
-- TOC entry 200 (class 1259 OID 16416)
-- Name: lancamento; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.lancamento (
    id integer NOT NULL,
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
-- TOC entry 199 (class 1259 OID 16414)
-- Name: lancamento_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.lancamento_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.lancamento_id_seq OWNER TO postgres;

--
-- TOC entry 2962 (class 0 OID 0)
-- Dependencies: 199
-- Name: lancamento_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.lancamento_id_seq OWNED BY public.lancamento.id;


--
-- TOC entry 202 (class 1259 OID 16428)
-- Name: log; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.log (
    id integer NOT NULL,
    cpf_pessoa public.cpf NOT NULL,
    data public.data_completa DEFAULT now() NOT NULL,
    sql text NOT NULL
);


ALTER TABLE public.log OWNER TO postgres;

--
-- TOC entry 201 (class 1259 OID 16426)
-- Name: log_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.log_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.log_id_seq OWNER TO postgres;

--
-- TOC entry 2963 (class 0 OID 0)
-- Dependencies: 201
-- Name: log_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.log_id_seq OWNED BY public.log.id;


--
-- TOC entry 203 (class 1259 OID 16435)
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
    estado boolean DEFAULT true NOT NULL,
    administrador boolean DEFAULT false NOT NULL
);


ALTER TABLE public.pessoa OWNER TO postgres;

--
-- TOC entry 204 (class 1259 OID 16443)
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
-- TOC entry 2790 (class 2604 OID 16490)
-- Name: detalhe_lancamento id_lancamento; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.detalhe_lancamento ALTER COLUMN id_lancamento SET DEFAULT nextval('public.detalhe_lancamento_id_lancamento_seq'::regclass);


--
-- TOC entry 2794 (class 2604 OID 16491)
-- Name: lancamento id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.lancamento ALTER COLUMN id SET DEFAULT nextval('public.lancamento_id_seq'::regclass);


--
-- TOC entry 2795 (class 2604 OID 16492)
-- Name: log id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.log ALTER COLUMN id SET DEFAULT nextval('public.log_id_seq'::regclass);


--
-- TOC entry 2947 (class 0 OID 16397)
-- Dependencies: 196
-- Data for Name: conta; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.conta (nome, nome_tipo_conta, codigo, conta_pai, comentario, data_criacao, data_modificacao, estado) FROM stdin;
ativos	ativo	1	\N	\N	2020-06-17 18:00:15.127816+00	2020-06-17 18:00:15.127816+00	t
despesas	despesa	2	\N	\N	2020-06-17 18:00:15.127816+00	2020-06-17 18:00:15.127816+00	t
líquidos	líquido	3	\N	\N	2020-06-17 18:00:15.127816+00	2020-06-17 18:00:15.127816+00	t
passivos	passivo	4	\N	\N	2020-06-17 18:00:15.127816+00	2020-06-17 18:00:15.127816+00	t
receitas	receita	5	\N	\N	2020-06-17 18:00:15.127816+00	2020-06-17 18:00:15.127816+00	t
conta corrente	banco	6	ativos	\N	2020-06-17 18:00:15.127816+00	2020-06-17 18:00:15.127816+00	t
conta poupança	banco	7	ativos	\N	2020-06-17 18:00:15.127816+00	2020-06-17 18:00:15.127816+00	t
dinheiro em carteira	carteira	8	ativos	\N	2020-06-17 18:00:15.127816+00	2020-06-17 18:00:15.127816+00	t
nu conta	banco	9	conta corrente	\N	2020-06-17 18:00:15.127816+00	2020-06-17 18:00:15.127816+00	t
bradesco	banco	10	conta poupança	\N	2020-06-17 18:00:15.127816+00	2020-06-17 18:00:15.127816+00	t
caixa econômica	banco	11	conta poupança	\N	2020-06-17 18:00:15.127816+00	2020-06-17 18:00:15.127816+00	t
internet	despesa	12	despesas	\N	2020-06-17 18:00:15.127816+00	2020-06-17 18:00:15.127816+00	t
telefone	despesa	13	despesas	\N	2020-06-17 18:00:15.127816+00	2020-06-17 18:00:15.127816+00	t
serviços	despesa	14	despesas	\N	2020-06-17 18:00:15.127816+00	2020-06-17 18:00:15.127816+00	t
eletricidade	despesa	15	serviços	\N	2020-06-17 18:00:15.127816+00	2020-06-17 18:00:15.127816+00	t
refeições fora	despesa	16	despesas	\N	2020-06-17 18:00:15.127816+00	2020-06-17 18:00:15.127816+00	t
computador	despesa	17	despesas	\N	2020-06-17 18:00:15.127816+00	2020-06-17 18:00:15.127816+00	t
\.


--
-- TOC entry 2949 (class 0 OID 16407)
-- Dependencies: 198
-- Data for Name: detalhe_lancamento; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.detalhe_lancamento (id_lancamento, nome_conta, debito, credito) FROM stdin;
1039	nu conta	\N	142.000
1039	receitas	142.000	\N
1040	conta corrente	122.000	\N
1040	dinheiro em carteira	\N	122.000
1041	bradesco	133.000	\N
1041	internet	\N	133.000
1042	bradesco	\N	144.000
1042	telefone	144.000	\N
1043	conta corrente	\N	512.000
1043	serviços	512.000	\N
\.


--
-- TOC entry 2951 (class 0 OID 16416)
-- Dependencies: 200
-- Data for Name: lancamento; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.lancamento (id, cpf_pessoa, data, numero, descricao, data_criacao, data_modificacao, estado) FROM stdin;
1039	00000000000	0001-01-01 00:00:00+00	010101	Teste 010101	2020-12-28 20:46:29.725947+00	2020-12-28 20:46:29.725947+00	t
1040	00000000000	0001-01-01 00:00:00+00	02020202	Teste 02020202	2020-12-28 20:51:43.224183+00	2020-12-28 20:51:43.224184+00	t
1041	00000000000	0001-01-01 00:00:00+00	0303	teste 03	2020-12-28 20:56:12.057275+00	2020-12-28 20:56:12.057276+00	t
1042	00000000000	2020-12-04 00:00:00+00	0404	teste 04	2020-12-28 21:04:02.591694+00	2020-12-28 21:04:02.591694+00	t
1043	00000000000	2020-12-05 00:00:00+00	0512	teste 0512	2020-12-28 21:13:21.481368+00	2020-12-28 21:13:21.481368+00	t
\.


--
-- TOC entry 2953 (class 0 OID 16428)
-- Dependencies: 202
-- Data for Name: log; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.log (id, cpf_pessoa, data, sql) FROM stdin;
\.


--
-- TOC entry 2954 (class 0 OID 16435)
-- Dependencies: 203
-- Data for Name: pessoa; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.pessoa (cpf, nome_completo, usuario, senha, email, data_criacao, data_modificacao, estado, administrador) FROM stdin;
12345678910	Paulo C Silva Jr	paulo	8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92	pauluscave@gmail.com	2020-01-03 00:05:28.253164+00	2020-01-03 00:05:28.253164+00	t	f
00000000000	Administrador	admin	8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918	meuemail@email.com	2020-06-15 20:35:32.092031+00	2020-06-15 21:09:35.447619+00	t	t
11111111111	teste 01	teste01	8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92	teste01@gmail.com	2020-01-03 00:05:28.248303+00	2021-05-05 00:20:04.222207+00	t	t
33333333333	João Alcântara	joao02	8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92	joaoa@gmail.com	2020-01-03 00:05:28.254942+00	2021-05-05 00:20:04.22598+00	f	f
00000001000	Teste 10 novo	teste10	72c6f86adbda726ae212d6d521d449436dcdcefd0584fb5cc20e93c34cb6205d	teste10@gmail.com	2021-01-28 23:51:30.940113+00	2021-01-28 23:51:35.764849+00	t	f
00000000019	Teste de usuário número 19	teste19	49dc52e6bf2abe5ef6e2bb5b0f1ee2d765b922ae6cc8b95d39dc06c21c848f8c	teste19@email.com	2021-03-02 12:41:12.800977+00	2021-03-03 12:47:13.550573+00	t	t
\.


--
-- TOC entry 2955 (class 0 OID 16443)
-- Dependencies: 204
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
-- TOC entry 2964 (class 0 OID 0)
-- Dependencies: 197
-- Name: detalhe_lancamento_id_lancamento_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.detalhe_lancamento_id_lancamento_seq', 1, false);


--
-- TOC entry 2965 (class 0 OID 0)
-- Dependencies: 199
-- Name: lancamento_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.lancamento_id_seq', 1276, true);


--
-- TOC entry 2966 (class 0 OID 0)
-- Dependencies: 201
-- Name: log_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.log_id_seq', 1, false);


--
-- TOC entry 2803 (class 2606 OID 16452)
-- Name: conta codigo_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.conta
    ADD CONSTRAINT codigo_uq UNIQUE (codigo);


--
-- TOC entry 2805 (class 2606 OID 16454)
-- Name: conta conta_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.conta
    ADD CONSTRAINT conta_pk PRIMARY KEY (nome);


--
-- TOC entry 2807 (class 2606 OID 16456)
-- Name: detalhe_lancamento detalhe_lancamento_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.detalhe_lancamento
    ADD CONSTRAINT detalhe_lancamento_pk PRIMARY KEY (id_lancamento, nome_conta);


--
-- TOC entry 2813 (class 2606 OID 16503)
-- Name: pessoa email_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pessoa
    ADD CONSTRAINT email_uq UNIQUE (email);


--
-- TOC entry 2809 (class 2606 OID 16458)
-- Name: lancamento lancamento_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.lancamento
    ADD CONSTRAINT lancamento_pk PRIMARY KEY (id);


--
-- TOC entry 2811 (class 2606 OID 16460)
-- Name: log log_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.log
    ADD CONSTRAINT log_pk PRIMARY KEY (id);


--
-- TOC entry 2815 (class 2606 OID 16462)
-- Name: pessoa pessoa_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pessoa
    ADD CONSTRAINT pessoa_pk PRIMARY KEY (cpf);


--
-- TOC entry 2819 (class 2606 OID 16464)
-- Name: tipo_conta tipo_conta_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tipo_conta
    ADD CONSTRAINT tipo_conta_pk PRIMARY KEY (nome);


--
-- TOC entry 2817 (class 2606 OID 16501)
-- Name: pessoa usuario_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pessoa
    ADD CONSTRAINT usuario_uq UNIQUE (usuario);


--
-- TOC entry 2822 (class 2606 OID 16465)
-- Name: detalhe_lancamento conta_detalhe_lancamento_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.detalhe_lancamento
    ADD CONSTRAINT conta_detalhe_lancamento_fk FOREIGN KEY (nome_conta) REFERENCES public.conta(nome) ON UPDATE CASCADE ON DELETE RESTRICT;


--
-- TOC entry 2821 (class 2606 OID 16493)
-- Name: conta conta_pai_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.conta
    ADD CONSTRAINT conta_pai_fk FOREIGN KEY (conta_pai) REFERENCES public.conta(nome) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- TOC entry 2823 (class 2606 OID 16470)
-- Name: detalhe_lancamento lancamento_detalhe_lancamento_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.detalhe_lancamento
    ADD CONSTRAINT lancamento_detalhe_lancamento_fk FOREIGN KEY (id_lancamento) REFERENCES public.lancamento(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- TOC entry 2824 (class 2606 OID 16475)
-- Name: lancamento pessoa_lancamento_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.lancamento
    ADD CONSTRAINT pessoa_lancamento_fk FOREIGN KEY (cpf_pessoa) REFERENCES public.pessoa(cpf) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- TOC entry 2825 (class 2606 OID 16480)
-- Name: log pessoa_log_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.log
    ADD CONSTRAINT pessoa_log_fk FOREIGN KEY (cpf_pessoa) REFERENCES public.pessoa(cpf) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- TOC entry 2820 (class 2606 OID 16485)
-- Name: conta tipo_conta_conta_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.conta
    ADD CONSTRAINT tipo_conta_conta_fk FOREIGN KEY (nome_tipo_conta) REFERENCES public.tipo_conta(nome) ON UPDATE CASCADE ON DELETE RESTRICT;


-- Completed on 2021-05-13 23:04:22 UTC

--
-- PostgreSQL database dump complete
--

