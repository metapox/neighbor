package writer

import (
	"database/sql"
	"fmt"
)

type PostgresWriter struct {
	db *sql.DB
	config Config
}
