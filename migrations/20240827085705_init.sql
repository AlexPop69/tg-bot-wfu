-- +goose Up
-- +goose StatementBegin

-- Create table for users
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    telegram_username VARCHAR(255) UNIQUE NOT NULL,
    phone_number VARCHAR(20),
    number_of_orders INT DEFAULT 1,
    added_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create table for shops
CREATE TABLE IF NOT EXISTS shops (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    address VARCHAR(255),
    open_time TIME,
    close_time TIME
);

-- Create table for orders
CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    shop VARCHAR(255) NOT NULL,
    items JSONB NOT NULL,
    order_amount NUMERIC(10, 2) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create table for admins
CREATE TABLE IF NOT EXISTS admins (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    telegram_username VARCHAR(255) UNIQUE NOT NULL,
    shop_id INT REFERENCES shops(id) ON DELETE SET NULL,
    is_superadmin BOOLEAN DEFAULT FALSE
);

-- Create table for menu items
CREATE TABLE IF NOT EXISTS menu (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price NUMERIC(10, 2) NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS admins;
DROP TABLE IF EXISTS menu;
DROP TABLE IF EXISTS shops;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
