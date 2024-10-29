DROP TABLE IF EXISTS "product_option";
DROP TABLE IF EXISTS "product";
DO $$ BEGIN DROP TYPE IF EXISTS product_status;
EXCEPTION
WHEN undefined_object THEN NULL;
END;
$$;