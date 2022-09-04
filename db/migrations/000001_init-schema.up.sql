CREATE TYPE "Currency" AS ENUM (
  'USD',
  'Tshs',
  'EUR'
);

CREATE TABLE "account" (
  "id" BIGSERIAL PRIMARY KEY,
  "created" timestamptz DEFAULT (now()),
  "code" int,
  "owner" varchar,
  "balance" bigint,
  "currency" varchar NOT NULL
);

CREATE TABLE "entry" (
  "id" BIGSERIAL PRIMARY KEY,
  "amount" bigint NOT NULL,
  "account" bigint,
  "created" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transfer" (
  "id" bigint PRIMARY KEY,
  "sender" bigint,
  "receiver" bigint,
  "created" timestamptz DEFAULT (now())
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
