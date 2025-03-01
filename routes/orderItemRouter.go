package routes

import (
	"GourmetOps/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemsRoutes(app *gin.Engine) {
	app.GET("/orderItems", controllers.GetOrderItems())
	app.GET("/orderItems/:orderItem_id", controllers.GetOrderItems())
	app.GET("/orderItems-order/:order_id", controllers.GetOrderItemsByOrder())
	app.POST("/orderItems", controllers.CreateOrderItem())
	app.PATCH("/orderItems/:orderItem_id", controllers.UpdateOrderItem())

}
