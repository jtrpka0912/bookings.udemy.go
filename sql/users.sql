CREATE TABLE users (
    id SERIAL(11) NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    access_level INT(11),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);