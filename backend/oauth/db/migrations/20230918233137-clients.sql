-- +migrate Up
CREATE TABLE clients
(
    id           BIGSERIAL PRIMARY KEY,
    client_id    BIGINT,
    name         VARCHAR(255),
    redirect_url VARCHAR(255),
    created_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +migrate Down
DROP TABLE clients;