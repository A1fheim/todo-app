-- +goose Up
CREATE TABLE todos (
                       id SERIAL PRIMARY KEY,
                       user_id INT NOT NULL,
                       title TEXT NOT NULL,
                       description TEXT,
                       status TEXT NOT NULL DEFAULT 'todo',
                       due_date TIMESTAMP NULL,
                       created_at TIMESTAMP NOT NULL DEFAULT NOW(),
                       updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE todos;
