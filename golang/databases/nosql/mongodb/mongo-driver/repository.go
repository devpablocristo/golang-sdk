package mongodbdriver

import (
	"context"
	"fmt"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	ports "github.com/devpablocristo/sdk/golang/sdk/databases/nosql/mongodb/mongo-driver/ports"
)

var (
	instance  ports.Repository
	once      sync.Once
	initError error
)

type repository struct {
	db *mongo.Database
}

func newRepository(c ports.Config) (ports.Repository, error) {
	once.Do(func() {
		instance = &repository{}
		initError = instance.Connect(c)
		if initError != nil {
			instance = nil
		}
	})
	return instance, initError
}

func (r *repository) Connect(c ports.Config) error {
	dsn := c.DSN()
	clientOptions := options.Client().ApplyURI(dsn)

	conn, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Verificar la conexión con un ping
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = conn.Ping(ctx, nil); err != nil {
		return fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	r.db = conn.Database(c.Database())
	return nil
}

func (r *repository) Close() {
	if r.db != nil {
		r.db.Client().Disconnect(context.TODO())
	}
}

func (r *repository) DB() *mongo.Database {
	return r.db
}
