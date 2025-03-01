package routes

import (
	"GourmetOps/controllers"

	"github.com/gin-gonic/gin"
)

func TableRoutes(app *gin.Engine) {
	app.GET("/tables", controllers.GetTables())
	app.GET("/tables/table_:id", controllers.GetTable())
	app.POST("/tables", controllers.CreateTables())
	app.PATCH("/tables/table_:id", controllers.UpdateTable())
}
