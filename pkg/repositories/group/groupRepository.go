package group

import (
	"context"
	"database/sql"
	"hey-notes-api/models"
)

type GroupRepository interface {
	Create(ctx context.Context, tx *sql.Tx, group models.Group) (*models.Group, error)
}

type GroupRepositoryImpl struct {
}

func NewGroupRepositoryImpl() GroupRepository {
	return &GroupRepositoryImpl{}
}

func (repository *GroupRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, group models.Group) (*models.Group, error) {
	SQL := "insert into groups(user_id, title, is_public, pin, created_at, updated_at) values (?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, group.UserId, group.Title, group.IsPublic, group.Pin, group.UpdatedAt, group.UpdatedAt)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	group.Id = int(id)

	return &group, nil
}