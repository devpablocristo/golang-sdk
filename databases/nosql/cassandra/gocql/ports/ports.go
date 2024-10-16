package ports

import (
	"github.com/gocql/gocql"
)

type Config interface {
	GetHosts() []string
	SetHosts(hosts []string)
	GetKeyspace() string
	SetKeyspace(keyspace string)
	GetUsername() string
	SetUsername(username string)
	GetPassword() string
	SetPassword(password string)
	Validate() error
}

type Repository interface {
	Connect(config Config) error
	Close()
	GetSession() *gocql.Session
}
