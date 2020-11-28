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
DROP TABLE IF EXISTS attachments;

CREATE TABLE comments (
    commentID SERIAL PRIMARY KEY,
    message TEXT,
    taskID INTEGER REFERENCES tasks(taskID)  ON DELETE CASCADE,
    commentOrder INTEGER,
    userID INTEGER
);

-- ID  ДОСКИ!!!!
CREATE TABLE tags (
    tagID SERIAL PRIMARY KEY,
    boardID  INTEGER REFERENCES boards(boardID)  ON DELETE CASCADE,
    name TEXT,
    color TEXT
);

CREATE TABLE task_members (
    taskID INTEGER REFERENCES tasks(taskID) ON DELETE CASCADE,
    userID INTEGER
);

CREATE TABLE task_tags (
    taskID INTEGER REFERENCES tasks(taskID) ON DELETE CASCADE,
    tagID INTEGER REFERENCES tags(tagID) ON DELETE CASCADE
);
ALTER TABLE IF EXISTS task_tags ADD CONSTRAINT uniq UNIQUE (taskID, tagID);

CREATE TABLE checklists (
    checklistID SERIAL PRIMARY KEY,
    taskID INTEGER REFERENCES tasks(taskID) ON DELETE CASCADE,
    name TEXT,
    items JSONB
);

CREATE TABLE attachments (
    attachmentID SERIAL PRIMARY KEY,
    taskID INTEGER REFERENCES tasks(taskID) ON DELETE CASCADE,
    filename TEXT,
    filepath TEXT
);