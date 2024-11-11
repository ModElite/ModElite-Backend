CREATE TABLE IF NOT EXISTS "tag_group" (
  id SERIAL PRIMARY KEY,
  label VARCHAR(255) NOT NULL,
  "show" BOOLEAN NOT NULL DEFAULT FALSE
);
CREATE TABLE IF NOT EXISTS "tag" (
  id SERIAL PRIMARY KEY,
  tag_group_id SERIAL NOT NULL,
  label VARCHAR(255) NOT NULL,
  FOREIGN KEY (tags_group_id) REFERENCES tag_group(id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS "product_tag" (
  tag_id SERIAL NOT NULL,
  product_id UUID NOT NULL,
  PRIMARY KEY (tag_id, product_id),
  FOREIGN KEY (tag_id) REFERENCES tag(id) ON DELETE CASCADE,
  FOREIGN KEY (product_id) REFERENCES product(id) ON DELETE CASCADE
);