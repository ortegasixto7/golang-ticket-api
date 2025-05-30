-- +goose Up
-- +goose StatementBegin
-- First, alter the column type to UUID, converting the existing data
ALTER TABLE tickets 
    ALTER COLUMN id TYPE UUID USING id::UUID;

ALTER TABLE apps 
    ALTER COLUMN id TYPE UUID USING id::UUID;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Revert the column type back to VARCHAR
ALTER TABLE tickets 
    ALTER COLUMN id TYPE VARCHAR(100) USING id::VARCHAR;

ALTER TABLE apps 
    ALTER COLUMN id TYPE VARCHAR(100) USING id::VARCHAR;
-- +goose StatementEnd 