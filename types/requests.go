package sdktypes

type LoginCredentials struct {
	Username     string `json:"username" binding:"required"`
	PasswordHash string `json:"password" binding:"required"`
}
