package models

import "time"

type Category struct {
	ID        uuid
	Name      string
	CreateAat *time.Time
	UpdatedAt *time.Time
	UriName   string
}

type NewCategory struct {
	Name    string
	UriName string
}
