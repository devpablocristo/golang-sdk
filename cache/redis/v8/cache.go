package sdkredis

import (
	"context"
	"fmt"
	"sync"
	"time"

	ports "github.com/devpablocristo/golang/sdk/pkg/cache/redis/v8/ports"
	"github.com/go-redis/redis/v8"
)

var (
	instance  ports.Cache
	once      sync.Once
	initError error
)

type Cache struct {
	client *redis.Client
}

// newCache inicializa la instancia del cache Redis utilizando el patrón singleton
func newCache(c ports.Config) (ports.Cache, error) {
	once.Do(func() {
		client := &Cache{}
		initError = client.connect(c)
		if initError != nil {
			instance = nil
		} else {
			instance = client
		}
	})
	return instance, initError
}

// connect conecta al servidor Redis utilizando los getters de la configuración
func (ch *Cache) connect(c ports.Config) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.GetAddress(),
		Password: c.GetPassword(),
		DB:       c.GetDB(),
	})

	// Contexto con timeout para limitar el tiempo de conexión
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Verificar si Redis está accesible mediante Ping
	if err := rdb.Ping(ctx).Err(); err != nil {
		return fmt.Errorf("failed to connect to redis: %w", err)
	}
	ch.client = rdb
	return nil
}

// Close cierra la conexión con el servidor Redis
func (ch *Cache) Close() {
	if ch.client != nil {
		ch.client.Close()
	}
}

// Client devuelve el cliente Redis
func (ch *Cache) Client() *redis.Client {
	return ch.client
}
