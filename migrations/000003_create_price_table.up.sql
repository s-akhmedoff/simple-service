CREATE TABLE IF NOT EXISTS "price" (
    "id" int PRIMARY KEY NOT NULL,
    "sale_price" int NOT NULL,
    "factory_price" int NOT NULL,
    "discount_price" int NOT NULL,
    "created_at" timestamp DEFAULT now(),
    "updated_at" timestamp DEFAULT now(),
    "is_active" bool NOT NULL,
    "product_id" uuid NOT NULL,
    FOREIGN KEY("product_id") REFERENCES product("id") ON DELETE CASCADE
);