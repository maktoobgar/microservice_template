
-- +migrate Up
CREATE TABLE users (
    id INTEGER NOT NULL UNIQUE PRIMARY KEY AUTOINCREMENT,
    email varchar(64) UNIQUE,
    phone_number varchar(64) NOT NULL UNIQUE,
    password varchar(255) NOT NULL,
    email_confirmed boolean DEFAULT false,
    phone_number_confirmed boolean DEFAULT false,
    role varchar(2) NOT NULL DEFAULT 00,
    joined_date DATETIME NOT NULL
);
-- +migrate Down
DROP TABLE users;
