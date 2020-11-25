DROP TABLE IF EXISTS boards CASCADE;
DROP TABLE IF EXISTS cards CASCADE;
DROP TABLE IF EXISTS tasks CASCADE;
DROP TABLE IF EXISTS board_members;

CREATE TABLE boards (
    boardID      SERIAL PRIMARY KEY,
    adminID      INTEGER,
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
);

DROP TABLE IF EXISTS comments;
DROP TABLE IF EXISTS checklists;
DROP TABLE IF EXISTS task_members;
DROP TABLE IF EXISTS task_tags;
DROP TABLE IF EXISTS tags CASCADE;
DROP TABLE IF EXISTS attaches;

CREATE TABLE comments (
    commentID SERIAL PRIMARY KEY,
    message TEXT,
    taskID INTEGER REFERENCES tasks(taskID)
);

-- ID  ДОСКИ!!!!
CREATE TABLE tags (
    tagID SERIAL PRIMARY KEY,
    boardID  INTEGER REFERENCES boards(boardID),
    name TEXT,
    color TEXT
);

CREATE TABLE task_members (
    taskID INTEGER REFERENCES tasks(taskID),
    userID INTEGER
   -- userID INTEGER REFERENCES users(userID)
);

CREATE TABLE task_tags (
    taskID INTEGER REFERENCES tasks(taskID),
    tagID INTEGER REFERENCES tags(tagID)
);

CREATE TABLE checklists (
    taskID INTEGER REFERENCES tasks(taskID),
    name TEXT,
    items JSONB
);

CREATE TABLE attachments (
    taskID INTEGER REFERENCES tasks(taskID),
    filename TEXT,
    filepath TEXT
);