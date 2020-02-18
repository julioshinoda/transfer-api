
CREATE SEQUENCE IF NOT EXISTS account_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	CACHE 1
	NO CYCLE;

CREATE TABLE IF NOT EXISTS accounts (
	id bigint NOT NULL DEFAULT nextval('account_id_seq'::regclass),
	"name" varchar NULL,
	cpf varchar NULL,
	ballance int4 NULL,
	created_at date NULL,
	CONSTRAINT accounts_pk PRIMARY KEY (id),
	CONSTRAINT accounts_un UNIQUE (cpf)
);

CREATE SEQUENCE IF NOT EXISTS transfers_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	CACHE 1
	NO CYCLE;

CREATE TABLE IF NOT EXISTS transfers (
	id bigint NOT NULL DEFAULT nextval('transfers_id_seq'::regclass),
	account_origin_id bigint NULL,
	account_destination_id bigint NULL,
	amount int4 NULL,
	created_at date NULL,
    CONSTRAINT transfers_pk PRIMARY KEY (id)
);
