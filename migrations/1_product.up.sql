CREATE TYPE "product_category" AS ENUM (
    'boots',
    'sandals',
    'sneakers'
    );

CREATE TABLE "product"
(
    "id"         BIGSERIAL PRIMARY KEY,
    "name"       varchar          NOT NULL,
    "sku"        varchar          NOT NULL,
    "price"      bigint           NOT NULL,
    "category"   product_category NOT NULL,
    "created_at" timestamp DEFAULT (now()),
    "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "discount"
(
    "id"         BIGSERIAL PRIMARY KEY,
    "field"      varchar NOT NULL,
    "value"      varchar NOT NULL,
    "amount"     bigint  NOT NULL,
    "priority"   bigint  NOT NULL,
    "created_at" timestamp DEFAULT (now()),
    "updated_at" timestamp DEFAULT (now())
);
