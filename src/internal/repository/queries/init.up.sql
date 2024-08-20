CREATE TABLE IF NOT EXISTS users (
    id                    serial         NOT NULL PRIMARY KEY,
    login                 varchar(32)    UNIQUE NOT NULL,
    password              varchar(64),
    email                 varchar(32)    UNIQUE NOT NULL,
    name                  varchar(16)    NOT NULL,
    surname               varchar(32)    NOT NULL,

    chat_lists            text           default '{"Favorites": [], "Friends": [], "Other": []}',
    raw_chats             int []         default ARRAY[]::int[] []
    );

CREATE TABLE IF NOT EXISTS chats(
    id serial NOT NULL PRIMARY KEY,
    creator_id integer,
    users_id integer[],
    messages_id integer[],
    name varchar(32),
)