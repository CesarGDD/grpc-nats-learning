-- name: GetBlogs :many
SELECT * FROM blogs
ORDER BY id;

-- name: GetBlog :one
SELECT * FROM blogs
WHERE id = $1;

-- name: DeleteBlog :one
DELETE FROM blogs
WHERE id = $1
RETURNING *;

-- name: CreateBlog :one
INSERT INTO blogs (username, content)
VALUES ($1, $2)
RETURNING *;