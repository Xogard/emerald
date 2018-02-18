--
-- PostgreSQL database dump
--

-- Dumped from database version 10.1
-- Dumped by pg_dump version 10.1

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

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: clientes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE clientes (
    id uuid NOT NULL,
    nome character varying(255),
    dtnascimento date NOT NULL,
    cpf character varying(11) NOT NULL,
    rg character varying(255) NOT NULL,
    orgaoexp character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    fone character varying(15) NOT NULL,
    celular character varying(15) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE clientes OWNER TO postgres;

--
-- Name: contratos; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE contratos (
    id uuid NOT NULL,
    clienteid uuid NOT NULL,
    exercicio smallint NOT NULL,
    dtcontrato date NOT NULL,
    duracao smallint NOT NULL,
    valor numeric NOT NULL,
    cancelado boolean NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    diavencimento smallint DEFAULT '10'::smallint NOT NULL
);


ALTER TABLE contratos OWNER TO postgres;

--
-- Name: enderecos; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE enderecos (
    id uuid NOT NULL,
    enderecoid uuid NOT NULL,
    logradouro text,
    complemento character varying(255) NOT NULL,
    cep character varying(255) NOT NULL,
    bairro character varying(255) NOT NULL,
    cidade character varying(255) NOT NULL,
    estado character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    clienteid uuid NOT NULL
);


ALTER TABLE enderecos OWNER TO postgres;

--
-- Name: escolas; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE escolas (
    id uuid NOT NULL,
    nome character varying(255) NOT NULL,
    endereco character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE escolas OWNER TO postgres;

--
-- Name: faturas; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE faturas (
    id uuid NOT NULL,
    contratoid uuid NOT NULL,
    clienteid uuid NOT NULL,
    pagamentoid uuid,
    valor numeric,
    dtvencimento date,
    dtpagamento date,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE faturas OWNER TO postgres;

--
-- Name: filhos; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE filhos (
    id uuid NOT NULL,
    nome character varying(255) NOT NULL,
    dtnascimento date NOT NULL,
    escolaid uuid NOT NULL,
    clienteid uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE filhos OWNER TO postgres;

--
-- Name: pagamentos; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE pagamentos (
    id uuid NOT NULL,
    contratoid uuid NOT NULL,
    clienteid uuid NOT NULL,
    valor numeric,
    dtpagamento date,
    observacao character varying(255),
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE pagamentos OWNER TO postgres;

--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE schema_migration (
    version character varying(255) NOT NULL
);


ALTER TABLE schema_migration OWNER TO postgres;

--
-- Name: clientes clientes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY clientes
    ADD CONSTRAINT clientes_pkey PRIMARY KEY (id);


--
-- Name: contratos contratos_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY contratos
    ADD CONSTRAINT contratos_pkey PRIMARY KEY (id);


--
-- Name: enderecos enderecos_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY enderecos
    ADD CONSTRAINT enderecos_pkey PRIMARY KEY (id);


--
-- Name: escolas escolas_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY escolas
    ADD CONSTRAINT escolas_pkey PRIMARY KEY (id);


--
-- Name: faturas faturas_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY faturas
    ADD CONSTRAINT faturas_pkey PRIMARY KEY (id);


--
-- Name: filhos filhos_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY filhos
    ADD CONSTRAINT filhos_pkey PRIMARY KEY (id);


--
-- Name: pagamentos pagamentos_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY pagamentos
    ADD CONSTRAINT pagamentos_pkey PRIMARY KEY (id);


--
-- Name: version_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX version_idx ON schema_migration USING btree (version);


--
-- Name: enderecos enderecos_clientes_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY enderecos
    ADD CONSTRAINT enderecos_clientes_id_fk FOREIGN KEY (clienteid) REFERENCES clientes(id) ON DELETE CASCADE;


--
-- Name: faturas faturas_clienteid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY faturas
    ADD CONSTRAINT faturas_clienteid_fkey FOREIGN KEY (clienteid) REFERENCES clientes(id) ON DELETE CASCADE;


--
-- Name: faturas faturas_contratoid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY faturas
    ADD CONSTRAINT faturas_contratoid_fkey FOREIGN KEY (contratoid) REFERENCES contratos(id) ON DELETE CASCADE;


--
-- Name: filhos filhos_clienteid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY filhos
    ADD CONSTRAINT filhos_clienteid_fkey FOREIGN KEY (clienteid) REFERENCES clientes(id) ON DELETE CASCADE;


--
-- Name: filhos filhos_escolaid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY filhos
    ADD CONSTRAINT filhos_escolaid_fkey FOREIGN KEY (escolaid) REFERENCES escolas(id) ON DELETE SET NULL;


--
-- Name: pagamentos pagamentos_clienteid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY pagamentos
    ADD CONSTRAINT pagamentos_clienteid_fkey FOREIGN KEY (clienteid) REFERENCES clientes(id) ON DELETE CASCADE;


--
-- Name: pagamentos pagamentos_contratoid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY pagamentos
    ADD CONSTRAINT pagamentos_contratoid_fkey FOREIGN KEY (contratoid) REFERENCES contratos(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

