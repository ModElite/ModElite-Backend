CREATE TABLE IF NOT EXISTS "address" (
  id SERIAL PRIMARY KEY,
  user_id UUID NOT NULL,
  first_name VARCHAR(255) NOT NULL,
  last_name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  phone VARCHAR(255) NOT NULL,
  label VARCHAR(255) NOT NULL,
  "default" BOOLEAN DEFAULT FALSE NOT NULL,
  "address" VARCHAR(255) NOT NULL,
  sub_district VARCHAR(255) NOT NULL,
  district VARCHAR(255) NOT NULL,
  province VARCHAR(255) NOT NULL,
  zip_code VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_user_id_address FOREIGN KEY (user_id) REFERENCES "users"(id) ON DELETE CASCADE
);