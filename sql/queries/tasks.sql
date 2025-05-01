-- name: CreateTask :one
INSERT INTO tasks (id, created_at, updated_at, name, summary, due_at, time_estimate_seconds, priority, enthusiasm)
VALUES (?, datetime('now'), datetime('now'), ?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: UpdateTask :one
UPDATE tasks SET updated_at = datetime('now'), name = ?, summary = ?, due_at = ?, time_estimate_seconds = ?, priority = ?, enthusiasm = ?
WHERE id = ? RETURNING *;

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

-- name: CompleteTask :one
UPDATE tasks SET is_complete = true WHERE id = ? RETURNING *;

