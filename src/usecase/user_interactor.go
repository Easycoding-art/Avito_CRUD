package usecase

import "Avito_CRUD/src/domain"

type UserInteractor struct {
	UserRepository
}

func (interactor *UserInteractor) Add(u domain.User) {
	interactor.UserRepository.Store(u)
}

func (interactor *UserInteractor) GetInfo() []domain.User {
	return interactor.UserRepository.Select()
}

func (interactor *UserInteractor) Delete(id string) {
	interactor.UserRepository.Delete(id)
}