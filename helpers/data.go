package helpers

import "database/sql"

func ValidateRange(value int) sql.NullInt64 {
	if value == 0 {
		return sql.NullInt64{
			Valid: false,
		}
	}
	if value > 10 {
		value = 10
	} else if value < 1 {
		value = 1
	}
	return sql.NullInt64{
		Valid: true,
		Int64: int64(value),
	}
}
