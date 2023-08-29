package infrastructure

import (
	"Avito_CRUD/internal/domain"
	"Avito_CRUD/internal/interfaces/database"

	"golang.org/x/exp/slices"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

func (handler *SqlHandler) Update(obj interface{}, id string, data_add []string, data_delete []string) {
	rows, _ := handler.db.Select("id", "data").Rows() // (*sql.Rows, error)
	defer rows.Close()

	var user domain.User
	var data []string = make([]string, 0)
	for rows.Next() {
		// ScanRows сканирует строку в user
		handler.db.ScanRows(rows, &user)
		if user.ID == id {
			for i := range user.Data {
				if !(slices.Contains(data_delete, user.Data[i])) {
					data = append(data, user.Data[i])
				}
			}
			break
		}
	}
	var data_real = append(data, data_add...)
	user_real := domain.User{ID: id, Data: data_real}
	handler.db.Save(&user_real)
}

/*
func (handler *SqlHandler) UpdateByPercent(obj interface{}, data_add []string, data_delete []string, percent float64) {
	rows, _ := handler.db.Select("id", "data").Rows() // (*sql.Rows, error)
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
*/
