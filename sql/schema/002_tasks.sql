-- +goose Up
CREATE TABLE tasks (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  name TEXT NOT NULL,
  description TEXT,
  due_at TIMESTAMP,
  time_estimate_seconds INTEGER,
  time_spent_seconds INTEGER DEFAULT 0,
  priority INTEGER,
  enthusiasm INTEGER,
  user_id UUID REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE tasks;
