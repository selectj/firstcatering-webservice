package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getCardBalance(w http.ResponseWriter, r *http.Request) {
	card := getCardFromRequestParams(w, r)

	if card.ID == "-1" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "Card is not registered. Please register card.}`))
		return
	}

	if !validatePinFromRequestParams(w, r, card.PIN) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "incorrect PIN"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(card)
}

func topupCardBalance(w http.ResponseWriter, r *http.Request) {
	amount := getFloatFromRequestParams(w, r, "amount")
	card := getCardFromRequestParams(w, r)
	if card.ID == "-1" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "Card is not registered. Please register card.}`))
		return
	}

	if !validatePinFromRequestParams(w, r, card.PIN) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "incorrect PIN"}`))
		return
	}

	if amount <= 0.00 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "Please enter a valid amount above 0"}`))
		return
	}

	currentBalance := card.Balance
	newBalance := currentBalance + amount
	card.Balance = newBalance
	status := updateCard(card)

	if !status {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "failed to update card balance"}`))
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"message": "success"}`))
}

func processPurchase(w http.ResponseWriter, r *http.Request) {
	card := getCardFromRequestParams(w, r)
	cost := getFloatFromRequestParams(w, r, "cost")

	if card.ID == "-1" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "Card is not registered. Please register card.}`))
		return
	}

	if !validatePinFromRequestParams(w, r, card.PIN) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "incorrect PIN"}`))
		return
	}

	if cost <= 0.00 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "Please enter a valid amount above 0"}`))
		return
	}

	currentBalance := card.Balance
	newBalance := currentBalance - cost
	if newBalance < 0 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "not enough balance"}`))
		return
	}

	card.Balance = newBalance
	status := updateCard(card)

	if !status {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "failed to update card balance"}`))
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"message": "success"}`))
}

func registerCustomer(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var customer NewCustomer
	err := decoder.Decode(&customer)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "bad request"}`))
		return
	}

	status := newCustomer(customer)

	if !status {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "could not create customer"}`))
		return
	}

	w.WriteHeader(http.StatusAccepted)
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
	r.HandleFunc("/cards/balance/{cardID}/{pin}", getCardBalance).Methods(http.MethodGet)
	r.HandleFunc("/cards/topup/{cardID}/{amount}", topupCardBalance).Methods(http.MethodPost)
	r.HandleFunc("/cards/purchase/{cardID}/{cost}", processPurchase).Methods(http.MethodPost)
	r.HandleFunc("/customers/register/{customer}", registerCustomer).Methods(http.MethodPut)
}

func getCardFromRequestParams(w http.ResponseWriter, r *http.Request) DataCard {
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	cardID := ""
	if value, ok := params["cardID"]; ok {
		cardID = value
	}
	card := getCard(cardID)
	return card
}

func validatePinFromRequestParams(w http.ResponseWriter, r *http.Request, cardPIN string) bool {
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	pin := ""
	if value, ok := params["pin"]; ok {
		pin = value
	}
	return pin != "" && pin == cardPIN
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
