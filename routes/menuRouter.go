package routes

import (
	"GourmetOps/controllers"

	"github.com/gin-gonic/gin"
)

func MenuRoutes(app *gin.Engine) {
	app.GET("/menus", controllers.GetMenus())
	app.GET("/menus/:menu_id", controllers.GetMenu())
	app.GET("/menus", controllers.CreateMenu())
	app.GET("/menus/:menu_id", controllers.UpdateMenu())
}
