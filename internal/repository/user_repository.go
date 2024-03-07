package repo

import (
	"context"
	"errors"
	"net/http"

	errs "github.com/Zoe-2Fu/ps-tag-onboarding-go-pr/internal/constants"
	models "github.com/Zoe-2Fu/ps-tag-onboarding-go-pr/internal/data"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection = "userdetails"

type UserRepository struct {
	db *mongo.Database
}

func NewUserRepo(db *mongo.Database) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Find(ctx echo.Context, id string) (models.User, error) {
	var user models.User

	objID, _ := primitive.ObjectIDFromHex(id)
	result := r.db.Collection(userCollection).FindOne(ctx.Request().Context(), bson.M{"_id": objID})
	if result == nil {
		return models.User{}, errors.New("can't find user from database")
	}

	err := result.Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.User{}, mongo.ErrNoDocuments
		} else {
			return models.User{}, errors.New("failed to decode user from database")
		}
	}
	return user, nil
}

func (r *UserRepository) Save(ctx context.Context, user models.User) (primitive.ObjectID, error) {
	userBSON, err := bson.Marshal(user)
	if err != nil {
		return primitive.NilObjectID, echo.NewHTTPError(http.StatusInternalServerError, errs.ErrorMessage{
			Error:   errs.ErrorInternalServerError,
			Details: []string{"Failed to marshaling BSON"},
		})
	}
	result, err := r.db.Collection(userCollection).InsertOne(ctx, userBSON)
	if err != nil {
		return primitive.NilObjectID, echo.NewHTTPError(http.StatusInternalServerError, errs.ErrorMessage{
			Error:   errs.ErrorInternalServerError,
			Details: []string{"Failed to save user"},
		})
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, echo.NewHTTPError(http.StatusInternalServerError, errs.ErrorMessage{
			Error:   errs.ErrorInternalServerError,
			Details: []string{"Failed to find objectID"},
		})
	}

	return insertedID, err
}

func (r *UserRepository) ValidateUserExisted(user models.User) (bool, error) {
	filter := bson.M{"firstname": user.FirstName, "lastname": user.LastName}

	count, err := r.db.Collection(userCollection).CountDocuments(context.Background(), filter)
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}
