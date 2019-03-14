package main

type Order struct {
	ID    int
	Items []Coffee
}

type Coffee struct {
	coffeType               string
	toppings                string
	personalizedIngredients PersonalizedIngredients
}

type PersonalizedIngredients struct {
	milk         string
	coffeeStyle  string
	coffeeShoots int
}
