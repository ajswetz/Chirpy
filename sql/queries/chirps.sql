-- name: CreateChirp :one
INSERT INTO chirps (id, created_at, updated_at, body, user_id)
VALUES (
        gen_random_uuid(),
        NOW(),
        NOW(),
        $1,
        $2
    )
RETURNING *;
-- name: GetSingleChirp :one
SELECT *
from chirps
WHERE id = $1;
-- name: GetAllChirpsAsc :many
SELECT *
FROM chirps
ORDER BY created_at ASC;
-- name: GetAllChirpsDesc :many
SELECT *
FROM chirps
ORDER BY created_at DESC;
-- name: GetChirpsForGivenAuthorAsc :many
SELECT *
FROM chirps
WHERE user_id = $1
ORDER BY created_at ASC;
-- name: GetChirpsForGivenAuthorDesc :many
SELECT *
FROM chirps
WHERE user_id = $1
ORDER BY created_at DESC;
-- name: DeleteChirp :exec
DELETE FROM chirps
WHERE id = $1;