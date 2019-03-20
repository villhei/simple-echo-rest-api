package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/villhei/rest-api/models"
)

func Get(repo models.Repo) func(c echo.Context) error {
	return func(c echo.Context) error {
		entity := repo.GetOne(1)
		return c.JSON(http.StatusOK, entity)
	}
}

func GetAll(repo models.Repo) func(c echo.Context) error {
	return func(c echo.Context) error {
		entities := repo.GetAll()
		return c.JSON(http.StatusOK, entities)
	}
}
