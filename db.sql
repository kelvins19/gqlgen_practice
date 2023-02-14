CREATE TABLE "categories" (
    "id" SERIAL PRIMARY KEY,
    "name" text NOT NULL,
    "description" text,
    "created_at" timestamptz DEFAULT (now()),
    "updated_at" timestamptz DEFAULT (now())
);
INSERT INTO categories (name,description) VALUES ('Jackets','Jackets');
INSERT INTO categories (name,description) VALUES ('Shoes','Shoes');
INSERT INTO categories (name,description) VALUES ('Sneakers','Sneakers');
INSERT INTO categories (name,description) VALUES ('Socks','Socks');

CREATE TABLE "products" (
    "id" SERIAL PRIMARY KEY,
    "name" text NOT NULL,
    "description" text,
    "categories" int[],
    "price" numeric,
    "created_at" timestamptz DEFAULT (now()),
    "updated_at" timestamptz DEFAULT (now())
);