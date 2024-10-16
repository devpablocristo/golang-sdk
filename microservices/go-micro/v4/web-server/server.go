package sdkgomicro

import (
	"fmt"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/web"

	ports "github.com/devpablocristo/golang-sdk/microservices/go-micro/v4/web-server/ports"
)

var (
	instance  ports.Server
	once      sync.Once
	initError error
)

type server struct {
	s web.Service
}

func newServer(config ports.Config) (ports.Server, error) {
	once.Do(func() {
		setupLogger()

		instance = &server{
			s: setupServer(config),
		}

		err := instance.SetRouter(config.GetRouter())
		if err != nil {
			initError = fmt.Errorf("error setting web router: %w", err)
			return
		}
	})

	if initError != nil {
		return nil, initError
	}

	return instance, nil
}

func setupServer(config ports.Config) web.Service {
	Server := web.NewService(
		web.Name(config.GetServerName()),
		web.Id(config.GetServerID()),
		web.Address(config.GetServerAddress()),
		web.Registry(setupRegistry(config)),
	)

	return Server
}

func setupRegistry(config ports.Config) registry.Registry {
	consulReg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{config.GetConsulAddress()}
	})
	return consulReg
}

func (s *server) SetRouter(router interface{}) error {
	switch r := router.(type) {
	case *gin.Engine:
		s.s.Handle("/", r)
	default:
		return fmt.Errorf("unsupported router type")
	}
	return nil
}

func (s *server) Run() error {
	return s.s.Run()
}

func setupLogger() {
	logger.DefaultLogger = logger.NewLogger(
		logger.WithLevel(logger.InfoLevel),
		logger.WithOutput(os.Stdout),
	)
}
