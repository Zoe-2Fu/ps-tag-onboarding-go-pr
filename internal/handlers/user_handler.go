package handler

import (
	"context"
	"github.com/Zoe-2Fu/ps-tag-onboarding-go-pr/internal/interfaces"
	"net/http"
	"time"

	errs "github.com/Zoe-2Fu/ps-tag-onboarding-go-pr/internal/constants"
	models "github.com/Zoe-2Fu/ps-tag-onboarding-go-pr/internal/data"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
	UserRepo  interfaces.UserRepo
	Validator interfaces.UserValidator
}

func (h *UserHandler) Find(c echo.Context) error {
	id := c.Param("id")
	var user models.User

	user, err := h.UserRepo.Find(c, id)
	if err != nil {
		errMsg := errs.NewErrorMessage(errs.ResponseUserNotFound, "User not found")
		return echo.NewHTTPError(http.StatusBadRequest, errMsg)
	}
	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) Save(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user := new(models.User)
	if err := c.Bind(user); err != nil {
		errMsg := errs.NewErrorMessage(errs.ErrorBadRequest, "Missing some user details/Invalid input format")
		return echo.NewHTTPError(http.StatusBadRequest, errMsg)
	}

	if validationErr := h.Validator.ValidateUserDetails(*user); validationErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, validationErr)
	}

	insertedID, err := h.UserRepo.Save(ctx, *user)
	if err != nil {
		return err
	}

	if insertedID == primitive.NilObjectID {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to save user")
	}

	user.ID = insertedID

	return c.JSON(http.StatusCreated, user)
}
