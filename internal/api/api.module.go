package api

import (
	"context"
	"sturdy-winner-api/config"
	"sturdy-winner-api/internal/api/controller"
	"sturdy-winner-api/internal/api/route"
	"sturdy-winner-api/internal/api/service"
	"sturdy-winner-api/internal/repository"
	"sturdy-winner-api/lib"

	"github.com/dollarsignteam/go-logger"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type FiberApp struct {
	Server *fiber.App
}

var Module = fx.Options(
	fx.Provide(NewLoggerOptions),
	fx.Provide(NewMongoDBOptions),
	fx.Provide(config.NewAPIConfig),
	fx.Provide(lib.NewHttpHandler),
	fx.Provide(logger.NewLogger),
	fx.Provide(lib.NewMongoDB),
	route.Module,
	controller.Module,
	service.Module,
	repository.Module,
	fx.Invoke(Run),
)

func NewLoggerOptions(config *config.APIConfig) logger.LoggerOptions {
	return logger.LoggerOptions{
		Level: config.LogLevel,
		Name:  "sturdy-winner-api",
	}
}

func NewMongoDBOptions(config *config.APIConfig) lib.MongoDBOptions {
	return lib.MongoDBOptions{
		DSN: config.MongoURI,
	}
}

func Run(
	lifeCycle fx.Lifecycle,
	routes route.Routes,
	httpHandler *lib.HttpHandler,
	config *config.APIConfig,
	log *logger.Logger,
	mongo *lib.MongoDB,
) {
	lifeCycle.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			go func() {
				routes.Setup()
				if err := httpHandler.Engine.Listen(config.ListenAddr()); err != nil {
					log.Fatalf("start server error : %v\n", err)
				}
			}()
			return nil
		},
		OnStop: func(_ context.Context) error {
			log.Warnf("stopping server ...")
			mongo.Disconnect()
			if err := httpHandler.ShutDown(); err != nil {
				log.Warn(err)
			}
			log.Warnf("stop server success")
			return nil
		},
	})
}
