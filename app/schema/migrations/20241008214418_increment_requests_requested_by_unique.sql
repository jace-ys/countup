-- +goose Up
-- modify "increment_requests" table
ALTER TABLE "public"."increment_requests" ADD CONSTRAINT "increment_requests_requested_by_key" UNIQUE ("requested_by");

-- +goose Down
-- reverse: modify "increment_requests" table
ALTER TABLE "public"."increment_requests" DROP CONSTRAINT "increment_requests_requested_by_key";
