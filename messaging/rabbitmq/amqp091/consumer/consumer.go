package sdkrabbit

import (
	"context"
	"fmt"
	"sync"

	"github.com/rabbitmq/amqp091-go"

	"github.com/devpablocristo/golang-sdk/messaging/rabbitmq/amqp091/consumer/ports"
)

var (
	instance  ports.Consumer
	once      sync.Once
	initError error
)

// consumer implementa la interfaz ports.Consumer para RabbitMQ.
type consumer struct {
	conn    *amqp091.Connection
	channel *amqp091.Channel
	config  ports.Config
}

// newConsumer crea una nueva instancia de consumidor de RabbitMQ utilizando el patrón Singleton.
func newConsumer(config ports.Config) (ports.Consumer, error) {
	once.Do(func() {
		connString := fmt.Sprintf("amqp://%s:%s@%s:%d%s",
			config.GetUser(), config.GetPassword(), config.GetHost(), config.GetPort(), config.GetVHost())

		conn, err := amqp091.Dial(connString)
		if err != nil {
			initError = fmt.Errorf("failed to connect to RabbitMQ: %w", err)
			return
		}

		ch, err := conn.Channel()
		if err != nil {
			initError = fmt.Errorf("failed to open a channel: %w", err)
			conn.Close()
			return
		}

		instance = &consumer{
			conn:    conn,
			channel: ch,
			config:  config,
		}
	})

	return instance, initError
}

func (c *consumer) GetConnection() *amqp091.Connection {
	return c.conn
}

// Consume procesa los mensajes de la cola especificada.
// Retorna el primer mensaje que coincide con el corrID proporcionado.
// Si corrID está vacío, retorna el primer mensaje recibido.
func (c *consumer) Consume(ctx context.Context, queueName, corrID string) ([]byte, string, error) {
	msgs, err := c.channel.Consume(
		queueName, "", c.config.GetAutoAck(), c.config.GetExclusive(),
		c.config.GetNoLocal(), c.config.GetNoWait(), nil,
	)
	if err != nil {
		return nil, "", fmt.Errorf("failed to consume from RabbitMQ: %w", err)
	}

	for {
		select {
		case <-ctx.Done():
			return nil, "", ctx.Err()
		case msg, ok := <-msgs:
			if !ok {
				return nil, "", fmt.Errorf("message channel closed")
			}
			if corrID == "" || msg.CorrelationId == corrID {
				return msg.Body, msg.CorrelationId, nil
			}
		}
	}
}

// SetupExchangeAndQueue configura el intercambio y la cola en RabbitMQ.
func (c *consumer) SetupExchangeAndQueue(exchangeName, exchangeType, queueName, routingKey string) error {
	if err := c.channel.ExchangeDeclare(
		exchangeName, // Nombre del intercambio
		exchangeType, // Tipo de intercambio (direct, topic, fanout, etc.)
		true,         // Durable
		false,        // Auto-deleted
		false,        // Internal
		false,        // No-wait
		nil,          // Arguments
	); err != nil {
		return fmt.Errorf("failed to declare exchange: %w", err)
	}

	if _, err := c.channel.QueueDeclare(
		queueName, // Nombre de la cola
		true,      // Durable
		false,     // Delete when unused
		false,     // Exclusive
		false,     // No-wait
		nil,       // Arguments
	); err != nil {
		return fmt.Errorf("failed to declare queue: %w", err)
	}

	if err := c.channel.QueueBind(
		queueName,    // Nombre de la cola
		routingKey,   // Clave de enrutamiento
		exchangeName, // Nombre del intercambio
		false,        // No-wait
		nil,          // Arguments
	); err != nil {
		return fmt.Errorf("failed to bind queue: %w", err)
	}

	return nil
}

// Close cierra de manera segura el consumidor.
func (c *consumer) Close() error {
	var errs []error

	if err := c.channel.Close(); err != nil {
		errs = append(errs, fmt.Errorf("failed to close RabbitMQ channel: %w", err))
	}

	if err := c.conn.Close(); err != nil {
		errs = append(errs, fmt.Errorf("failed to close RabbitMQ conn: %w", err))
	}

	if len(errs) > 0 {
		return fmt.Errorf("errors while closing consumer: %v", errs)
	}

	return nil
}

// Channel devuelve el canal actual de RabbitMQ.
func (c *consumer) Channel() (*amqp091.Channel, error) {
	if c.channel == nil {
		return nil, fmt.Errorf("RabbitMQ channel is not initialized")
	}
	return c.channel, nil
}
