-- +migrate Up

create table users (
    id bigserial primary key,
    address CHAR(42)
);

create table nonce (
    id bigserial primary key,
    message text not null,
    expiresAt timestamp without time zone not null,
    address CHAR(42)
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
drop table users;