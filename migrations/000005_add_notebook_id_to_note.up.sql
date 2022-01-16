ALTER TABLE notes
ADD notebook_id integer REFERENCES notebooks (id) not null;
