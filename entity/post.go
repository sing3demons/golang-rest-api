package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title string             `json:"title" bson:"title"`
	Text  string             `json:"text" bson:"text"`
}
