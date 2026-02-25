# CS367-Gin-API
# Student API - Lab 4

## 📌 Description

RESTful API for managing student information developed using Go (Golang) and Gin Framework.

---

## Database

This project uses **SQLite** with the following driver:

```
modernc.org/sqlite
```

(No CGO required)

## How to Run

```bash
go mod tidy
go run main.go
```

Server runs at:

```
http://localhost:8080
```

---

## API Endpoints

by using postman http://localhost:8080 to test API

| Method | Endpoint      | Description       |
| ------ | ------------- | ----------------- |
| GET    | /students     | Get all students  |
| GET    | /students/:id | Get student by ID |
| POST   | /students     | Create student    |
| PUT    | /students/:id | Update student    |
| DELETE | /students/:id | Delete student    |

---

## Example Request (POST)

```json
{
  "id": "6609650350",
  "name": "Natthida Bunsuea",
  "major": "CS",
  "gpa": 2.75
}
```

---

## Student Information

Name: Natthida Bunsuea

Student ID: 6609650350

## API Testing Result

Example testing of each endpoints using Postman.

<img width="1919" height="1079" alt="image" src="https://github.com/user-attachments/assets/2ac43d85-670d-4953-a318-1774eb1ac573" />


