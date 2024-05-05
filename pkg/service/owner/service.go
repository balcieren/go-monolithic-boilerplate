package owner

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
	ErrOwnerNotFound     = errors.New("owner not found")
	ErrOwnerNameRequired = errors.New("owner name is required")
	ErrOwnerListFailed   = errors.New("failed to list owners")
	ErrOwnerCreateFailed = errors.New("failed to create owner")
	ErrOwnerUpdateFailed = errors.New("failed to update owner")
	ErrOwnerDeleteFailed = errors.New("failed to delete owner")
)

type Servicer interface {
	ListOwners(ctx context.Context, req *ListOwnersRequest) (*ListOwnersResponse, error)
	GetOwner(ctx context.Context, req *GetOwnerRequest) (*GetOwnerResponse, error)
	CreateOwner(ctx context.Context, req *CreateOwnerRequest) (*CreateOwnerResponse, error)
	UpdateOwner(ctx context.Context, req *UpdateOwnerRequest) (*UpdateOwnerResponse, error)
	DeleteOwner(ctx context.Context, req *DeleteOwnerRequest) (*DeleteOwnerResponse, error)
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

func (s *Service) ListOwners(ctx context.Context, req *ListOwnersRequest) (*ListOwnersResponse, error) {
	oq := s.query.Owner

	owners, count, err := oq.WithContext(ctx).FindByPage((req.Page-1)*req.PerPage, req.PerPage)
	if err != nil {
		return nil, fail.New(fiber.StatusNotFound, err.Error())
	}

	rows := make([]GetOwnerResponse, 0)
	for _, owner := range owners {
		resp := GetOwnerResponse{
			Base: owner.Base,
			Name: owner.Name,
		}

		rows = append(rows, resp)
	}

	return &ListOwnersResponse{
		Rows:       rows,
		Page:       req.Page,
		PerPage:    req.PerPage,
		Total:      int(count),
		TotalPages: (int(count) / req.PerPage) + 1,
	}, nil

}

func (s *Service) GetOwner(ctx context.Context, req *GetOwnerRequest) (*GetOwnerResponse, error) {
	id, err := uuid.Parse(req.ID)
	if err != nil {
		return nil, fail.New(fiber.StatusBadRequest, err.Error())
	}

	oq := s.query.Owner

	owner, err := oq.WithContext(ctx).Where(oq.ID.Eq(id)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fail.New(fiber.StatusNotFound, err.Error())
		}
		return nil, fail.New(fiber.StatusInternalServerError, err.Error())
	}

	return &GetOwnerResponse{
		Base: owner.Base,
		Name: owner.Name,
	}, nil
}

func (s *Service) CreateOwner(ctx context.Context, req *CreateOwnerRequest) (*CreateOwnerResponse, error) {
	if len(req.Name) == 0 {
		return nil, fail.New(fiber.StatusBadRequest, "name is required")
	}

	owner := entity.Owner{
		Name: req.Name,
	}

	oq := s.query.Owner

	if err := oq.WithContext(ctx).Create(&owner); err != nil {
		return nil, fail.New(fiber.StatusInternalServerError, err.Error())
	}

	return &CreateOwnerResponse{
		Message: "owner created",
	}, nil
}

func (s *Service) UpdateOwner(ctx context.Context, req *UpdateOwnerRequest) (*UpdateOwnerResponse, error) {
	id, err := uuid.Parse(req.ID)
	if err != nil {
		return nil, fail.New(fiber.StatusBadRequest, err.Error())
	}

	fields := make([]field.AssignExpr, 0)

	if req.Name != nil {
		fields = append(fields, s.query.Owner.Name.Value(*req.Name))
	}

	oq := s.query.Owner

	_, err = oq.WithContext(ctx).Where(oq.ID.Eq(id)).UpdateSimple(fields...)
	if err != nil {
		return nil, fail.New(fiber.StatusInternalServerError, err.Error())
	}

	return &UpdateOwnerResponse{
		Message: "owner updated",
	}, nil
}

func (s *Service) DeleteOwner(ctx context.Context, req *DeleteOwnerRequest) (*DeleteOwnerResponse, error) {
	id, err := uuid.Parse(req.ID)
	if err != nil {
		return nil, fail.New(fiber.StatusBadRequest, err.Error())
	}

	if err := s.query.Transaction(func(tx *query.Query) error {
		oq := tx.Owner
		pq := tx.Pet

		_, err = oq.WithContext(ctx).Where(oq.ID.Eq(id)).Delete()
		if err != nil {
			return fail.New(fiber.StatusInternalServerError, err.Error())
		}

		_, err = pq.WithContext(ctx).Where(pq.OwnerID.Eq(id)).UpdateSimple(pq.OwnerID.Value(nil))
		if err != nil {
			return fail.New(fiber.StatusInternalServerError, err.Error())
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &DeleteOwnerResponse{
		Message: "owner deleted",
	}, nil
}
