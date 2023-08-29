package database

type SqlHandler interface {
	Create(object interface{})
	FindAll(object interface{})
	DeleteById(object interface{}, id string)
	Update(object interface{}, id string, data_add []string, data_delete []string)
}
