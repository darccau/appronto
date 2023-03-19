 ALTER TABLE users 
    ADD COLUMN version integer NOT NULL DEFAULT 1,
    ADD COLUMN activated boolean NOT NULL,
    ADD COLUMN created_at timestamp with time zone NOT NULL DEFAULT NOW();

    ALTER TABLE users ALTER COLUMN password TYPE bytea USING password::bytea;
