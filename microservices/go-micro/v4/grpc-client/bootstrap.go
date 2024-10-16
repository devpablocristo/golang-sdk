package sdkgomicro

import (
	"github.com/spf13/viper"

	ports "github.com/devpablocristo/golang-sdk/microservices/go-micro/v4/grpc-client/ports"
)

func Bootstrap() (ports.Client, error) {
	config := newConfig(
		viper.GetString("CONSUL_ADDRESS"),
		viper.GetString("GRPC_SERVER_NAME"),
	)

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return newClient(config)
}
