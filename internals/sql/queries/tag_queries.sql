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