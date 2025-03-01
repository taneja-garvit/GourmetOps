package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id" json:"id"`
	First_name    *string            `json:"first_name" validate:"required"`
	Last_name     *string            `json:"last_name" validate:"required"`
	Email         *string            `json:"email" validate:"required,email"`
	Password      *string            `json:"password" validate:"required,min=8"`
	Avatar        *string            `json:"avatar" `
	Phone         *string            `json:"phone" validate:"required"`
	Refresh_token *string            `json:"refresh_token"`
	User_id       string             `json:"user_id"`
	Created_at    *time.Time         `json:"created_at" `
	Updated_at    *time.Time         `json:"updated_at"`
}
