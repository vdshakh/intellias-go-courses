CREATE TABLE "users" (
    "id" varchar PRIMARY KEY,
    "full_name" varchar NOT NULL,
    "email" varchar NOT NULL,
    "password" varchar NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE INDEX ON "users" ("email");