CREATE TABLE "cart" (
    id serial not null primary key ,
    product_id int not null ,
    amount int not null,
    created_at timestamp,
    updated_at timestamp
)