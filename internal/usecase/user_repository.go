package usecase

import (
	"Avito_CRUD/internal/domain"
)

type UserRepository interface {
	Store(domain.User)
	Select() []domain.User
	Delete(id string)
	UpdateUser(id string, data []string)
	UpdateUsersByPercent(data []string, percent float64)
}
