package writer

import (
	"database/sql"
)

type PostgresWriter struct {
	db     *sql.DB
	config Config
}
