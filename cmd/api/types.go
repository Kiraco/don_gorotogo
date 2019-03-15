package main

//Order : the order that is saved
type Order struct {
	UUID  string   `json:"-"`
	Items []Coffee `json:"Items"`
}

// Coffee : item that is added to Items array
type Coffee struct {
	CoffeType               string      `json:"CoffeType"`
	Toppings                string      `json:"Toppings"`
	PersonalizedIngredients Ingredients `json:"PersonalizedIngredients"`
}

//Ingredients : modificable ingredientes in the coffee
type Ingredients struct {
	Milk         string `json:"Milk"`
	CoffeeStyle  string `json:"CoffeStyle"`
	CoffeeShoots int    `json:"CoffeShoots"`
}
