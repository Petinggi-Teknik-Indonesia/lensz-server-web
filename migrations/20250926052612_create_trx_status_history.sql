-- +goose Up
CREATE TABLE trx_status_history (
  id UUID PRIMARY KEY,
  id_glasses UUID NOT NULL REFERENCES glasses(id),
  id_user UUID NOT NULL REFERENCES users(id),
  id_scanner UUID NOT NULL REFERENCES scanner(id),
  status_change glasses_status NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE IF EXISTS trx_status_history;
