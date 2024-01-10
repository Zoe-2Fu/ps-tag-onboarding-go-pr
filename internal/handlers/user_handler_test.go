package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	errs "github.com/Zoe-2Fu/ps-tag-onboarding-go-pr/internal/constants"
	models "github.com/Zoe-2Fu/ps-tag-onboarding-go-pr/internal/data"
	repo "github.com/Zoe-2Fu/ps-tag-onboarding-go-pr/internal/repository"
	validator "github.com/Zoe-2Fu/ps-tag-onboarding-go-pr/internal/validators"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestSave_StatusCreated(t *testing.T) {
	// arrange
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/save", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	userRepoMock := new(repo.UserRepoMock)
	validatorMock := new(validator.UserValidatorMock)

	userHandler := &UserHandler{
		userRepo:  userRepoMock,
		validator: validatorMock,
	}

	userRepoMock.On("Save", mock.Anything, mock.Anything).Return(primitive.NewObjectID(), nil)
	validatorMock.On("ValidateUserDetails", mock.Anything, mock.Anything).Return((*errs.ErrorMessage)(nil))

	// act
	err := userHandler.Save(c)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
}

func TestFind_StatusOK(t *testing.T) {
	// arrange
	e := echo.New()

	userRepoMock := new(repo.UserRepoMock)
	validatorMock := new(validator.UserValidatorMock)

	userHandler := &UserHandler{
		userRepo:  userRepoMock,
		validator: validatorMock,
	}

	userID := primitive.NewObjectID()
	expectedUser := models.NewUser(userID, "John", "Doe", "a@a.a", 20)
	expectedReponseBody, _ := json.Marshal(expectedUser)

	userRepoMock.On("Find", mock.Anything, mock.Anything).Return(expectedUser, nil)
	validatorMock.On("ValidateUserDetails", mock.Anything, mock.Anything).Return((*errs.ErrorMessage)(nil))

	req := httptest.NewRequest(http.MethodPost, "/save/"+userID.Hex(), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// act
	err := userHandler.Find(c)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, string(expectedReponseBody), rec.Body.String())
}

func TestFind_StatusNotFound(t *testing.T) {
	// arrange
	e := echo.New()

	userRepoMock := new(repo.UserRepoMock)
	validatorMock := new(validator.UserValidatorMock)

	userHandler := &UserHandler{
		userRepo:  userRepoMock,
		validator: validatorMock,
	}

	userID := "123"
	err := echo.NewHTTPError(http.StatusNotFound, errs.ErrorMessage{
		Error:   errs.ErrorStatusNotFound,
		Details: []string{"User not found"},
	})
	userRepoMock.On("Find", mock.Anything, mock.Anything).Return(models.User{}, err)

	req := httptest.NewRequest(http.MethodPost, "/find/"+userID, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// act
	handlerErr := userHandler.Find(c)

	// assert
	assert.Error(t, handlerErr)
}
