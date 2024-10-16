package sdkrabbit

import (
	"github.com/spf13/viper"

	"github.com/devpablocristo/golang-sdk/messaging/rabbitmq/amqp091/consumer/ports"
)

// Bootstrap inicializa una nueva instancia de Consumer con configuración de Viper.
func Bootstrap() (ports.Consumer, error) {
	config := newConfig(
		viper.GetString("RABBITMQ_HOST"),
		viper.GetInt("RABBITMQ_PORT"),
		viper.GetString("RABBITMQ_USER"),
		viper.GetString("RABBITMQ_PASSWORD"),
		viper.GetString("RABBITMQ_VHOST"),
		viper.GetString("RABBITMQ_QUEUE"),
		viper.GetBool("RABBITMQ_AUTO_ACK"),
		viper.GetBool("RABBITMQ_EXCLUSIVE"),
		viper.GetBool("RABBITMQ_NO_LOCAL"),
		viper.GetBool("RABBITMQ_NO_WAIT"),
	)

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return newConsumer(config)
}
