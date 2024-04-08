package dto

type NoteRequest struct {
	Title string `json:"title" validate:"required,max=64"`
	Body  string `json:"body" validate:"required"`
}