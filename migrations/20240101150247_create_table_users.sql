-- +goose Up
-- +goose StatementBegin
CREATE TABLE "users" (
    "id" BIGSERIAL PRIMARY KEY,
    "name" text NOT NULL,
    "created_at" timestamp(0) DEFAULT now(),
    "updated_at" timestamp(0) DEFAULT now()
);

INSERT INTO "public"."users" ("id", "name", "created_at", "updated_at") VALUES
(1, 'Alice', '2024-01-01 15:07:38', '2024-01-01 15:07:38'),
(2, 'Bill', '2024-01-01 15:07:38', '2024-01-01 15:07:38'),
(3, 'Charles', '2024-01-01 15:07:38', '2024-01-01 15:07:38');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "users";
-- +goose StatementEnd
