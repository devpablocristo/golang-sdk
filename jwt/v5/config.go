package sdkjwt

import (
	"fmt"

	"github.com/devpablocristo/golang/sdk/pkg/jwt/v5/ports"
)

type config struct {
	secretKey string
}

func newConfig(secretKey string) (ports.Config, error) {
	return &config{
		secretKey: secretKey,
	}, nil
}

func (c *config) GetSecretKey() string {
	return c.secretKey
}

func (c *config) Validate() error {
	if c.secretKey == "" {
		return fmt.Errorf("JWT secret key is not configured")
	}
	return nil
}
