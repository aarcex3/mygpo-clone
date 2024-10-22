CREATE TABLE
    users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL,
        email TEXT NOT NULL UNIQUE
    );

CREATE TABLE
    tags (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        code TEXT NOT NULL UNIQUE,
        usage INTEGER NOT NULL DEFAULT 0
    );

CREATE TABLE
    podcasts (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        website TEXT NOT NULL,
        mygpo_link TEXT NOT NULL UNIQUE,
        description TEXT NOT NULL,
        subscribers INTEGER NOT NULL DEFAULT 0,
        title TEXT NOT NULL UNIQUE,
        author TEXT NOT NULL,
        url TEXT NOT NULL UNIQUE,
        logo_url TEXT DEFAULT ""
    );

CREATE TABLE
    episodes (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        url TEXT NOT NULL,
        podcast_title TEXT NOT NULL,
        podcast_url TEXT NOT NULL,
        description TEXT NOT NULL,
        website TEXT NOT NULL,
        released DATETIME NOT NULL,
        mygpo_link TEXT NOT NULL
    );