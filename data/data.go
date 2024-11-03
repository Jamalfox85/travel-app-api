package data

import "database/sql"

func NewNullInt32(value int32) sql.NullInt32 {
	return sql.NullInt32 {
		Int32: value,
		Valid: value != 0,
	}
}