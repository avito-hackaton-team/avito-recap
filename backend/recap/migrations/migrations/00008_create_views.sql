-- +goose Up
-- +goose StatementBegin
CREATE TABLE recap.views
(
    id         UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id    UUID        NOT NULL REFERENCES recap.users (id) ON DELETE CASCADE,
    listing_id UUID        NOT NULL REFERENCES recap.listings (id) ON DELETE CASCADE,
    viewed_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_views_user_viewed_at ON recap.views (user_id, viewed_at);
CREATE INDEX idx_views_listing_id ON recap.views (listing_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS recap.views;
-- +goose StatementEnd
