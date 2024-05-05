package v1

import (
	"github.com/balcieren/go-monolithic-boilerplate/pkg/fail"
	"github.com/balcieren/go-monolithic-boilerplate/pkg/service/pet"
	utils "github.com/balcieren/go-monolithic-boilerplate/pkg/util"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service pet.Servicer
}

func NewHandler(ps pet.Servicer) *Handler {
	return &Handler{
		service: ps,
	}
}

// @ID ListPets
// @Summary List Pets
// @Description List Pets
// @Tags pets v1
// @Accept json
// @Produce json
// @Param page query int false "Page"
// @Param per_page query int false "Per Page"
// @Param has_owner query bool false "Has Owner"
// @Param owner_id query string false "Owner ID"
// @Success 200 {object} pet.ListPetsResponse
// @Failure 400 {object} helper.ErrorResponse
// @Failure 500 {object} helper.ErrorResponse
// @Router /v1/pets [get]
func (h Handler) ListPets(c *fiber.Ctx) error {
	var hasOwner *bool = nil
	if c.Query("has_owner", "") != "" {
		hasOwner = utils.Ptr(c.Query("has_owner") == "true")
	}

	var ownerID *string = nil
	if c.Query("owner_id", "") != "" {
		ownerID = utils.Ptr(c.Query("owner_id"))
	}

	resp, err := h.service.ListPets(c.Context(), &pet.ListPetsRequest{
		Page:     c.QueryInt("page", 1),
		PerPage:  c.QueryInt("per_page", 10),
		HasOwner: hasOwner,
		OwnerID:  ownerID,
	})
	if err != nil {
		return fiber.NewError(fail.Convert(err))
	}

	return c.JSON(resp)
}

// @ID GetPet
// @Summary Get a pet
// @Description Get a pet
// @Tags pets v1
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} pet.GetPetResponse
// @Failure 400 {object} helper.ErrorResponse
// @Failure 404 {object} helper.ErrorResponse
// @Failure 500 {object} helper.ErrorResponse
// @Router /v1/pets/{id} [get]
func (h Handler) GetPet(c *fiber.Ctx) error {
	id := c.Params("id")

	resp, err := h.service.GetPet(c.Context(), &pet.GetPetRequest{
		ID: id,
	})
	if err != nil {
		return fiber.NewError(fail.Convert(err))
	}

	return c.JSON(resp)
}

// @ID CreatePet
// @Summary Create a pet
// @Description Create a pet
// @Tags pets v1
// @Accept json
// @Produce json
// @Param body body pet.CreatePetRequest true "Name"
// @Success 200 {object} pet.CreatePetResponse
// @Failure 400 {object} helper.ErrorResponse
// @Failure 500 {object} helper.ErrorResponse
// @Router /v1/pets [post]
func (h Handler) CreatePet(c *fiber.Ctx) error {
	var body pet.CreatePetRequest
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	resp, err := h.service.CreatePet(c.Context(), &body)
	if err != nil {
		return fiber.NewError(fail.Convert(err))
	}

	return c.Status(fiber.StatusCreated).JSON(resp)
}

// @ID UpdatePet
// @Summary Update a pet
// @Description Update a pet
// @Tags pets v1
// @Accept json
// @Produce json
// @Param body body pet.UpdatePetRequest true "Body"
// @Param owner_id body string false "Owner ID"
// @Success 200 {object} pet.UpdatePetResponse
// @Failure 400 {object} helper.ErrorResponse
// @Failure 500 {object} helper.ErrorResponse
// @Router /v1/pets/{id} [patch]
func (h Handler) UpdatePet(c *fiber.Ctx) error {
	body := pet.UpdatePetRequest{
		ID: c.Params("id"),
	}
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	resp, err := h.service.UpdatePet(c.Context(), &body)
	if err != nil {
		return fiber.NewError(fail.Convert(err))
	}

	return c.Status(fiber.StatusAccepted).JSON(resp)
}

// @ID DeletePet
// @Summary Delete a pet
// @Description Delete a pet
// @Tags pets v1
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} pet.DeletePetResponse
// @Failure 400 {object} helper.ErrorResponse
// @Failure 500 {object} helper.ErrorResponse
// @Router /v1/pets/{id} [delete]
func (h Handler) DeletePet(c *fiber.Ctx) error {
	id := c.Params("id")

	resp, err := h.service.DeletePet(c.Context(), &pet.DeletePetRequest{
		ID: id,
	})
	if err != nil {
		return fiber.NewError(fail.Convert(err))
	}

	return c.Status(fiber.StatusNoContent).JSON(resp)
}
