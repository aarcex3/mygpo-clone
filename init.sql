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

-- Insert multiple podcasts
INSERT INTO
    podcasts (
        website,
        mygpo_link,
        description,
        subscribers,
        title,
        author,
        url,
        logo_url
    )
VALUES
    (
        'http://coverville.com',
        'http://www.gpodder.net/podcast/16124',
        'The best cover songs, delivered to your ears two to three times a week!',
        19,
        'Coverville',
        'Brian Ibbott',
        'http://feeds.feedburner.com/coverville',
        'http://www.coverville.com/art/coverville_iTunes300.jpg'
    ),
    (
        'http://freakonomics.com',
        'http://www.gpodder.net/podcast/23456',
        'Discover the hidden side of everything with Stephen J. Dubner, co-author of the Freakonomics books.',
        15000,
        'Freakonomics Radio',
        'Stephen J. Dubner',
        'http://feeds.feedburner.com/freakonomics',
        'http://freakonomics.com/images/logo.png'
    ),
    (
        'http://serialpodcast.org',
        'http://www.gpodder.net/podcast/34567',
        'Serial is a podcast from the creators of This American Life, hosted by Sarah Koenig.',
        30000,
        'Serial',
        'Sarah Koenig',
        'http://feeds.serialpodcast.org/serialpodcast',
        'http://serialpodcast.org/images/logo.jpg'
    ),
    (
        'http://radiolab.org',
        'http://www.gpodder.net/podcast/45678',
        'Radiolab is a show about curiosity, where sound illuminates ideas and the boundaries blur between science, philosophy, and human experience.',
        25000,
        'Radiolab',
        'Jad Abumrad and Robert Krulwich',
        'http://feeds.radiolab.org/radiolab',
        'http://radiolab.org/images/logo.png'
    ),
    (
        'http://thetimferrissshow.com',
        'http://www.gpodder.net/podcast/56789',
        'The Tim Ferriss Show is the first business/interview podcast to exceed 100 million downloads.',
        45000,
        'The Tim Ferriss Show',
        'Tim Ferriss',
        'http://feeds.feedburner.com/timferriss',
        'http://timferriss.com/images/logo.png'
    );

-- Insert multiple podcast episodes
INSERT INTO
    episodes (
        title,
        url,
        podcast_title,
        podcast_url,
        description,
        website,
        released,
        mygpo_link
    )
VALUES
    (
        'TWiT 245: No Hitler For You',
        'http://www.podtrac.com/pts/redirect.mp3/aolradio.podcast.aol.com/twit/twit0245.mp3',
        'this WEEK in TECH - MP3 Edition',
        'http://leo.am/podcasts/twit',
        'A roundtable discussion about the latest trends in technology.',
        'http://www.podtrac.com/pts/redirect.mp3/aolradio.podcast.aol.com/twit/twit0245.mp3',
        '2010-12-25T00:30:00',
        'http://gpodder.net/episode/1046492'
    ),
    (
        'Coverville 123: Best of Beatles Covers',
        'http://www.coverville.com/audio/Coverville123.mp3',
        'Coverville',
        'http://feeds.feedburner.com/coverville',
        'An hour of some of the best Beatles covers from across the years.',
        'http://coverville.com',
        '2011-01-05T12:00:00',
        'http://gpodder.net/episode/54321'
    ),
    (
        'Freakonomics Episode 156: The Upside of Quitting',
        'http://freakonomics.com/audio/freakonomics156.mp3',
        'Freakonomics Radio',
        'http://feeds.feedburner.com/freakonomics',
        'Why quitting can sometimes be the best thing to do.',
        'http://freakonomics.com',
        '2012-03-15T14:00:00',
        'http://gpodder.net/episode/65432'
    ),
    (
        'Serial Episode 9: To Be Suspected',
        'http://serialpodcast.org/audio/serial09.mp3',
        'Serial',
        'http://feeds.serialpodcast.org/serialpodcast',
        'In this episode, we delve deeper into the investigation of the crime.',
        'http://serialpodcast.org',
        '2014-11-06T09:30:00',
        'http://gpodder.net/episode/76543'
    ),
    (
        'Radiolab: The Trust Engineers',
        'http://radiolab.org/audio/radiolab-trust.mp3',
        'Radiolab',
        'http://feeds.radiolab.org/radiolab',
        'Exploring the engineers who build systems of trust in the digital age.',
        'http://radiolab.org',
        '2015-02-19T16:00:00',
        'http://gpodder.net/episode/87654'
    );