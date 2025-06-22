package handler

import (
	"echo_start/internal/model"
	"echo_start/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CarHandler struct {
	service service.CarService
}

func NewCarHandler(s service.CarService) *CarHandler {
	return &CarHandler{service: s}
}

func (h *CarHandler) CreateCar(e echo.Context) error {
	var car model.Car

	if err := e.Bind(&car); err != nil {
		return e.JSON(http.StatusBadGateway, map[string]string{
			"error": err.Error(),
		})
	}

	if err := h.service.CreateCar(&car); err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return e.JSON(http.StatusCreated, car)
}

func (h *CarHandler) GetCars(e echo.Context) error {
	cars, err := h.service.GetCars()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error()})
	}

	return e.JSON(http.StatusOK, cars)
}
