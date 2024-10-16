
package mongodbdriver

import (
    "fmt"
)

type config struct {
    User         string
    Password     string
    Host         string
    Port         string
    DatabaseName string
}

// newConfig crea una nueva configuración para MongoDB
func newConfig(user, password, host, port, database string) *config {
    return &config{
        User:         user,
        Password:     password,
        Host:         host,
        Port:         port,
        DatabaseName: database,
    }
}

// GetUser devuelve el nombre de usuario
func (c *config) GetUser() string {
    return c.User
}

// GetPassword devuelve la contraseña
func (c *config) GetPassword() string {
    return c.Password
}

// GetHost devuelve el host
func (c *config) GetHost() string {
    return c.Host
}

// GetPort devuelve el puerto
func (c *config) GetPort() string {
    return c.Port
}

// GetDatabaseName devuelve el nombre de la base de datos
func (c *config) GetDatabaseName() string {
    return c.DatabaseName
}

// DSN devuelve la cadena de conexión para MongoDB
func (c *config) DSN() string {
    return fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",
        c.User, c.Password, c.Host, c.Port, c.DatabaseName)
}

// Database devuelve el nombre de la base de datos
func (c *config) Database() string {
    return c.DatabaseName
}

// Validate verifica que la configuración sea válida
func (c *config) Validate() error {
    if c.User == "" || c.Password == "" || c.Host == "" || c.Port == "" || c.DatabaseName == "" {
        return fmt.Errorf("incomplete MongoDB configuration")
    }
    return nil
}
