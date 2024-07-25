-- Create the Device table
CREATE TABLE IF NOT EXISTS Device (
    id TEXT PRIMARY KEY,
    caption TEXT NOT NULL UNIQUE,
    device_type TEXT NOT NULL,
    owner_id TEXT,
    FOREIGN KEY (owner_id) REFERENCES User(id)
);
-- Create the Episode table
CREATE TABLE IF NOT EXISTS Episode (
    id TEXT PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    audio_url TEXT NOT NULL,
    duration INTEGER NOT NULL,
    podcast_id TEXT,
    released_on TEXT NOT NULL,
    FOREIGN KEY (podcast_id) REFERENCES Podcast(id)
);
-- Create the ListeningHistory table
CREATE TABLE IF NOT EXISTS ListeningHistory (
    id TEXT PRIMARY KEY,
    user_id TEXT,
    episode_id TEXT,
    progress REAL NOT NULL,
    date TEXT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES User(id),
    FOREIGN KEY (episode_id) REFERENCES Episode(id)
);
-- Create the Podcast table
CREATE TABLE IF NOT EXISTS Podcast (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    description TEXT NOT NULL,
    website TEXT NOT NULL,
    xml_url TEXT NOT NULL,
    author TEXT NOT NULL
);
-- Create the PodcastList table
CREATE TABLE IF NOT EXISTS PodcastList (
    user_list_id TEXT NOT NULL,
    podcast_id TEXT NOT NULL,
    PRIMARY KEY (user_list_id, podcast_id),
    FOREIGN KEY (user_list_id) REFERENCES UserList(id),
    FOREIGN KEY (podcast_id) REFERENCES Podcast(id)
);
-- Create the PodcastTag table
CREATE TABLE IF NOT EXISTS PodcastTag (
    tag_id TEXT NOT NULL,
    podcast_id TEXT NOT NULL,
    PRIMARY KEY (tag_id, podcast_id),
    FOREIGN KEY (tag_id) REFERENCES Tag(id),
    FOREIGN KEY (podcast_id) REFERENCES Podcast(id)
);
-- Create the Subscription table
CREATE TABLE IF NOT EXISTS Subscription (
    device_id TEXT NOT NULL,
    podcast_id TEXT NOT NULL,
    PRIMARY KEY (device_id, podcast_id),
    FOREIGN KEY (device_id) REFERENCES Device(id),
    FOREIGN KEY (podcast_id) REFERENCES Podcast(id)
);
-- Create the Tag table
CREATE TABLE IF NOT EXISTS Tag (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    code TEXT NOT NULL UNIQUE,
    usage INTEGER NOT NULL
);
-- Create the UserList table
CREATE TABLE IF NOT EXISTS UserList (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    title TEXT NOT NULL,
    web TEXT NOT NULL,
    owner_id TEXT,
    FOREIGN KEY (owner_id) REFERENCES User(id)
);
-- Create the User table
CREATE TABLE IF NOT EXISTS User (
    id TEXT PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    email TEXT NOT NULL,
    UNIQUE (username)
);
-- Insert fake data into User table
INSERT INTO User (id, username, password, email)
VALUES (
        '1d9f1c10-9e34-11e9-b475-0800273b35d5',
        'user1',
        '$2b$12$.DTVBitY7gba6JFeLcQhm.pNUmeJZJQqcX25qH4xpcolzFhLomBUu',
        'user1@example.com'
    ),
    (
        '1d9f1c11-9e34-11e9-b475-0800273b35d5',
        'user2',
        '$2b$12$x7uJEQOr1z5yBMr7ofKVp.FNN18.ux84PLICSEzXOVtMz.q97CY/q',
        'user2@example.com'
    );
-- Insert fake data into Device table
INSERT INTO Device (id, caption, device_type, owner_id)
VALUES (
        '2d9f1c10-9e34-11e9-b475-0800273b35d5',
        'User1 Phone',
        'MOBILE',
        '1d9f1c10-9e34-11e9-b475-0800273b35d5'
    ),
    (
        '2d9f1c11-9e34-11e9-b475-0800273b35d5',
        'User2 Laptop',
        'DESKTOP',
        '1d9f1c11-9e34-11e9-b475-0800273b35d5'
    );
-- Insert fake data into Podcast table
INSERT INTO Podcast (id, name, description, website, xml_url, author)
VALUES (
        '3d9f1c10-9e34-11e9-b475-0800273b35d5',
        'Tech Talk',
        'A podcast about tech',
        'http://techtalk.com',
        'http://techtalk.com/rss',
        'Tech Talk Team'
    ),
    (
        '3d9f1c11-9e34-11e9-b475-0800273b35d5',
        'History Buffs',
        'A podcast about history',
        'http://historybuffs.com',
        'http://historybuffs.com/rss',
        'History Buffs Team'
    ),
    (
        '331bda0c-49e0-11ef-9454-0242ac120002',
        'History land',
        'A podcast about history of the countries',
        'http://historyland.com',
        'http://historyland.com/rss',
        'History Land Team'
    );
-- Insert fake data into Episode table
INSERT INTO Episode (
        id,
        title,
        description,
        audio_url,
        duration,
        podcast_id,
        released_on
    )
VALUES (
        '4d9f1c10-9e34-11e9-b475-0800273b35d5',
        'Episode 1',
        'First episode of Tech Talk',
        'http://techtalk.com/ep1.mp3',
        3600,
        '3d9f1c10-9e34-11e9-b475-0800273b35d5',
        '2023-01-01T10:00:00'
    ),
    (
        '4d9f1c11-9e34-11e9-b475-0800273b35d5',
        'Episode 2',
        'First episode of History Buffs',
        'http://historybuffs.com/ep1.mp3',
        4200,
        '3d9f1c11-9e34-11e9-b475-0800273b35d5',
        '2023-01-01T11:00:00'
    );
-- Insert fake data into ListeningHistory table
INSERT INTO ListeningHistory (id, user_id, episode_id, progress, date)
VALUES (
        '5d9f1c10-9e34-11e9-b475-0800273b35d5',
        '1d9f1c10-9e34-11e9-b475-0800273b35d5',
        '4d9f1c10-9e34-11e9-b475-0800273b35d5',
        50.0,
        '2023-01-01T12:00:00'
    ),
    (
        '5d9f1c11-9e34-11e9-b475-0800273b35d5',
        '1d9f1c11-9e34-11e9-b475-0800273b35d5',
        '4d9f1c11-9e34-11e9-b475-0800273b35d5',
        75.0,
        '2023-01-01T13:00:00'
    );
-- Insert fake data into Tag table
INSERT INTO Tag (id, name, code, usage)
VALUES (
        '6d9f1c10-9e34-11e9-b475-0800273b35d5',
        'Technology',
        'technology',
        5
    ),
    (
        '6d9f1c11-9e34-11e9-b475-0800273b35d5',
        'History',
        'history',
        10
    );
-- Insert fake data into PodcastTag table
INSERT INTO PodcastTag (tag_id, podcast_id)
VALUES (
        '6d9f1c10-9e34-11e9-b475-0800273b35d5',
        '3d9f1c10-9e34-11e9-b475-0800273b35d5'
    ),
    (
        '6d9f1c11-9e34-11e9-b475-0800273b35d5',
        '3d9f1c11-9e34-11e9-b475-0800273b35d5'
    ),
    (
        '6d9f1c11-9e34-11e9-b475-0800273b35d5',
        '331bda0c-49e0-11ef-9454-0242ac120002'
    );
-- Insert fake data into UserList table
INSERT INTO UserList (id, name, title, web, owner_id)
VALUES (
        '7d9f1c10-9e34-11e9-b475-0800273b35d5',
        'User1 Favorites',
        'Favorites',
        'http://favorites.com',
        '1d9f1c10-9e34-11e9-b475-0800273b35d5'
    ),
    (
        '7d9f1c11-9e34-11e9-b475-0800273b35d5',
        'User2 Must Listen',
        'Must Listen',
        'http://mustlisten.com',
        '1d9f1c11-9e34-11e9-b475-0800273b35d5'
    );
-- Insert fake data into PodcastList table
INSERT INTO PodcastList (user_list_id, podcast_id)
VALUES (
        '7d9f1c10-9e34-11e9-b475-0800273b35d5',
        '3d9f1c10-9e34-11e9-b475-0800273b35d5'
    ),
    (
        '7d9f1c11-9e34-11e9-b475-0800273b35d5',
        '3d9f1c11-9e34-11e9-b475-0800273b35d5'
    );
-- Insert fake data into Subscription table
INSERT INTO Subscription (device_id, podcast_id)
VALUES (
        '2d9f1c10-9e34-11e9-b475-0800273b35d5',
        '3d9f1c10-9e34-11e9-b475-0800273b35d5'
    ),
    (
        '2d9f1c11-9e34-11e9-b475-0800273b35d5',
        '3d9f1c11-9e34-11e9-b475-0800273b35d5'
    );