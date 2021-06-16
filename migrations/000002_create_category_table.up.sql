CREATE TABLE IF NOT EXISTS "category" (
    "id" int UNIQUE NOT NULL,
    "name" varchar UNIQUE NOT NULL,
    "created_at" timestamp DEFAULT now(),
    "updated_at" timestamp DEFAULT now(),
    "uri_name" varchar NOT NULL,
    "product_id" uuid NOT NULL,
    FOREIGN KEY("product_id") REFERENCES product("id") ON DELETE CASCADE ,
    PRIMARY KEY("id", "product_id")
);