-- +goose Up
CREATE TABLE scanner (
  id UUID PRIMARY KEY,
  name TEXT NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS scanner;
