CREATE DATABASE IF NOT EXISTS transactions;
USE transactions;

CREATE TABLE companies (
    id int not NULL,
    ibans VARCHAR(255),
    name VARCHAR(255),
    address VARCHAR(255),
    UNIQUE (id)
);

INSERT into companies (id, ibans, name, address) VALUES (1, "1234", "Heloo", "Address is here")