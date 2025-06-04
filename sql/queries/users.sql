-- name: CreateUser :one
INSERT INTO users(id, created_at, updated_at)
VALUES (?, datetime('now'), datetime('now'))
RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE id = ?;

-- name: GetFirstUser :one
SELECT * FROM users ORDER BY updated_at;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = ?;

-- name: SetUserWeights :one
UPDATE users SET updated_at = datetime('now'), priority_weight = ?, enthusiasm_weight = ?, first_task_weight = ?, due_date_daily_weight = ?
WHERE id = ? RETURNING *;

-- name: SetUserAllocatedTime :one
UPDATE users SET updated_at = datetime('now'), allocated_time_seconds = ?
WHERE id = ? RETURNING *;

-- name: SetUserTimeEstimateBuffer :one
UPDATE users SET updated_at = datetime('now'), time_estimate_buffer_percent = ?
WHERE id = ? RETURNING *;
