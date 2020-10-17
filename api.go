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
	customer := getCustomerFromRequestParams(w, r)

	if customer.ID == -1 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "Customer with customer ID not found"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customer)
}

func topupCustomerBalance(w http.ResponseWriter, r *http.Request) {
	amount := getFloatFromRequestParams(w, r, "amount")
	customer := getCustomerFromRequestParams(w, r)
	if customer.ID == -1 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "Customer with customer ID not found"}`))
		return
	}

	if amount <= 0.00 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "Please enter a valid amount above 0"}`))
		return
	}

	currentBalance := customer.Balance
	newBalance := currentBalance + amount
	customer.Balance = newBalance
	status := updateCustomer(customer)

	if !status {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "failed to update customer balance"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "success"}`))
}

func processPurchase(w http.ResponseWriter, r *http.Request) {
	customer := getCustomerFromRequestParams(w, r)
	cost := getFloatFromRequestParams(w, r, "cost")

	if customer.ID == -1 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "Customer with customer ID not found"}`))
		return
	}

	if cost <= 0.00 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "Please enter a valid amount above 0"}`))
		return
	}

	currentBalance := customer.Balance
	newBalance := currentBalance - cost
	if newBalance < 0 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "not enough balance"}`))
		return
	}

	customer.Balance = newBalance
	status := updateCustomer(customer)

	if !status {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "failed to update customer balance"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "success"}`))
}

func initAPI() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/").Subrouter()
	defineEndpoints(api)
	fmt.Println("FirstCatering web service listening..")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func defineEndpoints(r *mux.Router) {
	r.HandleFunc("/customer/balance/{customerID}", getCustomerBalance).Methods(http.MethodGet)
	r.HandleFunc("/customer/topup/{customerID}/{amount}", topupCustomerBalance).Methods(http.MethodPost)
	r.HandleFunc("/customer/purchase/{customerID}/{cost}", topupCustomerBalance).Methods(http.MethodPost)
}

func getCustomerFromRequestParams(w http.ResponseWriter, r *http.Request) Customer {
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	custID := -1
	var err error
	if value, ok := params["customerID"]; ok {
		custID, err = strconv.Atoi(value)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "Please specify a valid customer ID"}`))
		}
	}
	customer := getCustomer(custID)

	return customer
}

func getFloatFromRequestParams(w http.ResponseWriter, r *http.Request, paramName string) float64 {
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	amount := 0.00
	var err error
	if value, ok := params[paramName]; ok {
		amount, err = strconv.ParseFloat(value, 32)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "Please specify a valid amount"}`))
		}
	}
	return amount
}
