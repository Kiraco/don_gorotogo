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
	var order Order
	_ = json.NewDecoder(r.Body).Decode(&order)
	orders = append(orders, order)
	json.NewEncoder(w).Encode(order)
	fmt.Fprint(w, "Order added.")

}

func getOrders(w http.ResponseWriter, r *http.Request) {
	if len(orders) == 0 {
		fmt.Fprint(w, "No orders at the moment.")
	} else {
		json.NewEncoder(w).Encode(orders)
		fmt.Fprint(w, "All orders.")
	}
}

func getOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]
	for _, n := range orders {
		if n.UUID == uuid {
			json.NewEncoder(w).Encode(n)
			fmt.Fprint(w, "Requested order.")
			break
		}
	}
	fmt.Fprint(w, "Requested order retrieved.")
}

func deleteOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]
	for i, n := range orders {
		if n.UUID == uuid {
			orders = append(orders[:i], orders[i+1:]...)
			break
		}
	}
	fmt.Fprint(w, "Requested order deleted.")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/ping", ping).Methods("GET")
	router.HandleFunc("/orders", getOrders).Methods("GET")
	router.HandleFunc("/order/{uuid}", getOrder).Methods("GET")
	router.HandleFunc("/order/{uuid}", deleteOrder).Methods("DELETE")
	router.HandleFunc("/order", addOrder).Methods("POST")

	log.Fatal(http.ListenAndServe(":5000", router))
}
