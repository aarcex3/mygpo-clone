-- Insert a single user
INSERT INTO
    users (username, password, email)
VALUES
    (
        'testuser',
        '$2a$10$cCg4I596LKbghcph7pdFN.JoJQ5sEKCHgEUP/njhrcHhQ2x8yg.Cu',
        'test@mail.com'
    );

-- Insert multiple tags
INSERT INTO
    tags (title, code, usage)
VALUES
    ('Technology', 'technology', 530),
    ('Science', 'science', 410),
    ('Health', 'health', 325),
    ('Education', 'education', 275),
    ('Finance', 'finance', 600),
    ('Sports', 'sports', 475),
    ('Travel', 'travel', 200),
    ('Food', 'food', 150),
    ('Art', 'art', 100),
    ('History', 'history', 50);