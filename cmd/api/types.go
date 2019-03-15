package main

type Order struct {
	ID    int
	Items []Coffee
}

type Coffee struct {
	CoffeType               string
	Toppings                string
	PersonalizedIngredients Ingredients
}

type Ingredients struct {
	Milk         string
	CoffeeStyle  string
	CoffeeShoots int
}
