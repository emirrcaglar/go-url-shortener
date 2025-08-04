DROP TABLE IF EXISTS urls;
CREATE TABLE urls (
    id INT AUTO_INCREMENT NOT NULL,
    long_url varchar(255) NOT NULL,
    userID INT,
    FOREIGN KEY (userID) REFERENCES users(id),
    PRIMARY KEY (id)
) AUTO_INCREMENT = 1000000 ;

DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id INT AUTO_INCREMENT NOT NULL,
    username varchar(255) UNIQUE NOT NULL,
    userpass varchar(255) NOT NULL,
    PRIMARY KEY (id)
);