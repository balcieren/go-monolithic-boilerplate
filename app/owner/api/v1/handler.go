package v1

import (
	"github.com/balcieren/go-monolithic-boilerplate/pkg/fail"
	"github.com/balcieren/go-monolithic-boilerplate/pkg/service/owner"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service owner.Servicer
}

func NewHandler(os owner.Servicer) *Handler {
	return &Handler{
		service: os,
	}
}

// @ID ListOwners
// @Summary List owners
// @Description List owners
// @Tags owners v1
// @Accept json
// @Produce json
// @Param page query int false "Page"
// @Param per_page query int false "Per Page"
// @Success 200 {object} owner.ListOwnersResponse
// @Failure 400 {object} helper.ErrorResponse
// @Failure 500 {object} helper.ErrorResponse
// @Router /v1/owners [get]
func (h Handler) ListOwners(c *fiber.Ctx) error {
	resp, err := h.service.ListOwners(c.Context(), &owner.ListOwnersRequest{
		Page:    c.QueryInt("page", 1),
		PerPage: c.QueryInt("per_page", 10),
	})
	if err != nil {
		return fiber.NewError(fail.Convert(err))
	}

	return c.JSON(resp)
}

// @ID GetOwner
// @Summary Get a Owner
// @Description Get a Owner
// @Tags owners v1
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} owner.GetOwnerResponse
// @Failure 400 {object} helper.ErrorResponse
// @Failure 404 {object} helper.ErrorResponse
// @Failure 500 {object} helper.ErrorResponse
// @Router /v1/owners/{id} [get]
func (h Handler) GetOwner(c *fiber.Ctx) error {
	id := c.Params("id")

	resp, err := h.service.GetOwner(c.Context(), &owner.GetOwnerRequest{
		ID: id,
	})
	if err != nil {
		return fiber.NewError(fail.Convert(err))
	}

	return c.JSON(resp)
}

// @ID CreateOwner
// @Summary Create a owner
// @Description Create a owner
// @Tags owners v1
// @Accept json
// @Produce json
// @Param body body owner.CreateOwnerRequest true "Body"
// @Success 200 {object} owner.CreateOwnerResponse
// @Failure 400 {object} helper.ErrorResponse
// @Failure 500 {object} helper.ErrorResponse
// @Router /v1/owners [post]
func (h Handler) CreateOwner(c *fiber.Ctx) error {
	var body owner.CreateOwnerRequest

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	resp, err := h.service.CreateOwner(c.Context(), &body)
	if err != nil {
		return fiber.NewError(fail.Convert(err))
	}

	return c.Status(fiber.StatusCreated).JSON(resp)
}

// @ID UpdateOwner
// @Summary Update a owner
// @Description Update a owner
// @Tags owners v1
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Param body body owner.UpdateOwnerRequest true "Body"
// @Success 200 {object} owner.UpdateOwnerResponse
// @Failure 400 {object} helper.ErrorResponse
// @Failure 500 {object} helper.ErrorResponse
// @Router /v1/owners/{id} [patch]
func (h Handler) UpdateOwner(c *fiber.Ctx) error {
	body := owner.UpdateOwnerRequest{
		ID: c.Params("id"),
	}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	resp, err := h.service.UpdateOwner(c.Context(), &body)
	if err != nil {
		return fiber.NewError(fail.Convert(err))
	}

	return c.Status(fiber.StatusAccepted).JSON(resp)
}

// @ID DeleteOwner
// @Summary Delete a owner
// @Description Delete a owner
// @Tags owners v1
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} owner.DeleteOwnerResponse
// @Failure 400 {object} helper.ErrorResponse
// @Failure 500 {object} helper.ErrorResponse
// @Router /v1/owners/{id} [delete]
func (h Handler) DeleteOwner(c *fiber.Ctx) error {
	id := c.Params("id")

	resp, err := h.service.DeleteOwner(c.Context(), &owner.DeleteOwnerRequest{
		ID: id,
	})
	if err != nil {
		return fiber.NewError(fail.Convert(err))
	}

	return c.Status(fiber.StatusNoContent).JSON(resp)
}
