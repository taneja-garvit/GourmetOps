package controllers

import (
	database "GourmetOps/database"
	"GourmetOps/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food") // yha pr data is coming from db jo food db me h and storing infoodcollection
var validate = validator.New()

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		foodId := c.Param("food_id")                                                // ye to jo api se coming
		var food models.Food                                                        // now ye food name ka ek var h jiska type models.Food h means type vo h jo schema h
		err := foodCollection.FindOne(ctx, bson.M{"food_id": foodId}).Decode(&food) // yha pr food collection se data find kr rha h and decode kr rha h cz data mongodb me BSON format me store ho ta h so decode krna pdta h into a Go data structure like a struct
		if err != nil {
			c.JSON(404, gin.H{"error": "Food not found"})
			return
		}
		c.JSON(200, food)
	}
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var food models.Food
		var menu models.Menu
		if err := c.BindJSON(&food); err != nil { //means all the data which is coming from food struct is getting bind together for further process
			// in mux we would using this -> err := json.NewDecoder(r.Body).Decode(&food)
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}
		validaterErr := validate.Struct(food)
		if validaterErr != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}

		err := foodCollection.FindOne(ctx, bson.M{"menu_id": food.Menu_id}).Decode(&menu)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Menu not found"})
			return
		}
		food.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.ID = primitive.NewObjectID()
		food.Food_id = food.ID.Hex()
		// var num = toFixed(*food.Price, 2)
		// food.Price = &num

		result, inserr := foodCollection.InsertOne(ctx, food)
		if inserr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert food"})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)

	}
}
