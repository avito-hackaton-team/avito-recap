-- +goose Up
-- +goose StatementBegin
CREATE TABLE recap.categories
(
    id    UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title VARCHAR(128) NOT NULL,
    CONSTRAINT unique_category_title UNIQUE (title)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS recap.categories;
-- +goose StatementEnd
