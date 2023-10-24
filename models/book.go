package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Book struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	Author    string             `json:"author" bson:"author"`
	Genre     string             `json:"genre" bson:"genre"`
	Isbn      string             `json:"isbn" bson:"isbn"`
	Published bool               `bson:"published" bson:"published"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}
