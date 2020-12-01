DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS social_links;

CREATE TABLE users (
    userID      SERIAL PRIMARY KEY,
    email       TEXT,
    password    BYTEA,
    username    TEXT,
    fullname    TEXT,
    avatar      TEXT
);

/*CREATE TABLE social_links (
    userID      INTEGER REFERENCES users(userID) ON DELETE CASCADE,
    networkName TEXT,
    link TEXT
);
*/