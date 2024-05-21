package null

import "database/sql"

func NullFromPtrString(param *string) sql.NullString {
	var res sql.NullString
	if param != nil {
		res.String = *param
		res.Valid = true
		return res
	}
	res.Valid = false
	res.String = ""
	return res
}

func NullFromPtrStringOpt(param *string) sql.NullString {
	var res sql.NullString
	res.Valid = true
	if param != nil {
		res.String = *param
		return res
	}
	res.String = "shortest"
	return res
}

func NullFromPtrInt(param *int64) sql.NullInt64 {
	var res sql.NullInt64
	res.Valid = true
	if param != nil {
		res.Int64 = *param
		return res
	}
	res.Int64 = 0
	return res
}
