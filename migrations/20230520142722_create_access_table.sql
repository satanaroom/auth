-- +goose Up
CREATE TABLE IF NOT EXISTS accesses(
    id SERIAL PRIMARY KEY,
    endpoint_address text NOT NULL,
    role text NOT NULL,
    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp
);


-- +goose Down
DROP TABLE IF EXISTS accesses;