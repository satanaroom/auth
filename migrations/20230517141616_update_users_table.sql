-- +goose Up
ALTER TABLE users ADD COLUMN IF NOT EXISTS department jsonb;

-- +goose Down
ALTER TABLE users DROP COLUMN IF EXISTS department;