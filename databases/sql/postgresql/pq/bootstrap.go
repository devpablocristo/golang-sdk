package sdkpg

import (
	"github.com/spf13/viper"

	ports "github.com/devpablocristo/golang-sdk/databases/sql/postgresql/pq/ports"
)

func Bootstrap() (ports.Repository, error) {
	config := newConfig(
		viper.GetString("POSTGRES_USER"),
		viper.GetString("POSTGRES_PASSWORD"),
		viper.GetString("POSTGRES_HOST"),
		viper.GetString("POSTGRES_PORT"),
		viper.GetString("POSTGRES_DB"),
	)

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return newRepository(config)
}
