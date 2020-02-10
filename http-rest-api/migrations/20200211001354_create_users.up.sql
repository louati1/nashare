CREATE TABLE users (
    id bigserialnot null primary key,
    email varchar not null unique,
    encrypted_password varchar not null
);