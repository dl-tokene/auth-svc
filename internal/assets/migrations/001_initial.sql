-- +migrate Up

create table users (
    id bigserial primary key,
    address Bytea,
    createdat bigint not null
);

create table nonce (
    id bigserial primary key,
    message text not null,
    expiresat bigint not null,
    address Bytea

);

-- +migrate Down

drop table nonce;
drop table users;