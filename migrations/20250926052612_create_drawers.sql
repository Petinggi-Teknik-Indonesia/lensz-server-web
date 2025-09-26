-- +goose Up
CREATE TABLE drawers (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS drawers;
