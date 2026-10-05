CREATE TYPE "social_enums" AS ENUM (
  'GOOGLE',
  'KAKAO'
);

CREATE TYPE "one_depth_enums" AS ENUM (
  'ALL',
  'NORMAL',
  'DAILY',
  'WORKOUT',
  'BUSINESS',
  'TRAVEL',
  'FOOD'
);

CREATE TYPE  "access_enums" AS ENUM (
  'ACCESS',
  'DENIED'
);

CREATE TYPE  "access_role_enums" AS ENUM (
  'READ',
  'EDITABLE'
);

CREATE TABLE IF NOT EXISTS "users" (
  "id" uuid PRIMARY KEY,
  "created_at" TIMESTAMPTZ,
  "updated_at" TIMESTAMPTZ,
  "created_by" TEXT,
  "updated_by" TEXT,
  "email" TEXT NOT NULL,
  "social" social_enums NOT NULL,
  "provider_id" TEXT NOT NULL,
  "name" TEXT NOT NULL,
  "refresh_token" TEXT
);

CREATE TABLE IF NOT EXISTS "routines" (
  "id" uuid PRIMARY KEY,
  "created_at" TIMESTAMPTZ,
  "updated_at" TIMESTAMPTZ,
  "created_by" TEXT,
  "updated_by" TEXT,
  "routine_title" TEXT NOT NULL,
  "routine_dest" TEXT,
  "priorityOneDepth" integer,
  "category" one_depth_enums,
  "user_id" uuid
);

CREATE TABLE IF NOT EXISTS "routines_detail" (
  "id" uuid PRIMARY KEY,
  "created_at" TIMESTAMPTZ,
  "updated_at" TIMESTAMPTZ,
  "created_by" TEXT,
  "updated_by" TEXT,
  "priorityTwoDepth" integer,
  "pre_event_start_at" TIMESTAMPTZ NOT NULL,
  "pre_event_end_at" TIMESTAMPTZ NOT NULL,
  "start_at" TIMESTAMPTZ NOT NULL,
  "end_at" TIMESTAMPTZ NOT NULL,
  "routine_id" uuid
);

CREATE TABLE IF NOT EXISTS "routines_detail_desc" (
  "id" uuid PRIMARY KEY,
  "created_at" TIMESTAMPTZ,
  "updated_at" TIMESTAMPTZ,
  "created_by" TEXT,
  "updated_by" TEXT,
  "desc" TEXT NOT NULL,
  "routine_detail_id" uuid
);

CREATE TABLE IF NOT EXISTS "routine_detail_tags" (
  "id" uuid PRIMARY KEY,
  "created_at" TIMESTAMPTZ,
  "updated_at" TIMESTAMPTZ,
  "created_by" TEXT,
  "updated_by" TEXT,
  "tag" TEXT NOT NULL,
  "routine_detail_id" uuid
);

CREATE TABLE IF NOT EXISTS "routine_access" (
  "id" uuid PRIMARY KEY,
  "created_at" TIMESTAMPTZ,
  "updated_at" TIMESTAMPTZ,
  "created_by" TEXT,
  "updated_by" TEXT,
  "can_access" access_enums,
  "editable" access_role_enums,
  "routine_id" uuid,
  "access_user_id" uuid
);

ALTER TABLE "routines" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "routines_detail" ADD FOREIGN KEY ("routine_id") REFERENCES "routines" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "routines_detail_desc" ADD FOREIGN KEY ("routine_detail_id") REFERENCES "routines_detail" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "routine_detail_tags" ADD FOREIGN KEY ("routine_detail_id") REFERENCES "routines_detail" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "routine_access" ADD FOREIGN KEY ("routine_id") REFERENCES "routines" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "routine_access" ADD FOREIGN KEY ("access_user_id") REFERENCES "users" ("id") DEFERRABLE INITIALLY IMMEDIATE;
