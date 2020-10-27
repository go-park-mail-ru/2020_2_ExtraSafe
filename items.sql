DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS social_links;
DROP TABLE IF EXISTS boards;
DROP TABLE IF EXISTS columns;
DROP TABLE IF EXISTS tasks;

CREATE TABLE users (
  userID      BIGSERIAL PRIMARY KEY,
  email       TEXT,
  password    TEXT,
  username    TEXT,
  fullname    TEXT,
  avatar      TEXT
);

CREATE TABLE social_links (
  userID      BIGSERIAL,
  networkName TEXT,
  link TEXT
);

CREATE TABLE boards (
  boardID      SMALLSERIAL PRIMARY KEY,
  adminID      BIGSERIAL,
  boardName    TEXT,
  theme        TEXT
);

CREATE TABLE columns (
  columnID    SMALLSERIAL PRIMARY KEY,
  boardID     SMALLSERIAL,
  columnName    TEXT,
  columnOrder SMALLSERIAL
);

CREATE TABLE tasks (
  taskID    SMALLSERIAL PRIMARY KEY,
  columnID    SMALLSERIAL,
  taskName    TEXT,
  description TEXT,
  tasksOrder SMALLSERIAL,
  deadline TIMESTAMP
);
