-- create database hybrid


CREATE TABLE "users" (
  "id" BIGSERIAL PRIMARY KEY,
  "email" varchar(255) NOT NULL,
  "password" varchar(255) NOT NULL,
  "role_user" boolean NOT NULL,
  "branch_id" bigint,
  "department_id" bigint,
  "position_id" bigint,
  "image_id" bigint,
  "full_name" varchar(150),
  "phone" char(20),
  "birthday" varchar(50),
  "description" varchar(255),
  "facebook" varchar(200),
  "create_by" bigint,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "categories" (
  "id" BIGSERIAL PRIMARY KEY,
  "menu_id" bigint NOT NULL,
  "category_name" varchar(225) NOT NULL,
  "parent_name" varchar(225),
  "description" varchar(225) NOT NULL,
  "parent_id" bigint,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "posts" (
  "id" BIGSERIAL PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "title" varchar(255) NOT NULL,
  "description" varchar(255) NOT NULL,
  "category_id" bigint NOT NULL, --(Sub)
  "image_id" bigint NOT NULL,
  "content" text NOT NULL,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);


CREATE TABLE "branches" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" varchar(150) NOT NULL,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "departments" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" varchar(150) NOT NULL,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "images" (
  "id" BIGSERIAL PRIMARY KEY,
  "url" text NOT NULL,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "menus" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" varchar(150) NOT NULL,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "positions" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);


create table reset_passwords(
  "id" BIGSERIAL PRIMARY KEY,
  "user_id" BIGINT NOT NULL,
  "token" varchar(255) NOT NULL,
  "time" varchar(6) NOT NULL,
  "expired_at" timestamp,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

ALTER TABLE "users" ADD FOREIGN KEY ("branch_id") REFERENCES "branches" ("id");

ALTER TABLE "users" ADD FOREIGN KEY ("department_id") REFERENCES "departments" ("id");

ALTER TABLE "users" ADD FOREIGN KEY ("position_id") REFERENCES "positions" ("id");

ALTER TABLE "users" ADD FOREIGN KEY ("image_id") REFERENCES "images" ("id");

ALTER TABLE "categories" ADD FOREIGN KEY ("menu_id") REFERENCES "menus" ("id");

ALTER TABLE "posts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "posts" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "posts" ADD FOREIGN KEY ("image_id") REFERENCES "images" ("id");

ALTER TABLE "posts"  ALTER COLUMN image_id TYPE bigint,  ALTER COLUMN image_id SET NOT NULL


--IF ERROR RUN TWO ROW
-- ALTER TABLE public.users ALTER COLUMN full_name SET NOT NULL
-- ALTER TABLE categories ADD parent_name varchar(225)

