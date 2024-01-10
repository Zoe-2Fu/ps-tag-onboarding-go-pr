package main

import (
	handler "github.com/Zoe-2Fu/ps-tag-onboarding-go-pr/internal/handlers"
	repo "github.com/Zoe-2Fu/ps-tag-onboarding-go-pr/internal/repository"
	"github.com/Zoe-2Fu/ps-tag-onboarding-go-pr/internal/routes"
	validator "github.com/Zoe-2Fu/ps-tag-onboarding-go-pr/internal/validators"
	"github.com/labstack/echo/v4"
)

const userDB = "user"

func main() {
	e := echo.New()

	mongoClient := repo.ConnectMongoDB()
	db := mongoClient.NewMongoDB(userDB)

	userRepo := repo.NewUserRepo(db)
	userValidator := validator.NewUserValidator(userRepo)
	userHandler := handler.NewUserHandler(userRepo, userValidator)

	routes.UserRoute(e, *userHandler)
	e.Use(handler.HandleError)

	e.Logger.Fatal(e.Start(":6000"))
}
