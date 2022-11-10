package models

// type Ingredient struct {
// 	ID     int      `json:"id"`
// 	Amount int      `json:"amount"`
// 	Unit   string   `json:"unit"`
// 	Name   string   `json:"name"`
// 	Meta   []string `json:"meta"`
// 	Image  string   `json:"image"`
// }

//type Recipe struct {
//	ID                int          `json:"id"`
//	Title             string       `json:"title"`
//	Image             string       `json:"image"`
// MissedIngredients []Ingredient `json:"missedIngredients"`
// UsedIngredients   []Ingredient `json:"usedIngredients"`
// UnusedIngredients []Ingredient `json:"unusedIngredients"`
//}

// type Recipes struct {
// 	Results []Recipe
// }

type Recipe []struct {
	ID				int				`json:"id"`
	Title			string			`json:"title"`
}