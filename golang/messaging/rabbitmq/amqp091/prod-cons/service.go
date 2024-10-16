package sdkrabbit

import (
	"context"
	"fmt"
	"sync"

	"github.com/rabbitmq/amqp091-go"

	"github.com/devpablocristo/sdk/golang/sdk/messaging/rabbitmq/amqp091/prod-cons/ports"
)

var (
	instance  ports.Service
	once      sync.Once
	initError error
)

type service struct {
	conn    *amqp091.Connection
	channel *amqp091.Channel
	config  ports.Config
	mutex   sync.Mutex
}

// newService crea una nueva instancia de RabbitMQ que actúa como productor y consumidor.
func newService(config ports.Config) (ports.Service, error) {
	once.Do(func() {
		conn, err := amqp091.Dial(config.GetAddress())
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

		// Declarar el intercambio según la configuración
		err = ch.ExchangeDeclare(
			config.GetExchange(),
			config.GetExchangeType(),
			true,  // Durable
			false, // Auto-deleted
			false, // Internal
			config.GetNoWait(),
			nil, // Arguments adicionales
		)
		if err != nil {
			initError = fmt.Errorf("failed to declare exchange: %w", err)
			ch.Close()
			conn.Close()
			return
		}

		// Declarar la cola
		_, err = ch.QueueDeclare(
			config.GetQueue(),
			true,  // Durable
			false, // Delete when unused
			config.GetExclusive(),
			config.GetNoWait(),
			nil, // Arguments adicionales
		)
		if err != nil {
			initError = fmt.Errorf("failed to declare queue: %w", err)
			ch.Close()
			conn.Close()
			return
		}

		// Enlazar la cola al intercambio con la clave de enrutamiento
		err = ch.QueueBind(
			config.GetQueue(),
			config.GetRoutingKey(),
			config.GetExchange(),
			false, // No-wait
			nil,   // Arguments adicionales
		)
		if err != nil {
			initError = fmt.Errorf("failed to bind queue: %w", err)
			ch.Close()
			conn.Close()
			return
		}

		// Habilitar el modo de confirmación de publicador
		err = ch.Confirm(false)
		if err != nil {
			initError = fmt.Errorf("failed to enable confirm mode: %w", err)
			ch.Close()
			conn.Close()
			return
		}

		instance = &service{
			conn:    conn,
			channel: ch,
			config:  config,
		}
	})

	return instance, initError
}

// Publish envía un mensaje al intercambio especificado o directamente a una cola.
func (m *service) Publish(targetType, targetName, routingKey string, body []byte) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	var err error
	publishing := amqp091.Publishing{
		ContentType: "text/plain",
		Body:        body,
	}

	switch targetType {
	case "exchange":
		err = m.channel.Publish(
			targetName, // Exchange
			routingKey, // Routing key
			false,      // Mandatory
			false,      // Immediate
			publishing,
		)
	case "queue":
		err = m.channel.Publish(
			"",         // No exchange (direct to queue)
			targetName, // Queue name
			false,      // Mandatory
			false,      // Immediate
			publishing,
		)
	default:
		return fmt.Errorf("invalid target type: %s", targetType)
	}

	if err != nil {
		return fmt.Errorf("failed to publish message to %s: %w", targetType, err)
	}

	return nil
}

// Subscribe procesa los mensajes de un intercambio específico o una cola específica.
func (m *service) Subscribe(ctx context.Context, targetType, targetName, exchangeType, routingKey string) (<-chan amqp091.Delivery, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if targetType == "exchange" {
		if err := m.channel.ExchangeDeclare(
			targetName,
			exchangeType,
			true,  // Durable
			false, // Auto-deleted
			false, // Internal
			m.config.GetNoWait(),
			nil, // Arguments adicionales
		); err != nil {
			return nil, fmt.Errorf("failed to declare exchange: %w", err)
		}
	}

	queue, err := m.channel.QueueDeclare(
		targetName,
		true,  // Durable
		false, // Delete when unused
		m.config.GetExclusive(),
		m.config.GetNoWait(),
		nil, // Arguments adicionales
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare queue: %w", err)
	}

	if targetType == "exchange" {
		if err := m.channel.QueueBind(
			queue.Name,
			routingKey,
			targetName,
			false, // No-wait
			nil,   // Arguments adicionales
		); err != nil {
			return nil, fmt.Errorf("failed to bind queue: %w", err)
		}
	}

	msgs, err := m.channel.Consume(
		queue.Name,
		"", // Consumer
		m.config.GetAutoAck(),
		m.config.GetExclusive(),
		m.config.GetNoLocal(),
		m.config.GetNoWait(),
		nil, // Arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to consume from RabbitMQ: %w", err)
	}

	// Crear un canal para filtrar mensajes con cancelación de contexto
	filteredMsgs := make(chan amqp091.Delivery)

	go func() {
		defer close(filteredMsgs)
		for {
			select {
			case <-ctx.Done():
				return
			case msg, ok := <-msgs:
				if !ok {
					return
				}
				filteredMsgs <- msg
			}
		}
	}()

	return filteredMsgs, nil
}

// SetupExchangeAndQueue configura el intercambio y la cola en RabbitMQ.
func (m *service) SetupExchangeAndQueue(exchangeName, exchangeType, queueName, routingKey string) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if err := m.channel.ExchangeDeclare(
		exchangeName, // Nombre del intercambio
		exchangeType, // Tipo de intercambio (direct, topic, fanout, etc.)
		true,         // Durable
		false,        // Auto-deleted
		false,        // Internal
		m.config.GetNoWait(),
		nil, // Arguments adicionales
	); err != nil {
		return fmt.Errorf("failed to declare exchange: %w", err)
	}

	if _, err := m.channel.QueueDeclare(
		queueName, // Nombre de la cola
		true,      // Durable
		false,     // Delete when unused
		m.config.GetExclusive(),
		m.config.GetNoWait(),
		nil, // Arguments adicionales
	); err != nil {
		return fmt.Errorf("failed to declare queue: %w", err)
	}

	if err := m.channel.QueueBind(
		queueName,    // Nombre de la cola
		routingKey,   // Clave de enrutamiento
		exchangeName, // Nombre del intercambio
		false,        // No-wait
		nil,          // Arguments adicionales
	); err != nil {
		return fmt.Errorf("failed to bind queue: %w", err)
	}

	return nil
}

// Close cierra de manera segura la conexión de RabbitMQ.
func (m *service) Close() error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	var errs []error

	if err := m.channel.Close(); err != nil {
		errs = append(errs, fmt.Errorf("failed to close RabbitMQ channel: %w", err))
	}

	if err := m.conn.Close(); err != nil {
		errs = append(errs, fmt.Errorf("failed to close RabbitMQ connection: %w", err))
	}

	if len(errs) > 0 {
		return fmt.Errorf("errors while closing service: %v", errs)
	}

	return nil
}

// GetConnection devuelve la conexión actual de RabbitMQ.
func (m *service) GetConnection() *amqp091.Connection {
	return m.conn
}
