package sdkmapdb

import (
	"sync"

	ports "github.com/devpablocristo/golang/sdk/pkg/databases/in-memory/mapdb/ports"
)

var (
	instance ports.Repository
	once     sync.Once
)

type service struct {
	db map[string]any
}

func newRepository() ports.Repository {
	once.Do(func() {
		instance = &service{
			db: make(map[string]any),
		}
	})
	return instance
}

func (c *service) GetDb() map[string]any {
	return c.db
}
