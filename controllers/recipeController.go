package controllers

import (
	"fmt"
	"log"
	"os"
	"strconv"

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

		var recipes []models.Recipe

		err = helper.GetJSON(recipeURL, &recipes)
		if err != nil {
			fmt.Printf("error getting JSON respoonse: %s\n", err.Error())
			return
		}

		var recipesInformation []models.RecipeInformation

		for index := range recipes {
			recipeInformationURL := "https://api.spoonacular.com/recipes/" + strconv.Itoa(recipes[index].ID) + "/information?apiKey=" + apiKey
			
			var recipeInformation models.RecipeInformation
			err := helper.GetJSON(recipeInformationURL, &recipeInformation)
			if(err != nil){
				fmt.Printf("error getting recipe information JSON response: %s\n", err.Error())
				return
			}

			recipesInformation = append(recipesInformation, recipeInformation)
		}

		fmt.Printf("Recipe URL of first recipe is: %s\n", recipesInformation[0].SourceURL)
	}
}