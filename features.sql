DROP TABLE IF EXISTS comments;
DROP TABLE IF EXISTS checklists;
DROP TABLE IF EXISTS task_members;
DROP TABLE IF EXISTS task_tags;
DROP TABLE IF EXISTS tags;
DROP TABLE IF EXISTS attaches;

CREATE TABLE comments (
    commentID SERIAL PRIMARY KEY,
    message TEXT,
    taskID INTEGER REFERENCES tasks(taskID),
    commentOrder INTEGER,
    userID INTEGER
);

CREATE TABLE tags (
    tagID SERIAL PRIMARY KEY,
    color TEXT,
    title TEXT
);

CREATE TABLE task_members (
    taskID INTEGER REFERENCES tasks(taskID),
    userID INTEGER REFERENCES users(userID)
);

CREATE TABLE task_tags (
    taskID INTEGER REFERENCES tasks(taskID),
    tagID INTEGER REFERENCES tags(tagID)
);

CREATE TABLE checklists (
    checklistID SERIAL PRIMARY KEY,
    taskID INTEGER REFERENCES tasks(taskID),
    name TEXT,
    items JSONB
);

CREATE TABLE attachments (
    taskID INTEGER REFERENCES tasks(taskID),
    attachmentID SERIAL PRIMARY KEY,
    filename TEXT,
    filepath TEXT
);