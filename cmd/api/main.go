package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/xeipuuv/gojsonschema"
)

var orders []Order

/*
 * Verify if the application is alive.
 */
func ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "App is live.")
}

/*
 * Add an order to the orders collection.
 * For dev purposes instead of a DB is an inmemory array of orders that can be modified
 */
func addOrder(w http.ResponseWriter, r *http.Request) {
	//Verify there is a body
	bodyBytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Perform a schema validation,
	// need to verify all properties are set and does not include extra information
	schemaLoader := gojsonschema.NewStringLoader(Schema)
	dataLoader := gojsonschema.NewBytesLoader(bodyBytes)
	result, err := gojsonschema.Validate(schemaLoader, dataLoader)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "There was an error loading the schema and validating.")
		return
	}
	if result.Valid() {
		var order Order
		err = json.Unmarshal(bodyBytes, &order)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err)
			return
		}
		order.UUID = uuid.New().String()
		orders = append(orders, order)
		//Return the generated UUID
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, order.UUID)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "There is an error in the payload, verify the content has a valid structure.")
	}
}

/*
 * Retrieve all current orders.
 */
func getOrders(w http.ResponseWriter, r *http.Request) {
	if len(orders) == 0 {
		fmt.Fprint(w, "No orders at the moment.")
		w.WriteHeader(http.StatusNoContent)
	} else {
		json.NewEncoder(w).Encode(orders)
		w.WriteHeader(http.StatusOK)
	}
}

/*
 * Get an specific order
 */
func getOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]
	for _, n := range orders {
		if n.UUID == uuid {
			json.NewEncoder(w).Encode(n)
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	fmt.Fprint(w, "Requested order does not exists.")
	w.WriteHeader(http.StatusNoContent)
}

/*
* Cancel an order
 */
func deleteOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]
	for i, n := range orders {
		if n.UUID == uuid {
			orders = append(orders[:i], orders[i+1:]...)
			w.WriteHeader(http.StatusAccepted)
			return
		}
	}
	fmt.Fprint(w, "Requested order does not exists.")
	w.WriteHeader(http.StatusNoContent)
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
