-- +goose Up
CREATE TABLE companies (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS companies;
