-- +goose Up
-- +goose StatementBegin
CREATE TABLE recap.listings
(
    id             UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    seller_id      UUID         NOT NULL REFERENCES recap.users (id) ON DELETE CASCADE,
    title          VARCHAR(255) NOT NULL,
    description    TEXT,
    price          BIGINT       NOT NULL,
    category_id    UUID         NOT NULL REFERENCES recap.categories (id) ON DELETE RESTRICT,
    subcategory_id UUID,
    status         VARCHAR(16)  NOT NULL DEFAULT 'active',
    created_at     TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    closed_at      TIMESTAMPTZ,
    CONSTRAINT fk_listings_subcategory FOREIGN KEY (subcategory_id, category_id)
        REFERENCES recap.subcategories (id, category_id) ON DELETE RESTRICT,
    CONSTRAINT check_listings_price_non_negative CHECK (price >= 0),
    CONSTRAINT check_listings_status CHECK (status IN ('active', 'sold', 'closed')),
    CONSTRAINT check_listings_closed_at CHECK (closed_at IS NULL OR closed_at >= created_at)
);

CREATE INDEX idx_listings_seller_created_at ON recap.listings (seller_id, created_at DESC);
CREATE INDEX idx_listings_category_created_at ON recap.listings (category_id, created_at DESC);
CREATE INDEX idx_listings_subcategory_id ON recap.listings (subcategory_id);
CREATE INDEX idx_listings_status ON recap.listings (status) WHERE status = 'active';

COMMENT ON COLUMN recap.listings.price IS 'Цена в копейках';
COMMENT ON COLUMN recap.listings.status IS 'active | sold | closed';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS recap.listings;
-- +goose StatementEnd
