package v1

import "github.com/gofiber/fiber/v2"

type Router struct {
	handler *Handler
	root    fiber.Router
}

func NewRouter(app *fiber.App, h *Handler) *Router {
	return &Router{
		handler: h,
		root:    app.Group("/api/v1/pets"),
	}
}

func (r *Router) Setup() {
	r.root.Get("/", r.handler.ListPets)
	r.root.Get("/:id", r.handler.GetPet)
	r.root.Post("/", r.handler.CreatePet)
	r.root.Patch("/:id", r.handler.UpdatePet)
	r.root.Delete("/:id", r.handler.DeletePet)
}
