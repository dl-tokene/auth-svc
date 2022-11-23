-- +migrate Up


create table nonce (
    id bigserial primary key,
    message text not null,
    expiresat bigint not null,
    address Bytea

);

-- +migrate Down

drop table nonce;
