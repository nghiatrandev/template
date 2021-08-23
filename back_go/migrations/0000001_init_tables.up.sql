CREATE TABLE users (
                       id bigserial PRIMARY KEY,
                       user_name varchar(100) NOT NULL unique,
                       full_name varchar(255) NOT NULL,
                       email varchar(100) NOT NULL ,
                       password varchar(255) NOT NULL
)