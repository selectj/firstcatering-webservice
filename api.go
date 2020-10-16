package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getCustomerBalance(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	custID := -1
	var err error
	if value, ok := params["customerID"]; ok {
		custID, err = strconv.Atoi(value)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "Please specify a customer ID"}`))
			return
		}
	}

	if custID == -1 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Customer with customer ID not found"}`))
		return
	}

	customer := getCustomer(custID)

	json.NewEncoder(w).Encode(customer)
}

func initAPI() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/").Subrouter()

	api.HandleFunc("/customer/balance/{customerID}", getCustomerBalance).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Println("FirstCatering web service listening..")
}
