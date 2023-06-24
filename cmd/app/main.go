package main

import (
	"context"
	"github.com/devfurkankizmaz/quiz-app/api/handlers"
	"github.com/devfurkankizmaz/quiz-app/config"
	"github.com/devfurkankizmaz/quiz-app/repositories"
	"github.com/devfurkankizmaz/quiz-app/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	app := config.App()
	db := app.DB
	err := db.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	quizRepo := repositories.NewQuizRepository(db)
	quizService := services.NewQuizService(quizRepo)
	quizHandler := handlers.NewQuizHandler(quizService)

	e.GET("/test", quizHandler.GetSmt)
	e.Logger.Fatal(e.Start(":1323"))
}
