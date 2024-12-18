package entities

import (
	"database/sql"
	"strings"
)

func NullString(uuid string) sql.NullString {
	return sql.NullString{String: uuid, Valid: true}
}

func IsEmptyString(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}
