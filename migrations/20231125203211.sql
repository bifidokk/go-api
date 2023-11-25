-- Modify "notes" table
ALTER TABLE "notes" ADD COLUMN "user_id" bigint NULL, ADD
 CONSTRAINT "fk_users_notes" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
