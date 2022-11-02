CREATE TABLE IF NOT EXISTS users (
  id bigserial PRIMARY KEY,
  first_name text NOT NULL,
  last_name text NOT NULL,
  email text NOT NULL,
  password text NOT NULL
);
