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

// CreateCar godoc
// @Summary      Create a new car
// @Description  Adds a new car to the database
// @Tags         cars
// @Accept       json
// @Produce      json
// @Param        car  body  model.Car  true  "Car data"
// @Success      201  {object}  model.Car
// @Failure      400  {object}  map[string]string
// @Router       /car [post]
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

// GetCars godoc
// @Summary      Get all cars
// @Description  Returns a list of all cars
// @Tags         cars
// @Produce      json
// @Success      200  {array}  model.Car
// @Router       /cars [get]
func (h *CarHandler) GetCars(e echo.Context) error {
	cars, err := h.service.GetCars()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error()})
	}

	return e.JSON(http.StatusOK, cars)
}
