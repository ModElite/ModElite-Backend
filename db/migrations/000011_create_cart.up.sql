CREATE TABLE IF NOT EXISTS "cart" (
  user_id UUID NOT NULL,
  product_size_id UUID NOT NULL,
  quantity INT NOT NULL DEFAULT 1,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (user_id, product_size_id),
  FOREIGN KEY (user_id) REFERENCES "users"(id),
  FOREIGN KEY (product_size_id) REFERENCES "product_size"(id)
);