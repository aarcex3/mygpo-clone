// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: podcast_queries.sql

package database

import (
	"context"
	"database/sql"
)

const getPodcastByUrl = `-- name: GetPodcastByUrl :one
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
    url = ?
`

type GetPodcastByUrlRow struct {
	Title       string
	Website     sql.NullString
	MygpoLink   sql.NullString
	Description string
	Subscribers int64
	Author      string
	Url         string
	LogoUrl     sql.NullString
}

func (q *Queries) GetPodcastByUrl(ctx context.Context, url string) (GetPodcastByUrlRow, error) {
	row := q.db.QueryRowContext(ctx, getPodcastByUrl, url)
	var i GetPodcastByUrlRow
	err := row.Scan(
		&i.Title,
		&i.Website,
		&i.MygpoLink,
		&i.Description,
		&i.Subscribers,
		&i.Author,
		&i.Url,
		&i.LogoUrl,
	)
	return i, err
}
