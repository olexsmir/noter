CREATE TABLE users (
  id serial not null unique,
  name varchar(255) not null,
  email varchar(255) not null unique,
  password varchar(255) not null,
  registred_at time not null,
  last_visit_at time not null
);
