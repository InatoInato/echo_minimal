package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func CustomHTTPErrorHandler(err error, e echo.Context) {
	code := http.StatusInternalServerError
	msg := map[string]interface{}{
		"error": err.Error(),
	}

	e.Logger().Error(err)

	if !e.Response().Committed {
		e.JSON(code, msg)
	}
}
