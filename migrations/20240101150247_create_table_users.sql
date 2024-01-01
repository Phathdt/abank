-- +goose Up
-- +goose StatementBegin
CREATE TABLE "users" (
    "id" BIGSERIAL PRIMARY KEY,
    "name" text NOT NULL,
    "created_at" timestamp(0) DEFAULT now(),
    "updated_at" timestamp(0) DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "users";
-- +goose StatementEnd
