\c crypton_db

DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS tokens;

CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  username VARCHAR(255) UNIQUE,
  password VARCHAR(255)
);

CREATE INDEX ON users(id);

CREATE TABLE tokens (
  id SERIAL PRIMARY KEY,
  block BIGINT
)
