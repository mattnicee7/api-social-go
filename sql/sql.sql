CREATE DATABASE IF NOT EXISTS mattnicee7;

USE mattnicee7;
DROP TABLE IF EXISTS USUARIOS;
CREATE TABLE usuarios(
    id int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(200) not null unique,
    criadoEm timestamp default current_timestamp()
) ENGINE=INNODB;