package user

import (
	"context"
	"database/sql"
	globalHelper "hey-notes-api/helpers"
	"hey-notes-api/models"
)

type UserRepository interface {
	Create(ctx context.Context, tx *sql.Tx, user models.User) (*models.User, error)
	EmailExist(ctx context.Context, db *sql.DB, email string) bool
	FindByEmail(ctx context.Context, db *sql.DB, email string) (*models.User, error)
}

type UserRepositoryImpl struct {
}

func NewUserRepositoryImpl() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, user models.User) (*models.User, error) {
	SQL := "insert into users(username, email, password, created_at, updated_at) values (?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, user.Username, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	user.Id = int(id)

	return &user, nil
}

func (repository *UserRepositoryImpl) EmailExist(ctx context.Context, db *sql.DB, email string) bool {
	SQL := "SELECT * FROM users WHERE email = ?"
	rows, err := db.QueryContext(ctx, SQL, email)
	globalHelper.PanicIfError(err)
	defer rows.Close()

	return rows.Next()
}

func (repository *UserRepositoryImpl) FindByEmail(ctx context.Context, db *sql.DB, email string) (*models.User, error) {
	SQL := "SELECT id, username, email, created_at, updated_at FROM users WHERE email = ? LIMIT 1"
	row, err := db.QueryContext(ctx, SQL, email)

	if err != nil {
		return nil, err
	}

	if !row.Next() {
		return nil, nil
	}

	var user models.User

	row.Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	return &user, nil
}