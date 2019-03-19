package main

//Schema : schema that POST needs to adhere for adding an order.
var Schema = `
{
	"$schema": "http://json-schema.org/draft-04/schema#",
	"type": "object",
	"required": [
	  "items"
	],
	"additionalProperties": false,
	"properties": {
	  "UUID": {
		"type": "string"
	  },
	  "items": {
		"type": "array",
		"minItems": 1,
		"items": {
		  "type": "object",
		  "required": [
			"coffe_type",
			"toppings",
			"personalized_ingredients"
		  ],
		  "additionalProperties": false,
		  "properties": {
			"coffe_type": {
			  "type": "string",
			  "enum": [
				"Americano",
				"Latte",
				"Capuccino",
				"Espresso"
			  ]
			},
			"toppings": {
			  "type": "string",
			  "enum": [
				"Crema Batida",
				"Chispas de chocolate"
			  ]
			},
			"personalized_ingredients": {
			  "type": "object",
			  "required": [
				"milk",
				"coffee_style",
				"coffee_shots"
			  ],
			  "additionalProperties": false,
			  "properties": {
				"milk": {
				  "type": "string",
				  "enum": [
					"Light",
					"Almendra",
					"Deslactosada"
				  ]
				},
				"coffee_style": {
				  "type": "string",
				  "enum": [
					"Caliente",
					"Frio",
					"Frappe"
				  ]
				},
				"coffee_shots": {
				  "type": "integer",
				  "min": 1
				}
			  }
			}
		  }
		}
	  }
	}
  }
`
