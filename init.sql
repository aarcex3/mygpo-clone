-- Create the Device table
CREATE TABLE IF NOT EXISTS Device (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    caption TEXT NOT NULL UNIQUE,
    device_type TEXT NOT NULL,
    owner_id INTEGER,
    FOREIGN KEY (owner_id) REFERENCES User(id)
);

-- Create the Episode table
CREATE TABLE IF NOT EXISTS Episode (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    audio_url TEXT NOT NULL,
    duration INTEGER NOT NULL,
    podcast_id INTEGER,
    released_on TEXT NOT NULL,
    FOREIGN KEY (podcast_id) REFERENCES Podcast(id)
);

-- Create the ListeningHistory table
CREATE TABLE IF NOT EXISTS ListeningHistory (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER,
    episode_id INTEGER,
    progress REAL NOT NULL,
    date TEXT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES User(id),
    FOREIGN KEY (episode_id) REFERENCES Episode(id)
);

-- Create the Podcast table
CREATE TABLE IF NOT EXISTS Podcast (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE,
    description TEXT NOT NULL,
    website TEXT NOT NULL,
    xml_url TEXT NOT NULL,
    author TEXT NOT NULL,
    subscribers_count INTEGER NOT NULL
);

-- Create the PodcastList table
CREATE TABLE IF NOT EXISTS PodcastList (
    user_list_id INTEGER NOT NULL,
    podcast_id INTEGER NOT NULL,
    PRIMARY KEY (user_list_id, podcast_id),
    FOREIGN KEY (user_list_id) REFERENCES UserList(id),
    FOREIGN KEY (podcast_id) REFERENCES Podcast(id)
);

-- Create the PodcastTag table
CREATE TABLE IF NOT EXISTS PodcastTag (
    tag_id INTEGER NOT NULL,
    podcast_id INTEGER NOT NULL,
    PRIMARY KEY (tag_id, podcast_id),
    FOREIGN KEY (tag_id) REFERENCES Tag(id),
    FOREIGN KEY (podcast_id) REFERENCES Podcast(id)
);

-- Create the Subscription table
CREATE TABLE IF NOT EXISTS Subscription (
    device_id INTEGER NOT NULL,
    podcast_id INTEGER NOT NULL,
    PRIMARY KEY (device_id, podcast_id),
    FOREIGN KEY (device_id) REFERENCES Device(id),
    FOREIGN KEY (podcast_id) REFERENCES Podcast(id)
);

-- Create the Tag table
CREATE TABLE IF NOT EXISTS Tag (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE,
    code TEXT NOT NULL UNIQUE,
    usage INTEGER NOT NULL
);

-- Create the UserList table
CREATE TABLE IF NOT EXISTS UserList (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE,
    title TEXT NOT NULL,
    web TEXT NOT NULL,
    owner_id INTEGER,
    FOREIGN KEY (owner_id) REFERENCES User(id)
);

-- Create the User table
CREATE TABLE IF NOT EXISTS User (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    UNIQUE (username)
);

-- Insert fake data into User table
INSERT INTO User (username, password, email)
VALUES
    ('user1', '$2b$12$.DTVBitY7gba6JFeLcQhm.pNUmeJZJQqcX25qH4xpcolzFhLomBUu', 'user1@example.com'),
    ('user2', '$2b$12$x7uJEQOr1z5yBMr7ofKVp.FNN18.ux84PLICSEzXOVtMz.q97CY/q', 'user2@example.com');

-- Insert fake data into Device table
INSERT INTO Device (caption, device_type, owner_id)
VALUES
    ('User1 Phone', 'MOBILE', 1),
    ('User2 Laptop', 'DESKTOP', 2);

-- Insert fake data into Podcast table
INSERT INTO Podcast (name, description, website, xml_url, author, subscribers_count)
VALUES
    ('Tech Talk', 'A podcast about tech', 'http://techtalk.com', 'http://techtalk.com/rss', 'Tech Talk Team', 12340),
    ('History Buffs', 'A podcast about history', 'http://historybuffs.com', 'http://historybuffs.com/rss', 'History Buffs Team', 34),
    ('History Land', 'A podcast about history of the countries', 'http://historyland.com', 'http://historyland.com/rss', 'History Land Team', 1234);

-- Insert fake data into Episode table
INSERT INTO Episode (title, description, audio_url, duration, podcast_id, released_on)
VALUES
    ('Episode 1', 'First episode of Tech Talk', 'http://techtalk.com/ep1.mp3', 3600, 1, '2023-01-01T10:00:00'),
    ('Episode 2', 'First episode of History Buffs', 'http://historybuffs.com/ep1.mp3', 4200, 2, '2023-01-01T11:00:00');

-- Insert fake data into ListeningHistory table
INSERT INTO ListeningHistory (user_id, episode_id, progress, date)
VALUES
    (1, 1, 50.0, '2023-01-01T12:00:00'),
    (2, 2, 75.0, '2023-01-01T13:00:00');

-- Insert fake data into Tag table
INSERT INTO Tag (name, code, usage)
VALUES
    ('Technology', 'technology', 5),
    ('History', 'history', 10);

-- Insert fake data into PodcastTag table
INSERT INTO PodcastTag (tag_id, podcast_id)
VALUES
    (1, 1),
    (2, 2),
    (2, 3);

-- Insert fake data into UserList table
INSERT INTO UserList (name, title, web, owner_id)
VALUES
    ('favorites', 'Favorites', 'http://favorites.com', 1),
    ('must-listen', 'Must Listen', 'http://mustlisten.com', 2);

-- Insert fake data into PodcastList table
INSERT INTO PodcastList (user_list_id, podcast_id)
VALUES
    (1, 1),
    (2, 2);

-- Insert fake data into Subscription table
INSERT INTO Subscription (device_id, podcast_id)
VALUES
    (1, 1),
    (2, 2),
    (1, 2);
