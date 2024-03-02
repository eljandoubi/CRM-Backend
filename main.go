package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Customer struct representing each customer
type Customer struct {
	ID        int
	Name      string
	Role      string
	Email     string
	Phone     string
	Contacted bool
}

const (
	host     = "localhost"
	port     = 5432
	user     = "user"
	password = "password"
	dbname   = "db_customers"
)

func main() {
	// Connect to the PostgreSQL database
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	defer db.Close()

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	} else {
		log.Println("Successfully connected to PostgreSQL database")
	}

	cu := Customer{ID: 4, Name: "Bob", Role: "Johnson", Email: "bob.johnson@example.com", Phone: "098765432", Contacted: false}
	err = insertCustomer(db, cu)
	if err != nil {
		log.Fatal("Error insert customer: ", err)
	} else {
		log.Println("Successfully inserted customer")
	}

	ct, err := getCustomer(db, cu.ID)

	if err != nil {
		log.Fatal("Error get customer: ", err)
	} else {
		log.Println("ID:", ct.ID, " Name:", ct.Name, " Role:", ct.Role, "Email:",
			ct.Email, "Phone:", ct.Phone, "Contacted:", ct.Contacted)
	}

	cu.Contacted = true

	err = updateCustomer(db, cu)
	if err != nil {
		log.Fatal("Error update customer: ", err)
	}

	ct, err = getCustomer(db, cu.ID)

	if err != nil {
		log.Fatal("Error get customer: ", err)
	} else {
		log.Println("ID:", ct.ID, " Name:", ct.Name, " Role:", ct.Role, "Email:",
			ct.Email, "Phone:", ct.Phone, "Contacted:", ct.Contacted)
	}

	err = deleteCustomer(db, cu.ID)
	if err != nil {
		log.Fatal("Error delete customer: ", err)
	}

	// Retrieve all customers
	db_customers, err := getAllCustomers(db)
	if err != nil {
		log.Fatal("Error getting customers: ", err)
	}

	fmt.Println("All Customers:")
	for _, c := range db_customers {
		log.Println("ID:", c.ID, " Name:", c.Name, " Role:", c.Role, "Email:", c.Email, "Phone:", c.Phone, "Contacted:", c.Contacted)
	}
}

// updateCustomerTable updates a customer infos
func updateCustomer(db *sql.DB, c Customer) error {
	_, err := db.Exec(`
	UPDATE customers 
	SET name=$1,
		role=$2,
		email=$3,
		phone=$4,
		contacted=$5
	WHERE id=$6;
	`, c.Name, c.Role, c.Email, c.Phone, c.Contacted, c.ID)
	return err
}

// insertCustomer inserts a customer into the database
func insertCustomer(db *sql.DB, c Customer) error {
	_, err := db.Exec(`
		INSERT INTO customers (name, role, email, phone, contacted) VALUES ($1, $2, $3, $4, $5);
	`, c.Name, c.Role, c.Email, c.Phone, c.Contacted)
	return err
}

// deleteCustomer deletes a customer from the database
func deleteCustomer(db *sql.DB, id int) error {
	_, err := db.Exec(`
	DELETE FROM customers WHERE id = $1;
	`, id)
	return err
}

func getCustomer(db *sql.DB, id int) (Customer, error) {
	var cs Customer

	rows, err := db.Query("SELECT * FROM customers WHERE id=$1;", id)
	if err != nil {
		return cs, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&cs.ID, &cs.Name, &cs.Role, &cs.Email, &cs.Phone, &cs.Contacted)
	} else {
		// No rows found
		return cs, sql.ErrNoRows
	}

	if err != nil {
		return cs, err
	}

	return cs, nil
}

// getAllCustomers retrieves all customers from the database
func getAllCustomers(db *sql.DB) ([]Customer, error) {
	rows, err := db.Query("SELECT * FROM customers;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []Customer
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.ID, &c.Name, &c.Role, &c.Email, &c.Phone, &c.Contacted)
		if err != nil {
			return nil, err
		}
		customers = append(customers, c)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return customers, nil
}
