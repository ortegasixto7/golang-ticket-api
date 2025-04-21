-- +goose Up
-- +goose StatementBegin
CREATE TABLE tickets (
    id VARCHAR(100) PRIMARY KEY NOT NULL,
    title VARCHAR(100) NOT NULL,
    code VARCHAR(100) NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ,
    expires_at TIMESTAMPTZ
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tickets;
-- +goose StatementEnd