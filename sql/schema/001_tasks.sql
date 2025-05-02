-- +goose Up
CREATE TABLE tasks (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  name TEXT NOT NULL,
  summary TEXT NOT NULL,
  due_at TIMESTAMP NOT NULL,
  time_estimate_seconds INTEGER NOT NULL,
  time_spent_seconds INTEGER DEFAULT 0 NOT NULL,
  priority INTEGER DEFAULT 0 NOT NULL,
  enthusiasm INTEGER DEFAULT 0 NOT NULL,
  is_complete BOOLEAN DEFAULT false NOT NULL
  -- user_id UUID REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE tasks;
