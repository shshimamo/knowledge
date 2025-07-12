-- +goose Up
-- +goose StatementBegin
CREATE TABLE tags
(
    id         BIGSERIAL PRIMARY KEY,
    name       VARCHAR(50) NOT NULL UNIQUE,
    created_at TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE knowledge_tags
(
    id           BIGSERIAL PRIMARY KEY,
    knowledge_id BIGINT    NOT NULL REFERENCES knowledges (id) ON DELETE CASCADE,
    tag_id       BIGINT    NOT NULL REFERENCES tags (id) ON DELETE CASCADE,
    created_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (knowledge_id, tag_id)
);

CREATE INDEX idx_knowledge_tags_knowledge_id ON knowledge_tags (knowledge_id);
CREATE INDEX idx_knowledge_tags_tag_id ON knowledge_tags (tag_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_knowledge_tags_tag_id;
DROP INDEX IF EXISTS idx_knowledge_tags_knowledge_id;
DROP TABLE knowledge_tags;
DROP TABLE tags;
-- +goose StatementEnd