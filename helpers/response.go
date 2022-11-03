package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Respons structure
type response struct {
	Status  string `json:"status" xml:"status"`
	Data    any    `json:"data" xml:"data"`
	Message string `json:"message" xml:"message"`
}

// Send a success response
func Success(e echo.Context, data any) error {
	var u = &response{
		Status:  "ok",
		Data:    data,
		Message: "",
	}

	return e.JSON(http.StatusOK, u)
}

// Send an error response
func Error(e echo.Context, message string) error {
	return e.JSON(http.StatusBadRequest, &response{
		Status:  "error",
		Data:    nil,
		Message: message,
	})
}
