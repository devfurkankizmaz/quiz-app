package models

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type Question struct {
	ID        uuid.UUID `json:"id"`
	QuizID    uuid.UUID `json:"quiz_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateQuestionParams struct {
	ID      uuid.UUID
	QuizID  uuid.UUID
	Content string
}

type ListQuestionsParams struct {
	Limit  int32
	Offset int32
}

type UpdateQuestionParams struct {
	ID      uuid.UUID
	Content string
}

type QuestionRepository interface {
	CreateQuestion(ctx context.Context, arg CreateQuestionParams) (Question, error)
	DeleteQuestionByID(ctx context.Context, id uuid.UUID) error
	GetQuestionByID(ctx context.Context, id uuid.UUID) (Question, error)
	UpdateQuestionByID(ctx context.Context, arg UpdateQuestionParams) (Question, error)
	ListQuestions(ctx context.Context, arg ListQuestionsParams) ([]Question, error)
}

type QuestionService interface {
	CreateQuestion(ctx context.Context, arg CreateQuestionParams) (Question, error)
}
