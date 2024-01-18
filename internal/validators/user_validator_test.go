package validator

import (
	"testing"

	errs "github.com/Zoe-2Fu/ps-tag-onboarding-go-pr/internal/constants"
	models "github.com/Zoe-2Fu/ps-tag-onboarding-go-pr/internal/data"
	repo "github.com/Zoe-2Fu/ps-tag-onboarding-go-pr/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestValidateUserDetails_ValidUserDetails(t *testing.T) {
	user := models.User{ID: primitive.NilObjectID, FirstName: "John", LastName: "Doe", Email: "good@example.com", Age: 25}

	userRepoMock := new(repo.UserRepoMock)
	validator := &UserValidator{
		userRepo: userRepoMock,
	}

	expectedOutput := (*errs.ErrorMessage)(nil)

	userRepoMock.On("ValidateUserExisted", mock.Anything, mock.Anything).Return(false)

	result := validator.ValidateUserDetails(user)

	assert.Nil(t, result, expectedOutput)
}

func TestValidateUserDetails_UserIsExisted(t *testing.T) {
	user := models.User{ID: primitive.NilObjectID, FirstName: "John", LastName: "Doe", Email: "a@a.a", Age: 20}

	userRepoMock := new(repo.UserRepoMock)
	validator := &UserValidator{
		userRepo: userRepoMock,
	}

	expectedOutput := errs.NewErrorMessage(errs.ResponseValidationFailed, errs.ErrorNameUnique)
	expectedOutputPointer := &expectedOutput

	userRepoMock.On("ValidateUserExisted", mock.Anything, mock.Anything).Return(true)

	result := validator.ValidateUserDetails(user)

	assert.Equal(t, expectedOutputPointer, result)
}

func TestValidateUserDetails_UserNameIsMissing(t *testing.T) {
	user := models.User{ID: primitive.NilObjectID, FirstName: "", LastName: "Doe", Email: "a@a.a", Age: 20}

	userRepoMock := new(repo.UserRepoMock)
	validator := &UserValidator{
		userRepo: userRepoMock,
	}

	expectedOutput := errs.NewErrorMessage(errs.ResponseValidationFailed, errs.ErrorNameRequired)
	expectedOutputPointer := &expectedOutput

	userRepoMock.On("ValidateUserExisted", mock.Anything, mock.Anything).Return(false)

	result := validator.ValidateUserDetails(user)

	assert.Equal(t, expectedOutputPointer, result)
}

func TestValidateUserDetails_UserEmailIsMissing(t *testing.T) {
	user := models.User{ID: primitive.NewObjectID(), FirstName: "John", LastName: "Doe", Email: "", Age: 20}

	userRepoMock := new(repo.UserRepoMock)
	validator := &UserValidator{
		userRepo: userRepoMock,
	}

	expectedOutput := errs.NewErrorMessage(errs.ResponseValidationFailed, errs.ErrorEmailRequired)
	expectedOutputPointer := &expectedOutput

	userRepoMock.On("ValidateUserExisted", mock.Anything, mock.Anything).Return(false)

	result := validator.ValidateUserDetails(user)

	assert.Equal(t, expectedOutputPointer, result)
}

func TestValidateUserDetails_InvalidUserEmailFormat(t *testing.T) {
	user := models.User{ID: primitive.NewObjectID(), FirstName: "John", LastName: "Doe", Email: "aa.a", Age: 20}

	userRepoMock := new(repo.UserRepoMock)
	validator := &UserValidator{
		userRepo: userRepoMock,
	}

	expectedOutput := errs.NewErrorMessage(errs.ResponseValidationFailed, errs.ErrorEmailFormat)
	expectedOutputPointer := &expectedOutput

	userRepoMock.On("ValidateUserExisted", mock.Anything, mock.Anything).Return(false)

	result := validator.ValidateUserDetails(user)

	assert.Equal(t, expectedOutputPointer, result)
}

func TestValidateUserDetails_InvalidUserAge(t *testing.T) {
	user := models.User{ID: primitive.NilObjectID, FirstName: "John", LastName: "Doe", Email: "aa.a", Age: 16}

	userRepoMock := new(repo.UserRepoMock)
	validator := &UserValidator{
		userRepo: userRepoMock,
	}

	expectedOutput := errs.NewErrorMessage(errs.ResponseValidationFailed, errs.ErrorAgeMinimum)
	expectedOutputPointer := &expectedOutput

	userRepoMock.On("ValidateUserExisted", mock.Anything, mock.Anything).Return(false)

	result := validator.ValidateUserDetails(user)

	assert.Equal(t, expectedOutputPointer, result)
}

func TestValidateUserDetails_MultipleUserDetailsErrors(t *testing.T) {
	user := models.User{ID: primitive.NilObjectID, FirstName: "", LastName: "Doe", Email: "aa.a", Age: 20}

	userRepoMock := new(repo.UserRepoMock)
	validator := &UserValidator{
		userRepo: userRepoMock,
	}

	expectedOutput := errs.ErrorMessage{
		Error:   errs.ResponseValidationFailed,
		Details: []string{errs.ErrorNameRequired, errs.ErrorEmailFormat},
	}
	expectedOutputPointer := &expectedOutput

	userRepoMock.On("ValidateUserExisted", mock.Anything, mock.Anything).Return(false)

	result := validator.ValidateUserDetails(user)

	assert.Equal(t, expectedOutputPointer, result)
}
