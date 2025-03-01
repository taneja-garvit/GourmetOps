package controllers

import (
	database "GourmetOps/databse"
	"GourmetOps/models"
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.second)
		foodId := c.Param("food_id")
		var food models.Food
		foodCollection.FindOne(ctx, bson.M{"food_id"}).Decode(&food)

	}
}
