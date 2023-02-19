CREATE TABLE "product" (
    id serial not null primary key,
    name varchar(255) not null ,
    category_id int not null,
    price float not null,
    created_at timestamp,
    updated_at timestamp
)