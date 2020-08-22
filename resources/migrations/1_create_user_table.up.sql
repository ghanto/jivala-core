CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table if not exists "user"
(
    id uuid     default uuid_generate_v4() constraint user_pk primary key,
    login       varchar(255) not null,
    password    varchar(255) not null,
    email       varchar(255) not null,
    created_at timestamp default now()
);

create unique index user_email_uindex on "user" (email);

