package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Customer struct {
	Id        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Role      string `json:"role,omitempty"`
	Email     string `json:"email,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Contacted bool   `json:"contacted"`
}

var customers = []Customer{
	{Id: 1, Name: "Monkey D Luffy", Role: "Founder", Email: "luffy@gmail.com", Phone: "+1 123 456 999", Contacted: true},
	{Id: 2, Name: "Roronoa Zoro", Role: "CEO", Email: "zoro@google.com", Phone: "+1 234 567 123", Contacted: true},
	{Id: 3, Name: "Sanji", Role: "Vice CEO", Email: "sanji@gmail.com", Phone: "+1 345 678 666", Contacted: false},
}

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(customers)
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idParam := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParam)

	for _, customer := range customers {
		if customer.Id == id {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(customer)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func addCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	reqBody, _ := ioutil.ReadAll(r.Body)
	newCustomer := Customer{}
	json.Unmarshal(reqBody, &newCustomer)

	maxId := 0
	for _, customer := range customers {
		if customer.Id > maxId {
			maxId = customer.Id
		}
	}
	newCustomer.Id = maxId + 1

	customers = append(customers, newCustomer)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCustomer)
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idParam := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParam)

	updatedCustomer := Customer{}
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &updatedCustomer)
	updatedCustomer.Id = id

	for index, customer := range customers {
		if customer.Id == id {
			customers[index] = updatedCustomer

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(updatedCustomer)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idParam := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParam)

	for index, customer := range customers {
		if customer.Id == id {
			customers = append(customers[:index], customers[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", index)
	router.HandleFunc("/customers", getCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/customers", addCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")

	fmt.Println("Starting the CRM backend server at http://localhost:3000")
	http.ListenAndServe(":3000", router)
}
