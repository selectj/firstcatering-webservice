package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func startCardSession(w http.ResponseWriter, r *http.Request) {
	card := getCardFromRequestParams(w, r)

	if card.ID == "-1" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "Card is not registered. Please register card."}`))
		return
	}

	if !validatePinFromRequestParams(w, r, card.PIN) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "incorrect PIN"}`))
		return
	}
	sessionCard = card
	customer := card.getCardOwner()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Welcome ` + customer.Name + `!"}`))
}

func endCardSession(w http.ResponseWriter, r *http.Request) {
	if sessionCard.ID == "-1" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "Card is not registered. Please register card."}`))
		return
	}
	sessionCard.ID = "-1"

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Goodbye!"}`))
}

func getCardBalance(w http.ResponseWriter, r *http.Request) {
	if sessionCard.ID == "-1" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "There is no active session"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sessionCard)
}

func topupCardBalance(w http.ResponseWriter, r *http.Request) {
	amount := getFloatFromRequestParams(w, r, "amount")
	if sessionCard.ID == "-1" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "There is no active session"}`))
		return
	}

	if amount <= 0.00 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "Please enter a valid amount above 0"}`))
		return
	}

	currentBalance := sessionCard.Balance
	newBalance := currentBalance + amount
	sessionCard.Balance = newBalance
	status := saveCard(sessionCard)

	if !status {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "failed to update card balance"}`))
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"message": "success"}`))
}

func processPurchase(w http.ResponseWriter, r *http.Request) {
	cost := getFloatFromRequestParams(w, r, "cost")

	if sessionCard.ID == "-1" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "Card is not registered. Please register card."}`))
		return
	}

	if cost <= 0.00 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "Please enter a valid amount above 0"}`))
		return
	}

	currentBalance := sessionCard.Balance
	newBalance := currentBalance - cost
	if newBalance < 0 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "not enough balance"}`))
		return
	}

	sessionCard.Balance = newBalance
	status := saveCard(sessionCard)

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
	r.HandleFunc("/login/{cardID}/{pin}", startCardSession).Methods(http.MethodPost)
	r.HandleFunc("/logout", endCardSession).Methods(http.MethodPost)
	r.HandleFunc("/cards/balance", getCardBalance).Methods(http.MethodGet)
	r.HandleFunc("/cards/topup/{amount}", topupCardBalance).Methods(http.MethodPost)
	r.HandleFunc("/cards/purchase/{cost}", processPurchase).Methods(http.MethodPost)
	r.HandleFunc("/customers/register", registerCustomer).Methods(http.MethodPut)
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
