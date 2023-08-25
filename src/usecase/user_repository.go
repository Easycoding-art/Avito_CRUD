package usecase

import (
	"Avito_CRUD/src/domain"
)

type UserRepository interface {
	Store(domain.User)
	Select() []domain.User
	Delete(id string)
}
