package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/zzats/rest-api/controllers"
	"github.com/zzats/rest-api/models"
)
func main() {
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	shareRepo := new(models.ShareRepo)
	shareRepo.Initialize()
	dividendRepo := new(models.DividendRepo)
	dividendRepo.Initialize()
	e.GET("/shares/", controllers.GetAll(shareRepo))
	e.GET("/shares/:id", controllers.Get(shareRepo))

	e.GET("/dividends/", controllers.GetAll(dividendRepo))
	e.GET("/dividends/:id", controllers.Get(dividendRepo))

	// Server
	e.Start(":1323")
}
