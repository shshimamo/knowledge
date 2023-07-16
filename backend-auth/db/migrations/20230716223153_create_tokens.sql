-- +goose Up
-- +goose StatementBegin
CREATE TABLE tokens (
                       id SERIAL PRIMARY KEY,
                       auth_user_id bigint NOT NULL,
                       token VARCHAR(255) NOT NULL,
                       active BOOLEAN NOT NULL DEFAULT FALSE,
                       expires_at TIMESTAMP NOT NULL,
                       created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tokens;
-- +goose StatementEnd
