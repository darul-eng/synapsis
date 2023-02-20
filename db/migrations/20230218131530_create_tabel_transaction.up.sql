CREATE TABLE "transaction" (
    id serial not null primary key ,
    user_id int not null ,
    amount float not null,
    payment_url varchar(255),
    status varchar(255) not null,
    created_at timestamp,
    updated_at timestamp
)