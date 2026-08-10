-- +goose Up
-- +goose StatementBegin
CREATE TABLE recap.favorites
(
    user_id    UUID        NOT NULL REFERENCES recap.users (id) ON DELETE CASCADE,
    listing_id UUID        NOT NULL REFERENCES recap.listings (id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT pk_favorites PRIMARY KEY (user_id, listing_id)
);

CREATE INDEX idx_favorites_user_created_at ON recap.favorites (user_id, created_at);
CREATE INDEX idx_favorites_listing_id ON recap.favorites (listing_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS recap.favorites;
-- +goose StatementEnd
