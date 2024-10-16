package ports

import (
	"context"

	"github.com/rabbitmq/amqp091-go"
)

// Producer define las operaciones específicas para un productor de RabbitMQ.
type Producer interface {
	// Channel devuelve el canal actual de RabbitMQ.
	Channel() (*amqp091.Channel, error)

	// Close cierra de manera segura el productor de RabbitMQ.
	Close() error

	// Produce envía un mensaje a la cola especificada con una opción de reply-to y ID de correlación.
	Produce(context.Context, string, string, string, any) (string, error)

	// ProduceWithRetry envía un mensaje con reintentos en caso de fallo.
	ProduceWithRetry(context.Context, string, string, string, any, int) (string, error)

	GetConnection() *amqp091.Connection
}

// Config define la configuración específica para un productor de RabbitMQ.
type Config interface {
	GetHost() string
	SetHost(string)

	GetPort() int
	SetPort(int)

	GetUser() string
	SetUser(string)

	GetPassword() string
	SetPassword(string)

	GetVHost() string
	SetVHost(string)

	GetExchange() string
	SetExchange(string)

	GetExchangeType() string
	SetExchangeType(string)

	IsDurable() bool
	SetDurable(bool)

	IsAutoDelete() bool
	SetAutoDelete(bool)

	IsInternal() bool
	SetInternal(bool)

	IsNoWait() bool
	SetNoWait(bool)

	Validate() error
}
