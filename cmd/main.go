package main

import (
	"echo_start/internal/config"
	"echo_start/internal/handler"
	"echo_start/internal/middleware"
	"echo_start/internal/repository"
	"echo_start/internal/service"

	"github.com/labstack/echo/v4"
	echowm "github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "echo_start/docs" // путь к docs — он должен соответствовать
)
// @title Cars API
// @version 1.0
// @description API for managing cars
// @host localhost:8080
// @BasePath /

func main() {
	config.LoadConfig()
	config.ConnectDB()

	repo := repository.NewCarRepository()
	svc := service.NewServiceCar(repo)
	h := handler.NewCarHandler(svc)

	e := echo.New()
	e.HTTPErrorHandler = middleware.CustomHTTPErrorHandler

	e.Use(echowm.Logger())
	e.Use(echowm.Recover())
	e.Use(echowm.CORSWithConfig(echowm.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	e.POST("/car", h.CreateCar)
	e.GET("/cars", h.GetCars)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":" + config.Cfg.AppPort))
}
