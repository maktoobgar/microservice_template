
-- +migrate Up
CREATE TABLE users (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    email VARCHAR(64),
    phone_number VARCHAR(64) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    email_confirmed BOOLEAN NOT NULL DEFAULT(FALSE),
    phone_number_confirmed BOOLEAN NOT NULL DEFAULT(FALSE),
    role VARCHAR(2) NOT NULL DEFAULT('00'),
    joined_date DATETIME NOT NULL
);
-- +migrate Down
DROP TABLE users;
