package sdkrabbit

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/rabbitmq/amqp091-go"

	"github.com/devpablocristo/golang-sdk/messaging/rabbitmq/amqp091/producer/ports"
)

var (
	instance  ports.Producer
	once      sync.Once
	initError error
)

// producer implementa la interfaz ports.Producer para RabbitMQ.
type producer struct {
	conn     *amqp091.Connection
	channel  *amqp091.Channel
	exchange string
}

// newProducer crea una nueva instancia de RabbitMQ que actúa como productor.
func newProducer(config ports.Config) (ports.Producer, error) {
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

		// Declarar el intercambio según la configuración
		err = ch.ExchangeDeclare(
			config.GetExchange(),     // Nombre del intercambio
			config.GetExchangeType(), // Tipo de intercambio (direct, topic, fanout, etc.)
			config.IsDurable(),       // Durable
			config.IsAutoDelete(),    // Auto-eliminado
			config.IsInternal(),      // Interno
			config.IsNoWait(),        // No-wait
			nil,                      // Argumentos adicionales
		)
		if err != nil {
			initError = fmt.Errorf("failed to declare exchange: %w", err)
			ch.Close()
			conn.Close()
			return
		}

		// Habilitar el modo de confirmación de publicador
		err = ch.Confirm(false)
		if err != nil {
			initError = fmt.Errorf("failed to put channel into confirm mode: %w", err)
			ch.Close()
			conn.Close()
			return
		}

		instance = &producer{
			conn:     conn,
			channel:  ch,
			exchange: config.GetExchange(),
		}
	})

	return instance, initError
}

// GetInstance devuelve la instancia única de RabbitMQ como productor.
func GetInstance() (ports.Producer, error) {
	if instance == nil {
		return nil, fmt.Errorf("rabbitmq sdkrabbit instance is not initialized")
	}
	return instance, nil
}

// Produce envía un mensaje a la cola especificada con una opción de reply-to y ID de correlación.
func (p *producer) Produce(ctx context.Context, queueName, replyTo, corrID string, message any) (string, error) {
	// Convertir el mensaje a JSON
	body, err := json.Marshal(message)
	if err != nil {
		return "", fmt.Errorf("failed to marshal message: %w", err)
	}

	// Generar un ID de correlación si no se proporciona
	if corrID == "" {
		corrID = fmt.Sprintf("%d", time.Now().UnixNano())
	}

	// Publicar el mensaje con contexto
	err = p.channel.PublishWithContext(ctx,
		p.exchange, // Nombre del intercambio
		queueName,  // Clave de enrutamiento
		false,      // Mandatory
		false,      // Immediate
		amqp091.Publishing{
			ContentType:   "application/json",
			Body:          body,
			CorrelationId: corrID,
			ReplyTo:       replyTo,
		})
	if err != nil {
		return "", fmt.Errorf("failed to publish message to RabbitMQ: %w", err)
	}

	// Confirmar que el mensaje fue recibido por RabbitMQ
	confirms := p.channel.NotifyPublish(make(chan amqp091.Confirmation, 1))

	select {
	case confirmation, ok := <-confirms:
		if !ok {
			return "", fmt.Errorf("confirmation channel closed")
		}
		if confirmation.Ack {
			log.Println("Message acknowledged by RabbitMQ")
		} else {
			return "", fmt.Errorf("message not acknowledged by RabbitMQ")
		}
	case <-ctx.Done():
		return "", ctx.Err()
	}

	return corrID, nil
}

// ProduceWithRetry envía un mensaje con reintentos en caso de fallo.
func (p *producer) ProduceWithRetry(ctx context.Context, queueName, replyTo, corrID string, message any, maxRetries int) (string, error) {
	var err error
	for i := 0; i < maxRetries; i++ {
		corrID, err = p.Produce(ctx, queueName, replyTo, corrID, message)
		if err == nil {
			return corrID, nil
		}
		log.Printf("Retry %d/%d failed: %v", i+1, maxRetries, err)
		time.Sleep(time.Duration(i+1) * time.Second) // Exponencial backoff
	}
	return "", fmt.Errorf("max retries reached: %w", err)
}

// Channel devuelve el canal actual de RabbitMQ.
func (p *producer) Channel() (*amqp091.Channel, error) {
	if p.channel == nil {
		return nil, fmt.Errorf("RabbitMQ channel is not initialized")
	}
	return p.channel, nil
}

// Close cierra de manera segura el productor de RabbitMQ.
func (p *producer) Close() error {
	var errs []error

	if err := p.channel.Close(); err != nil {
		errs = append(errs, fmt.Errorf("failed to close RabbitMQ channel: %w", err))
	}

	if err := p.conn.Close(); err != nil {
		errs = append(errs, fmt.Errorf("failed to close RabbitMQ conn: %w", err))
	}

	if len(errs) > 0 {
		return fmt.Errorf("errors while closing sdkrabbit: %v", errs)
	}

	return nil
}

func (p *producer) GetConnection() *amqp091.Connection {
	return p.conn
}
