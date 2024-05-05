package pet

import (
	"context"
	"errors"

	"github.com/balcieren/go-monolithic-boilerplate/pkg/entity"
	"github.com/balcieren/go-monolithic-boilerplate/pkg/fail"
	"github.com/balcieren/go-monolithic-boilerplate/pkg/query"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

var (
	ErrPetNotFound     = errors.New("pet not found")
	ErrPetTypeInvalid  = errors.New("invalid pet type")
	ErrPetNameRequired = errors.New("pet name is required")
	ErrPetListFailed   = errors.New("failed to list pets")
	ErrPetCreateFailed = errors.New("failed to create pet")
	ErrPetUpdateFailed = errors.New("failed to update pet")
	ErrPetDeleteFailed = errors.New("failed to delete pet")
)

type Servicer interface {
	ListPets(ctx context.Context, req *ListPetsRequest) (*ListPetsResponse, error)
	GetPet(ctx context.Context, req *GetPetRequest) (*GetPetResponse, error)
	CreatePet(ctx context.Context, req *CreatePetRequest) (*CreatePetResponse, error)
	UpdatePet(ctx context.Context, req *UpdatePetRequest) (*UpdatePetResponse, error)
	DeletePet(ctx context.Context, req *DeletePetRequest) (*DeletePetResponse, error)
}

var _ Servicer = (*Service)(nil)

type Service struct {
	query *query.Query
}

func NewService(q *query.Query) *Service {
	return &Service{
		query: q,
	}
}

func (s *Service) ListPets(ctx context.Context, req *ListPetsRequest) (*ListPetsResponse, error) {
	pq := s.query.Pet
	query := pq.WithContext(ctx)

	if req.HasOwner != nil {
		if *req.HasOwner {
			query = query.Where(pq.OwnerID.IsNotNull())
		} else {
			query = query.Where(pq.OwnerID.IsNull())
		}
	}

	if req.OwnerID != nil {
		ownerID, err := uuid.Parse(*req.OwnerID)
		if err != nil {
			return nil, fail.New(fiber.StatusBadRequest, err.Error())
		}

		query = query.Where(pq.OwnerID.Eq(ownerID))
	}

	pets, count, err := query.FindByPage(
		(req.Page-1)*req.PerPage,
		req.PerPage,
	)
	if err != nil {
		return nil, fail.New(fiber.StatusNotFound, ErrPetListFailed.Error())
	}

	rows := make([]GetPetResponse, 0)
	if len(pets) > 0 {
		for _, pet := range pets {
			row := GetPetResponse{
				Base: pet.Base,
				Name: pet.Name,
				Type: pet.Type,
			}

			if pet.OwnerID != nil {
				ownerID := pet.OwnerID.String()
				row.OwnerID = &ownerID
			}

			rows = append(rows, row)
		}
	}

	return &ListPetsResponse{
		Rows:       rows,
		Page:       req.Page,
		PerPage:    req.PerPage,
		Total:      int(count),
		TotalPages: (int(count) / req.PerPage) + 1,
	}, nil

}

func (s *Service) GetPet(ctx context.Context, req *GetPetRequest) (*GetPetResponse, error) {
	id, err := uuid.Parse(req.ID)
	if err != nil {
		return nil, fail.New(fiber.StatusBadRequest, err.Error())
	}

	pq := s.query.Pet

	pet, err := pq.WithContext(ctx).Where(pq.ID.Eq(id)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fail.New(fiber.StatusNotFound, ErrPetNotFound.Error())
		}
		return nil, err
	}

	resp := GetPetResponse{
		Base: pet.Base,
		Name: pet.Name,
		Type: pet.Type,
	}

	if pet.OwnerID != nil {
		ownerID := pet.OwnerID.String()
		resp.OwnerID = &ownerID
	}

	return &resp, nil
}

func (s *Service) CreatePet(ctx context.Context, req *CreatePetRequest) (*CreatePetResponse, error) {
	pet := entity.Pet{}

	if len(req.Name) == 0 {
		return nil, fail.New(fiber.StatusBadRequest, ErrPetNameRequired.Error())
	}
	pet.Name = req.Name

	if !entity.PetType(req.Type).IsValid() {
		return nil, fail.New(fiber.StatusBadRequest, ErrPetTypeInvalid.Error())
	}
	pet.Type = entity.PetType(req.Type)

	if req.OwnerID != nil {
		ownerID, err := uuid.Parse(*req.OwnerID)
		if err != nil {
			return nil, fail.New(fiber.StatusBadRequest, err.Error())
		}
		pet.OwnerID = &ownerID
	}

	pq := s.query.Pet

	if err := pq.WithContext(ctx).Create(&pet); err != nil {
		return nil, fail.New(fiber.StatusInternalServerError, ErrPetCreateFailed.Error())
	}

	return &CreatePetResponse{
		Message: "pet created",
	}, nil
}

func (s *Service) UpdatePet(ctx context.Context, req *UpdatePetRequest) (*UpdatePetResponse, error) {
	id, err := uuid.Parse(req.ID)
	if err != nil {
		return nil, fail.New(fiber.StatusBadRequest, err.Error())
	}

	fields := make([]field.AssignExpr, 0)

	if req.Name != nil {
		if len(*req.Name) == 0 {
			return nil, fail.New(fiber.StatusBadRequest, ErrPetNameRequired.Error())
		}

		fields = append(fields, s.query.Pet.Name.Value(*req.Name))
	}

	if req.Type != nil {
		if !entity.PetType(*req.Type).IsValid() {
			return nil, fail.New(fiber.StatusBadRequest, ErrPetTypeInvalid.Error())
		}

		fields = append(fields, s.query.Pet.Type.Value(*req.Type))
	}

	if req.OwnerID != nil {
		ownerID, err := uuid.Parse(*req.OwnerID)
		if err != nil {
			return nil, fail.New(fiber.StatusBadRequest, err.Error())
		}

		fields = append(fields, s.query.Pet.OwnerID.Value(ownerID))
	}

	pq := s.query.Pet

	_, err = pq.WithContext(ctx).Where(pq.ID.Eq(id)).UpdateSimple(fields...)
	if err != nil {
		return nil, fail.New(fiber.StatusInternalServerError, ErrPetUpdateFailed.Error())
	}

	return &UpdatePetResponse{
		Message: "pet updated",
	}, nil
}

func (s *Service) DeletePet(ctx context.Context, req *DeletePetRequest) (*DeletePetResponse, error) {
	id, err := uuid.Parse(req.ID)
	if err != nil {
		return nil, fail.New(fiber.StatusBadRequest, err.Error())
	}

	pq := s.query.Pet

	_, err = pq.WithContext(ctx).Where(pq.ID.Eq(id)).Delete()
	if err != nil {
		return nil, fail.New(fiber.StatusInternalServerError, ErrPetDeleteFailed.Error())
	}

	return &DeletePetResponse{
		Message: "pet deleted",
	}, nil
}
