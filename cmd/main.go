package main

import (
	"echo_start/internal/config"
	"echo_start/internal/handler"
	"echo_start/internal/repository"
	"echo_start/internal/service"

	"github.com/labstack/echo/v4"
)

func main() {
	config.LoadConfig()
	config.ConnectDB()

	repo := repository.NewCarRepository()
	svc := service.NewServiceCar(repo)
	h := handler.NewCarHandler(svc)

	e := echo.New()

	e.POST("/car", h.CreateCar)
	e.GET("/cars", h.GetCars)

	e.Logger.Fatal(e.Start(":" + config.Cfg.AppPort))
}
