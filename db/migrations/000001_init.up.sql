drop schema if exists loyalty_system cascade;
create schema loyalty_system;

create table loyalty_system.users
(
    id serial primary key,
    name varchar not null,
    email varchar not null,
    password varchar not null,
    admin bool default false,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
);