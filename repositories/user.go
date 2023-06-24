package repositories

import (
	"context"
	"github.com/devfurkankizmaz/quiz-app/models"
)

type userRepository struct {
	db DBTX
}

func NewUserRepository(db DBTX) models.UserRepository {
	return &userRepository{db}
}

func (r *userRepository) CreateUser(ctx context.Context, arg models.CreateUserParams) (models.User, error) {
	query := `INSERT INTO users (id, full_name, email, password) 
VALUES ($1, $2, $3, $4) RETURNING id, full_name, email, password`
	row := r.db.QueryRow(ctx, query, arg.ID, arg.FullName, arg.Email, arg.Password)
	var i models.User
	err := row.Scan(&i.ID, &i.FullName, &i.Email, &i.Password, &i.CreatedAt)
	if err != nil {
		return models.User{}, err
	}
	return i, nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	query := `SELECT id, full_name, email, password, created_at FROM users
WHERE email = $1 LIMIT 1`
	row := r.db.QueryRow(ctx, query, email)
	var i models.User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
	)
	return i, err
}
