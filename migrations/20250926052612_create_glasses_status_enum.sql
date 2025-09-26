-- +goose Up
CREATE TYPE glasses_status AS ENUM (
  'Tersedia',
  'Terjual',
  'Rusak',
  'Terpinjam',
  'Lainnya'
);

-- +goose Down
DROP TYPE IF EXISTS glasses_status;
