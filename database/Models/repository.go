package Models

import "database/sql"

// Repository struct for db connection
type Repository struct {
	Conn *sql.DB
}

type ErrSql struct {
	code    uint
	message string
}
