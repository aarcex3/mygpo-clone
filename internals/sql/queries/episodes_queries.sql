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