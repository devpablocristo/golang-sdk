package sdkmwr

import (
	"net/http"

	"github.com/gin-gonic/gin"

	sdktypes "github.com/devpablocristo/golang-sdk/types"
)

// Constantes para los mensajes de error
const (
	errMissingCredentials = "username and password are required"
)

// ValidateCredentials middleware para validar el payload del login
func ValidateCredentials() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var loginRequest sdktypes.LoginCredentials

		// Manejo del binding y retorno de error en caso de fallo
		if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": errMissingCredentials})
			ctx.Abort()
			return
		}

		// Guardar los datos validados en el contexto para el siguiente handler
		ctx.Set("loginRequest", loginRequest)
		ctx.Next()
	}
}
