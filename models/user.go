package models

import (
	"github.com/google/uuid"
	"golang.org/x/net/context"
	"time"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateUserParams struct {
	ID       uuid.UUID
	FullName string
	Email    string
	Password string
}

type UserRepository interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
}

type UserService interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
}
