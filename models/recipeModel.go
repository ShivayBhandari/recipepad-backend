package models

type Recipe struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Image string `json:"image"`
}

// type Ingredient []struct {
// 	ID				int				`json:"id"`
// 	Name			string			`json:"name"`
// 	Amount			int				`json:"amount"`
// 	Unit			string			`json:"unit"`
// }

// type RecipeInformation struct {
// 	Vegetarian		bool			`json:"vegetarian"`
// 	Vegan			bool			`json:"vegan"`
// 	GlutenFree		bool			`json:"glutenFree"`
// 	DairyFree		bool			`json:"dairyFree"`
// 	VeryHealthy		bool			`json:"veryHealthy"`
// 	ID				int				`json:"id"`
// 	Title			string			`json:"title"`
// 	SourceURL		string			`json:"sourceUrl"`
// 	ExtendedIngredients		[]interface{}	`json:"ingredients"`
// }

type RecipeInformation struct {
	Vegetarian           bool                   `json:"vegetarian"`
	Vegan                bool                   `json:"vegan"`
	GlutenFree           bool                   `json:"glutenFree"`
	DairyFree            bool                   `json:"dairyFree"`
	VeryHealthy          bool                   `json:"veryHealthy"`
	Cheap                bool                   `json:"cheap"`
	VeryPopular          bool                   `json:"veryPopular"`
	Sustainable          bool                   `json:"sustainable"`
	LowFodmap            bool                   `json:"lowFodmap"`
	ExtendedIngredients  []ExtendedIngredients  `json:"extendedIngredients"`
	ID                   int                    `json:"id"`
	Title                string                 `json:"title"`
	ReadyInMinutes       int                    `json:"readyInMinutes"`
	Servings             int                    `json:"servings"`
	SourceURL            string                 `json:"sourceUrl"`
	SpoonacularSourceURL string                 `json:"spoonacularSourceUrl"`
	Image                string                 `json:"image"`
	Summary              string                 `json:"summary"`
	Instructions         string                 `json:"instructions"`
	AnalyzedInstructions []AnalyzedInstructions `json:"analyzedInstructions"`
}

type ExtendedIngredients struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
	Unit   string  `json:"unit"`
}

type Steps struct {
	Number int    `json:"number"`
	Step   string `json:"step"`
}
type AnalyzedInstructions struct {
	Steps []Steps `json:"steps"`
}

// For Recipe Search (normal)
type RecipeSearch struct {
	Results []Results `json:"results"`
}
type Results struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Image string `json:"image"`
}

// For Random Recipe Search
type RandomRecipe struct {
	Recipes []Recipes `json:"recipes"`
}
type Recipes struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Image string `json:"image"`
}

// For Recipe Nutrients
type RecipeNutrients struct {
	Calories string `json:"calories"`
	Carbs    string `json:"carbs"`
	Fat      string `json:"fat"`
	Protein  string `json:"protein"`
}