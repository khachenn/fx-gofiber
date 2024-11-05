package route

import (
	"net/http"
	"sturdy-winner-api/config"
	"sturdy-winner-api/lib"

	"sturdy-winner-api/internal/api/doc"

	"github.com/dollarsignteam/go-logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/swaggo/swag"
)

// @securityDefinitions.apikey Bearer
// @in                         header
// @name                       Authorization
// @schemes                    http https
// @basePath                   /
type SwaggerRoute struct {
	logger  *logger.Logger
	handler *lib.HttpHandler
}

func NewSwaggerRoute(
	logger *logger.Logger,
	handler *lib.HttpHandler,
) SwaggerRoute {
	return SwaggerRoute{
		logger:  logger,
		handler: handler,
	}
}

func (r SwaggerRoute) Setup() {
	doc.SwaggerInfoAPI.Title = "Sturdy winner API"
	doc.SwaggerInfoAPI.Version = config.Version
	swag.Register(swag.Name, doc.SwaggerInfoAPI)

	r.logger.Info("Setting up swagger routes")
	r.handler.Engine.Get("/swagger/*", swagger.HandlerDefault)
	r.handler.Engine.Get("/swagger", func(c *fiber.Ctx) error {
		return c.Redirect("/swagger/index.html", http.StatusMovedPermanently)
	})
}
