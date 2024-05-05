package v1

import (
	"github.com/balcieren/go-monolithic-boilerplate/pkg/service/pet"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"pet-api-v1",
	fx.Provide(fx.Annotate(
		pet.NewService,
		fx.As(new(pet.Servicer)),
	)),
	fx.Provide(NewHandler),
	fx.Provide(NewRouter),
	fx.Invoke(func(lc fx.Lifecycle, r *Router) {
		r.Setup()
	}),
)
