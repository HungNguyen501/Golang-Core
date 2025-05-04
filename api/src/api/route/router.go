package route

import (
	"golang-core/api/src/api/controller"
	echoutils "golang-core/api/src/common/echo_utils"

	"github.com/labstack/echo/v4"
)

func NewRouter(cts *controller.Controllers, validators *echoutils.Validators) *echo.Echo {
	ec := echo.New()
	ec.HidePort = true
	ec.Validator = validators

	healthGroup := ec.Group("/api/v1")
	healthGroup.GET("/health", cts.HealthController.HealthCheck)

	userGroup := ec.Group("/api/v1/user")
	userGroup.POST("/insert", cts.UserController.Insert)
	userGroup.POST("/update", cts.UserController.Update)
	userGroup.DELETE("/:id", cts.UserController.Delete)
	userGroup.GET("/:id", cts.UserController.FindById)
	userGroup.GET("/list", cts.UserController.ListByIds)

	return ec
}
