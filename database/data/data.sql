CREATE DATABASE IF NOT EXISTS transactions;
USE transactions;

CREATE TABLE companies (
    id int not NULL,
    ibans VARCHAR(255),
    name VARCHAR(255),
    address VARCHAR(255),
    UNIQUE (id)
);

CREATE TABLE Companies (
    ID int not NULL,
    Ibans VARCHAR(255),
    Name VARCHAR(255),
    Address VARCHAR(255),
    UNIQUE (ID)
);

INSERT into companies (id, ibans, name, address) VALUES 
(1, "1234", "Heloo", "Address is here"),(2, "2211", "There", "Here Here!!");


INSERT into Companies (ID, Ibans, Name, Address) VALUES 
(1, "1234", "Heloo", "Address is here"),(2, "2211", "There", "Here Here!!");