package main

import (
	_ "github.com/balcieren/go-monolithic-boilerplate/docs"
	"github.com/balcieren/go-monolithic-boilerplate/pkg/infrastructure"
	"go.uber.org/fx"

	ownerApiV1 "github.com/balcieren/go-monolithic-boilerplate/app/owner/api/v1"
	petApiV1 "github.com/balcieren/go-monolithic-boilerplate/app/pet/api/v1"
)

// @title  Go-Monolithic-Boilerplate API Documentation
// @version 1.0
// @description This is a boilerplate for a monolithic application using Go.
// @host      localhost:8000
// @BasePath  /api
// @schemes http

func main() {
	fx.New(
		infrastructure.HTTPModule("go-monolithic-boilerplate"),
		infrastructure.SwaggerModule, petApiV1.Module, ownerApiV1.Module,
		fx.Invoke(infrastructure.LaunchHTTPServer),
	).Run()
}
