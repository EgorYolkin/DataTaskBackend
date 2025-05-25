CREATE TABLE notification
(
    id                 SERIAL PRIMARY KEY,
    owner_id           INTEGER      REFERENCES users (id) ON DELETE SET NULL,

    title VARCHAR(255) NOT NULL,
    description TEXT,
    
    is_read BOOLEAN DEFAULT FALSE,

    created_at         TIMESTAMP DEFAULT NOW(),
    updated_at         TIMESTAMP DEFAULT NOW()
);