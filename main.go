package main

import (
	"GourmetOps/database"
	"GourmetOps/middlewares"
	"GourmetOps/routes"
	"os"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}
	router := gin.New()      //  This creates a new Gin engine (router instance). This engine is responsible for registering routes and starting the server.
	router.Use(gin.Logger()) // This adds the Logger middleware to the Gin engine. The logger middleware automatically logs incoming requests, including information like the HTTP method, route, status code, and response time.
	routes.UserRoutes(router)
	router.Use(middlewares.Authentication)

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemsRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":", port)

}
