// utils/postgresSQL.go

package utils

import (
	"database/sql"
)

// updateCustomerTable updates a customer infos
func UpdateCustomer(db *sql.DB, c Customer) error {
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
func InsertCustomer(db *sql.DB, c Customer) error {
	_, err := db.Exec(`
		INSERT INTO customers (id, name, role, email, phone, contacted) VALUES ($1, $2, $3, $4, $5, $6);
	`, c.ID, c.Name, c.Role, c.Email, c.Phone, c.Contacted)
	return err
}

func CreateTable(db *sql.DB) error {
	// Create the table if it doesn't exist
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS customers (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		role VARCHAR(50) NOT NULL,
		email VARCHAR(100) NOT NULL,
		phone VARCHAR(20) NOT NULL,
		contacted BOOLEAN);`)

	if err != nil {
		return err
	}

	// Insert data into the table
	db.Exec(`
	INSERT INTO customers (id, name, role, email, phone, contacted) 
	VALUES 
		(1, 'John Doe',  'Customer', 'john@example.com', '123456789', false), 
		(2, 'Jane Smith', 'Customer', 'jane@example.com', '987654321', true),
		(3, 'Alice Johnson', 'Customer', 'alice@example.com', '555123456', false);
	`)

	return nil
}

// deleteCustomer deletes a customer from the database
func DeleteCustomer(db *sql.DB, id int) error {
	_, err := db.Exec(`
	DELETE FROM customers WHERE id = $1;
	`, id)
	return err
}

func GetCustomer(db *sql.DB, id int) (Customer, error) {
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
func GetAllCustomers(db *sql.DB) ([]Customer, error) {
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
