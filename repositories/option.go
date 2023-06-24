package repositories

import (
	"context"
	"github.com/devfurkankizmaz/quiz-app/models"
	"github.com/google/uuid"
)

type optionRepository struct {
	db DBTX
}

func NewOptionRepository(db DBTX) models.OptionRepository {
	return &optionRepository{db}
}

func (r *optionRepository) CreateOption(ctx context.Context, arg models.CreateOptionParams) (models.Option, error) {
	query := `INSERT INTO options (id, question_id, content, is_correct) VALUES ($1,$2,$3,$4) RETURNING id, question_id, content, is_correct, created_at`
	row := r.db.QueryRow(ctx, query,
		arg.ID,
		arg.QuestionID,
		arg.Content,
		arg.IsCorrect,
	)
	var i models.Option
	err := row.Scan(
		&i.ID,
		&i.QuestionID,
		&i.Content,
		&i.IsCorrect,
		&i.CreatedAt,
	)
	return i, err
}

func (r *optionRepository) DeleteOptionByID(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM options WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}
func (r *optionRepository) UpdateOptionByID(ctx context.Context, arg models.UpdateOptionParams) (models.Option, error) {
	query := `UPDATE options SET is_correct = $3, content = $2 WHERE id = $1 RETURNING id, content, is_correct, created_at`
	row := r.db.QueryRow(ctx, query, arg.ID, arg.Content, arg.IsCorrect)
	var i models.Option
	err := row.Scan(&i.ID, &i.QuestionID, &i.Content, &i.IsCorrect, &i.CreatedAt)
	return i, err
}

func (r *optionRepository) GetOptionByID(ctx context.Context, id uuid.UUID) (models.Option, error) {
	query := `SELECT id, question_id, content, is_correct, created_at FROM options WHERE id = $1 LIMIT 1`
	row := r.db.QueryRow(ctx, query, id)
	var i models.Option
	err := row.Scan(
		&i.ID,
		&i.QuestionID,
		&i.Content,
		&i.IsCorrect,
		&i.CreatedAt,
	)
	return i, err
}

func (r *optionRepository) ListOptions(ctx context.Context, arg models.ListOptionsParams) ([]models.Option, error) {
	query := `SELECT id, question_id, content, is_correct, created_at FROM options ORDER BY created_at LIMIT $1 OFFSET $2`
	rows, err := r.db.Query(ctx, query, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.Option
	for rows.Next() {
		var i models.Option
		if err := rows.Scan(
			&i.ID,
			&i.QuestionID,
			&i.Content,
			&i.IsCorrect,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
