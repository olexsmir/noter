CREATE TABLE sessions (
  id serial not null unique,
  user_id  integer REFERENCES users (id) unique,
  refresh_token varchar(255) not null,
  expires_at time not null
);
