package sdkgomicro

import (
	"github.com/spf13/viper"

	ports "github.com/devpablocristo/golang-sdk/microservices/go-micro/v4/grpc-server/ports"
)

func Bootstrap() (ports.Server, error) {
	config := newConfig(
		viper.GetString("GRPC_SERVER_NAME"),
		viper.GetString("GRPC_SERVER_HOST"),
		viper.GetInt("GRPC_SERVER_PORT"),
	)

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return newServer(config)
}
