package sdkrabbit

import (
	"fmt"

	"github.com/spf13/viper"

	"github.com/devpablocristo/sdk/golang/sdk/messaging/rabbitmq/amqp091/prod-cons/ports"
)

// Bootstrap inicializa una nueva instancia de Messaging con configuración de Viper.
func Bootstrap() (ports.Service, error) {
	config := newConfig(
		viper.GetString("RABBITMQ_HOST"),
		viper.GetInt("RABBITMQ_PORT"),
		viper.GetString("RABBITMQ_USER"),
		viper.GetString("RABBITMQ_PASSWORD"),
		viper.GetString("RABBITMQ_VHOST"),
		viper.GetString("RABBITMQ_QUEUE"),
		viper.GetString("RABBITMQ_EXCHANGE"),
		viper.GetString("RABBITMQ_EXCHANGE_TYPE"),
		viper.GetString("RABBITMQ_ROUTING_KEY"),
		viper.GetBool("RABBITMQ_AUTO_ACK"),
		viper.GetBool("RABBITMQ_EXCLUSIVE"),
		viper.GetBool("RABBITMQ_NO_LOCAL"),
		viper.GetBool("RABBITMQ_NO_WAIT"),
	)

	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("configuration validation failed: %w", err)
	}

	return newService(config)
}

// RABBITMQ_HOST=localhost
// RABBITMQ_PORT=5672
// RABBITMQ_USER=guest
// RABBITMQ_PASSWORD=guest
// RABBITMQ_VHOST=/
// RABBITMQ_QUEUE=example.queue
// RABBITMQ_EXCHANGE=example.exchange
// RABBITMQ_EXCHANGE_TYPE=topic
// RABBITMQ_ROUTING_KEY=example.key
// RABBITMQ_AUTO_ACK=true
// RABBITMQ_EXCLUSIVE=false
// RABBITMQ_NO_LOCAL=false
// RABBITMQ_NO_WAIT=false
