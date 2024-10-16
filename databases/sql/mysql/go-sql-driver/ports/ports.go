package ports

import "database/sql"

type Repository interface {
	Close()
	DB() *sql.DB
}
