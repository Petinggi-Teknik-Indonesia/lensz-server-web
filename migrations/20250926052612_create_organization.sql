-- +goose Up
CREATE TABLE organization (
  id UUID PRIMARY KEY,
  name TEXT NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS organization;
