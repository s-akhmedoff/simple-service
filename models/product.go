package models

import "time"

type Product struct {
	ID          string
	SKI         string
	Name        string
	ProductType string
	URI         string
	Description string
	IsActive    bool
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}

type NewProduct struct {
	SKI         string
	Name        string
	ProductType string
	URI         string
	Description string
	IsActive    bool
}
