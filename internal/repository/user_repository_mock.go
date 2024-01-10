package repo

import (
	"context"

	models "github.com/Zoe-2Fu/ps-tag-onboarding-go-pr/internal/data"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepoMock struct {
	mock.Mock
}

func (m *UserRepoMock) Find(c echo.Context, id string) (models.User, error) {
	args := m.Called(c, id)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *UserRepoMock) Save(c context.Context, user models.User) (primitive.ObjectID, error) {
	args := m.Called(c, user)
	return args.Get(0).(primitive.ObjectID), args.Error(1)
}

func (m *UserRepoMock) ValidateUserExisted(user models.User) bool {
	args := m.Called(user)
	return args.Bool(0)
}
