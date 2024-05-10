package ptr

import "database/sql"

func PtrFromString(nullString sql.NullString) *string {
	if nullString.Valid {
		return &nullString.String
	}
	return nil
}
