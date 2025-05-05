-- name: CreateFeedFollow :one
INSERT INTO feed_follows (id, feed_id, user_id)
VALUES ($1, $2, $3)
RETURNING *;


-- name: GetFeedFollows :many
SELECT * FROM feed_follows
WHERE user_id = $1;
