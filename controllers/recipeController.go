package controllers

import (
	"fmt"
	"log"
	"os"

	helper "github.com/ShivayBhandari/recipepad-backend/helpers"
	"github.com/ShivayBhandari/recipepad-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func GetRecipe() gin.HandlerFunc{
	return func(c *gin.Context){
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading env file")
		}
		apiKey := os.Getenv("APIKEY")

		var recipeURL = "https://api.spoonacular.com/recipes/findByIngredients?apiKey=" + apiKey + "&ingredients="
		var ingredients, present = c.GetQueryArray("ingredients")
		if(present){
			for index, value := range ingredients{
				if(index == len(ingredients)-1){
					recipeURL = recipeURL + value
					continue
				}
				recipeURL = recipeURL + value + ","
			}
		}

		fmt.Println(recipeURL)

		var recipe models.Recipe

		err = helper.GetJSON(recipeURL, &recipe)
		if err != nil {
			fmt.Printf("error getting JSON respoonse: %s\n", err.Error())
			return
		} else {
			fmt.Printf("First recipe name: %s\n", recipe[0].Title)
		}
	}
}