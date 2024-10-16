package sdkcassandra

import (
	"github.com/spf13/viper"

	ports "github.com/devpablocristo/golang/sdk/pkg/databases/nosql/cassandra/gocql/ports"
)

func Bootstrap() (ports.Repository, error) {
	config := newConfig(
		viper.GetStringSlice("CASSANDRA_HOSTS"),
		viper.GetString("CASSANDRA_KEYSPACE"),
		viper.GetString("CASSANDRA_USERNAME"),
		viper.GetString("CASSANDRA_PASSWORD"),
	)

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return newRepository(config)
}
