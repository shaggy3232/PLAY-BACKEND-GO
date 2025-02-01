-- +goose Up
-- +goose StatementBegin
CREATE TYPE role AS ENUM ('ref', 'admin', 'org');

CREATE TABLE users (
    id UUID PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    phone_number VARCHAR(11),
    type role NOT NULL,
    created_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE availabilities (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users (id),
    price MONEY NOT NULL,
    start_time TIMESTAMPTZ NOT NULL,
    end_time TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE bookings (
    id UUID PRIMARY KEY,
    referee_id UUID REFERENCES users (id) NOT NULL,
    organizer_id UUID REFERENCES users (id) NOT NULL,
    price MONEY NOT NULL,
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
