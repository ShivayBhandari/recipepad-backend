package routes

import (
	controller "github.com/ShivayBhandari/recipepad-backend/controllers"
	"github.com/gin-gonic/gin"
)

func RecipeSearch(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/recipe", controller.GetRecipe())
}