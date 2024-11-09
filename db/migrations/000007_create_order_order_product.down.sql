DROP TABLE IF EXISTS "order_product";
DROP TABLE IF EXISTS "voucher";
DROP TABLE IF EXISTS "order";
DO $$ BEGIN DROP TYPE IF EXISTS order_status;
EXCEPTION
WHEN undefined_object THEN NULL;
END;
$$;
DO $$ BEGIN DROP TYPE IF EXISTS order_product_status;
EXCEPTION
WHEN undefined_object THEN NULL;
END;
$$;