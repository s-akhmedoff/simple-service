package models

import "time"

type Product struct {
	ID          UUID       `db:"id"`
	SKU         string     `db:"sku"`
	Name        string     `db:"name"`
	ProductType string     `db:"type"`
	URI         string     `db:"uri"`
	Description string     `db:"description"`
	IsActive    bool       `db:"is_active"`
	CreatedAt   *time.Time `db:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
}

type NewProduct struct {
	SKI         string
	Name        string
	ProductType string
	URI         string
	Description string
	IsActive    bool
}
