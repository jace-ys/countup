-- +goose Up
-- create "users" table
CREATE TABLE "public"."users" ("id" character varying(24) NOT NULL, "email" text NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "users_email_key" UNIQUE ("email"));

-- +goose Down
-- reverse: create "users" table
DROP TABLE "public"."users";
