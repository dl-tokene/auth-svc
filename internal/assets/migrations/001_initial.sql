-- +migrate Up

create table users (
    id bigserial primary key,
    options text
);

create table address (
    address text not null primary key,
    user_id bigint not null references users(id) on delete cascade
);

create table nonce (
    id bigserial primary key,
    message text not null,
    expires timestamp without time zone not null,
    address text not null
);

create table log (
    id bigserial primary key,
    user_id bigint not null references users(id) on delete cascade,
    log_time timestamp without time zone not null,
    type text not null,
    notes jsonb not null
);


-- +migrate Down

drop table log;
drop table nonce;
drop table address;
drop table users;