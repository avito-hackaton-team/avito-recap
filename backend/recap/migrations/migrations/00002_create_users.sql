-- +goose Up
-- +goose StatementBegin
CREATE TABLE recap.users
(
    id         UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name       VARCHAR(64)  NOT NULL,
    surname    VARCHAR(64)  NOT NULL,
    avatar_url VARCHAR(512),
    hint       VARCHAR(255),
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

COMMENT ON COLUMN recap.users.hint IS 'Короткое описание тестового профиля для витрины выбора';
COMMENT ON COLUMN recap.users.created_at IS 'Дата регистрации, отдается в API как registeredAt';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS recap.users;
-- +goose StatementEnd
