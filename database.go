package main

import (
	"database/sql"
	"log"
	"strconv"
)

const DB_DRIVER = "mysql"
const DB_HOST = "localhost"
const DB_PORT = "3306"
const DB_NAME = "firstcatering"
const DB_USER = "firstcatering"
const DB_PASS = "knL4nC2jNJ378vVT"

var db *sql.DB

func openConnection() {
	db, err := sql.Open(DB_DRIVER, DB_USER+":"+DB_PASS+"@tcp("+DB_HOST+":"+DB_PORT+")/"+DB_NAME)

	if err != nil {
		log.Panicln("Failed to create connection to database")
		panic(err.Error())
	}

	log.Println("Successfully established database connection")
	defer db.Close()
}

func getCustomer(custID int) Customer {
	var customer Customer
	err := db.QueryRow(("SELECT id, balance FROM customers WHERE id = "+strconv.Itoa(custID)), 2).Scan(&customer.ID, &customer.Balance)

	if err != nil {
		panic(err.Error())
	}

	return customer
}
