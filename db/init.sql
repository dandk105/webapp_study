-- Connect to the database
\c mydatabase

-- Create tables and define schema here
-- init.sql

-- Enable UUID support
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- USERS Table
CREATE TABLE USERS (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4() UNIQUE NOT NULL,
    name TEXT NOT NULL CHECK (length(name) > 0 AND length(name) <= 50),
    birth_day DATE
);

-- ROOMS Table
CREATE TABLE ROOMS (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4() UNIQUE NOT NULL,
    name TEXT NOT NULL,
    capacity SMALLINT NOT NULL
);

-- RESERVATIONS Table
CREATE TABLE RESERVATIONS (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4() UNIQUE NOT NULL,
    reserved_user UUID REFERENCES USERS(id) NOT NULL,
    reserved_room UUID REFERENCES ROOMS(id) NOT NULL,
    start_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP NOT NULL
);


-- Add sample data to USERS
INSERT INTO USERS (id, name, birth_day) VALUES
    (uuid_generate_v4(), 'Alice', '1990-01-01'),
    (uuid_generate_v4(), 'Bob', '1992-05-15'),
    (uuid_generate_v4(), 'Charlie', '1988-11-23');

-- Add sample data to ROOMS
INSERT INTO ROOMS (id, name, capacity) VALUES
    (uuid_generate_v4(), 'Meeting Room 101', 10),
    (uuid_generate_v4(), 'Conference Hall', 100),
    (uuid_generate_v4(), 'Small Room A', 4);