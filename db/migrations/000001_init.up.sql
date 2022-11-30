create table organization
(
    id serial primary key,
    name varchar(255) not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
);

create table "user"
(
    id serial primary key,
    first_name varchar(255) not null,
    last_name varchar(255) not null,
    email varchar(255) unique not null,
    password_hash varchar(255) not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
);

create table organization_user
(
    id serial primary key,
    organization_id int not null,
    user_id int not null,
    constraint fk_organization_id foreign key(organization_id) references organization(id) on delete cascade,
    constraint fk_user_id foreign key(user_id) references "user"(id) on delete cascade
);

create table admin_organization
(
    id serial primary key,
    organization_id int not null,
    user_id int not null,
    constraint fk_organization_id foreign key(organization_id) references organization(id) on delete cascade,
    constraint fk_user_id foreign key(user_id) references "user"(id) on delete cascade
)