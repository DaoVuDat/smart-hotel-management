-- +goose Up
ALTER TABLE users
ADD COLUMN date_of_birth DATE;

-- +goose Down
ALTER TABLE users
DROP COLUMN date_of_birth;


