package repositories

import (
	"context"
	"github.com/devfurkankizmaz/quiz-app/models"
	"github.com/google/uuid"
)

type questionRepository struct {
	db DBTX
}

func NewQuestionRepository(db DBTX) models.QuestionRepository {
	return &questionRepository{db}
}

func (r *questionRepository) CreateQuestion(ctx context.Context, arg models.CreateQuestionParams) (models.Question, error) {
	query := `INSERT INTO questions (id, quiz_id, content) VALUES ($1,$2,$3) RETURNING id, quiz_id, content, created_at`
	row := r.db.QueryRow(ctx, query, arg.ID, arg.QuizID, arg.Content)
	var i models.Question
	err := row.Scan(
		&i.ID,
		&i.QuizID,
		&i.Content,
		&i.CreatedAt,
	)
	return i, err
}

func (r *questionRepository) DeleteQuestionByID(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM questions WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}

func (r *questionRepository) GetQuestionByID(ctx context.Context, id uuid.UUID) (models.Question, error) {
	query := `SELECT id, quiz_id, content, created_at FROM questions WHERE id = $1 LIMIT 1`
	row := r.db.QueryRow(ctx, query, id)
	var i models.Question
	err := row.Scan(
		&i.ID,
		&i.QuizID,
		&i.Content,
		&i.CreatedAt,
	)
	return i, err
}

func (r *questionRepository) UpdateQuestionByID(ctx context.Context, arg models.UpdateQuestionParams) (models.Question, error) {
	query := `UPDATE questions SET content = $2 WHERE id = $1 RETURNING id, quiz_id, content, created_at`
	row := r.db.QueryRow(ctx, query, arg.ID, arg.Content)
	var i models.Question
	err := row.Scan(
		&i.ID,
		&i.QuizID,
		&i.Content,
		&i.CreatedAt,
	)
	return i, err
}

func (r *questionRepository) ListQuestions(ctx context.Context, arg models.ListQuestionsParams) ([]models.Question, error) {
	query := `SELECT id, quiz_id, content, created_at FROM questions ORDER BY created_at LIMIT $1 OFFSET $2`
	rows, err := r.db.Query(ctx, query, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.Question
	for rows.Next() {
		var i models.Question
		if err := rows.Scan(
			&i.ID,
			&i.QuizID,
			&i.Content,
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
