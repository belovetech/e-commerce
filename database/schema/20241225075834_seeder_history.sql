-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE seeding_history (
    id SERIAL PRIMARY KEY,
    seeder_name TEXT UNIQUE NOT NULL,
    executed_at TIMESTAMP NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE seeding_history;
-- +goose StatementEnd
