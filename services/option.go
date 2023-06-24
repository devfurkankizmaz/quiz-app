package services

import (
	"context"
	"github.com/devfurkankizmaz/quiz-app/models"
)

type optionService struct {
	optionRepository models.OptionRepository
}

func NewOptionService(optionRepository models.OptionRepository) models.OptionService {
	return &optionService{optionRepository: optionRepository}
}

func (s *optionService) CreateOption(ctx context.Context, arg models.CreateOptionParams) (models.Option, error) {
	option, err := s.optionRepository.CreateOption(ctx, arg)
	if err != nil {
		return models.Option{}, err
	}
	return option, nil
}
