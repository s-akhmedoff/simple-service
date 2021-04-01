package models

import "time"

type Price struct {
	ID            uuid
	SalePrice     int
	FactoryPrice  int
	DiscountPrice int
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
	IsActive      bool
	ProductID     uuid
}

type NewPrice struct {
	SalePrice     int
	FactoryPrice  int
	DiscountPrice int
	IsActive      bool
	ProductID     uuid
}
