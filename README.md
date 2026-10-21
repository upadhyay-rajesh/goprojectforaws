# Employee Management - Gin + GORM + PostgreSQL + GraphQL

Simple version of the Employee Management project.

## Included

- Go
- Gin
- GORM
- PostgreSQL
- GraphQL
- GraphiQL / Playground
- Controller
- Service
- Repository
- Dependency Injection
- Basic validation
- REST + GraphQL

## Not included

- JWT
- Authentication
- Authorization
- Swagger
- OpenAPI
- Pagination
- Filtering
- Sorting

## Architecture

Client
  -> Gin
  -> REST Controller OR GraphQL Resolver
  -> Employee Service
  -> Employee Repository
  -> GORM
  -> PostgreSQL

## Database

```sql
CREATE DATABASE employee_db;
```

Update `.env` with your PostgreSQL password.

## Run

```powershell
go mod tidy
go run .
```

## GraphQL UI

Open:

http://localhost:8080/graphql

or:

http://localhost:8080/playground

There is NO token and NO HTTP header required.

## Query all employees

```graphql
query {
  employees {
    id
    name
    email
    department
    salary
  }
}
```

## Query one employee

```graphql
query {
  employee(id: 1) {
    id
    name
    email
    department
    salary
  }
}
```

## Create

```graphql
mutation {
  createEmployee(input: {
    name: "Amit"
    email: "amit@example.com"
    department: "IT"
    salary: 75000
  }) {
    id
    name
    email
    department
    salary
  }
}
```

## Update

```graphql
mutation {
  updateEmployee(id: 1, input: {
    name: "Amit Updated"
    email: "amit.updated@example.com"
    department: "HR"
    salary: 85000
  }) {
    id
    name
    department
    salary
  }
}
```

## Delete

```graphql
mutation {
  deleteEmployee(id: 1)
}
```

## REST endpoints

GET    /employees
POST   /employees
GET    /employees/:id
PUT    /employees/:id
DELETE /employees/:id

## Test

```powershell
go test ./...
```
