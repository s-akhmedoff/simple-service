package postgres

import "errors"

var (
	errNilPointerReference = errors.New("[SQL Storage]: argument got nil pointer reference")
	errEmptyArgument       = errors.New("[SQL Storage]: empty argument has been passed")
	errNoRowsAffected      = errors.New("[SQL Storage]: no rows was affected")
)
