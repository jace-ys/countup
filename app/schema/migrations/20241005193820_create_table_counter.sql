-- +goose Up
-- create "counter" table
CREATE TABLE "public"."counter" ("id" serial NOT NULL, "count" integer NOT NULL, "last_increment_by" text NULL, "last_increment_at" timestamptz NULL, "next_finalize_at" timestamptz NULL, PRIMARY KEY ("id"), CONSTRAINT "counter_id_check" CHECK (id = 1));

-- +goose Down
-- reverse: create "counter" table
DROP TABLE "public"."counter";
