-- +goose Up 
CREATE TABLE users (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  name TEXT NOT NULL,
  time_estimate_buffer INTERVAL NOT NULL,
  allocated_work_time INTERVAL -- possible issue, if another value is updated, might skip daily allocation check
);

-- +goose Down
DROP TABLE users;
