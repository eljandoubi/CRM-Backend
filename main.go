// main.go

package main

import (
	"CRM-Backend/utils"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "user"
	password = "password"
	dbname   = "db_customers"
)

func main() {
	// Establish a database connection
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	} else {
		log.Println("Successfully connected to PostgreSQL database")
	}

	err = utils.CreateTable(db)

	if err != nil {
		log.Fatal("Error creating customers table: ", err)
	}

	// Initialize the Gorilla Mux router
	router := utils.InitializeRouter(db)

	// Start the HTTP server
	log.Println("Starting the CRM backend server at http://localhost:3000")
	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatal(err)
	}

}
