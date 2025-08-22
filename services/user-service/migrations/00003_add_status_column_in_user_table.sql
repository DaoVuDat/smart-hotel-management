-- +goose Up
CREATE TYPE user_status AS ENUM ('active', 'deleted', 'banned', 'inactive');

ALTER TABLE users
ADD COLUMN status user_status NOT NULL DEFAULT 'inactive';

-- +goose Down
ALTER TABLE users
DROP COLUMN status;

DROP TYPE IF EXISTS user_status;
