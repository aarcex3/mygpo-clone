-- Create the person table
CREATE TABLE IF NOT EXISTS person (
    id UUID PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    is_active BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    sa_orm_sentinel BIGINT
);

-- Create the device table
CREATE TABLE IF NOT EXISTS device (
    id UUID PRIMARY KEY,
    caption TEXT NOT NULL UNIQUE,
    device_type TEXT NOT NULL,
    owner_id UUID,
    FOREIGN KEY (owner_id) REFERENCES person(id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    sa_orm_sentinel BIGINT
);

-- Create the podcast table
CREATE TABLE IF NOT EXISTS podcast (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    description TEXT NOT NULL,
    website TEXT NOT NULL,
    xml_url TEXT NOT NULL,
    author TEXT NOT NULL,
    subscribers_count INTEGER NOT NULL,
    created_at TIMESTAMPTZ  DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ  DEFAULT CURRENT_TIMESTAMP,
    sa_orm_sentinel BIGINT
);

-- Create the episode table
CREATE TABLE IF NOT EXISTS episode (
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    audio_url TEXT NOT NULL,
    duration INTEGER NOT NULL,
    podcast_id UUID,
    released_on TIMESTAMPTZ NOT NULL,
    FOREIGN KEY (podcast_id) REFERENCES podcast(id),
    created_at TIMESTAMPTZ  DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ  DEFAULT CURRENT_TIMESTAMP,
    sa_orm_sentinel BIGINT
);

-- Create the listening_history table
CREATE TABLE IF NOT EXISTS listening_history (
    id UUID PRIMARY KEY,
    person_id UUID,
    episode_id UUID,
    progress REAL NOT NULL,
    date TIMESTAMPTZ NOT NULL,
    FOREIGN KEY (person_id) REFERENCES person(id),
    FOREIGN KEY (episode_id) REFERENCES episode(id),
    created_at TIMESTAMPTZ  DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ  DEFAULT CURRENT_TIMESTAMP,
    sa_orm_sentinel BIGINT
);

-- Create the tag table
CREATE TABLE IF NOT EXISTS tag (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    code TEXT NOT NULL UNIQUE,
    usage INTEGER NOT NULL,
    created_at TIMESTAMPTZ  DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ  DEFAULT CURRENT_TIMESTAMP,
    sa_orm_sentinel BIGINT
);

-- Create the person_list table
CREATE TABLE IF NOT EXISTS person_list (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    title TEXT NOT NULL,
    web TEXT NOT NULL,
    owner_id UUID,
    FOREIGN KEY (owner_id) REFERENCES person(id),
    created_at TIMESTAMPTZ  DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ  DEFAULT CURRENT_TIMESTAMP,
    sa_orm_sentinel BIGINT
);

-- Create the podcast_list table
CREATE TABLE IF NOT EXISTS podcast_list (
    person_list_id UUID NOT NULL,
    podcast_id UUID NOT NULL,
    PRIMARY KEY (person_list_id, podcast_id),
    FOREIGN KEY (person_list_id) REFERENCES person_list(id),
    FOREIGN KEY (podcast_id) REFERENCES podcast(id),
    created_at TIMESTAMPTZ  DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ  DEFAULT CURRENT_TIMESTAMP,
    sa_orm_sentinel BIGINT
);

-- Create the podcast_tag table
CREATE TABLE IF NOT EXISTS podcast_tag (
    tag_id UUID NOT NULL,
    podcast_id UUID NOT NULL,
    PRIMARY KEY (tag_id, podcast_id),
    FOREIGN KEY (tag_id) REFERENCES tag(id),
    FOREIGN KEY (podcast_id) REFERENCES podcast(id),
    created_at TIMESTAMPTZ  DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ  DEFAULT CURRENT_TIMESTAMP,
    sa_orm_sentinel BIGINT
);

-- Create the subscription table
CREATE TABLE IF NOT EXISTS subscription (
    device_id UUID NOT NULL,
    podcast_id UUID NOT NULL,
    PRIMARY KEY (device_id, podcast_id),
    FOREIGN KEY (device_id) REFERENCES device(id),
    FOREIGN KEY (podcast_id) REFERENCES podcast(id),
    created_at TIMESTAMPTZ  DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ  DEFAULT CURRENT_TIMESTAMP,
    sa_orm_sentinel BIGINT
);

-- Insert fake data into person table
INSERT INTO person (id, username, password, email, sa_orm_sentinel)
VALUES
    ('550e8400-e29b-41d4-a716-446655440000', 'user1', '$2b$12$.DTVBitY7gba6JFeLcQhm.pNUmeJZJQqcX25qH4xpcolzFhLomBUu', 'user1@example.com', 100),
    ('550e8400-e29b-41d4-a716-446655440001', 'user2', '$2b$12$x7uJEQOr1z5yBMr7ofKVp.FNN18.ux84PLICSEzXOVtMz.q97CY/q', 'user2@example.com', 100);

-- Insert fake data into device table
INSERT INTO device (id, caption, device_type, owner_id, sa_orm_sentinel)
VALUES
    ('550e8400-e29b-41d4-a716-446655440002', 'User1 Phone', 'MOBILE', '550e8400-e29b-41d4-a716-446655440000', 100),
    ('550e8400-e29b-41d4-a716-446655440003', 'User2 Laptop', 'DESKTOP', '550e8400-e29b-41d4-a716-446655440001', 100);

-- Insert fake data into podcast table
INSERT INTO podcast (id, name, description, website, xml_url, author, subscribers_count, sa_orm_sentinel)
VALUES
    ('550e8400-e29b-41d4-a716-446655440004', 'Tech Talk', 'A podcast about tech', 'http://techtalk.com', 'http://techtalk.com/rss', 'Tech Talk Team', 12340, 100),
    ('550e8400-e29b-41d4-a716-446655440005', 'History Buffs', 'A podcast about history', 'http://historybuffs.com', 'http://historybuffs.com/rss', 'History Buffs Team', 34, 100),
    ('550e8400-e29b-41d4-a716-446655440006', 'History Land', 'A podcast about history of the countries', 'http://historyland.com', 'http://historyland.com/rss', 'History Land Team', 1234, 100);

-- Insert fake data into episode table
INSERT INTO episode (id, title, description, audio_url, duration, podcast_id, released_on, sa_orm_sentinel)
VALUES
    ('550e8400-e29b-41d4-a716-446655440007', 'Episode 1', 'First episode of Tech Talk', 'http://techtalk.com/ep1.mp3', 3600, '550e8400-e29b-41d4-a716-446655440004', '2023-01-01T10:00:00Z', 100),
    ('550e8400-e29b-41d4-a716-446655440008', 'Episode 2', 'First episode of History Buffs', 'http://historybuffs.com/ep1.mp3', 4200, '550e8400-e29b-41d4-a716-446655440005', '2023-01-01T11:00:00Z', 100);

-- Insert fake data into listening_history table
INSERT INTO listening_history (id, person_id, episode_id, progress, date, sa_orm_sentinel)
VALUES
    ('550e8400-e29b-41d4-a716-446655440009', '550e8400-e29b-41d4-a716-446655440000', '550e8400-e29b-41d4-a716-446655440007', 50.0, '2023-01-01T12:00:00Z', 100),
    ('550e8400-e29b-41d4-a716-446655440010', '550e8400-e29b-41d4-a716-446655440001', '550e8400-e29b-41d4-a716-446655440008', 75.0, '2023-01-01T13:00:00Z', 100);

-- Insert fake data into tag table
INSERT INTO tag (id, name, code, usage, sa_orm_sentinel)
VALUES
    ('550e8400-e29b-41d4-a716-446655440011', 'Technology', 'technology', 5, 100),
    ('550e8400-e29b-41d4-a716-446655440012', 'History', 'history', 10, 100);

-- Insert fake data into podcast_tag table
INSERT INTO podcast_tag (tag_id, podcast_id, sa_orm_sentinel)
VALUES
    ('550e8400-e29b-41d4-a716-446655440011', '550e8400-e29b-41d4-a716-446655440004', 100),
    ('550e8400-e29b-41d4-a716-446655440012', '550e8400-e29b-41d4-a716-446655440005', 100),
    ('550e8400-e29b-41d4-a716-446655440012', '550e8400-e29b-41d4-a716-446655440006', 100);

-- Insert fake data into person_list table
INSERT INTO person_list (id, name, title, web, owner_id, sa_orm_sentinel)
VALUES
    ('550e8400-e29b-41d4-a716-446655440013', 'favorites', 'Favorites', 'http://favorites.com', '550e8400-e29b-41d4-a716-446655440000', 100),
    ('550e8400-e29b-41d4-a716-446655440014', 'must-listen', 'Must Listen', 'http://mustlisten.com', '550e8400-e29b-41d4-a716-446655440001', 100);

-- Insert fake data into podcast_list table
INSERT INTO podcast_list (person_list_id, podcast_id, sa_orm_sentinel)
VALUES
    ('550e8400-e29b-41d4-a716-446655440013', '550e8400-e29b-41d4-a716-446655440004', 100),
    ('550e8400-e29b-41d4-a716-446655440014', '550e8400-e29b-41d4-a716-446655440005', 100);

-- Insert fake data into subscription table
INSERT INTO subscription (device_id, podcast_id, sa_orm_sentinel)
VALUES
    ('550e8400-e29b-41d4-a716-446655440002', '550e8400-e29b-41d4-a716-446655440004', 100),
    ('550e8400-e29b-41d4-a716-446655440003', '550e8400-e29b-41d4-a716-446655440005', 100);
