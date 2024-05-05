package v1

import (
	"github.com/balcieren/go-monolithic-boilerplate/pkg/service/owner"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"owner-api-v1",
	fx.Provide(fx.Annotate(
		owner.NewService,
		fx.As(new(owner.Servicer)),
	)),
	fx.Provide(NewHandler),
	fx.Provide(NewRouter),
	fx.Invoke(func(lc fx.Lifecycle, r *Router) {
		r.Setup()
	}),
)
