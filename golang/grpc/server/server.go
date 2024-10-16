package sdkgrpcserver

import (
	"fmt"
	"net"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"

	ports "github.com/devpablocristo/sdk/golang/sdk/grpc/server/ports"
)

var (
	instance ports.Server
	once     sync.Once
	initErr  error
)

type Server struct {
	server   *grpc.Server
	listener net.Listener
}

func newServer(config ports.Config) (ports.Server, error) {
	once.Do(func() {
		var opts []grpc.ServerOption
		if config.GetTLSConfig() != nil {
			tlsConfig, err := loadTLSConfig(config.GetTLSConfig())
			if err != nil {
				initErr = fmt.Errorf("failed to load TLS config: %v", err)
				return
			}
			creds := credentials.NewTLS(tlsConfig)
			opts = append(opts, grpc.Creds(creds))
		}

		address := fmt.Sprintf("%s:%d", config.GetHost(), config.GetPort())
		listener, err := net.Listen("tcp", address)
		if err != nil {
			initErr = fmt.Errorf("failed to listen: %v", err)
			return
		}

		server := grpc.NewServer(opts...)
		reflection.Register(server) // Registro de reflexión gRPC

		instance = &Server{server: server, listener: listener}
	})
	return instance, initErr
}

func (s *Server) Start() error {
	return s.server.Serve(s.listener)
}

func (s *Server) Stop() error {
	s.server.GracefulStop()
	return s.listener.Close()
}

func (s *Server) RegisterService(serviceDesc any, impl any) {
	sd, ok := serviceDesc.(*grpc.ServiceDesc)
	if !ok {
		panic("serviceDesc must be of type *grpc.ServiceDesc")
	}

	// Registrar el servicio con el servidor gRPC
	s.server.RegisterService(sd, impl)
}
