package route

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewRoutes),
	fx.Provide(NewSwaggerRoute),
	fx.Provide(NewUserRoute),
)

type Route interface {
	Setup()
}

type Routes []Route

func NewRoutes(
	swagger SwaggerRoute,
	user UserRoute,
) Routes {
	return Routes{
		swagger,
		user,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
