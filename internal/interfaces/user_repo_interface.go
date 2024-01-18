package interfaces

import (
	"context"
	models "github.com/Zoe-2Fu/ps-tag-onboarding-go-pr/internal/data"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepo interface {
	Find(ctx echo.Context, id string) (models.User, error)
	Save(ctx context.Context, user models.User) (primitive.ObjectID, error)
}
