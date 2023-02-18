CREATE TABLE "transaction" (
    id serial not null primary key ,
    user_id int not null ,
    amount float not null
)