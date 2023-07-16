-- +goose Up
-- +goose StatementBegin
CREATE TABLE auth_users (
                       id SERIAL PRIMARY KEY,
                       email VARCHAR(255),
                       password_digest VARCHAR(255),
                       created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE auth_users;
-- +goose StatementEnd
