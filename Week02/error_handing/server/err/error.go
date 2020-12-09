package err

import "errors"

var (
	ErrNoRowsSql = errors.New("sql: no rows in result set")
)
