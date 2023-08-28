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

type ByPercent struct {
	Data    []string
	Percent float64
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

func (controller *UserController) GetUser() []domain.User {
	res := controller.Interactor.GetInfo()
	return res
}

func (controller *UserController) Delete(id string) {
	controller.Interactor.Delete(id)
}
func (controller *UserController) Update(c echo.Context) {
	jsonData, err := io.ReadAll(c.Request().Body)
	if err != nil {
		log.Fatal(err)
	}
	var user domain.User
	err2 := json.Unmarshal(jsonData, &user)

	if err2 != nil {

		log.Fatal(err2)
	}
	controller.Interactor.Update(user.ID, user.Data)
}
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
