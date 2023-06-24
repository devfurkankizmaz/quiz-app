package handlers

import (
	"context"
	"github.com/devfurkankizmaz/quiz-app/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

type QuizHandler struct {
	quizService models.QuizService
}

func NewQuizHandler(quizService models.QuizService) *QuizHandler {
	return &QuizHandler{quizService: quizService}
}

func (h *QuizHandler) GetSmt(c echo.Context) error {
	return c.String(http.StatusOK, "Some Text")
}

func (h *QuizHandler) Create(c echo.Context) error {
	var payload models.Quiz
	err := c.Bind(&payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "fail", "message": err.Error()})
	}
	quizParam := models.CreateQuizParams{
		ID:    uuid.New(),
		Title: payload.Title,
	}
	_, err = h.quizService.CreateQuiz(context.Background(), quizParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "fail", "message": err.Error()})
	}
	return nil
}
