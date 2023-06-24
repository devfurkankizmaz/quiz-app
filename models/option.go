package models

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type Option struct {
	ID         uuid.UUID `json:"id"`
	QuestionID uuid.UUID `json:"question_id"`
	Content    string    `json:"content"`
	IsCorrect  bool      `json:"is_correct"`
	CreatedAt  time.Time `json:"created_at"`
}

type CreateOptionParams struct {
	ID         uuid.UUID
	QuestionID uuid.UUID
	Content    string
	IsCorrect  bool
}

type ListOptionsParams struct {
	Limit  int32
	Offset int32
}

type UpdateOptionParams struct {
	ID        uuid.UUID
	Content   string
	IsCorrect bool
}

type OptionRepository interface {
	CreateOption(ctx context.Context, arg CreateOptionParams) (Option, error)
	DeleteOptionByID(ctx context.Context, id uuid.UUID) error
	UpdateOptionByID(ctx context.Context, arg UpdateOptionParams) (Option, error)
	GetOptionByID(ctx context.Context, id uuid.UUID) (Option, error)
	ListOptions(ctx context.Context, arg ListOptionsParams) ([]Option, error)
}

type OptionService interface {
	CreateOption(ctx context.Context, arg CreateOptionParams) (Option, error)
}
