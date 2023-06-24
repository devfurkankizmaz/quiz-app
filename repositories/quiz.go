package repositories

import (
	"context"
	"github.com/devfurkankizmaz/quiz-app/models"
	"github.com/google/uuid"
)

type quizRepository struct {
	db DBTX
}

func NewQuizRepository(db DBTX) models.QuizRepository {
	return &quizRepository{db}
}

func (r *quizRepository) CreateQuiz(ctx context.Context, arg models.CreateQuizParams) (models.Quiz, error) {
	query := `INSERT INTO quizzes (id, user_id, title) VALUES ($1, $2, $3) RETURNING id, user_id, title`
	row := r.db.QueryRow(ctx, query, arg.ID, arg.UserID, arg.Title)
	var i models.Quiz
	err := row.Scan(&i.ID, &i.UserID, &i.Title, &i.CreatedAt)
	if err != nil {
		return models.Quiz{}, err
	}
	return i, nil
}

func (r *quizRepository) DeleteQuizByID(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM quizzes WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}

func (r *quizRepository) UpdateQuizByID(ctx context.Context, arg models.UpdateQuizParams) (models.Quiz, error) {
	query := `UPDATE quizzes SET title = $2 WHERE id = $1 RETURNING id, title, created_at`
	row := r.db.QueryRow(ctx, query, arg.ID, arg.Title)
	var i models.Quiz
	err := row.Scan(&i.ID, &i.Title, &i.CreatedAt)
	return i, err
}

func (r *quizRepository) GetQuizByID(ctx context.Context, id uuid.UUID) (models.Quiz, error) {
	query := `SELECT id, title, created_at FROM quizzes WHERE id = $1 LIMIT 1`
	row := r.db.QueryRow(ctx, query, id)
	var i models.Quiz
	err := row.Scan(&i.ID, &i.Title, &i.CreatedAt)
	return i, err
}

func (r *quizRepository) ListQuizzes(ctx context.Context, arg models.ListQuizzesParams) ([]models.Quiz, error) {
	query := `SELECT id, title, created_at FROM quizzes ORDER BY created_at LIMIT $1 OFFSET $2`
	rows, err := r.db.Query(ctx, query, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.Quiz
	for rows.Next() {
		var i models.Quiz
		if err := rows.Scan(&i.ID, &i.Title, &i.CreatedAt); err != nil {
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
