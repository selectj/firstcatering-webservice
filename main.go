package main

//Customer structure (already registered customer data model)
type Customer struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
}

//NewCustomer structure for newly registered customers (new struct used to pass PIN in req body)
type NewCustomer struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
	PIN       string `json:"pin"`
}

//DataCard structure
type DataCard struct {
	ID         string  `json:"id"`
	PIN        string  `json:"pin"`
	Balance    float64 `json:"balance"`
	CustomerID int     `json:"cusomerId"`
}

func main() {
	openConnection()
	initAPI()
}
