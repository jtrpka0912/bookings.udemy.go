CREATE TABLE reservations (
    id SERIAL NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    phone VARCHAR(50) NOT NULL,
    start_date DATE,
    end_date DATE,
    room_id INT(11) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

-- Add foreign key for room_id to rooms