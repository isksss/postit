-- database
CREATE DATABASE IF NOT EXISTS `postit`;
use postit;

-- users
DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(256) NOT NULL,
    email VARCHAR(256) NOT NULL UNIQUE,
    password VARCHAR(256) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- postit
drop table if exists postits;
create table postits(
	id integer primary key auto_increment,
    user_id integer not null,
    `text` text not null,
    created_at timestamp not null default current_timestamp,
    
    foreign key(user_id) references users(id) on update cascade on delete cascade
);

-- session
drop table if exists session;
create table sessions(
	id integer primary key auto_increment,
    user_id integer not null,
    `key` varchar(256) not null,
    created_at timestamp not null default current_timestamp,
    deleted_at timestamp default null,
    
    foreign key(user_id) references users(id) on update cascade on delete cascade
);