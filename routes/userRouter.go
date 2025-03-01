package routes

import (
	"GourmetOps/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(app *gin.Engine) {
	app.GET("/users", controllers.GetUsers())	
	app.GET("/users/:user_id", controllers.GetUser())	
	app.GET("/users/signup", controllers.Signup())	
	app.GET("/users/login", controllers.Login())	
}
