-- +goose Up
CREATE TABLE organization_members (
  id UUID PRIMARY KEY,
  id_organization UUID NOT NULL REFERENCES organization(id),
  id_user UUID NOT NULL REFERENCES users(id)
);

-- +goose Down
DROP TABLE IF EXISTS organization_members;
