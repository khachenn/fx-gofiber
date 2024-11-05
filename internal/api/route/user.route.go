package route

import (
	"sturdy-winner-api/internal/api/controller"
	"sturdy-winner-api/lib"

	"github.com/dollarsignteam/go-logger"
)

type UserRoute struct {
	logger  *logger.Logger
	handler *lib.HttpHandler
	ctrl    controller.UserController
}

func NewUserRoute(
	logger *logger.Logger,
	handler *lib.HttpHandler,
	ctrl controller.UserController,
) UserRoute {
	return UserRoute{
		logger:  logger,
		handler: handler,
		ctrl:    ctrl,
	}
}

func (r UserRoute) Setup() {
	r.logger.Info("Setting up user route")

	api := r.handler.Engine.Group("/api")
	api.Get("/users", r.ctrl.GetUserList)
}
