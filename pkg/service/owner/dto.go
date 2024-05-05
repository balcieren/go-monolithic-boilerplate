package owner

import (
	"github.com/balcieren/go-monolithic-boilerplate/pkg/entity"
)

type ListOwnersRequest struct {
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
}

type ListOwnersResponse struct {
	Rows       []GetOwnerResponse `json:"rows"`
	Page       int                `json:"page"`
	PerPage    int                `json:"per_page"`
	Total      int                `json:"total"`
	TotalPages int                `json:"total_pages"`
}

type GetOwnerRequest struct {
	ID string `json:"id"`
}

type GetOwnerResponse struct {
	entity.Base
	Name string `json:"name"`
}

type CreateOwnerRequest struct {
	Name string `json:"name"`
}

type CreateOwnerResponse struct {
	Message string `json:"message"`
}

type UpdateOwnerRequest struct {
	ID   string  `json:"-"`
	Name *string `json:"name"`
}

type UpdateOwnerResponse struct {
	Message string `json:"message"`
}

type DeleteOwnerRequest struct {
	ID string `json:"id"`
}

type DeleteOwnerResponse struct {
	Message string `json:"message"`
}
