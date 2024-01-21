create database cube_inventory;
use cube_inventory;

create table USERS(
    id int not null auto_increment,
    email varchar(255) not null,
    name varchar(255) not null,
    password varchar(255) not null,
    primary key(id)
);

create table CUBES(
    id int not null auto_increment,
    name varchar(255) not null,
    brand varchar(255) not null,
    shape varchar(255) not null,
    primary key(id)
);