package sdkpostgresql

import (
	"github.com/spf13/viper"

	"github.com/devpablocristo/sdk/golang/sdk/databases/sql/postgresql/pgxpool/ports"
)

func Bootstrap() (ports.Repository, error) {
	config := newConfig(
		viper.GetString("POSTGRES_USER"),
		viper.GetString("POSTGRES_PASSWORD"),
		viper.GetString("POSTGRES_HOST"),
		viper.GetString("POSTGRES_PORT"),
		viper.GetString("POSTGRES_NAME"),
	)

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return newRepository(config)
}
