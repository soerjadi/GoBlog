-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id      BIGSERIAL   NOT NULL PRIMARY KEY,
    username    VARCHAR(20) NOT NULL,
    email       VARCHAR(55) NOT NULL,
    fullname    VARCHAR(155) NOT NULL,
    passhash    VARCHAR(80) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
