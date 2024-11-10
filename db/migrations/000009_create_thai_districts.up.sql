CREATE TABLE IF NOT EXISTS "thai_districts" (
  id UUID NOT NULL PRIMARY KEY,
  name_th VARCHAR(150) NOT NULL,
  province_id UUID NOT NULL,
  FOREIGN KEY (province_id) REFERENCES thai_provinces(id)
);