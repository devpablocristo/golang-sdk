package ports

import (
	"context"

	"github.com/rabbitmq/amqp091-go"
)

// Messaging define las operaciones para un sistema de mensajería RabbitMQ.
type Service interface {
	Publish(string, string, string, []byte) error
	Subscribe(context.Context, string, string, string, string) (<-chan amqp091.Delivery, error)
	SetupExchangeAndQueue(string, string, string, string) error
	Close() error
	GetConnection() *amqp091.Connection
}

// Config define la configuración específica para RabbitMQ.
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

	GetExchange() string
	SetExchange(string)

	GetExchangeType() string
	SetExchangeType(string)

	GetRoutingKey() string
	SetRoutingKey(string)

	GetAutoAck() bool
	SetAutoAck(bool)

	GetExclusive() bool
	SetExclusive(bool)

	GetNoLocal() bool
	SetNoLocal(bool)

	GetNoWait() bool
	SetNoWait(bool)

	GetAddress() string

	Validate() error
}
