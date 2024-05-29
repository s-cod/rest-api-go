CREATE TABLE users(
  id bigserial not null primary key,
  amail varchar not null unique,
  emcript_password varchar not null
);
