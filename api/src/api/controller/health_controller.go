package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"golang-core/api/src/client/response"
)

type HealthController struct {
	log *zap.SugaredLogger
}

var healthResponse = response.ToSuccessResponse(
	response.HealthResponse{
		Status: "OK",
	},
)

func NewHealthController(log *zap.SugaredLogger) *HealthController {
	return &HealthController{
		log: log,
	}
}

func (c *HealthController) HealthCheck(ec echo.Context) error {
	return ec.JSON(http.StatusOK, healthResponse)
}
