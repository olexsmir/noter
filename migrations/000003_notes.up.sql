CREATE TABLE notes (
  id serial not null unique,
  author_id integer REFERENCES users (id),
  title varchar(255) not null,
  content text not null,
  created_at time not null,
  updated_at time not null
);
