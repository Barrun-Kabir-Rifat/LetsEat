
-- Migration: 002_menu_management.sql
-- Purpose: Create categories and menu_items tables.

CREATE TABLE IF NOT EXISTS categories (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT categories_name_not_empty
        CHECK (LENGTH(TRIM(name)) > 0),

    CONSTRAINT categories_name_unique
        UNIQUE (name)
);

CREATE TABLE IF NOT EXISTS menu_items (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,

    restaurant_id BIGINT NOT NULL,
    category_id BIGINT NOT NULL,

    name VARCHAR(150) NOT NULL,
    description TEXT,
    price NUMERIC(10, 2) NOT NULL,
    is_available BOOLEAN NOT NULL DEFAULT TRUE,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT menu_items_name_not_empty
        CHECK (LENGTH(TRIM(name)) > 0),

    CONSTRAINT menu_items_price_nonnegative
        CHECK (price >= 0),

    CONSTRAINT menu_items_restaurant_fk
        FOREIGN KEY (restaurant_id)
        REFERENCES restaurants(id)
        ON DELETE CASCADE,

    CONSTRAINT menu_items_category_fk
        FOREIGN KEY (category_id)
        REFERENCES categories(id)
        ON DELETE RESTRICT
);

CREATE INDEX IF NOT EXISTS idx_menu_items_restaurant_id
    ON menu_items(restaurant_id);

CREATE INDEX IF NOT EXISTS idx_menu_items_category_id
    ON menu_items(category_id);
