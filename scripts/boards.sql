DROP TABLE IF EXISTS boards CASCADE;
DROP TABLE IF EXISTS cards CASCADE;
DROP TABLE IF EXISTS tasks;
DROP TABLE IF EXISTS board_members;

CREATE TABLE boards (
    boardID      SERIAL PRIMARY KEY,
    adminID      INTEGER,
                      --  adminID      INTEGER REFERENCES users(userID) ON DELETE CASCADE,
    boardName    TEXT,
    theme        TEXT,
    star         BOOLEAN
);

CREATE TABLE cards (
    cardID      SERIAL PRIMARY KEY,
    boardID     INTEGER REFERENCES boards(boardID) ON DELETE CASCADE,
    cardName    TEXT,
    cardOrder   SMALLINT
);

CREATE TABLE tasks (
                       taskID      SERIAL PRIMARY KEY,
                       cardID      INTEGER REFERENCES cards(cardID) ON DELETE CASCADE,
                       taskName    TEXT,
                       description TEXT,
                       tasksOrder  SMALLINT,
                       deadline    TIMESTAMP
);

CREATE TABLE board_members (
                               boardID  INTEGER REFERENCES boards(boardID) ON DELETE CASCADE,
                               userID  INTEGER
                              -- userID  INTEGER REFERENCES users(userID) ON DELETE CASCADE
);