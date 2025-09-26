-- +goose Up
CREATE TABLE glasses (
  id UUID PRIMARY KEY,
  type TEXT NOT NULL,
  color TEXT NOT NULL,
  status glasses_status NOT NULL,
  description TEXT NOT NULL,
  id_drawer INT REFERENCES drawers(id),
  id_brand INT REFERENCES brands(id),
  id_company INT REFERENCES companies(id),
  id_organization UUID NOT NULL REFERENCES organization(id),
  created_at TIMESTAMP NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE IF EXISTS glasses;
