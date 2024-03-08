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

func TestValidateUserDetails(t *testing.T) {
	nilErrMsg := (*errs.ErrorMessage)(nil)
	userExistedErrMsg := errs.NewErrorMessage(errs.ResponseValidationFailed, errs.ErrorNameUnique)
	userNameMissingErrMsg := errs.NewErrorMessage(errs.ResponseValidationFailed, errs.ErrorNameRequired)
	userEmailMissingErrMsg := errs.NewErrorMessage(errs.ResponseValidationFailed, errs.ErrorEmailRequired)
	userEmailFormatErrMsg := errs.NewErrorMessage(errs.ResponseValidationFailed, errs.ErrorEmailFormat)
	userAgeErrMsg := errs.NewErrorMessage(errs.ResponseValidationFailed, errs.ErrorAgeMinimum)
	multiUserDetailErrMsg := errs.ErrorMessage{
		Error:   errs.ResponseValidationFailed,
		Details: []string{errs.ErrorNameRequired, errs.ErrorEmailFormat},
	}

	testCases := []struct {
		name           string
		user           models.User
		isValidDetails bool
		isUserExisted  bool
		expectedOutput *errs.ErrorMessage
	}{
		{
			name:           "Valid user details",
			user:           models.User{ID: primitive.NilObjectID, FirstName: "John", LastName: "Doe", Email: "good@example.com", Age: 25},
			isValidDetails: true,
			isUserExisted:  false,
			expectedOutput: nilErrMsg,
		}, {
			name:           "User is existed",
			user:           models.User{ID: primitive.NilObjectID, FirstName: "John", LastName: "Doe", Email: "a@a.a", Age: 20},
			isValidDetails: false,
			isUserExisted:  true,
			expectedOutput: &userExistedErrMsg,
		}, {
			name:           "User name is missing",
			user:           models.User{ID: primitive.NilObjectID, FirstName: "", LastName: "Doe", Email: "a@a.a", Age: 20},
			isValidDetails: false,
			isUserExisted:  false,
			expectedOutput: &userNameMissingErrMsg,
		}, {
			name:           "User email is missing",
			user:           models.User{ID: primitive.NewObjectID(), FirstName: "John", LastName: "Doe", Email: "", Age: 20},
			isValidDetails: false,
			isUserExisted:  false,
			expectedOutput: &userEmailMissingErrMsg,
		}, {
			name:           "Invalid user email format",
			user:           models.User{ID: primitive.NewObjectID(), FirstName: "John", LastName: "Doe", Email: "aa.a", Age: 20},
			isValidDetails: false,
			isUserExisted:  false,
			expectedOutput: &userEmailFormatErrMsg,
		}, {
			name:           "Invalid user age",
			user:           models.User{ID: primitive.NilObjectID, FirstName: "John", LastName: "Doe", Email: "a@a.a", Age: 16},
			isValidDetails: false,
			isUserExisted:  false,
			expectedOutput: &userAgeErrMsg,
		}, {
			name:           "Multiple user details errors",
			user:           models.User{ID: primitive.NilObjectID, FirstName: "", LastName: "Doe", Email: "aa.a", Age: 20},
			isValidDetails: false,
			isUserExisted:  false,
			expectedOutput: &multiUserDetailErrMsg,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			userRepoMock := new(repo.UserRepoMock)
			validator := &UserValidator{
				userRepo: userRepoMock,
			}

			userRepoMock.On("ValidateUserExisted", mock.Anything, mock.Anything).Return(tc.isUserExisted, nil)

			result := validator.ValidateUserDetails(tc.user)

			if tc.isValidDetails {
				assert.Nil(t, result, tc.expectedOutput)
			} else {
				assert.Equal(t, tc.expectedOutput, result)
			}
		})
	}
}
