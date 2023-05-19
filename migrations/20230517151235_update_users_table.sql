-- +goose Up
ALTER TABLE users ADD COLUMN IF NOT EXISTS department_type integer;

-- +goose Down
ALTER TABLE users DROP COLUMN IF EXISTS department_type;