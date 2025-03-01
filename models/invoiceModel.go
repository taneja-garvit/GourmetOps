package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Invoice struct {
	ID primitive.ObjectID `bson:"_id"`
	Invoice_id string `json:"invoice_id"`
	Order_id   string `json:"order_id"`
	Payment_Method *string `json:"payment_method" validate:"eq=CARD|eq=CASH|eq=""`
	Payment_Status *string `json:"payment_status"`
	Payment_due_date time.Time `json:"payment_due_date"`
	Created_at     time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}