-- Database generated with pgModeler (PostgreSQL Database Modeler).
-- pgModeler version: 1.0.3
-- PostgreSQL version: 15.0
-- Project Site: pgmodeler.io
-- Model Author: ---

-- Database creation must be performed outside a multi lined SQL file. 
-- These commands were put in this file only as a convenience.
-- 
-- object: todo | type: DATABASE --
-- DROP DATABASE IF EXISTS todo;
CREATE DATABASE todo;
-- ddl-end --


-- object: public.task | type: TABLE --
-- DROP TABLE IF EXISTS public.task CASCADE;
CREATE TABLE public.task (
	id uuid NOT NULL,
	description varchar(255) NOT NULL,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,
	id_task_status uuid NOT NULL,
	CONSTRAINT task_pk PRIMARY KEY (uuid)
);
-- ddl-end --
ALTER TABLE public.task OWNER TO postgres;
-- ddl-end --

-- object: public.task_status | type: TABLE --
-- DROP TABLE IF EXISTS public.task_status CASCADE;
CREATE TABLE public.task_status (
	id uuid NOT NULL,
	name varchar(20) NOT NULL,
	CONSTRAINT task_status_pk PRIMARY KEY (id)
);
-- ddl-end --
ALTER TABLE public.task_status OWNER TO postgres;
-- ddl-end --

-- object: task_status_fk | type: CONSTRAINT --
-- ALTER TABLE public.task DROP CONSTRAINT IF EXISTS task_status_fk CASCADE;
ALTER TABLE public.task ADD CONSTRAINT task_status_fk FOREIGN KEY (id_task_status)
REFERENCES public.task_status (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --


-- INIT DATA
insert into task_status (id, name) values ('31fa0b36-80b9-45d9-9c97-75bab30ffcf4', 'todo');
insert into task_status (id, name) values ('0a086e48-dd50-4714-b6a1-606043596ba5', 'doing');
insert into task_status (id, name) values ('9145295e-f8b4-4e06-ab6f-ae741fa2ba5a', 'finish');
