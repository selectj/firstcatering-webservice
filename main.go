package main

//Customer structure
type Customer struct {
	ID      int     `json:"id"`
	Balance float32 `json:"balance"`
}

func main() {
	openConnection()
	initAPI()
}
