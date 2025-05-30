-- +goose Up
-- +goose StatementBegin
ALTER TABLE integrations RENAME TO apps;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE apps RENAME TO integrations;
-- +goose StatementEnd 