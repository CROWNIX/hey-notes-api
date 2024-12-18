package auth

import (
	"context"
	"database/sql"
	"time"

	"hey-notes-api/database"
	"hey-notes-api/helpers"
	"hey-notes-api/models"
	"hey-notes-api/pkg/dto"
	"hey-notes-api/pkg/repositories/user"

	"hey-notes-api/internal/api/http/exception"

	"github.com/go-playground/validator/v10"
)

type AuthService interface {
	Register(ctx context.Context, req *dto.RegisterRequest) (*models.User, error)
	Login(ctx context.Context, request *dto.LoginRequest) (*dto.LoginResponse, error)
}

type AuthServiceImpl struct {
	UserRepo   user.UserRepository
	DbImpl     *database.DbImpl
	Validation *validator.Validate
}

func NewAuthServiceImpl(
	userRepo user.UserRepository,
	dbImpl *database.DbImpl,
	validation *validator.Validate,
) *AuthServiceImpl {
	return &AuthServiceImpl{
		UserRepo:   userRepo,
		DbImpl:     dbImpl,
		Validation: validation,
	}
}

func (service *AuthServiceImpl) Register(ctx context.Context, request *dto.RegisterRequest) (*models.User, error) {
	err := service.Validation.Struct(request)
	if err != nil {
		return nil, &exception.BadRequest{Message: err.Error()}
	}

	if emailExists := service.UserRepo.EmailExist(ctx, service.DbImpl.DB, request.Email); emailExists {
		return nil, &exception.BadRequest{Message: "email already registered"}
	}

	passwordHash, err := helpers.HashPassword(request.Password)
	if err != nil {
		return nil, &exception.InternalServer{Message: err.Error()}
	}

	var userEntity *models.User

	err = service.DbImpl.RunWithTransaction(ctx, &sql.TxOptions{ReadOnly: false}, func(tx *sql.Tx) error {
		result, err := service.UserRepo.Create(ctx, tx, models.User{Username: request.Username, Email: request.Email, Password: passwordHash, CreatedAt: time.Now(), UpdatedAt: time.Now()})
		if err != nil {
			return err
		}

		userEntity = result
		return nil
	})

	if err != nil {
		return nil, &exception.InternalServer{Message: err.Error()}
	}

	return userEntity, nil
}

func (service *AuthServiceImpl) Login(ctx context.Context, request *dto.LoginRequest) (*dto.LoginResponse, error) {
	err := service.Validation.Struct(request)
	if err != nil {
		return nil, &exception.BadRequest{Message: err.Error()}
	}

	user, err := service.UserRepo.FindByEmail(ctx, service.DbImpl.DB, request.Email)
	if err != nil {
		return nil, &exception.InternalServer{Message: err.Error()}
	}

	if user == nil {
		return nil, &exception.NotFound{Message: "wrong email or password"}
	}

	if err := helpers.VerifyPassword(user.Password, request.Password); err != nil {
		return nil, &exception.NotFound{Message: err.Error()}
	}

	token, err := helpers.GenerateToken(user)
	if err != nil {
		return nil, &exception.InternalServer{Message: err.Error()}
	}

	response := dto.LoginResponse{
		User: user,
		Token: token,
	}

	return &response, nil
}
