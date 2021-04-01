package models

import "time"

type Shop struct {
	ID           uuid
	Name         string
	CreatedAT    *time.Time
	UpdatedAt    *time.Time
	Address      string
	Longitude    string
	Latitude     string
	WorkingHours workingHours
}

type NewShop struct {
	Name         string
	Address      string
	Longitude    string
	Latitude     string
	WorkingHours workingHours
}
