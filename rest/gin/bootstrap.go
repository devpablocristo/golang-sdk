package sdkgin

import (
	"github.com/spf13/viper"

	ports "github.com/devpablocristo/golang/sdk/pkg/rest/gin/ports"
)

func Bootstrap() (ports.Server, error) {
	config := newConfig(
		viper.GetString("WEB_SERVER_PORT"),
		viper.GetString("API_VERSION"),
	)

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return newServer(config)
}
