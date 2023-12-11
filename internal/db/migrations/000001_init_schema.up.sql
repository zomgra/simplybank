CREATE TABLE "users" (
  id bigserial PRIMARY KEY,
  "username" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "currency" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transfer" (
  "id" bigserial PRIMARY KEY,
  "to_account_id" bigint,
  "from_account_id" bigint,
  "status" varchar NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "entries" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint,
  "amount" bigint NOT NULL,
  "status" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
;

CREATE INDEX ON "transfer" ("to_account_id");

CREATE INDEX ON "transfer" ("from_account_id");

CREATE INDEX ON "transfer" ("to_account_id", "from_account_id");

CREATE INDEX ON "entries" ("account_id");

ALTER TABLE "transfer" ADD FOREIGN KEY ("to_account_id") REFERENCES "users" ("id");

ALTER TABLE "transfer" ADD FOREIGN KEY ("from_account_id") REFERENCES "users" ("id");

ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "users" ("id");
