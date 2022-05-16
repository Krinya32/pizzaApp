CREATE TABLE pizzas
(
    id serial not null unique,
    title varchar(255) not null,
    price integer not null,
    description varchar(255),
    spicy bool,
    available bool
)