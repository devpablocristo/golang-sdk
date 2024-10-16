package sdkjwt

import (
	"time"
)

// Token representa un token de autenticación JWT
type Token struct {
	AccessToken  string
	RefreshToken string
	ExpiresAt    time.Time
}

// LoginCredentials representa las credenciales de inicio de sesión del usuario
type LoginCredentials struct {
	Username     string
	PasswordHash string
}

type TokenClaims struct {
	Subject   string
	ExpiresAt time.Time
	IssuedAt  time.Time
}

//type MapClaims jwt.MapClaims
