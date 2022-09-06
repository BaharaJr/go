-- FUNCTION: public.uuid_generate_v4()

-- DROP FUNCTION IF EXISTS public.uuid_generate_v4();

CREATE OR REPLACE FUNCTION public.uuid_generate_v4(
	)
    RETURNS uuid
    LANGUAGE 'c'
    COST 1
    VOLATILE STRICT PARALLEL SAFE 
AS '$libdir/uuid-ossp', 'uuid_generate_v4'
;

ALTER FUNCTION public.uuid_generate_v4()
    OWNER TO postgres;

CREATE TYPE "Currency" AS ENUM (
  'USD',
  'Tshs',
  'EUR'
);

CREATE TABLE "account" (
  "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  "created" timestamptz NOT NULL DEFAULT (now()),
  "code" varchar,
  "owner" varchar NOT NULL,
  "balance" BIGSERIAL NOT NULL,
  "currency" varchar NOT NULL
);

CREATE TABLE "entry" (
  "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  "amount" BIGSERIAL NOT NULL,
  "account" uuid,
  "created" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transfer" (
  "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  "sender" uuid,
  "receiver" uuid,
  "created" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "account" ("owner");

CREATE INDEX ON "entry" ("amount");

CREATE INDEX ON "entry" ("account");

CREATE INDEX ON "transfer" ("sender");

CREATE INDEX ON "transfer" ("receiver");

CREATE INDEX ON "transfer" ("sender", "receiver");

ALTER TABLE "entry" ADD FOREIGN KEY ("account") REFERENCES "account" ("id");

ALTER TABLE "transfer" ADD FOREIGN KEY ("sender") REFERENCES "account" ("id");

ALTER TABLE "transfer" ADD FOREIGN KEY ("receiver") REFERENCES "account" ("id");
