package sdkviper

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

func LoadConfig(configPaths ...string) error {
	configureViper()

	if len(configPaths) == 0 {
		configPaths = []string{"."}
	}

	for _, configPath := range configPaths {
		fmt.Printf("Searching for config files: %s\n", configPath)

		fileName := filepath.Base(configPath)
		fileExtension := filepath.Ext(fileName)

		if fileExtension == "" {
			return fmt.Errorf("the file path must contain a file name with extension")
		}

		fileExtension = strings.TrimPrefix(fileExtension, ".")
		fileNameWithoutExt := strings.TrimSuffix(fileName, "."+fileExtension)

		viper.SetConfigName(fileNameWithoutExt)
		viper.SetConfigType(fileExtension)
		viper.AddConfigPath(filepath.Dir(configPath))

		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				fmt.Printf("No config file found at dir '%s'\n", configPath)
			} else {
				return fmt.Errorf("error reading config file (%s): %w", configPath, err)
			}
		} else {
			fmt.Printf("Configuration successfully loaded from %s\n", viper.ConfigFileUsed())
		}
	}

	return nil
}

func configureViper() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

func UnmarshalConfig(cfg interface{}) error {
	if err := viper.Unmarshal(cfg); err != nil {
		return fmt.Errorf("unable to decode configuration into struct: %w", err)
	}
	return nil
}
