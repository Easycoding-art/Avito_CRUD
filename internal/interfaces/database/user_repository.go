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
func (db *UserRepository) UpdateUser(id string, data []string) {
	user := []domain.User{}
	db.Update(&user, id, data)
}
func (db *UserRepository) UpdateUsersByPercent(data []string, percent float64) {
	user := []domain.User{}
	db.UpdateByPercent(&user, data, percent)
}
