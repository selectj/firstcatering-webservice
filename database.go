package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//Database driver
const dbDriver = "mysql"

//Database host
const dbHost = "localhost"

//Database port
const dbPort = "3306"

//Database name
const dbName = "firstcatering"

//Database username
const dbUser = "firstcatering"

//Database user password
const dbPass = "knL4nC2jNJ378vVT"

var db *sql.DB

func openConnection() {
	connectionString := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName
	con, err := sql.Open(dbDriver, connectionString)

	if err != nil {
		panic(err.Error())
	}

	db = con
	fmt.Println("Successfully established database connection")
	// defer db.Close()
}

func getCustomer(custID int) Customer {
	var customer Customer
	err := db.QueryRow("SELECT id, name, email, telephone FROM customers WHERE id = ?", custID).Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Telephone)

	if err != nil {
		customer.ID = -1
	}

	return customer
}

func getCard(cardID string) DataCard {
	var card DataCard
	err := db.QueryRow("SELECT id, pin, balance, customerID FROM cards WHERE id = ?", cardID).Scan(&card.ID, &card.PIN, &card.Balance, &card.CustomerID)

	if err != nil {
		card.ID = "-1"
	}

	return card
}

func newCustomer(customer NewCustomer) bool {
	toIns, err := db.Prepare("INSERT INTO customers(id, name, email, telephone) VALUES(?, ? , ?, ?)")
	if err != nil {
		return false
	}

	toIns.Exec(customer.ID, customer.Name, customer.Email, customer.Telephone)

	toIns, err = db.Prepare("INSERT INTO cards(id, pin, balance, customerID) VALUES(?, ? , ?, ?)")
	if err != nil {
		return false
	}

	newID := randomString(16)
	toIns.Exec(newID, customer.PIN, 0, customer.ID)
	return true
}

func (card DataCard) save() bool {
	toIns, err := db.Prepare("UPDATE cards SET balance=? WHERE id=?")
	if err != nil {
		return false
	}

	toIns.Exec(card.Balance, card.ID)
	return true
}
