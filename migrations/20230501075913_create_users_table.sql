-- +goose Up
CREATE TABLE IF NOT EXISTS users(
    id serial PRIMARY KEY,
    username text NOT NULL UNIQUE,
    email text,
    password text,
    role integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


-- +goose Down
DROP TABLE IF EXISTS users;