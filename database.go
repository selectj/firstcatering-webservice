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
	err := db.QueryRow("SELECT id, balance FROM customers WHERE id = ?", custID).Scan(&customer.ID, &customer.Balance)

	if err != nil {
		customer.ID = -1
		panic(err.Error())
	}

	return customer
}
