-- +goose Up
-- +goose StatementBegin
CREATE TYPE role AS ENUM ('ref', 'admin', 'org');

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    phone_number VARCHAR(11),
    user_role role NOT NULL,
    created_at TIMESTAMPTZ default NOW()
);

CREATE TABLE availabilities (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users (id),
    price NUMERIC(10,5) NOT NULL,
    start_time TIMESTAMPTZ NOT NULL,
    end_time TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ default NOW()
);

CREATE TABLE bookings (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    referee_id UUID REFERENCES users (id) NOT NULL,
    organizer_id UUID REFERENCES users (id) NOT NULL,
    price NUMERIC(10,5) NOT NULL,
    start_time TIMESTAMPTZ NOT NULL,
    end_time TIMESTAMPTZ NOT NULL,
    location VARCHAR(100) NOT NULL,
    accepted BOOL NOT NULL,
    cancelled BOOL NOT NULL,
    last_updated TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
DROP TABLE availabilities
DROP TABLE bookings
-- +goose StatementEnd
