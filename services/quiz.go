package services

import (
	"context"
	"github.com/devfurkankizmaz/quiz-app/models"
)

type quizService struct {
	quizRepository models.QuizRepository
}

func NewQuizService(quizRepository models.QuizRepository) models.QuizService {
	return &quizService{quizRepository: quizRepository}
}

func (s *quizService) CreateQuiz(ctx context.Context, arg models.CreateQuizParams) (models.Quiz, error) {
	quiz, err := s.quizRepository.CreateQuiz(ctx, arg)
	if err != nil {
		return models.Quiz{}, err
	}
	return quiz, nil
}
