package app

import (
	"log"
	"net/http"

	"github.com/bhatsubhas/banking/domain"
	"github.com/bhatsubhas/banking/service"
	"github.com/gorilla/mux"
)

func Start() {

	router := mux.NewRouter()
	// wiring
	ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryDb())}
	// define all our routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	// starting the server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
