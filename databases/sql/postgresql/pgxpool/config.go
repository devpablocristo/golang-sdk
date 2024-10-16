package sdkpostgresql

import (
	"fmt"

	"github.com/devpablocristo/golang/sdk/pkg/databases/sql/postgresql/pgxpool/ports"
)

type config struct {
	Host     string
	User     string
	Password string
	DbName   string
	Port     string
}

// newConfig crea una nueva configuración con los valores proporcionados
func newConfig(user, password, host, port, dbName string) ports.Config {
	return &config{
		Host:     host,
		User:     user,
		Password: password,
		DbName:   dbName,
		Port:     port,
	}
}

// DNS genera la cadena de conexión para PostgreSQL
func (c *config) DNS() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", c.User, c.Password, c.Host, c.Port, c.DbName)
}

// Validate valida que los campos necesarios estén presentes
func (c *config) Validate() error {
	if c.User == "" {
		return fmt.Errorf("POSTGRES_USER is required")
	}
	if c.Password == "" {
		return fmt.Errorf("POSTGRES_PASSWORD is required")
	}
	if c.Host == "" {
		return fmt.Errorf("POSTGRES_HOST is required")
	}
	if c.Port == "" {
		return fmt.Errorf("POSTGRES_PORT is required")
	}
	if c.DbName == "" {
		return fmt.Errorf("POSTGRES_NAME is required")
	}
	return nil
}
