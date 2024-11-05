package repository

import (
	"sturdy-winner-api/lib"

	"github.com/dollarsignteam/go-logger"
	"go.uber.org/fx"
)

type Handler struct {
	mongo  *lib.MongoDB
	logger *logger.Logger
}

func NewHandler(
	mongo *lib.MongoDB,
	logger *logger.Logger,
) *Handler {
	return &Handler{
		mongo:  mongo,
		logger: logger,
	}
}

var Module = fx.Options(
	fx.Provide(NewHandler),
)
