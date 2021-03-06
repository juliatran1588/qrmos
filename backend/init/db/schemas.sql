-- Initial SQL script to initialize the database for QRMOS.

-- For PRODUCTION: need to have database named `qrmos` CREATED
--   to be able to run this script.

-- For DEVELOPMENT: this script will be executed when database 
--   docker-compose container run for the first time

USE qrmos;

CREATE TABLE users (
  username VARCHAR(255) NOT NULL PRIMARY KEY,
  password VARCHAR(255) NOT NULL,
  password_salt VARCHAR(255) NOT NULL,
  full_name VARCHAR(30) NOT NULL,
  role VARCHAR(10),
  active BOOLEAN default true
);

CREATE TABLE delivery_destinations (
  name VARCHAR(255) NOT NULL PRIMARY KEY,
  security_code VARCHAR(255) NOT NULL
);