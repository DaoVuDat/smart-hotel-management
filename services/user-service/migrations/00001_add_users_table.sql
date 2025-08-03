-- +goose Up
CREATE TYPE user_role AS ENUM ('guest', 'staff', 'manager', 'admin');

CREATE TABLE users (
   id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
   email VARCHAR(255) UNIQUE NOT NULL,
   name VARCHAR(255) NOT NULL,
   phone VARCHAR(50),
   role user_role NOT NULL DEFAULT 'guest',
   created_at TIMESTAMP DEFAULT NOW(),
   updated_at TIMESTAMP DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS users;

DROP TYPE IF EXISTS user_role;
