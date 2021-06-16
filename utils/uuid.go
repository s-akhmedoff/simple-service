package utils

import (
	"github.com/google/uuid"
)

//GenUUID - ...
func GenUUID() string {
	data, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}

	return data.String()
}
