-- +goose Up
-- +goose StatementBegin
CREATE TABLE recap.recaps
(
    id           UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id      UUID        NOT NULL REFERENCES recap.users (id) ON DELETE CASCADE,
    year         SMALLINT    NOT NULL,
    archetype    VARCHAR(32) NOT NULL,
    slides       JSONB       NOT NULL DEFAULT '[]'::jsonb,
    generated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT unique_recap_per_user_year UNIQUE (user_id, year),
    CONSTRAINT check_recaps_year CHECK (year BETWEEN 2015 AND 2100),
    CONSTRAINT check_recaps_slides_is_array CHECK (jsonb_typeof(slides) = 'array'),
    CONSTRAINT check_recaps_archetype CHECK (archetype IN
        ('collector', 'dealmaker', 'negotiator', 'explorer'))
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS recap.recaps;
-- +goose StatementEnd
