CREATE TABLE customers (id int NOT NULL, name VARCHAR(255), email VARCHAR(255), telephone VARCHAR(255),  PRIMARY KEY (id));

CREATE TABLE cards (id VARCHAR(16) NOT NULL, pin VARCHAR(4), balance decimal, customerID int, PRIMARY KEY(id), FOREIGN KEY (customerID) REFERENCES customers(id));
