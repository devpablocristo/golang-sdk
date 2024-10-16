package ports

type Repository interface {
	GetDb() map[string]any
}
