-- +goose Up
-- create "increment_requests" table
CREATE TABLE "public"."increment_requests" ("requested_by" text NOT NULL, "requested_at" timestamptz NOT NULL);

-- +goose Down
-- reverse: create "increment_requests" table
DROP TABLE "public"."increment_requests";
