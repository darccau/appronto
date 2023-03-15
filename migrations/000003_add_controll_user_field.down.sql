ALTER TABLE users 
DROP COLUMN version,
DROP COLUMN activated,
DROP COLUMN created_at;

ALTER TABLE users ALTER COLUMN password TYPE TEXT USING password::TEXT;
