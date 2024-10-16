package ports

import "github.com/go-redis/redis/v8"

type Cache interface {
	Client() *redis.Client
	Close()
}

// Config define el puerto para la configuración de Redis
type Config interface {
	GetAddress() string
	GetPassword() string
	GetDB() int
	Validate() error
}
