package routes

import (
	"GourmetOps/controllers"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(app *gin.Engine) {
	app.GET("/foods", controllers.GetFoods())
	app.GET("/food/:food_id", controllers.GetFood())
	app.POST("/foods", controllers.CreateFood())
	app.PATCH("/foods/:food_id", controllers.UpdateFood())
}
