package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	ID         primitive.ObjectID `bson:"_id"`
	Text       string             `json:"text"`
	Title      string             `json:"title"`
	Note_id    string             `json:"note_id"`
	Created_at *time.Time         `json:"created_at" `
	Updated_at *time.Time         `json:"updated_at"`
}
