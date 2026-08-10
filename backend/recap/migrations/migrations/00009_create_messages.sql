-- +goose Up
-- +goose StatementBegin
CREATE TABLE recap.messages
(
    id         UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    buyer_id   UUID        NOT NULL REFERENCES recap.users (id) ON DELETE CASCADE,
    seller_id  UUID        NOT NULL REFERENCES recap.users (id) ON DELETE CASCADE,
    listing_id UUID        NOT NULL REFERENCES recap.listings (id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT check_messages_different_users CHECK (buyer_id <> seller_id)
);

CREATE INDEX idx_messages_buyer_created_at ON recap.messages (buyer_id, created_at);
CREATE INDEX idx_messages_seller_created_at ON recap.messages (seller_id, created_at);
CREATE INDEX idx_messages_listing_id ON recap.messages (listing_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS recap.messages;
-- +goose StatementEnd
