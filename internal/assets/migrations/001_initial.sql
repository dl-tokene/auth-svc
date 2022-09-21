-- +migrate Up

create table users (
    id bigserial primary key,
    created_at timestamp without time zone not null,
    address Bytea
);

create table nonce (
    id bigserial primary key,
    message text not null,
    expiresat timestamp without time zone not null,
    address Bytea
);

-- +migrate Down

drop table nonce;
drop table users;