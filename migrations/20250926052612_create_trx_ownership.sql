-- +goose Up
CREATE TABLE trx_ownership (
  id UUID PRIMARY KEY,
  id_organization UUID NOT NULL REFERENCES organization(id),
  id_scanner UUID NOT NULL REFERENCES scanner(id)
);

-- +goose Down
DROP TABLE IF EXISTS trx_ownership;
