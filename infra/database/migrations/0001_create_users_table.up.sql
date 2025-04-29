BEGIN;

CREATE TABLE users
(
    id              SERIAL PRIMARY KEY,
    email           varchar(255) UNIQUE,

    hashed_password   TEXT,
    salt              TEXT,

    name            varchar(255),
    surname         varchar(255),
    avatar_url      TEXT      DEFAULT NULL,

    created_at      TIMESTAMP DEFAULT NOW(),
    updated_at      TIMESTAMP DEFAULT NOW()
);

CREATE OR REPLACE FUNCTION update_updated_at_column()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER set_updated_at
    BEFORE UPDATE
    ON users
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

COMMIT;