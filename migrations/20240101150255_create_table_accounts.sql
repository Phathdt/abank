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

INSERT INTO "public"."accounts" ("id", "user_id", "name", "balance", "created_at", "updated_at") VALUES
(1, 1, 'Alice first account', 10000, '2024-01-01 15:08:23', '2024-01-01 15:08:23'),
(2, 1, 'Alice second account', 10000, '2024-01-01 15:08:23', '2024-01-01 15:08:23'),
(3, 1, 'Alice third account', 10000, '2024-01-01 15:08:23', '2024-01-01 15:08:23'),
(4, 2, 'Bill first account', 10000, '2024-01-01 15:09:02', '2024-01-01 15:09:02'),
(5, 2, 'Bill second account', 10000, '2024-01-01 15:09:02', '2024-01-01 15:09:02'),
(6, 2, 'Bill third account', 10000, '2024-01-01 15:09:02', '2024-01-01 15:09:02'),
(7, 3, 'Charles first account', 10000, '2024-01-01 15:09:02', '2024-01-01 15:09:02'),
(8, 3, 'Charles second account', 10000, '2024-01-01 15:09:02', '2024-01-01 15:09:02'),
(9, 3, 'Charles third account', 10000, '2024-01-01 15:09:02', '2024-01-01 15:09:02');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "accounts";
-- +goose StatementEnd
