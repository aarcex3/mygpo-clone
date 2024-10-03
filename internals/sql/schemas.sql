CREATE TABLE
    users (
        id BIGSERIAL PRIMARY KEY,
        username text NOT NULL,
        password text NOT NULL,
        email text NOT NULL
    );