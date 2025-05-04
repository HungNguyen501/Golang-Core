package controller

import (
	"go.uber.org/zap"

	"golang-core/api/src/manager"
)

type Controllers struct {
	HealthController *HealthController
	UserController   *UserController
}

func NewControllers(log *zap.SugaredLogger, managers *manager.Managers) *Controllers {
	return &Controllers{
		HealthController: NewHealthController(log),
		UserController:   NewUserController(managers.UserManager),
	}
}
