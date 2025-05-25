BEGIN;

CREATE TABLE users
(
    id              SERIAL PRIMARY KEY,
    email           VARCHAR(255) UNIQUE,
    hashed_password TEXT,
    salt            TEXT,
    name            VARCHAR(255),
    surname         VARCHAR(255),
    avatar_url      TEXT      DEFAULT NULL,
    created_at      TIMESTAMP DEFAULT NOW(),
    updated_at      TIMESTAMP DEFAULT NOW()
);

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

CREATE TABLE projects
(
    id                SERIAL PRIMARY KEY,
    owner_id          INTEGER      REFERENCES users (id) ON DELETE SET NULL,
    name              VARCHAR(255) NOT NULL,
    description       TEXT,
    color             VARCHAR(255),
    parent_project_id INTEGER      REFERENCES projects (id) ON DELETE SET NULL,
    created_at        TIMESTAMP DEFAULT NOW(),
    updated_at        TIMESTAMP DEFAULT NOW()
);

CREATE TABLE kanban
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(255) NOT NULL,
    project_id INTEGER REFERENCES projects (id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE task
(
    id           SERIAL PRIMARY KEY,
    title        VARCHAR(255) NOT NULL,
    kanban_id    INTEGER REFERENCES kanban (id) ON DELETE CASCADE,
    owner_id    INTEGER REFERENCES users (id) ON DELETE CASCADE,
    description  TEXT,
    is_completed BOOLEAN   DEFAULT FALSE,
    created_at   TIMESTAMP DEFAULT NOW(),
    updated_at   TIMESTAMP DEFAULT NOW()
);

CREATE TABLE comment
(
    id         SERIAL PRIMARY KEY,
    author     INTEGER REFERENCES users (id) ON DELETE CASCADE,
    text       TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE comment_task
(
    task_id    INTEGER REFERENCES task (id) ON DELETE CASCADE,
    comment_id INTEGER REFERENCES comment (id) ON DELETE CASCADE,
    PRIMARY KEY (task_id, comment_id)
);

CREATE TABLE task_users
(
    task_id INTEGER REFERENCES task (id) ON DELETE CASCADE,
    user_id INTEGER REFERENCES users (id) ON DELETE CASCADE,
    PRIMARY KEY (task_id, user_id)
);

CREATE TABLE kanban_tasks
(
    kanban_id       INTEGER REFERENCES kanban (id) ON DELETE CASCADE,
    task_id         INTEGER REFERENCES task (id) ON DELETE CASCADE,
    order_in_kanban INTEGER,
    PRIMARY KEY (kanban_id, task_id)
);

CREATE TABLE project_tasks
(
    project_id INTEGER REFERENCES projects (id) ON DELETE CASCADE,
    task_id    INTEGER REFERENCES task (id) ON DELETE CASCADE,
    PRIMARY KEY (project_id, task_id)
);

CREATE TABLE project_topics
(
    project_id INTEGER REFERENCES projects (id) ON DELETE CASCADE,
    topic_id   INTEGER REFERENCES projects (id) ON DELETE CASCADE,
    PRIMARY KEY (project_id, topic_id)
);

CREATE TABLE project_users
(
    project_id         INTEGER REFERENCES projects (id) ON DELETE CASCADE,
    user_id            INTEGER REFERENCES users (id) ON DELETE CASCADE,
    permission         VARCHAR(50) NOT NULL, -- 'read', 'edit', 'owner'
    invited_by_user_id INTEGER     REFERENCES users (id) ON DELETE SET NULL,
    invited_at         TIMESTAMP DEFAULT NOW(),
    joined_at          TIMESTAMP,
    PRIMARY KEY (project_id, user_id)
);

-- Триггерная функция
CREATE OR REPLACE FUNCTION update_updated_at_column()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Триггеры
CREATE TRIGGER set_updated_at_users
    BEFORE UPDATE
    ON users
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER set_updated_at_projects
    BEFORE UPDATE
    ON projects
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER set_updated_at_tasks
    BEFORE UPDATE
    ON task
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER set_updated_at_kanban
    BEFORE UPDATE
    ON kanban
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

COMMIT;
