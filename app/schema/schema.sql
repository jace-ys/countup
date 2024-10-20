CREATE TABLE users (
  id CHAR(27) PRIMARY KEY,
  email TEXT NOT NULL UNIQUE
);

CREATE TABLE counter (
  id SERIAL PRIMARY KEY CHECK (id = 1),
  count INTEGER NOT NULL,
  last_increment_by TEXT REFERENCES users(email),
  last_increment_at TIMESTAMPTZ,
  next_finalize_at TIMESTAMPTZ
);

CREATE UNLOGGED TABLE increment_requests (
  requested_by TEXT UNIQUE NOT NULL REFERENCES users(email),
  requested_at TIMESTAMPTZ NOT NULL
);