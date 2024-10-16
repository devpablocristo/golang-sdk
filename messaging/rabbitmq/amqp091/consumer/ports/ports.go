package ports

import (
	"context"

	"github.com/rabbitmq/amqp091-go"
)

// Consumer define las operaciones específicas para un consumidor de RabbitMQ.
type Consumer interface {
	Channel() (*amqp091.Channel, error)
	Close() error
	Consume(context.Context, string, string) ([]byte, string, error)
	SetupExchangeAndQueue(string, string, string, string) error
	GetConnection() *amqp091.Connection
}

// Config define la configuración específica para un consumidor de RabbitMQ.
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

	GetQueue() string
	SetQueue(string)

	GetAutoAck() bool
	SetAutoAck(bool)

	GetExclusive() bool
	SetExclusive(bool)

	GetNoLocal() bool
	SetNoLocal(bool)

	GetNoWait() bool
	SetNoWait(bool)

	Validate() error
}
