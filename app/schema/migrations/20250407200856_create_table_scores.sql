-- +goose Up
-- create "scores" table
CREATE TABLE "public"."scores" ("user_email" text NOT NULL, "score" integer NOT NULL, CONSTRAINT "scores_user_email_fkey" FOREIGN KEY ("user_email") REFERENCES "public"."users" ("email") ON UPDATE NO ACTION ON DELETE NO ACTION);

-- +goose Down
-- reverse: create "scores" table
DROP TABLE "public"."scores";
