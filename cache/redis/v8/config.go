package sdkredis

import (
	"fmt"

	ports "github.com/devpablocristo/golang/sdk/pkg/cache/redis/v8/ports"
)

type config struct {
	Address  string
	Password string
	DB       int
}

// newConfig crea una nueva configuración de Redis
func newConfig(address, password string, db int) ports.Config {
	return &config{
		Address:  address,
		Password: password,
		DB:       db,
	}
}

// Validate verifica que la configuración de Redis sea válida
func (c *config) Validate() error {
	if c.Address == "" {
		return fmt.Errorf("REDIS_ADDRESS is required")
	}
	if c.DB < 0 {
		return fmt.Errorf("REDIS_DB must be a non-negative integer")
	}
	return nil
}

// GetAddress devuelve la dirección de Redis
func (c *config) GetAddress() string {
	return c.Address
}

// GetPassword devuelve la contraseña de Redis
func (c *config) GetPassword() string {
	return c.Password
}

// GetDB devuelve el número de la base de datos de Redis
func (c *config) GetDB() int {
	return c.DB
}
