-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE SEQUENCE products_id_seq START 100 MINVALUE 100;

-- products table
CREATE TABLE products (
    id INT PRIMARY KEY DEFAULT NEXTVAL('products_id_seq'),
    name TEXT NOT NULL,
    description TEXT,
    price NUMERIC(10, 2) NOT NULL,
    stock INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE products;
-- +goose StatementEnd
