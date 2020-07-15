CREATE EXTENSION IF NOT EXISTS CITEXT;
DROP TABLE IF EXISTS users, groups CASCADE;

CREATE TABLE users (
  id        INTEGER NOT NULL,
  nickname  CITEXT UNIQUE NOT NULL,
  user_chat_id INTEGER NOT NULL,
  is_user_busy   BOOL NOT NULL
);

CREATE TABLE groups (
  group_id        SERIAL PRIMARY KEY NOT NULL,
  is_group_busy   BOOL NOT NULL,
  first_user_id   INTEGER,
  second_user_id  INTEGER
);
