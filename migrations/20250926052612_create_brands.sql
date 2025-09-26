-- +goose Up
CREATE TABLE brands (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS brands;
