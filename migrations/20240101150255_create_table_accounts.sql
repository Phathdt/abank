-- +goose Up
-- +goose StatementBegin
CREATE TABLE "accounts" (
    "id" BIGSERIAL PRIMARY KEY,
    "user_id" bigint NOT NULL,
    "name" text NOT NULL,
    "balance" bigint,
    "created_at" timestamp(0) DEFAULT now(),
    "updated_at" timestamp(0) DEFAULT now()
);

ALTER TABLE "accounts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "accounts";
-- +goose StatementEnd
