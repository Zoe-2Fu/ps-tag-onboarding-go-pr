package main

import (
	models "github.com/Zoe-2Fu/ps-tag-onboarding-go-pr/internal/data"
	handler "github.com/Zoe-2Fu/ps-tag-onboarding-go-pr/internal/handlers"
	"github.com/Zoe-2Fu/ps-tag-onboarding-go-pr/internal/mongo"
	repo "github.com/Zoe-2Fu/ps-tag-onboarding-go-pr/internal/repository"
	"github.com/Zoe-2Fu/ps-tag-onboarding-go-pr/internal/routes"
	validator "github.com/Zoe-2Fu/ps-tag-onboarding-go-pr/internal/validators"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db := mongo.ConnectMongoDB(models.UserDB)

	userRepo := repo.NewUserRepo(db)
	userValidator := validator.NewUserValidator(userRepo)
	userHandler := handler.UserHandler{UserRepo: userRepo, Validator: userValidator}

	routes.UserRoute(e, userHandler)

	e.Logger.Fatal(e.Start(":6000"))
}
