package sdkredis

import (
	"github.com/spf13/viper"

	ports "github.com/devpablocristo/sdk/golang/sdk/cache/redis/v8/ports"
)

func Bootstrap() (ports.Cache, error) {
	config := newConfig(
		viper.GetString("REDIS_ADDRESS"),
		viper.GetString("REDIS_PASSWORD"),
		viper.GetInt("REDIS_DB"),
	)

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return newCache(config)
}
