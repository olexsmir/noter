CREATE TABLE notebooks (
  id serial not null unique,
  author_id integer REFERENCES users (id),
  name varchar(255) not null unique,
  description varchar(255),
  created_at time not null,
  updated_at time not null
);
