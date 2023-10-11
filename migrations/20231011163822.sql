-- Create "notes" table
CREATE TABLE "notes" (
  "id" bigserial NOT NULL,
  "note_title" character varying(255) NULL,
  "note_description" text NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
