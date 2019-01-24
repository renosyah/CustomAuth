DROP DATABASE custom_auth_db;

CREATE DATABASE custom_auth_db;

USE custom_auth_db;

CREATE TABLE user(
    id varchar(250),
    name varchar(50),
    email varchar(50),
    username varchar(50),
    password varchar(100)
);

ALTER TABLE user ADD PRIMARY KEY (id);