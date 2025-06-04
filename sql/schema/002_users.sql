-- +goose Up 
CREATE TABLE users (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  allocated_time_seconds INTEGER DEFAULT 27900 NOT NULL,
  time_estimate_buffer_percent FLOAT DEFAULT 0.2 NOT NULL,
  priority_weight FLOAT DEFAULT 0.3 NOT NULl,
  enthusiasm_weight FLOAT DEFAULT 0.2 NOT NULL,
  first_task_weight FLOAT DEFAULT 0.4 NOT NULL,
  due_date_daily_weight FLOAT DEFAULT 0.2 NOT NULL
);

-- +goose Down
DROP TABLE users;
