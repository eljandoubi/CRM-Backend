# CRM-Backend
Simple CRM backend written in Go

## Set up postgres via docker
``` bash 
docker run --name my_postgres -e POSTGRES_USER=user -e POSTGRES_PASSWORD=password -p 5432:5432 -d postgres
```
wait couple of minutes for the postgreSQL to start and then run
```bash
docker exec -it my_postgres psql -U user

CREATE DATABASE db_customers;

CREATE TABLE IF NOT EXISTS customers (
			id SERIAL UNIQUE PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			role VARCHAR(50) NOT NULL,
			email VARCHAR(100) NOT NULL,
			phone VARCHAR(20) NOT NULL,
			contacted BOOLEAN
		);

INSERT INTO customers (id, name, role, email, phone, contacted) 
VALUES 
    (1, 'John Doe',  'Customer', 'john@example.com', '123456789', false), 
    (2, 'Jane Smith', 'Customer', 'jane@example.com', '987654321', true),
    (3, 'Alice Johnson', 'Customer', 'alice@example.com', '555123456', false);

\q
```
