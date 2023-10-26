CREATE TABLE room_restrictions (
    id SERIAL NOT NULL,
    start_date DATE,
    end_date DATE,
    room_id INT(11) NOT NULL,
    reservation_id INT(11) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    restriction_id INT(11) NOT NULL
);

-- Create the constraints for room_id and reservation_id