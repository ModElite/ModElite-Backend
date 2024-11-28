CREATE TABLE IF NOT EXISTS "seller" (
  id UUID PRIMARY KEY,
  "name" VARCHAR(256) NOT NULL,
  "description" TEXT NOT NULL,
  logo_url TEXT NOT NULL,
  "location" TEXT NOT NULL,
  bank_account_name TEXT NOT NULL,
  bank_account_number TEXT NOT NULL,
  bank_account_provider TEXT NOT NULL,
  phone VARCHAR(16) NOT NULL,
  owner_id UUID NOT NULL,
  is_verify BOOLEAN NOT NULL DEFAULT FALSE,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_seller_owner FOREIGN KEY (owner_id) REFERENCES "users"(id)
);