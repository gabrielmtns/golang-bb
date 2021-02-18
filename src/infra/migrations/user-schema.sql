CREATE DATABASE IF NOT EXISTS go_test_gabriel_bubble

CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    name  VARCHAR,
    lastname VARCHAR,
    username VARCHAR UNIQUE,
    password VARCHAR,
    salt VARCHAR,
    createdAt TIMESTAMP NOT NULL DEFAULT NOW()
)