package models

import "time"

type UUID string

type workingHours struct {
	OpenAt  *time.Time
	CloseAt *time.Time
}
