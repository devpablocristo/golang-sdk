package sdkmysql

import (
	"context"
	"database/sql"
	"fmt"
	"sync"

	"github.com/devpablocristo/sdk/golang/sdk/databases/sql/mysql/go-sql-driver/ports"
	_ "github.com/go-sql-driver/mysql"
)

var (
	instance  ports.Repository
	once      sync.Once
	initError error
)

type Repository struct {
	db *sql.DB
}

// newRepository crea una nueva instancia de Repository con configuración proporcionada.
func newRepository(c config) (ports.Repository, error) {
	once.Do(func() {
		client := &Repository{}
		initError = client.connect(c)
		if initError != nil {
			instance = nil
		} else {
			instance = client
		}
	})
	return instance, initError
}

// connect establece la conexión a la base de datos MySQL.
func (r *Repository) connect(c config) error {
	dsn := c.dsn()
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("failed to connect to MySQL: %w", err)
	}
	if err := conn.Ping(); err != nil {
		return fmt.Errorf("failed to ping MySQL: %w", err)
	}
	r.db = conn
	return nil
}

// Ping verifica la conexión a la base de datos.
func (r *Repository) Ping(ctx context.Context) error {
	return r.db.PingContext(ctx)
}

// Close cierra la conexión a la base de datos.
func (r *Repository) Close() {
	if r.db != nil {
		r.db.Close()
	}
}

// DB devuelve la instancia *sql.DB.
func (r *Repository) DB() *sql.DB {
	return r.db
}
