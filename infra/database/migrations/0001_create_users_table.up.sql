BEGIN;

CREATE TABLE users
(
    id              SERIAL PRIMARY KEY,
    email           varchar(255) UNIQUE,
    hashed_password text,

    name            varchar(255),
    surname         varchar(255),
    avatar_url      TEXT      DEFAULT NULL,

    created_at      TIMESTAMP DEFAULT NOW(),
    updated_at      TIMESTAMP DEFAULT NOW()
);

COMMIT;