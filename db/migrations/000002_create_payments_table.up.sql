create table payment (
    id serial,
    account_from_id integer not null references account(id),
    account_to_id integer not null references account(id),
    created  timestamp not null default now(),
    amount numeric(12, 2) not null check (amount >=0)
);