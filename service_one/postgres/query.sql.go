// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: query.sql

package postgres

import (
	"context"
)

const createBlog = `-- name: CreateBlog :one
INSERT INTO blogs (username, content)
VALUES ($1, $2)
RETURNING id, username, content
`

type CreateBlogParams struct {
	Username string
	Content  string
}

func (q *Queries) CreateBlog(ctx context.Context, arg CreateBlogParams) (Blog, error) {
	row := q.db.QueryRowContext(ctx, createBlog, arg.Username, arg.Content)
	var i Blog
	err := row.Scan(&i.Id, &i.Username, &i.Content)
	return i, err
}

const deleteBlog = `-- name: DeleteBlog :one
DELETE FROM blogs
WHERE id = $1
RETURNING id, username, content
`

func (q *Queries) DeleteBlog(ctx context.Context, id int32) (Blog, error) {
	row := q.db.QueryRowContext(ctx, deleteBlog, id)
	var i Blog
	err := row.Scan(&i.Id, &i.Username, &i.Content)
	return i, err
}

const getBlog = `-- name: GetBlog :one
SELECT id, username, content FROM blogs
WHERE id = $1
`

func (q *Queries) GetBlog(ctx context.Context, id int32) (Blog, error) {
	row := q.db.QueryRowContext(ctx, getBlog, id)
	var i Blog
	err := row.Scan(&i.Id, &i.Username, &i.Content)
	return i, err
}

const getBlogs = `-- name: GetBlogs :many
SELECT id, username, content FROM blogs
ORDER BY id
`

func (q *Queries) GetBlogs(ctx context.Context) ([]Blog, error) {
	rows, err := q.db.QueryContext(ctx, getBlogs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Blog
	for rows.Next() {
		var i Blog
		if err := rows.Scan(&i.Id, &i.Username, &i.Content); err != nil {
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