package group

import (
	"context"
	"database/sql"

	"hey-notes-api/database"
	"hey-notes-api/internal/api/http/exception"
	"hey-notes-api/models"
	"hey-notes-api/pkg/dto"
	"hey-notes-api/pkg/repositories/group"

	"github.com/go-playground/validator/v10"
)

type GroupService interface {
	Create(ctx context.Context, req *dto.GroupRequest) (*dto.GroupResponse, error)
}

type GroupServiceImpl struct {
	GroupRepo   group.GroupRepository
	DbImpl     *database.DbImpl
	Validation *validator.Validate
}

func NewGroupServiceImpl(
	groupRepo group.GroupRepository,
	dbImpl *database.DbImpl,
	validation *validator.Validate,
) *GroupServiceImpl {
	return &GroupServiceImpl{
		GroupRepo:   groupRepo,
		DbImpl:     dbImpl,
		Validation: validation,
	}
}

func (service *GroupServiceImpl) Create(ctx context.Context, request *dto.GroupRequest) (*dto.GroupResponse, error) {
	err := service.Validation.Struct(request)
	if err != nil {
		return nil, &exception.BadRequest{Message: err.Error()}
	}

	var groupEntity *models.Group

	err = service.DbImpl.RunWithTransaction(ctx, &sql.TxOptions{ReadOnly: false}, func(tx *sql.Tx) error {
		result, err := service.GroupRepo.Create(ctx, tx, *dto.ToGroupModel(request))
		if err != nil {
			return err
		}

		groupEntity = result
		return nil
	})

	if err != nil {
		return nil, &exception.InternalServer{Message: err.Error()}
	}

	return dto.ToGroupResponse(groupEntity), nil
}
