-- +migrate Up

create table users (
    id bigserial primary key,
    address CHAR(42)
);

create table nonce (
    id bigserial primary key,
    message text not null,
    expiresat timestamp without time zone not null,
    address CHAR(42)
);

-- +migrate Down

drop table nonce;
drop table users;