package sdkgomicro

import (
	"github.com/spf13/viper"

	ports "github.com/devpablocristo/golang-sdk/microservices/go-micro/v4/web-server/ports"
)

func Bootstrap(webRouter interface{}) (ports.Server, error) {
	config := newConfig(
		webRouter,
		viper.GetString("WEB_SERVER_NAME"),
		viper.GetString("CONSUL_ADDRESS"),
		viper.GetString("WEB_SERVER_HOST"),
		viper.GetInt("WEB_SERVER_PORT"),
	)

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return newServer(config)
}
