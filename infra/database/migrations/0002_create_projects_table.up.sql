BEGIN;

CREATE TABLE users
(
    id         SERIAL PRIMARY KEY,
    email      VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now()
);

CREATE TABLE task
(
    id           SERIAL PRIMARY KEY,
    title        VARCHAR(255) NOT NULL,
    description  TEXT,
    is_completed BOOLEAN                     DEFAULT FALSE,
    created_at   TIMESTAMP WITHOUT TIME ZONE DEFAULT now(),
    updated_at   TIMESTAMP WITHOUT TIME ZONE DEFAULT now()
);

CREATE TABLE task_users
(
    task_id INTEGER REFERENCES task (id) ON DELETE CASCADE,
    user_id INTEGER REFERENCES users (id) ON DELETE CASCADE,
    PRIMARY KEY (task_id, user_id)
);

CREATE TABLE kanban
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now()
);

CREATE TABLE kanban_tasks
(
    kanban_id       INTEGER REFERENCES kanban (id) ON DELETE CASCADE,
    task_id         INTEGER REFERENCES task (id) ON DELETE CASCADE,
    order_in_kanban INTEGER,
    PRIMARY KEY (kanban_id, task_id)
);

CREATE TABLE projects
(
    id          SERIAL PRIMARY KEY,
    owner_id    INTEGER      REFERENCES users (id) ON DELETE SET NULL,
    name        VARCHAR(255) NOT NULL,
    description TEXT,
    color       VARCHAR(255),
    created_at  TIMESTAMP WITHOUT TIME ZONE DEFAULT now()
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

COMMIT;