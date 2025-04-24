-- name: CreateTask :one
INSERT INTO tasks (id, created_at, updated_at, name, description, due_at, time_estimate_seconds, priority, enthusiasm, user_id)
VAULES ($1, NOW(), NOW(), $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetTaskByName :one
SELECT * FROM tasks WHERE name = $1;

-- name: GetTaskByID :one
SELECT * FROM tasks WHERE id = $1;

-- name: DeleteTask :exec
DELETE FROM tasks WHERE id = $1;

-- name: UpdateDueAt :one
UPDATE tasks SET due_at = $2 WHERE id = $1 RETURNING *;

-- name: AddTimeSpent: one
UPDATE tasks SET time_spent_seconds = time_spent_seconds + $2 WHERE id = $1 RETURNING *;

