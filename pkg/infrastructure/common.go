package infrastructure

import (
	"github.com/balcieren/go-monolithic-boilerplate/pkg/config"
	"github.com/balcieren/go-monolithic-boilerplate/pkg/log"
	"go.uber.org/fx"
)

func CommonModule() fx.Option {
	return fx.Module(
		"common",
		fx.Provide(config.NewEnv),
		fx.Provide(config.NewCommon),
		fx.Provide(log.New),
	)
}
