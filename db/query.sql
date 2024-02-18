-- name: DeleteCollection :exec
DELETE FROM collections
WHERE id = $1;

-- name: GetCollection :one
SELECT * FROM collections
WHERE id = $1 LIMIT 1;