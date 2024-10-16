package sdkrabbit

import (
	"fmt"

	"github.com/spf13/viper"

	"github.com/devpablocristo/golang-sdk/messaging/rabbitmq/amqp091/producer/ports"
)

// Bootstrap inicializa una nueva instancia de Producer con configuración de Viper.
func Bootstrap() (ports.Producer, error) {
	config := newConfig(
		viper.GetString("RABBITMQ_HOST"),
		viper.GetInt("RABBITMQ_PORT"),
		viper.GetString("RABBITMQ_USER"),
		viper.GetString("RABBITMQ_PASSWORD"),
		viper.GetString("RABBITMQ_VHOST"),

		viper.GetString("RABBITMQ_EXCHANGE"),
		viper.GetString("RABBITMQ_EXCHANGE_TYPE"),
		viper.GetBool("RABBITMQ_DURABLE"),
		viper.GetBool("RABBITMQ_AUTO_DELETE"),
		viper.GetBool("RABBITMQ_INTERNAL"),
		viper.GetBool("RABBITMQ_NO_WAIT"),
	)

	// Validar la configuración
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("configuration validation failed: %w", err)
	}

	return newProducer(config)
}
