package postgres

import (
	"github.com/jackc/pgx/v5/pgconn"
)

// postgres errors define follow below link
// https://github.com/lib/pq/blob/master/error.go#L78C4-L78C4
const (
	UniqueViolation  key = "23505"
	NotNullViolation key = "23502"
)

type key string

func ErrorIs(targetErr error, postgresErrKey key) bool {
	if err, ok := targetErr.(*pgconn.PgError); ok {
		if err.Code == string(postgresErrKey) {
			return true
		}
	}
	return false
}
