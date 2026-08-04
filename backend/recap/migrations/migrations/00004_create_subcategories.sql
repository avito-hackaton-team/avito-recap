-- +goose Up
-- +goose StatementBegin
CREATE TABLE recap.subcategories
(
    id          UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    category_id UUID         NOT NULL REFERENCES recap.categories (id) ON DELETE CASCADE,
    title       VARCHAR(128) NOT NULL,
    CONSTRAINT unique_subcategory_title_per_category UNIQUE (category_id, title),
    CONSTRAINT unique_subcategory_id_category UNIQUE (id, category_id)
);

CREATE INDEX idx_subcategories_category_id ON recap.subcategories (category_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS recap.subcategories;
-- +goose StatementEnd
