package exception

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"golang-core/api/src/client/response"
)

type ErrorModel response.GeneralResponse[any]

func newError(httpCode int, message string, err error) *echo.HTTPError {
	msg := ErrorModel{
		Code:        httpCode,
		Message:     message,
		ErrorDetail: err.Error(),
	}
	return echo.NewHTTPError(httpCode, msg)
}

func NewBadRequestError(err error) *echo.HTTPError {
	return newError(http.StatusBadRequest, "Bad request", err)
}
