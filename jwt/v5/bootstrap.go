package sdkjwt

import (
	"github.com/spf13/viper"

	"github.com/devpablocristo/golang/sdk/pkg/jwt/v5/ports"
)

func Bootstrap() (ports.Service, error) {
	config, err := newConfig(
		viper.GetString("JWT_SECRET_KEY"),
	)
	if err != nil {
		return nil, err
	}

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return newService(config)
}
