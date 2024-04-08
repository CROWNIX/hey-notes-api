package note

import (
	"context"
	"database/sql"
	"time"

	"hey-notes-api/database"
	globalHelper "hey-notes-api/helpers"
	"hey-notes-api/internal/api/http/exception"
	"hey-notes-api/models"
	"hey-notes-api/pkg/dto"
	"hey-notes-api/pkg/repositories/note"

	"github.com/go-playground/validator/v10"
)

type NoteService interface {
	GetAllNotes(ctx context.Context) (*[]models.Note, error)
	GetArchivedNotes(ctx context.Context) (*[]models.Note, error)
	Create(ctx context.Context, req *dto.NoteRequest) (*models.Note, error)
	Archived(ctx context.Context, slug string) (bool, error)
	Unarchived(ctx context.Context, slug string) (bool, error)
	Delete(ctx context.Context, slug string) (bool, error)
}

type NoteServiceImpl struct {
	NoteRepo   note.NoteRepository
	DbImpl     *database.DbImpl
	Validation *validator.Validate
}

func NewNoteServiceImpl(
	userRepo note.NoteRepository,
	dbImpl *database.DbImpl,
	validation *validator.Validate,
) *NoteServiceImpl {
	return &NoteServiceImpl{
		NoteRepo:   userRepo,
		DbImpl:     dbImpl,
		Validation: validation,
	}
}

func (service *NoteServiceImpl) GetAllNotes(ctx context.Context) (*[]models.Note, error) {
	notes, err := service.NoteRepo.GetAllNotes(ctx, service.DbImpl.DB)

	if err != nil {
		return nil, &exception.InternalServer{Message: err.Error()}
	}

	if len(*notes) == 0 {
		return nil, &exception.NotFound{Message: "Notes Not Found"}
	}

	return notes, nil
}

func (service *NoteServiceImpl) GetArchivedNotes(ctx context.Context) (*[]models.Note, error) {
	notes, err := service.NoteRepo.GetArchivedNotes(ctx, service.DbImpl.DB)

	if err != nil {
		return nil, &exception.InternalServer{Message: err.Error()}
	}

	if len(*notes) == 0 {
		return nil, &exception.NotFound{Message: "Notes Not Found"}
	}

	return notes, nil
}

func (service *NoteServiceImpl) Create(ctx context.Context, request *dto.NoteRequest) (*models.Note, error) {
	err := service.Validation.Struct(request)
	if err != nil {
		return nil, &exception.BadRequest{Message: err.Error()}
	}

	slug := globalHelper.GenerateSlug(ctx, service.DbImpl.DB, "notes", request.Title)

	var noteEntity *models.Note

	err = service.DbImpl.RunWithTransaction(ctx, &sql.TxOptions{ReadOnly: false}, func(tx *sql.Tx) error {
		result, err := service.NoteRepo.Create(ctx, tx, models.Note{Title: request.Title, Slug: slug, Body: request.Body, CreatedAt: time.Now(), UpdatedAt: time.Now()})
		if err != nil {
			return err
		}

		noteEntity = result
		return nil
	})

	if err != nil {
		return nil, &exception.InternalServer{Message: err.Error()}
	}

	return noteEntity, nil
}

func (service *NoteServiceImpl) Archived(ctx context.Context, slug string) (bool, error) {
	note, err := service.NoteRepo.FindBySlug(ctx, service.DbImpl.DB, slug)

	if err != nil {
		return false, &exception.InternalServer{Message: err.Error()}
	}

	if note == nil {
		return false, &exception.NotFound{Message: "Note Not Found"}
	}

	result, err := service.NoteRepo.Archived(ctx, service.DbImpl.DB, slug)

	if err != nil {
		return false, &exception.InternalServer{Message: err.Error()}
	}

	return result, nil
}

func (service *NoteServiceImpl) Unarchived(ctx context.Context, slug string) (bool, error) {
	note, err := service.NoteRepo.FindBySlug(ctx, service.DbImpl.DB, slug)

	if err != nil {
		return false, &exception.InternalServer{Message: err.Error()}
	}

	if note == nil {
		return false, &exception.NotFound{Message: "Note Not Found"}
	}

	result, err := service.NoteRepo.Unarchived(ctx, service.DbImpl.DB, slug)

	if err != nil {
		return false, &exception.InternalServer{Message: err.Error()}
	}

	return result, nil
}

func (service *NoteServiceImpl) Delete(ctx context.Context, slug string) (bool, error) {
	note, err := service.NoteRepo.FindBySlug(ctx, service.DbImpl.DB, slug)

	if err != nil {
		return false, &exception.InternalServer{Message: err.Error()}
	}

	if note == nil {
		return false, &exception.NotFound{Message: "Note Not Found"}
	}

	result, err := service.NoteRepo.Delete(ctx, service.DbImpl.DB, slug)

	if err != nil {
		return false, &exception.InternalServer{Message: err.Error()}
	}

	return result, nil
}
