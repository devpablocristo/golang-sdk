package sdkgomicro

import (
	"fmt"

	"github.com/devpablocristo/golang-sdk/microservices/go-micro/v4/grpc-client/ports"
)

type config struct {
	consulAddress string
	serverName    string
}

func newConfig(ca, sn string) ports.Config {
	return &config{
		consulAddress: ca,
		serverName:    sn,
	}
}

func (c *config) GetConsulAddress() string {
	return c.consulAddress
}

func (c *config) GetServerName() string {
	return c.serverName
}

func (c *config) Validate() error {
	if c.consulAddress == "" {
		return fmt.Errorf("missing consul address")
	}
	if c.serverName == "" {
		return fmt.Errorf("missing service name")
	}
	return nil
}
