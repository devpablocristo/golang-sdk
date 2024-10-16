package sdkgodotenv

import (
	"fmt"

	"github.com/joho/godotenv"
)

// LoadConfig carga múltiples archivos desde las rutas completas proporcionadas.
// Cada string en configPaths debe representar la ruta completa, incluyendo el nombre del archivo.
// Retorna un error en caso de que no se pueda cargar algún archivo.
func LoadConfig(configPaths ...string) error {
	// Verificar que se hayan proporcionado rutas
	if len(configPaths) == 0 {
		return fmt.Errorf("no config paths provided")
	}

	// Cargar cada archivo desde las rutas completas especificadas.
	for _, filePath := range configPaths {
		fmt.Printf("Attempting to load file from: %s\n", filePath) // Debugging
		if err := godotenv.Load(filePath); err != nil {
			return fmt.Errorf("failed to load file from path %s: %w", filePath, err)
		}
	}

	fmt.Println("Files loaded successfully.")
	return nil
}
