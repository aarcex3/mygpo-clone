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