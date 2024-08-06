CREATE TABLE counter (
  id SERIAL PRIMARY KEY CHECK (id = 1),
  count INTEGER NOT NULL,
  last_increment_by TEXT,
  last_increment_at TIMESTAMPTZ,
  next_finalize_at TIMESTAMPTZ
);

CREATE UNLOGGED TABLE increment_requests (
  requested_by TEXT NOT NULL UNIQUE,
  requested_at TIMESTAMPTZ NOT NULL
);