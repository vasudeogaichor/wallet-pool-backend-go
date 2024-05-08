-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR NOT NULL,
    last_name VARCHAR,
    username VARCHAR NOT NULL UNIQUE,
    email VARCHAR NOT NULL UNIQUE,
    mobile VARCHAR UNIQUE,
    password VARCHAR NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE users;
-- +goose StatementEnd
