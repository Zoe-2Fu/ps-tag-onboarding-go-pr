package interfaces

import (
	errs "github.com/Zoe-2Fu/ps-tag-onboarding-go-pr/internal/constants"
	models "github.com/Zoe-2Fu/ps-tag-onboarding-go-pr/internal/data"
)

type UserValidator interface {
	ValidateUserDetails(user models.User) *errs.ErrorMessage
}
