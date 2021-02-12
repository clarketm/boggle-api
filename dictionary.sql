-- Create a `db` database.
CREATE DATABASE IF NOT EXISTS db;

-- Drop `dictionary` table.
DROP TABLE db.dictionary;

-- Create a `dictionary` table.
CREATE TABLE db.dictionary
(
    word VARCHAR(50) NOT NULL PRIMARY KEY
);

-- docker cp ./dictionary.txt mysql:/var/lib/mysql

-- Load word data into the `dictionary` table.
# LOAD DATA INFILE './dictionary.txt' INTO TABLE db.dictionary LINES TERMINATED BY '\n';
