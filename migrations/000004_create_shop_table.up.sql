CREATE TABLE IF NOT EXISTS "shop" (
    "id" int UNIQUE NOT NULL,
    "name" varchar NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT now(),
    "updated_at" timestamp NOT NULL DEFAULT now(),
    "address" varchar NOT NULL,
    "lan" varchar,
    "lat" varchar,
    "working_hours" varchar,
    "product_id" uuid NOT NULL,
    FOREIGN KEY("product_id") REFERENCES product("id") ON DELETE CASCADE,
    PRIMARY KEY ("id", "product_id")
);

/*id 	int 	да 	да
name 	varchar 	да 	нет
create_at 	timestamp 	да 	нет
updated_at 	timestamp 	да 	нет
address 	varchar 	да 	нет
lon 	varchar 	нет 	нет
lat 	varchar 	нет 	нет
working_hours 	varchar 	нет 	нет
 */