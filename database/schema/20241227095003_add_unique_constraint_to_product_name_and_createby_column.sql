-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE products ADD COLUMN created_by INT NOT NULL DEFAULT 0;
ALTER TABLE products ADD CONSTRAINT unique_product_name UNIQUE (name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE products DROP CONSTRAINT unique_product_name;
ALTER TABLE products DROP COLUMN created_by;
-- +goose StatementEnd
