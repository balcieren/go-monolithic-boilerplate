package infrastructure

import (
	"context"

	"net"

	"github.com/balcieren/go-monolithic-boilerplate/pkg/config"
	"github.com/balcieren/go-monolithic-boilerplate/pkg/database"
	"github.com/balcieren/go-monolithic-boilerplate/pkg/helper"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.uber.org/fx"
)

func HTTPModule(name AppName) fx.Option {
	return fx.Module(
		"http",
		fx.Provide(func() *AppName {
			return &name
		}),
		fx.Provide(database.NewPostgreSQL),
		fx.Provide(func(an *AppName, env *config.Env) *fiber.App {
			app := fiber.New(fiber.Config{
				AppName:      an.String(),
				ErrorHandler: helper.ErrorHandler,
			})

			app.Use(logger.New())

			return app
		}),
	)
}

func LaunchHTTPServer(lc fx.Lifecycle, app *fiber.App, cmn *config.Common) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go app.Listen(net.JoinHostPort("", cmn.API.PORT))
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return app.Shutdown()
		},
	})
}
