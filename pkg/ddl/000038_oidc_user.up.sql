BEGIN;

ALTER TABLE auth_users
    ADD COLUMN IF NOT EXISTS external_id TEXT UNIQUE;
END;
