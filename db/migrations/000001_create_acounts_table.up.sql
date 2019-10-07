create table account (
    id serial unique,
    name varchar(256) not null unique,
    balance numeric(12, 2) not null check (balance >= 0),
    currency char(3) not null
);