CREATE TABLE IF NOT EXISTS "thai_subdistricts" (
  id UUID NOT NULL PRIMARY KEY,
  zip_code INT NOT NULL,
  name_th VARCHAR(150) NOT NULL,
  district_id UUID NOT NULL,
  FOREIGN KEY (district_id) REFERENCES thai_districts(id)
);