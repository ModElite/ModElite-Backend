DROP TABLE IF EXISTS "session";
DROP TABLE IF EXISTS "user";
DO $$ BEGIN DROP TYPE IF EXISTS user_role;
EXCEPTION
WHEN undefined_object THEN NULL;
END;
$$;