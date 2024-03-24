create database if not exists follow; 

-- 创建用户表
create table user (
    id int auto_increment primary key,
    username varchar(50) not null,
    password varchar(50) not null,
    email varchar(100) not null,
    phone varchar(11) not null,
    status int not null,
    is_admin int not null,
    create_time int not null
);


-- 创建脚本表
create table script (
    id bigint auto_increment primary key,
    username varchar(50) not null,
    name varchar(50) not null unique ,
    type varchar(50) not null,
    language varchar(50) not null,
    Code varchar(150) not null,
    cycle int default 360,
    status bool default false,
    create_time bigint not null,
    update_time bigint not null,
    description varchar(150) not null
);

