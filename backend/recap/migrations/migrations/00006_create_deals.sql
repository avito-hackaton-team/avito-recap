-- +goose Up
-- +goose StatementBegin
CREATE TABLE recap.deals
(
    id           UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    listing_id   UUID        NOT NULL REFERENCES recap.listings (id) ON DELETE CASCADE,
    buyer_id     UUID        NOT NULL REFERENCES recap.users (id) ON DELETE CASCADE,
    price        BIGINT      NOT NULL,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    completed_at TIMESTAMPTZ,
    CONSTRAINT check_deals_price_non_negative CHECK (price >= 0),
    CONSTRAINT check_deals_completed_at CHECK (completed_at IS NULL OR completed_at >= created_at)
);

CREATE UNIQUE INDEX unique_completed_deal_per_listing
    ON recap.deals (listing_id) WHERE completed_at IS NOT NULL;

CREATE INDEX idx_deals_buyer_completed_at ON recap.deals (buyer_id, completed_at DESC);
CREATE INDEX idx_deals_listing_id ON recap.deals (listing_id);

COMMENT ON COLUMN recap.deals.price IS 'Цена сделки в копейках';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS recap.deals;
-- +goose StatementEnd
