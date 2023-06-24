package services

import (
	"context"
	"github.com/devfurkankizmaz/quiz-app/models"
)

type questionService struct {
	questionRepository models.QuestionRepository
}

func NewQuestionService(questionRepository models.QuestionRepository) models.QuestionService {
	return &questionService{questionRepository: questionRepository}
}

func (s *questionService) CreateQuestion(ctx context.Context, arg models.CreateQuestionParams) (models.Question, error) {
	question, err := s.questionRepository.CreateQuestion(ctx, arg)
	if err != nil {
		return models.Question{}, err
	}
	return question, nil
}
