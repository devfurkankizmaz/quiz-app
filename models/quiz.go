package models

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type Quiz struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateQuizParams struct {
	ID     uuid.UUID
	UserID uuid.UUID
	Title  string
}

type ListQuizzesParams struct {
	Limit  int32
	Offset int32
}

type UpdateQuizParams struct {
	ID    uuid.UUID
	Title string
}

type QuizRepository interface {
	CreateQuiz(ctx context.Context, arg CreateQuizParams) (Quiz, error)
	DeleteQuizByID(ctx context.Context, id uuid.UUID) error
	UpdateQuizByID(ctx context.Context, arg UpdateQuizParams) (Quiz, error)
	GetQuizByID(ctx context.Context, id uuid.UUID) (Quiz, error)
	ListQuizzes(ctx context.Context, arg ListQuizzesParams) ([]Quiz, error)
}

type QuizService interface {
	CreateQuiz(ctx context.Context, arg CreateQuizParams) (Quiz, error)
}
