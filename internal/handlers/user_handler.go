package handler

import (
	"context"
	"net/http"
	"time"

	errs "github.com/Zoe-2Fu/ps-tag-onboarding-go-pr/internal/constants"
	models "github.com/Zoe-2Fu/ps-tag-onboarding-go-pr/internal/data"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userRepo interface {
	Find(ctx echo.Context, id string) (models.User, error)
	Save(ctx context.Context, user models.User) (primitive.ObjectID, error)
}

type userValidator interface {
	ValidateUserDetails(user models.User) *errs.ErrorMessage
}

type UserHandler struct {
	userRepo  userRepo
	validator userValidator
}

func NewUserHandler(repo userRepo, validator userValidator) *UserHandler {
	return &UserHandler{userRepo: repo, validator: validator}
}

func (h *UserHandler) Find(c echo.Context) error {
	id := c.Param("id")
	var user models.User

	user, err := h.userRepo.Find(c, id)
	if err != nil {
		errMsg := errs.NewErrorMessage(errs.ErrorBadRequest, "User not found")
		return echo.NewHTTPError(http.StatusBadRequest, errMsg)
	}
	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) Save(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user := new(models.User)
	if err := c.Bind(user); err != nil {
		errMsg := errs.NewErrorMessage(errs.ErrorBadRequest, "Can't bind values")
		return echo.NewHTTPError(http.StatusBadRequest, errMsg)
	}

	if validationErr := h.validator.ValidateUserDetails(*user); validationErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, validationErr)
	}

	insertedID, err := h.userRepo.Save(ctx, *user)
	if err != nil {
		return err
	}

	if insertedID == primitive.NilObjectID {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to save user")
	}

	user.ID = insertedID

	return c.JSON(http.StatusCreated, user)
}
