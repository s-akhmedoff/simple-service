package models

import "time"

type Price struct {
	ID            UUID
	SalePrice     int
	FactoryPrice  int
	DiscountPrice int
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
	IsActive      bool
	ProductID     UUID
}

type NewPrice struct {
	SalePrice     int
	FactoryPrice  int
	DiscountPrice int
	IsActive      bool
	ProductID     UUID
}
