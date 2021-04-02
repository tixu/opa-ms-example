package main

import (
	"encoding/json"
	"log"
	"net/http"
    "fmt"
	"github.com/gorilla/mux"
)

type customer struct {
	CustomerID   string `json:"custid"`
	CompanyName  string `json:"companyname"`
	Description  string `json:"description"`
	CustomerType string `json:"customertype"`
}

type allCustomers []customer

var customers = allCustomers{
	{
		CustomerID:   "1",
		CompanyName:  "Metrics Corporation",
		Description:  "A large pharma company",
		CustomerType: "LE",
	},
	{
		CustomerID:   "2",
		CompanyName:  "GeoMeta Corporation",
		Description:  "A large geolocation services company",
		CustomerType: "LE",
	},
	{
		CustomerID:   "3",
		CompanyName:  "OfficeSpace.com",
		Description:  "A large Office space company",
		CustomerType: "LE",
	},
	{
		CustomerID:   "4",
		CompanyName:  "Smals",
		Description:  "A smals Office space company",
		CustomerType: "LE",
	},
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	customerID := mux.Vars(r)["id"]

	for _, singleCustomer := range customers {
		if singleCustomer.CustomerID == customerID {
			json.NewEncoder(w).Encode(singleCustomer)
		}
	}
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(customers)
}

func getInfo(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Application Customer")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/customers", getAllCustomers).Methods("GET")
	router.HandleFunc("/customer/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/", getInfo).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
