-- name: CreateFeedFollow :one
INSERT INTO feed_follows (id, feed_id, user_id)
VALUES ($1, $2, $3)
RETURNING *;
