-- name: GetEpisodeByUrl :one
SELECT
    title,
    url,
    podcast_title,
    podcast_url,
    description,
    website,
    released,
    mygpo_link
FROM
    episodes
where
    url = ?;

-- name: GetPodcastByUrl :one
SELECT
    title,
    website,
    mygpo_link,
    description,
    subscribers,
    author,
    url,
    logo_url
FROM
    podcasts
where
    url = ?;

-- name: ListTopTags :many
SELECT
    title,
    code,
    usage
FROM
    tags
ORDER BY
    usage DESC
LIMIT
    ?;

-- name: GetUserById :one
SELECT
    *
FROM
    users
WHERE
    id = ?
LIMIT
    1;

-- name: GetUserByUsername :one
SELECT
    *
FROM
    users
WHERE
    username = ?
LIMIT
    1;

-- name: CreateUser :exec
INSERT INTO
    users (username, password, email)
values
    (?, ?, ?);

-- name: UserExists :one
SELECT
    count(*)
FROM
    users
WHERE
    username = ?
    OR email = ?;