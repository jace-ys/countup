-- +goose Up
-- modify "counter" table
ALTER TABLE "public"."counter" ADD CONSTRAINT "counter_last_increment_by_fkey" FOREIGN KEY ("last_increment_by") REFERENCES "public"."users" ("email") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- modify "increment_requests" table
ALTER TABLE "public"."increment_requests" ADD CONSTRAINT "increment_requests_requested_by_fkey" FOREIGN KEY ("requested_by") REFERENCES "public"."users" ("email") ON UPDATE NO ACTION ON DELETE NO ACTION;

-- +goose Down
-- reverse: modify "increment_requests" table
ALTER TABLE "public"."increment_requests" DROP CONSTRAINT "increment_requests_requested_by_fkey";
-- reverse: modify "counter" table
ALTER TABLE "public"."counter" DROP CONSTRAINT "counter_last_increment_by_fkey";
