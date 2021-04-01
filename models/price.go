package models

import "time"

type Price struct {
	ID            string
	SalePrice     int
	FactoryPrice  int
	DiscountPrice int
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
	IsActive      bool
	ProductID     int
}

type NewPrice struct {
	SalePrice     int
	FactoryPrice  int
	DiscountPrice int
	IsActive      bool
	ProductID     string
}
