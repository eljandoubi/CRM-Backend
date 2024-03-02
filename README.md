# CRM-Backend

A simple CRM backend written in Go.

## Set up PostgreSQL via Docker

```bash 
docker run --name my_postgres -e POSTGRES_USER=user -e POSTGRES_PASSWORD=password -p 5432:5432 -d postgres
```

Wait a couple of minutes for PostgreSQL to start, and then run the following commands:

```bash
docker exec -it my_postgres psql -U user

CREATE DATABASE db_customers;

\q
```

## Set up the Environment

Clone the repository in your `GOPATH`:

```bash
go env GOPATH
cd {put your GOPATH here}/src
git clone https://github.com/eljandoubi/CRM-Backend.git
cd CRM-Backend
```

Then set up packages and start the server:

```bash
go mod init
go get github.com/lib/pq
go get github.com/gorilla/mux
go run main.go
```

## Test the API

List all customers:

```bash
curl -is http://localhost:3000/customers
```

Add a customer:

```bash
curl -is -X POST -H "Content-Type: application/json" -d '{"ID":4,"Name": "eljandoubi", "Role": "Data Scientist", "Email": "abdel@example.com", "Phone": "123456789", "Contacted": true}' http://localhost:3000/customers
```

Update a customer:

```bash
curl -is -X PUT -H "Content-Type: application/json" -d '{"Name": "eljandoubi", "Role": "Data Scientist", "Email": "abdel@example.com", "Phone": "123456789", "Contacted": false}' http://localhost:3000/customers/4
```

Get a customer:

```bash
curl -is http://localhost:3000/customers/4
```

Delete a customer:

```bash
curl -is -X DELETE http://localhost:3000/customers/4
```
