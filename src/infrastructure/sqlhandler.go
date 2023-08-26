package infrastructure

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"Avito_CRUD/src/domain"
	"Avito_CRUD/src/interfaces/database"
)

type SqlHandler struct {
	db *gorm.DB
}

func NewSqlHandler() database.SqlHandler {
	dsn := "root:password@tcp(127.0.0.1:3306)/go_sample?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.db = db
	return sqlHandler
}

func (handler *SqlHandler) Create(obj interface{}) {
	handler.db.Create(obj)
}

func (handler *SqlHandler) FindAll(obj interface{}) {
	handler.db.Find(obj)
}

func (handler *SqlHandler) DeleteById(obj interface{}, id string) {
	handler.db.Delete(obj, id)
}

func (handler *SqlHandler) Update(obj interface{}, id string, data []string) {
	user := domain.User{ID: id, Data: data}
	handler.db.Save(&user)
}
func (handler *SqlHandler) UpdateByPercent(obj interface{}, data []string, percent float64) {
	rows, _ := handler.db.Select("name, age, email").Rows() // (*sql.Rows, error)
	defer rows.Close()

	var user domain.User
	var count int64
	handler.db.Count(&count)
	people := int64(float64(count) * percent)
	i := 0
	for rows.Next() {
		// ScanRows сканирует строку в user
		handler.db.ScanRows(rows, &user)
		user.Data = data
		handler.db.Save(&user)
		i++
		if i == int(people) {
			break
		}
	}

}
