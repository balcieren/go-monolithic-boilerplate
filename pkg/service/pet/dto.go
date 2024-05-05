package pet

import (
	"github.com/balcieren/go-monolithic-boilerplate/pkg/entity"
)

type ListPetsRequest struct {
	Page     int     `json:"page"`
	PerPage  int     `json:"per_page"`
	HasOwner *bool   `json:"has_owner"`
	OwnerID  *string `json:"owner_id"`
}

type ListPetsResponse struct {
	Rows       []GetPetResponse `json:"rows"`
	Page       int              `json:"page"`
	PerPage    int              `json:"per_page"`
	Total      int              `json:"total"`
	TotalPages int              `json:"total_pages"`
}

type GetPetRequest struct {
	ID string `json:"id"`
}

type GetPetResponse struct {
	entity.Base
	Name    string         `json:"name"`
	OwnerID *string        `json:"owner_id"`
	Type    entity.PetType `json:"type"`
}

type CreatePetRequest struct {
	Name    string  `json:"name"`
	Type    string  `json:"type"`
	OwnerID *string `json:"owner_id"`
}

type CreatePetResponse struct {
	Message string `json:"message"`
}

type UpdatePetRequest struct {
	ID      string  `json:"-"`
	Name    *string `json:"name"`
	Type    *string `json:"type"`
	OwnerID *string `json:"owner_id"`
}

type UpdatePetResponse struct {
	Message string `json:"message"`
}

type DeletePetRequest struct {
	ID string `json:"id"`
}

type DeletePetResponse struct {
	Message string `json:"message"`
}
