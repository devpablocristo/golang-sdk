package sdkcgrpcclient

import (
	"github.com/spf13/viper"

	ports "github.com/devpablocristo/sdk/golang/sdk/grpc/client/ports"
)

func Bootstrap() (ports.Client, error) {
	config := newConfig(
		viper.GetString("GRPC_SERVER_HOST"),
		viper.GetInt("GRPC_SERVER_PORT"),
		nil, // Configuración TLS, si es necesario
	)

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return newClient(config)
}
