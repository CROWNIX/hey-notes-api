package dto

import (
	"hey-notes-api/models"
	"time"
)

type GroupRequest struct {
	UserId   int `json:"user_id" validate:"required"`
	Title    string `json:"title" validate:"required,max=64"`
	IsPublic bool `json:"is_public" validate:"required"`
	Pin      string `json:"pin" validate:"required"`
}

type GroupResponse struct {
	Id        int
	UserId    int
	Title     string
	IsPublic  bool
	Pin       string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func ToGroupModel(request *GroupRequest) *models.Group {
	return &models.Group{
		UserId: request.UserId,
		Title: request.Title,
		IsPublic: request.IsPublic,
		Pin: request.Pin,
	}
}

func ToGroupResponse(group *models.Group) *GroupResponse {
	return &GroupResponse{
		Id: group.Id,
		Title: group.Title,
		UserId: group.UserId,
		IsPublic: group.IsPublic,
		Pin: group.Pin,
		CreatedAt: group.CreatedAt,
		UpdatedAt: group.UpdatedAt,
	}
}