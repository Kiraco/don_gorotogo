package main

type Order struct {
	UUID  string   `json:"UUID"`
	Items []Coffee `json:"Items"`
}

type Coffee struct {
	CoffeType               string      `json:"CoffeType"`
	Toppings                string      `json:"Toppings"`
	PersonalizedIngredients Ingredients `json:"PersonalizedIngredients"`
}

type Ingredients struct {
	Milk         string `json:"Milk"`
	CoffeeStyle  string `json:"CoffeStyle"`
	CoffeeShoots int    `json:"CoffeShoots"`
}
