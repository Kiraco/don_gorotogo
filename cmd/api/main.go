package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var orders []Order

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "App is live.")
}

func addOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Order added.")
}

func getOrders(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(orders)
	fmt.Fprint(w, "All orders.")
}

func getOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Requested order retrieved.")
}

func deleteOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Requested order deleted.")
}

func main() {

	pIngredients := Ingredients{Milk: "Light", CoffeeStyle: "Caliente", CoffeeShoots: 4}
	cType := Coffee{CoffeType: "Americano", Toppings: "Crema batida", PersonalizedIngredients: pIngredients}
	order := Order{ID: 1, Items: []Coffee{cType}}
	orders = append(orders, order)

	router := mux.NewRouter()
	router.HandleFunc("/ping", ping).Methods("GET")
	router.HandleFunc("/orders", getOrders).Methods("GET")
	router.HandleFunc("/order/{id}", getOrder).Methods("GET")
	router.HandleFunc("/order/{id}", deleteOrder).Methods("DELETE")
	router.HandleFunc("/order", addOrder).Methods("POST")

	log.Fatal(http.ListenAndServe(":5000", router))
}
