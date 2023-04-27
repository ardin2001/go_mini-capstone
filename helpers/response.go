package helpers

import "github.com/labstack/echo/v4"

type ResponseModel struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Status  bool        `json:"status"`
}

func Response(e echo.Context, status_code int, response_data ResponseModel) error {
	return e.JSON(status_code, response_data)
}
