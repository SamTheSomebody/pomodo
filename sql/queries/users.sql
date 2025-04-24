-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES ( $1, NOW(), NOW(), $2)
RETURNING *;

-- name: GetUserByName :one
SELECT * FROM users WHERE name = $1;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;

-- name: UpdateAllocatedTime :one
UPDATE users SET allocated_work_time = $2 WHERE id = $1 RETURNING *;

-- name: UpdateTimeEstimateBuffer :one
UPDATE users SET time_estimate_buffer = $2 WHERE id = $1 RETURNING *;
