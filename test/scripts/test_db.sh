#!/bin/bash

docker run --name postgres-test -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -p 6432:5432 -d postgres:latest
echo "Postgres container is running on port 6432"
sleep 3

docker exec -it postgres-test psql -U postgres -d postgres -c "CREATE DATABASE product_db"
sleep 3
echo "Database product_db is created"

docker exec -it postgres-test psql -U postgres -d product_db -c "
CREATE TABLE IF NOT EXISTS products 
(
    id BIGSERIAL NOT NULL PRIMARY KEY, 
    name VARCHAR(255) NOT NULL, 
    price DECIMAL(10, 2) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    discount DECIMAL(10, 2) DEFAULT 0,
    store VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)"

sleep 3
echo "Table products is created"