package controllers

import (
	"Avito_CRUD/internal/domain"
	"Avito_CRUD/internal/interfaces/database"
	"Avito_CRUD/internal/usecase"
	"encoding/json"
	"io"
	"log"

	"github.com/labstack/echo"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

type Request struct {
	ID         string
	DataAdd    []string
	DataDelete []string
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) Create(c echo.Context) {
	u := domain.User{}
	c.Bind(&u)
	controller.Interactor.Add(u)
	createdUsers := controller.Interactor.GetInfo()
	c.JSON(201, createdUsers)
	return
}

func (controller *UserController) GetUser(c echo.Context) []domain.User {
	results := controller.Interactor.GetInfo()
	id := c.QueryParam("id")
	var res []string
	for i := range results {
		if results[i].ID == id {
			res = results[i].Data
			break
		}
	}
	var user domain.User = domain.User{ID: id, Data: res}

	c.JSON(201, user)
	return nil
}

func (controller *UserController) Delete(id string) {
	controller.Interactor.Delete(id)
}
func (controller *UserController) Update(c echo.Context) {
	jsonData, err := io.ReadAll(c.Request().Body)
	if err != nil {
		log.Fatal(err)
	}
	var data Request
	err2 := json.Unmarshal(jsonData, &data)

	if err2 != nil {

		log.Fatal(err2)
	}
	controller.Interactor.Update(data.ID, data.DataAdd, data.DataDelete)
}

/*
func (controller *UserController) UpdateByPercent(c echo.Context) {
	jsonData, err := io.ReadAll(c.Request().Body)
	if err != nil {
		log.Fatal(err)
	}
	var req ByPercent
	err2 := json.Unmarshal(jsonData, &req)

	if err2 != nil {

		log.Fatal(err2)
	}
	controller.Interactor.UpdateByPercent(req.Data, req.Percent)

}
*/
