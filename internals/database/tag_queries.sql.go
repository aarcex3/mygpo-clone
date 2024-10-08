// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: tag_queries.sql

package database

import (
	"context"
)

const listTopTags = `-- name: ListTopTags :many
SELECT
    title,
    code,
    usage
FROM
    tags
ORDER BY
    usage DESC
LIMIT
    ?
`

type ListTopTagsRow struct {
	Title string
	Code  string
	Usage int64
}

func (q *Queries) ListTopTags(ctx context.Context, limit int64) ([]ListTopTagsRow, error) {
	rows, err := q.db.QueryContext(ctx, listTopTags, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListTopTagsRow
	for rows.Next() {
		var i ListTopTagsRow
		if err := rows.Scan(&i.Title, &i.Code, &i.Usage); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
