package routes

import (
	controller "github.com/ShivayBhandari/recipepad-backend/controllers"
	"github.com/ShivayBhandari/recipepad-backend/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.Use(middleware.Authenticate())
	//incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
	//incomingRoutes.GET("/users/:user_id")
}