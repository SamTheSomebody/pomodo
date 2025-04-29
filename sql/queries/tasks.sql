-- name: CreateTask :one
INSERT INTO tasks (id, created_at, updated_at, name, description, due_at, time_estimate_seconds, priority, enthusiasm)
VALUES (?, NOW(), NOW(), ?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: GetTasks :many
SELECT * FROM tasks;

-- name: GetTaskByName :one
SELECT * FROM tasks WHERE name = ?;

-- name: GetTaskByID :one
SELECT * FROM tasks WHERE id = ?;

-- name: DeleteTask :exec
DELETE FROM tasks WHERE id = ?;

-- name: UpdateDueAt :one
UPDATE tasks SET due_at = ? WHERE id = ? RETURNING *;

-- name: AddTimeSpent :one
UPDATE tasks SET time_spent_seconds = time_spent_seconds + ? WHERE id = ? RETURNING *;

