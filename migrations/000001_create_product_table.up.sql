CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "product" (
    "id" 	       uuid NOT NULL PRIMARY KEY,
    "sku" 	       varchar NOT NULL,
    "name" 	       varchar NOT NULL,
    "type" 	       varchar NOT NULL,
    "uri" 	       varchar NOT NULL,
    "description"  varchar NOT NULL,
    "is_active"    bool NOT NULL,
    "created_at"   timestamp DEFAULT now(),
    "updated_at"   timestamp DEFAULT now()
);