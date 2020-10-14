package models

import "database/sql"

// Repository struct for db connection
type Repository struct {
	Conn *sql.DB
}
