-- name: CreatePost :one
INSERT INTO posts (id, title, description, published_at, url, feed_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;
