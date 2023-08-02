CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now() AT TIME ZONE 'Asia/Singapore')
);

CREATE TABLE "stores" (
  "id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now() AT TIME ZONE 'Asia/Singapore')
);

CREATE TABLE "products" (
  "id" bigserial PRIMARY KEY,
  "store_id" bigint NOT NULL,
  "name" varchar NOT NULL,
  "brand" varchar NOT NULL,
  "price" decimal(10, 2) NOT NULL,
  "quantity" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now() AT TIME ZONE 'Asia/Singapore')
);

CREATE INDEX ON "stores" ("name");

CREATE INDEX ON "products" ("name");

ALTER TABLE "stores" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");

ALTER TABLE "products" ADD FOREIGN KEY ("store_id") REFERENCES "stores" ("id");
