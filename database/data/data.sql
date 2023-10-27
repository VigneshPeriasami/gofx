CREATE DATABASE IF NOT EXISTS transactions;
USE transactions;

CREATE TABLE companies (
    id int not NULL,
    ibans VARCHAR(255),
    name VARCHAR(255),
    address VARCHAR(255),
    UNIQUE (id)
);

create TABLE transactions (
    id VARCHAR(50) not NULL,
    beneficiary VARCHAR(50) not NULL,
    sender VARCHAR(50),
    currency VARCHAR(10),
    transactionTime timestamp,
    amount DECIMAL(16, 4)
);

INSERT into companies (id, ibans, name, address) VALUES 
(1, "1234", "Heloo", "Address is here"),(2, "2211", "There", "Here Here!!");

