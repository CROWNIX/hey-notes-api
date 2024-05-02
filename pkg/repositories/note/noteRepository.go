package note

import (
	"context"
	"database/sql"
	"hey-notes-api/models"
	"hey-notes-api/pkg/dto"
)

type NoteRepository interface {
	GetAllNotes(ctx context.Context, db *sql.DB) (*[]dto.NoteResponse, error)
	GetArchivedNotes(ctx context.Context, db *sql.DB) (*[]dto.NoteResponse, error)
	FindBySlug(ctx context.Context, db *sql.DB, slug string) (*models.Note, error)
	Create(ctx context.Context, tx *sql.Tx, note models.Note) (*models.Note, error)
	Archived(ctx context.Context, db *sql.DB, slug string) (bool, error)
	Unarchived(ctx context.Context, db *sql.DB, slug string) (bool, error)
	Delete(ctx context.Context, db *sql.DB, slug string) (bool, error)
}

type NoteRepositoryImpl struct {
}

func NewNoteRepositoryImpl() NoteRepository {
	return &NoteRepositoryImpl{}
}

func (repository *NoteRepositoryImpl) GetAllNotes(ctx context.Context, db *sql.DB) (*[]dto.NoteResponse, error) {
	SQL := "SELECT * FROM notes"
	rows, err := db.QueryContext(ctx, SQL)
	
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var notes []dto.NoteResponse

	for rows.Next() {
		note := models.Note{}
		err := rows.Scan(&note.Id, &note.Title, &note.Slug, &note.Body, &note.Archived, &note.CreatedAt, &note.UpdatedAt)
		
		if err != nil {
			return nil, err
		}

		notes = append(notes, *dto.ToNoteResponse(&note))
	}

	return &notes, nil
}

func (repository *NoteRepositoryImpl) GetArchivedNotes(ctx context.Context, db *sql.DB) (*[]dto.NoteResponse, error) {
	SQL := "SELECT * FROM notes WHERE archived = ?"
	rows, err := db.QueryContext(ctx, SQL, true)
	
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var notes []dto.NoteResponse

	for rows.Next() {
		note := models.Note{}
		err := rows.Scan(&note.Id, &note.Title, &note.Slug, &note.Body, &note.Archived, &note.CreatedAt, &note.UpdatedAt)
		
		if err != nil {
			return nil, err
		}

		notes = append(notes, *dto.ToNoteResponse(&note))
	}

	return &notes, nil
}

func (repository *NoteRepositoryImpl) FindBySlug(ctx context.Context, db *sql.DB, slug string) (*models.Note, error) {
	SQL := "SELECT * FROM notes WHERE slug = ? LIMIT 1"
	row, err := db.QueryContext(ctx, SQL, slug)

	if err != nil {
		return nil, err
	}

	defer row.Close()

	if !row.Next() {
		return nil, nil
	}

	var note models.Note

	row.Scan(
		&note.Id,
		&note.Title,
		&note.Slug,
		&note.Body,
		&note.Archived,
		&note.CreatedAt,
		&note.UpdatedAt,
	)
	return &note, nil
}

func (repository *NoteRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, note models.Note) (*models.Note, error) {
	SQL := "INSERT INTO notes(title, slug, body, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, note.Title, note.Slug, note.Body, note.CreatedAt, note.UpdatedAt)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	note.Id = int(id)

	return &note, nil
}

func (repository *NoteRepositoryImpl) Archived(ctx context.Context, db *sql.DB, slug string) (bool, error) {

	SQL := "UPDATE notes SET archived = ? WHERE slug = ?"
	_, err := db.ExecContext(ctx, SQL, true, slug)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (repository *NoteRepositoryImpl) Unarchived(ctx context.Context, db *sql.DB, slug string) (bool, error) {

	SQL := "UPDATE notes SET archived = ? WHERE slug = ?"
	_, err := db.ExecContext(ctx, SQL, false, slug)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (repository *NoteRepositoryImpl) Delete(ctx context.Context, db *sql.DB, slug string) (bool, error) {

	SQL := "DELETE FROM notes WHERE slug = ?"
	_, err := db.ExecContext(ctx, SQL, slug)

	if err != nil {
		return false, err
	}

	return true, nil
}


