-- +goose Up
-- modify "increment_requests" table
ALTER TABLE "public"."increment_requests" SET UNLOGGED;

-- +goose Down
-- reverse: modify "increment_requests" table
ALTER TABLE "public"."increment_requests" SET LOGGED;
