-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE products ADD COLUMN is_available BOOLEAN DEFAULT TRUE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE products DROP COLUMN is_available;
-- +goose StatementEnd
