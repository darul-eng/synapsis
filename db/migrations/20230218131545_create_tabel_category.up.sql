CREATE TABLE "category" (
    id serial not null primary key ,
    name varchar(255) not null,
    created_at timestamp,
    updated_at timestamp
)