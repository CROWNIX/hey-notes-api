package dto

import (
	"hey-notes-api/models"
	"time"
)

type NoteRequest struct {
	Title string `json:"title" validate:"required,max=64"`
	Body  string `json:"body" validate:"required"`
}

type NoteResponse struct {
	Id int 
	Title string 
	Slug string 
	Body  string 
	Archived  bool 
	CreatedAt time.Time
	UpdatedAt time.Time
}

func ToNoteModel(request *NoteRequest, slug string) *models.Note {
	return &models.Note{
		Title: request.Title,
		Slug: slug,
		Body: request.Body,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
func ToNoteResponse(note *models.Note) *NoteResponse {
	return &NoteResponse{
		Id: note.Id,
		Title: note.Title,
		Slug: note.Slug,
		Body: note.Body,
		Archived: note.Archived,
		CreatedAt: note.CreatedAt,
		UpdatedAt: note.UpdatedAt,
	}
}