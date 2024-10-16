package mongodbdriver

import (
	"github.com/spf13/viper"

	ports "github.com/devpablocristo/golang-sdk/databases/nosql/mongodb/mongo-driver/ports"
)

func Bootstrap() (ports.Repository, error) {
	config := newConfig(
		viper.GetString("MONGO_USER"),
		viper.GetString("MONGO_PASSWORD"),
		viper.GetString("MONGO_HOST"),
		viper.GetString("MONGO_PORT"),
		viper.GetString("MONGO_DATABASE"),
	)

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return newRepository(config)
}
