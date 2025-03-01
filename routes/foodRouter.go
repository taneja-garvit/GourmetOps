package routes

import (
	"github.com/gin-gonic/gin"
	"GourmetOps/controllers"


)

func FoodRoutes(app *gin.Engine) {
	app.GET("/foods", controllers.GetFoods())
	app.GET("/food/:food_id", controllers.GetFood())
	app.POST("/foods", controllers.CreateFood())
	app.PATCH("/foods/:food_id", controllers.UpdateFood())
}
