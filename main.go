package main

//Active sessions array
var sessions []Session

//Customer structure (already registered customer data model)
type Customer struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
}

//Session structure for storing card sessions
type Session struct {
	ID           string
	Card         DataCard
	LastActivity int64
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

func (dc DataCard) getCardOwner() Customer {
	return getCustomer(dc.CustomerID)
}

func (session Session) end() {
	for i, s := range sessions {
		if s.ID == session.ID {
			removeFromSlice(sessions, i)
			return
		}
	}
}

func main() {
	openConnection()
	initAPI()
}
