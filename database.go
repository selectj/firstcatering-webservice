package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const DB_DRIVER = "mysql"
const DB_HOST = "localhost"
const DB_PORT = "3306"
const DB_NAME = "firstcatering"
const DB_USER = "firstcatering"
const DB_PASS = "knL4nC2jNJ378vVT"

var db *sql.DB

func openConnection() {
	connectionString := DB_USER + ":" + DB_PASS + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME
	con, err := sql.Open(DB_DRIVER, connectionString)

	if err != nil {
		panic(err.Error())
	}

	db = con
	fmt.Println("Successfully established database connection")
	// defer db.Close()
}

func getCustomer(custID int) Customer {
	var customer Customer
	err := db.QueryRow("SELECT id, balance FROM customers WHERE id = ?", custID).Scan(&customer.ID, &customer.Balance)

	if err != nil {
		panic(err.Error())
	}

	return customer
}
