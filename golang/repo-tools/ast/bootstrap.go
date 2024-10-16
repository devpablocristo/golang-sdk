package sdkast

import (
	"github.com/spf13/viper"

	ports "github.com/devpablocristo/sdk/golang/sdk/repo-tools/ast/ports"
)

// Bootstrap inicializa y valida la configuración del AST parser.
func Bootstrap() (ports.Service, error) {
	config := newConfig(
		viper.GetString("AST_ANALYZE_PATH"),
	)

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return newService(config)
}
