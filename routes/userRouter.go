package routes

import (
	"GourmetOps/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(app *gin.Engine) {
	app.GET("/users", controllers.GetUsers())
	app.GET("/users/:user_id", controllers.GetUser())
	app.POST("/users/signup", controllers.Signup())
	app.POST("/users/login", controllers.Login())
}
