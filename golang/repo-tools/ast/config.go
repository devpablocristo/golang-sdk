package sdkast

import (
	"fmt"

	ports "github.com/devpablocristo/sdk/golang/sdk/repo-tools/ast/ports"
)

type config struct {
	AnalyzePath string
}

func newConfig(analyzePath string) ports.Config {
	return &config{
		AnalyzePath: analyzePath,
	}
}

func (c *config) GetAnalyzePath() string {
	return c.AnalyzePath
}

func (c *config) SetAnalyzePath(analyzePath string) {
	c.AnalyzePath = analyzePath
}

func (c *config) Validate() error {
	if c.AnalyzePath == "" {
		return fmt.Errorf("analyze path is not configured")
	}
	return nil
}
