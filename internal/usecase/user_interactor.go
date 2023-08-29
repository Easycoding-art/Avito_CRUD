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

func (interactor *UserInteractor) Update(id string, data_add []string, data_delete []string) {
	interactor.UserRepository.UpdateUser(id, data_add, data_delete)
}

/*
func (interactor *UserInteractor) UpdateByPercent(data_add []string, data_delete []string, percent float64) {
	interactor.UserRepository.UpdateUsersByPercent(data, percent)
}
*/
