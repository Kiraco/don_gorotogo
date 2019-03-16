package main

//Order : the order that is saved
type Order struct {
	UUID  string   `json:"-"`
	Items []Coffee `json:"items"`
}

// Coffee : item that is added to Items array
type Coffee struct {
	CoffeType               string      `json:"coffe_type"`
	Toppings                string      `json:"toppings"`
	PersonalizedIngredients Ingredients `json:"personalized_ingredients"`
}

//Ingredients : modificable ingredientes in the coffee
type Ingredients struct {
	Milk        string `json:"milk"`
	CoffeeStyle string `json:"coffee_style"`
	CoffeeShots int    `json:"coffee_shots"`
}
