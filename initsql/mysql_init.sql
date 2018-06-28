CREATE DATABASE aria;

\c aria

CREATE TABLE users (
  id         SERIAL PRIMARY KEY NOT NULL,
  name       VARCHAR(256)       NOT NULL UNIQUE,
  email      VARCHAR(256)       NOT NULL UNIQUE,
  password   VARCHAR(64)        NOT NULL,
  created_at TIMESTAMP          NOT NULL DEFAULT current_timestamp,
  updated_at TIMESTAMP          NOT NULL DEFAULT current_timestamp
);
