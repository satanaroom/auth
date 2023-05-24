-- +goose Up
ALTER TABLE accesses DROP COLUMN IF EXISTS role;
ALTER TABLE accesses ADD COLUMN IF NOT EXISTS role integer[] NOT NULL DEFAULT array[]::integer[];

-- +goose Down
ALTER TABLE accesses DROP COLUMN IF EXISTS role;
ALTER TABLE accesses ADD COLUMN IF NOT EXISTS role text;
