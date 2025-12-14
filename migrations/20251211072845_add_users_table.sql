-- +goose Up
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username TEXT UNIQUE NOT NULL,
                       password_hash TEXT NOT NULL,
                       created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

ALTER TABLE todos
    ADD CONSTRAINT fk_todos_users
        FOREIGN KEY (user_id)
            REFERENCES users (id)
            ON DELETE CASCADE;

-- +goose Down
ALTER TABLE todos
DROP CONSTRAINT IF EXISTS fk_todos_users;

DROP TABLE users;
