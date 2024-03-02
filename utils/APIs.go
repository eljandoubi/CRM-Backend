// utils/restful.go

package utils

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// InitializeRouter initializes and returns a new Gorilla Mux router
func InitializeRouter(db *sql.DB) *mux.Router {
	router := mux.NewRouter()

	// Define RESTful endpoints
	router.HandleFunc("/", home)
	router.HandleFunc("/customers", getAllCustomersHandler(db)).Methods("GET")
	router.HandleFunc("/customers/{id}", getCustomerHandler(db)).Methods("GET")
	router.HandleFunc("/customers", createCustomerHandler(db)).Methods("POST")
	router.HandleFunc("/customers/{id}", updateCustomerHandler(db)).Methods("PUT")
	router.HandleFunc("/customers/{id}", deleteCustomerHandler(db)).Methods("DELETE")

	return router
}

func home(w http.ResponseWriter, r *http.Request) {
	filePath := "static/index.html"
	w.Header().Set("Content-Type", "text/html")
	log.Println("Serve", filePath)
	http.ServeFile(w, r, filePath)
}

func getAllCustomersHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		customers, err := GetAllCustomers(db)
		if err != nil {
			handleError(w, err, http.StatusInternalServerError)
			return
		} else {
			log.Println("Successfully get All Customers")
		}

		respondJSON(w, customers)
	}
}

func getCustomerHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			handleError(w, err, http.StatusBadRequest)
			return
		}

		customer, err := GetCustomer(db, id)
		if err != nil {
			if err == sql.ErrNoRows {
				handleError(w, fmt.Errorf("customer not found"), http.StatusNotFound)
			} else {
				handleError(w, err, http.StatusInternalServerError)
			}
			return
		} else {
			log.Println("Successfully get Customer")
		}

		respondJSON(w, customer)
	}
}

func createCustomerHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var customer Customer
		if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
			handleError(w, err, http.StatusBadRequest)
			return
		}

		if err := InsertCustomer(db, customer); err != nil {
			handleError(w, err, http.StatusInternalServerError)
			return
		} else {
			log.Println("Successfully create Customer")
		}

		respondJSON(w, customer)
	}
}

func updateCustomerHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			handleError(w, err, http.StatusBadRequest)
			return
		}

		var customer Customer
		if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
			handleError(w, err, http.StatusBadRequest)
			return
		}

		customer.ID = id
		if err := UpdateCustomer(db, customer); err != nil {
			handleError(w, err, http.StatusInternalServerError)
			return
		} else {
			log.Println("Successfully update Customer")
		}

		respondJSON(w, customer)
	}
}

func deleteCustomerHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			handleError(w, err, http.StatusBadRequest)
			return
		}

		if err := DeleteCustomer(db, id); err != nil {
			handleError(w, err, http.StatusInternalServerError)
			return
		} else {
			log.Println("Successfully delete Customer")
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func respondJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func handleError(w http.ResponseWriter, err error, status int) {
	log.Println("Error: %v", err)
	http.Error(w, http.StatusText(status), status)
}
