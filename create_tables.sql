CREATE DATABASE links;

USE links;

CREATE TABLE `links` (
    `slug` VARCHAR(64) NOT NULL DEFAULT '',
    `url` VARCHAR(255) NOT NULL DEFAULT '',
    PRIMARY KEY (`slug`)
);