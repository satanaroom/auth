-- +goose Up
CREATE TABLE IF NOT EXISTS users(
    id serial PRIMARY KEY,
    username varchar(50) NOT NULL UNIQUE,
    email varchar(50),
    password varchar(50),
    role integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


-- +goose Down
DROP TABLE IF EXISTS users;