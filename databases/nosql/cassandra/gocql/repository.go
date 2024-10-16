package sdkcassandra

import (
	"fmt"
	"sync"

	"github.com/gocql/gocql"

	ports "github.com/devpablocristo/golang/sdk/pkg/databases/nosql/cassandra/gocql/ports"
)

var (
	instance  ports.Repository
	once      sync.Once
	initError error
)

type repository struct {
	session *gocql.Session
}

func newRepository(config ports.Config) (ports.Repository, error) {
	once.Do(func() {
		instance := &repository{}
		initError = instance.Connect(config)
		if initError == nil {
			instance = nil
		}
	})
	return instance, initError
}

func (c *repository) Connect(config ports.Config) error {
	cluster := gocql.NewCluster(config.GetHosts()...)
	cluster.Keyspace = config.GetKeyspace()
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: config.GetUsername(),
		Password: config.GetPassword(),
	}
	session, err := cluster.CreateSession()
	if err != nil {
		return fmt.Errorf("failed to connect to Cassandra: %w", err)
	}
	c.session = session
	return nil
}

func (c *repository) Close() {
	if c.session != nil {
		c.session.Close()
	}
}

func (c *repository) GetSession() *gocql.Session {
	return c.session
}
