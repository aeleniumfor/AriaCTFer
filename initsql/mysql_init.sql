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

CREATE TABLE groups (
  id         SERIAL PRIMARY KEY NOT NULL,
  name       VARCHAR(256)       NOT NULL UNIQUE,
  email      VARCHAR(256)       NOT NULL UNIQUE,
  created_at TIMESTAMP          NOT NULL DEFAULT current_timestamp,
  updated_at TIMESTAMP          NOT NULL DEFAULT current_timestamp
);

CREATE TABLE groups_to_users (
  id         SERIAL PRIMARY KEY NOT NULL,
  group_id   INT REFERENCES groups (id) ON UPDATE CASCADE ON DELETE CASCADE,
  user_id    INT REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE,
  created_at TIMESTAMP          NOT NULL DEFAULT current_timestamp,
  updated_at TIMESTAMP          NOT NULL DEFAULT current_timestamp
);
