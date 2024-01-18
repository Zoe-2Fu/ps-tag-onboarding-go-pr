package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const UserDB = "user"

type User struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty" param:"id,omitempty"`
	FirstName string             `json:"firstname" bson:"firstname"`
	LastName  string             `json:"lastname" bson:"lastname"`
	Email     string             `json:"email" bson:"email"`
	Age       int                `json:"age" bson:"age"`
}
