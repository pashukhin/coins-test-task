create table account (
    id serial unique,
    name varchar(256) not null unique,
    balance numeric(12, 2) not null check (balance >= 0),
    currency char(3) not null
);

create table payment (
    id serial,
    account_from_id integer not null references account(id),
    account_to_id integer not null references account(id),
    created  timestamp not null default now(),
    amount numeric(12, 2) not null check (amount >=0)
);

insert into account (name, balance, currency) values
('Alice', 0, 'USD'),
('Bob', 10, 'USD'),
('Boris', 1000, 'RUR'),
('Eve', 1000000, 'USD'),
('Natasha', 10000, 'RUR'),
('Vladimir', 1000000, 'RUR');