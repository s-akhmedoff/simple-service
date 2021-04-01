package models

import "time"

type uuid string

type workingHours struct {
	OpenAt  *time.Time
	CloseAt *time.Time
}
