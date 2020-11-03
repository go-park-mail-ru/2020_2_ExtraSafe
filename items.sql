DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS social_links;
DROP TABLE IF EXISTS boards CASCADE;
DROP TABLE IF EXISTS columns;
DROP TABLE IF EXISTS cards CASCADE;
DROP TABLE IF EXISTS tasks;
DROP TABLE IF EXISTS board_members;
DROP TABLE IF EXISTS board_admins;

CREATE TABLE users (
  userID      BIGSERIAL PRIMARY KEY,
  email       TEXT,
  password    TEXT,
  username    TEXT,
  fullname    TEXT,
  avatar      TEXT
);

CREATE TABLE social_links (
  userID      BIGSERIAL REFERENCES users(userID),
  networkName TEXT,
  link TEXT
);

CREATE TABLE boards (
  boardID      BIGSERIAL PRIMARY KEY,
  adminID      BIGSERIAL REFERENCES users(userID),
  boardName    TEXT,
  theme        TEXT,
  star         BOOLEAN
);

CREATE TABLE cards (
  cardID    SMALLSERIAL PRIMARY KEY,
  boardID     BIGSERIAL REFERENCES boards(boardID),
  columnName    TEXT,
  columnOrder SMALLSERIAL
);

CREATE TABLE tasks (
  taskID    SMALLSERIAL PRIMARY KEY,
  cardID    SMALLSERIAL REFERENCES cards(cardID),
  taskName    TEXT,
  description TEXT,
  tasksOrder SMALLSERIAL,
  deadline TIMESTAMP
);

CREATE TABLE board_members (
    boardID  BIGSERIAL REFERENCES boards(boardID),
    userID  BIGSERIAL REFERENCES users(userID)
);

/*CREATE TABLE board_admins (
    boardID  BIGSERIAL REFERENCES boards(boardID),
    adminID  BIGSERIAL REFERENCES users(userID)
);*/