package routes

import (
	"GourmetOps/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(app *gin.Engine) {
	app.GET("/orders", controllers.GetOrders())
	app.GET("/orders/:id", controllers.GetOrder())
	app.POST("/orders", controllers.CreateOrder())
	app.PATCH("/orders/:id", controllers.UpdateOrder())

}
