CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "user" (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  username TEXT NOT NULL UNIQUE CHECK (char_length(username) > 0 AND char_length(username) <= 64),
  email TEXT NOT NULL UNIQUE CHECK (char_length(email) > 0 AND char_length(email) <= 64),
  password_hash TEXT NOT NULL,
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT NOW()
);
