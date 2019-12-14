CREATE DATABASE IF NOT EXISTS hello_world CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE hello_world;

CREATE TABLE hello (
    id int(10) NOT NULL AUTO_INCREMENT,
    name TEXT NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE world (
    id int(10) NOT NULL AUTO_INCREMENT,
    name TEXT NOT NULL,
    hello_id int(10) NOT NULL
    FOREIGN KEY (hello_id) REFERENCES hello (id),
    PRIMARY KEY (id)
);