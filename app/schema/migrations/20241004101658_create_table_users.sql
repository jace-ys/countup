-- +goose Up
-- create "users" table
CREATE TABLE "public"."users" ("id" character (27) NOT NULL, "email" text NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "users_email_key" UNIQUE ("email"));

-- +goose Down
-- reverse: create "users" table
DROP TABLE "public"."users";
