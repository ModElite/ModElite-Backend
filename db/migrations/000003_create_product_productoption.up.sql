DO $$ BEGIN BEGIN CREATE TYPE product_status AS ENUM ('ACTIVE', 'INACTIVE', 'DELETE');
EXCEPTION
WHEN duplicate_object THEN NULL;
END;
END;
$$;
CREATE TABLE IF NOT EXISTS "product" (
  id UUID PRIMARY KEY,
  seller_id UUID NOT NULL,
  name VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,
  feature TEXT NOT NULL,
  price DOUBLE PRECISION NOT NULL,
  "status" product_status NOT NULL DEFAULT 'ACTIVE',
  image_url TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP,
  CONSTRAINT fk_seller_product FOREIGN KEY (seller_id) REFERENCES "seller"(id)
);
CREATE TABLE IF NOT EXISTS "product_option" (
  id UUID PRIMARY KEY,
  product_id UUID NOT NULL,
  label VARCHAR(255) NOT NULL,
  image_url TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP,
  CONSTRAINT fk_seller_product_option FOREIGN KEY (product_id) REFERENCES "product"(id)
);
CREATE TABLE IF NOT EXISTS "size" (
  id UUID PRIMARY KEY,
  size VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS "product_size" (
  id UUID PRIMARY KEY,
  product_option_id UUID NOT NULL,
  size_id UUID NOT NULL,
  quantity INT NOT NULL DEFAULT 0,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP,
  CONSTRAINT fk_product_option_product_size FOREIGN KEY (product_option_id) REFERENCES "product_option"(id),
  CONSTRAINT fk_size_product_size FOREIGN KEY (size_id) REFERENCES "size"(id)
);