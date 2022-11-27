drop schema if exists loyalty_system cascade;
create schema loyalty_system;

create table loyalty_system.organization
(
    id serial primary key,
    name varchar(255) not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
);

create table loyalty_system.user
(
    id serial primary key,
    first_name varchar(255) not null,
    last_name varchar(255) not null,
    email varchar(255) unique not null,
    password_hash varchar(255) not null,
    is_admin boolean default false,
    organization_id int default null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    constraint fk_organization_id foreign key(organization_id) references loyalty_system.organization(id)
);