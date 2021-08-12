package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zip_code"`
}

func greet(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Hello, World! from Go App")
}

func getAllCustomers(rw http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"Subhas", "Udupi", "576226"},
		{"Karthik", "Bengaluru", "560001"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		rw.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(rw).Encode(customers)
	} else {
		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(customers)

	}
}

func createCustomer(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Post request received..")
}

func getCustomer(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(rw, vars["customer_id"])
}
