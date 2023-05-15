-- +goose Up

-- +goose StatementBegin
CREATE TABLE events
(
    id UUID NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP NOT NULL,
    description VARCHAR NOT NULL,
    type VARCHAR NOT NULL
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
