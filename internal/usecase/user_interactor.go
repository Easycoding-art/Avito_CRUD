package usecase

import "Avito_CRUD/internal/domain"

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

func (interactor *UserInteractor) Update(id string, data []string) {
	interactor.UserRepository.UpdateUser(id, data)
}

func (interactor *UserInteractor) UpdateByPercent(data []string, percent float64) {
	interactor.UserRepository.UpdateUsersByPercent(data, percent)
}
