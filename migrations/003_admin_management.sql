CREATE TABLE admins (
    id BIGSERIAL PRIMARY KEY,

    username VARCHAR(50) NOT NULL,

    password_hash TEXT NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX admins_username_unique
ON admins (LOWER(username));