package sdkrabbit

import (
	"fmt"

	"github.com/devpablocristo/golang-sdk/messaging/rabbitmq/amqp091/consumer/ports"
)

// config estructura que implementa la interfaz ports.Config para el consumidor de RabbitMQ.
type config struct {
	host      string
	port      int
	user      string
	password  string
	vhost     string
	queue     string
	autoAck   bool
	exclusive bool
	noLocal   bool
	noWait    bool
}

// newConfig crea una nueva configuración para el consumidor de RabbitMQ con opciones adicionales.
func newConfig(host string, port int, user, password, vhost, queue string, autoAck, exclusive, noLocal, noWait bool) ports.Config {
	return &config{
		host:      host,
		port:      port,
		user:      user,
		password:  password,
		vhost:     vhost,
		queue:     queue,
		autoAck:   autoAck,
		exclusive: exclusive,
		noLocal:   noLocal,
		noWait:    noWait,
	}
}

// Getters y Setters

func (c *config) GetHost() string     { return c.host }
func (c *config) SetHost(host string) { c.host = host }

func (c *config) GetPort() int     { return c.port }
func (c *config) SetPort(port int) { c.port = port }

func (c *config) GetUser() string     { return c.user }
func (c *config) SetUser(user string) { c.user = user }

func (c *config) GetPassword() string         { return c.password }
func (c *config) SetPassword(password string) { c.password = password }

func (c *config) GetVHost() string      { return c.vhost }
func (c *config) SetVHost(vhost string) { c.vhost = vhost }

func (c *config) GetQueue() string      { return c.queue }
func (c *config) SetQueue(queue string) { c.queue = queue }

func (c *config) GetAutoAck() bool        { return c.autoAck }
func (c *config) SetAutoAck(autoAck bool) { c.autoAck = autoAck }

func (c *config) GetExclusive() bool          { return c.exclusive }
func (c *config) SetExclusive(exclusive bool) { c.exclusive = exclusive }

func (c *config) GetNoLocal() bool        { return c.noLocal }
func (c *config) SetNoLocal(noLocal bool) { c.noLocal = noLocal }

func (c *config) GetNoWait() bool       { return c.noWait }
func (c *config) SetNoWait(noWait bool) { c.noWait = noWait }

// Validate verifica que todos los parámetros de configuración sean válidos.
func (c *config) Validate() error {
	if c.host == "" {
		return fmt.Errorf("rabbitmq host is not configured")
	}
	if c.port == 0 {
		return fmt.Errorf("rabbitmq port is not configured")
	}
	if c.user == "" {
		return fmt.Errorf("rabbitmq user is not configured")
	}
	if c.password == "" {
		return fmt.Errorf("rabbitmq password is not configured")
	}
	if c.vhost == "" {
		return fmt.Errorf("rabbitmq vhost is not configured")
	}
	if c.queue == "" {
		return fmt.Errorf("rabbitmq queue is not configured")
	}
	return nil
}
