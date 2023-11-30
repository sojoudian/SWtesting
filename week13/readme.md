#Setup Mysql
mysql -u root -p

CREATE USER 'newuser'@'localhost' IDENTIFIED BY 'password';

GRANT ALL PRIVILEGES ON *.* TO 'newuser'@'localhost' WITH GRANT OPTION;


#create database
CREATE DATABASE goAPI;
USE DATABASE goAPI;

#create table
CREATE TABLE time_table (
    id INT AUTO_INCREMENT PRIMARY KEY,
    time DATETIME NOT NULL
);


# install mysql driver

go get -u github.com/go-sql-driver/mysql