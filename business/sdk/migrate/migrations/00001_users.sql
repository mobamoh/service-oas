-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    user_id       UUID        NOT NULL,
    name          TEXT        NOT NULL,
    email         TEXT UNIQUE NOT NULL,
    roles         TEXT[]      NOT NULL,
    password_hash BYTEA       NOT NULL,
    department    TEXT NULL,
    enabled       BOOLEAN     NOT NULL,
    date_created  TIMESTAMP   NOT NULL,
    date_updated  TIMESTAMP   NOT NULL,

    PRIMARY KEY (user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users CASCADE;
-- +goose StatementEnd