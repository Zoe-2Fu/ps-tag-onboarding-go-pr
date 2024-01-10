package validator

import (
	errs "github.com/Zoe-2Fu/ps-tag-onboarding-go-pr/internal/constants"
	models "github.com/Zoe-2Fu/ps-tag-onboarding-go-pr/internal/data"
	"github.com/stretchr/testify/mock"
)

type UserValidatorMock struct {
	mock.Mock
}

func (m *UserValidatorMock) ValidateUserDetails(user models.User) *errs.ErrorMessage {
	args := m.Called(user)
	return args.Get(0).(*errs.ErrorMessage)
}
