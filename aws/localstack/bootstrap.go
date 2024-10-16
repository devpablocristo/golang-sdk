package sdkawslcl

import (
	"github.com/spf13/viper"

	"github.com/devpablocristo/golang-sdk/aws/localstack/ports"
)

func Bootstrap() (ports.Stack, error) {
	config := newConfig(
		viper.GetString("AWS_ACCESS_KEY_ID"),
		viper.GetString("AWS_SECRET_ACCESS_KEY"),
		viper.GetString("AWS_REGION"),
		viper.GetString("AWS_LOCALSTACK_ENDPOINT"),
	)

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return newStack(config)
}
