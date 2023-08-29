package database

import (
	"Avito_CRUD/internal/domain"
)

type UserRepository struct {
	SqlHandler
}

func (db *UserRepository) Store(u domain.User) {
	db.Create(&u)
}

func (db *UserRepository) Select() []domain.User {
	user := []domain.User{}
	db.FindAll(&user)
	return user
}
func (db *UserRepository) Delete(id string) {
	user := []domain.User{}
	db.DeleteById(&user, id)
}
func (db *UserRepository) UpdateUser(id string, data_add []string, data_delete []string) {
	user := []domain.User{}
	db.Update(&user, id, data_add, data_delete)
}

/*
func (db *UserRepository) UpdateUsersByPercent(data_add []string, data_delete []string, percent float64) {
	user := []domain.User{}
	db.UpdateByPercent(&user, data, percent)
}
*/
