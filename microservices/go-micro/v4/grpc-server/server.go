package sdkgomicro

import (
	"fmt"
	"sync"

	"github.com/go-micro/plugins/v4/server/grpc"
	gmserver "go-micro.dev/v4/server"

	ports "github.com/devpablocristo/golang-sdk/microservices/go-micro/v4/grpc-server/ports"
)

var (
	instance  ports.Server
	once      sync.Once
	initError error
)

type server struct {
	s gmserver.Server
}

func newServer(config ports.Config) (ports.Server, error) {
	once.Do(func() {
		srv, err := setupServer(config)
		if err != nil {
			initError = fmt.Errorf("error setting up server: %w", err)
			return
		}
		instance = &server{
			s: srv,
		}
	})

	if initError != nil {
		return nil, initError
	}

	return instance, nil
}

func setupServer(config ports.Config) (gmserver.Server, error) {
	s := grpc.NewServer(
		gmserver.Name(config.GetServerName()),
		gmserver.Id(config.GetServerID()),
		gmserver.Address(fmt.Sprintf("%s:%d", config.GetServerHost(), config.GetServerPort())),
	)

	return s, nil
}

func (s *server) GetServer() gmserver.Server {
	return s.s
}
