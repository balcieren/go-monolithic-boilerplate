# Go Monolithic Boilerplate

This boilerplate serves as a solid foundation for building robust, scalable, and maintainable monolithic applications using the Go programming language.

### Monolithic Architecture

In the world of software architecture, a monolithic application refers to a single-layer software application in which various components are combined into a single program. While microservices are popular for their scalability and modularity, monoliths still have their place, especially in scenarios where simplicity and ease of deployment are crucial.

> [!WARNING]
> Pet and owner were shown in the project only as an example.

### Aims of the Boilerplate

1. **Modularity**: Although it is a monolithic boilerplate, it promotes modular design principles and allows you to organize your code base into logical components. This process is made easier using [Uber FX](https://github.com/uber-go/fx) for the dependency injection in the boilerplate.
2. **Scalability**: While a monolithic application may not scale as easily as a microservice, this example provides best practices and patterns to help you scale your application efficiently as it grows.
3. **Ease of Deployment**: Application deployment and management become easier with a single deployable unit compared to a distributed microservices architecture.

### Extras of the Boilerplate

How to use are [Swagger](https://github.com/gofiber/swagger), [Gorm](https://gorm.io/index.html) and [Gorm Gen](https://gorm.io/gen/) mentioned in the boilerplate.

### The Biggest Advantage of the Boilerplate

The app can convertible easily to microservice from monolithic thanks to modularity. All we have to do is creating main.go files separately on the under of the app's folders.

```go
// ./app/pet/main.go
package main

import (
	"github.com/balcieren/go-monolithic-boilerplate/pkg/infrastructure"
	"go.uber.org/fx"

	petApiV1 "github.com/balcieren/go-monolithic-boilerplate/app/pet/api/v1"
)

func main() {
	fx.New(
		infrastructure.HTTPModule("go-monolithic-boilerplate/pet-api"),
		petApiV1.Module,
		fx.Invoke(infrastructure.LaunchHTTPServer),
	).Run()
}

```

### Packages

-   [Fiber](https://github.com/gofiber/fiber)
-   [Fiber Swagger](https://github.com/gofiber/swagger)
-   [Uber FX](https://github.com/uber-go/fx)
-   [Uber Zap](https://github.com/uber-go/zap)
-   [Gorm](https://gorm.io/index.html)
-   [Gorm Gen](https://gorm.io/gen/)

### Databases

-   [PostgreSQL](https://www.postgresql.org/)

### Makefile Commands

Run the app as development mode

```bash
make dev
```

Run the app as production mode

```bash
make prod
```

Generate the swagger document

```bash
make swagger
```

Gorm Gen

```bash
make gorm
```
