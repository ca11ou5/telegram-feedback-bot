package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Question struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Message string             `db:"message"`
}

type QA struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Question string             `db:"question"`
	Answer   string             `db:"answer"`
}
