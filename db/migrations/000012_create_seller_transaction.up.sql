CREATE TABLE IF NOT EXISTS "seller_transaction" (
  id UUID NOT NULL PRIMARY KEY,
  seller_id UUID NOT NULL,
  bank_account_name VARCHAR(255) NOT NULL,
  bank_account_number VARCHAR(255) NOT NULL,
  bank_account_provider VARCHAR(255) NOT NULL,
  bank_transaction_id VARCHAR(255) NOT NULL,
  bank_transaction_amount DOUBLE PRECISION NOT NULL,
  bank_transaction_fee DOUBLE PRECISION NOT NULL,
  bank_transaction_datetime TIMESTAMP NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (seller_id) REFERENCES "seller"(id)
);
CREATE TABLE IF NOT EXISTS "seller_transaction_order" (
  seller_transaction_id UUID NOT NULL,
  order_id UUID NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (seller_transaction_id, order_id),
  FOREIGN KEY (seller_transaction_id) REFERENCES "seller_transaction"(id),
  FOREIGN KEY (order_id) REFERENCES "order"(id)
);