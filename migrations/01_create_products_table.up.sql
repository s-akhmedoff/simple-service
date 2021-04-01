CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "products" (
    "id" 	       uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    "sku" 	       varchar,
    "name" 	       varchar,
    "type" 	       varchar,
    "uri" 	       varchar,
    "description"  varchar,
    "is_active"    bool,
    "created_at"   timestamp,
    "updated_at"   timestamp
);