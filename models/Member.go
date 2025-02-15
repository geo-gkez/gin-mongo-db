package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Member struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Name    string             `bson:"name,omitempty"`
	Surname string             `bson:"surname,omitempty"`
	Email   string             `bson:"email,omitempty"`
}
