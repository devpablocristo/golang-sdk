package sdkrabbit

import (
	"fmt"

	"github.com/devpablocristo/golang/sdk/pkg/messaging/rabbitmq/amqp091/producer/ports"
)

// config estructura que implementa la interfaz ports.Config para el productor de RabbitMQ.
type config struct {
	host         string
	port         int
	user         string
	password     string
	vhost        string
	exchange     string
	exchangeType string
	durable      bool
	autoDelete   bool
	internal     bool
	noWait       bool
}

// newConfig crea una nueva configuración para el productor de RabbitMQ con opciones adicionales.
func newConfig(host string, port int, user, password, vhost, exchange, exchangeType string, durable, autoDelete, internal, noWait bool) ports.Config {
	return &config{
		host:         host,
		port:         port,
		user:         user,
		password:     password,
		vhost:        vhost,
		exchange:     exchange,
		exchangeType: exchangeType,
		durable:      durable,
		autoDelete:   autoDelete,
		internal:     internal,
		noWait:       noWait,
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

func (c *config) GetExchange() string         { return c.exchange }
func (c *config) SetExchange(exchange string) { c.exchange = exchange }

func (c *config) GetExchangeType() string             { return c.exchangeType }
func (c *config) SetExchangeType(exchangeType string) { c.exchangeType = exchangeType }

func (c *config) IsDurable() bool         { return c.durable }
func (c *config) SetDurable(durable bool) { c.durable = durable }

func (c *config) IsAutoDelete() bool            { return c.autoDelete }
func (c *config) SetAutoDelete(autoDelete bool) { c.autoDelete = autoDelete }

func (c *config) IsInternal() bool          { return c.internal }
func (c *config) SetInternal(internal bool) { c.internal = internal }

func (c *config) IsNoWait() bool        { return c.noWait }
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
	if c.exchange == "" {
		return fmt.Errorf("rabbitmq exchange is not configured")
	}
	if c.exchangeType == "" {
		return fmt.Errorf("rabbitmq exchange type is not configured")
	}
	return nil
}
