CREATE TABLE IF NOT EXISTS users (
    id                    serial         NOT NULL PRIMARY KEY,
    login                 varchar(32)    UNIQUE NOT NULL,
    password              varchar(64),
    name                  varchar(16)    NOT NULL,
    surname               varchar(32)    NOT NULL,
    );

CREATE TABLE IF NOT EXISTS chats(
    id serial NOT NULL PRIMARY KEY,
    creator_id integer,
    users_id integer[],
    name varchar(32),
)