package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	//"strconv"

	helper "github.com/ShivayBhandari/recipepad-backend/helpers"
	"github.com/ShivayBhandari/recipepad-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func GetRecipesFromIngredients() gin.HandlerFunc{
	return func(c *gin.Context){
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading env file")
		}
		apiKey := os.Getenv("APIKEY")

		var recipeURL = "https://api.spoonacular.com/recipes/findByIngredients?number=20&apiKey=" + apiKey + "&ingredients="
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
		//fmt.Println(recipeURL)

		var recipes []models.Recipe

		err = helper.GetJSON(recipeURL, &recipes)
		if err != nil {
			fmt.Printf("error getting JSON respoonse: %s\n", err.Error())
			return
		}

		//fmt.Println(recipes)
		c.JSON(http.StatusOK, recipes)
	}
}

func GetRecipeInformation() gin.HandlerFunc{
	return func(c *gin.Context){
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading env file")
		}
		apiKey := os.Getenv("APIKEY")
		id := c.Param("id")
		recipeInformationURL := "https://api.spoonacular.com/recipes/" + id + "/information?apiKey=" + apiKey

		var recipeInformation models.RecipeInformation

		err = helper.GetJSON(recipeInformationURL, &recipeInformation)
		if(err != nil){
		 	fmt.Printf("error getting recipe information JSON response: %s\n", err.Error())
		 	return
		}

		//fmt.Println(recipeInformation)
		c.JSON(http.StatusOK, recipeInformation)
	}
}


func GetRecipeNutrients() gin.HandlerFunc{
	return func(c *gin.Context){
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading env file")
		}
		apiKey := os.Getenv("APIKEY")
		id := c.Param("id")
		recipeNutrientsURL := "https://api.spoonacular.com/recipes/" + id + "/nutritionWidget.json?apiKey=" + apiKey
		//println(recipeNutrientsURL)

		var recipeNutrients models.RecipeNutrients

		err = helper.GetJSON(recipeNutrientsURL, &recipeNutrients)
		if(err != nil){
		 	fmt.Printf("error getting recipe information JSON response: %s\n", err.Error())
		 	return
		}

		c.JSON(http.StatusOK, recipeNutrients)
	}
}


func GetRecipesFromSearch() gin.HandlerFunc{
	return func(c *gin.Context){
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading env file")
		}
		apiKey := os.Getenv("APIKEY")

		var query, present = c.GetQuery("query")
		if(present){
			var recipeURL = "https://api.spoonacular.com/recipes/complexSearch?number=20&apiKey=" + apiKey + "&query=" + query
			//fmt.Println(recipeURL)

			var recipes models.RecipeSearch
			err = helper.GetJSON(recipeURL, &recipes)
			if err != nil {
				fmt.Printf("error getting JSON respoonse: %s\n", err.Error())
				return
			}
			c.JSON(http.StatusOK, recipes)
		} else {
			var recipeURL = "https://api.spoonacular.com/recipes/random?number=20&apiKey=" + apiKey
			//fmt.Println(recipeURL)

			var recipes models.RandomRecipe
			err = helper.GetJSON(recipeURL, &recipes)
			if err != nil {
				fmt.Printf("error getting JSON response: %s\n", err.Error())
				return
			}
			c.JSON(http.StatusOK, recipes)
		}
		

	}
}