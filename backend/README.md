# Backend
This is the backend of the Edukita project. Tech stack used: only Go.

Libraries used:
- [UUID](https://pkg.go.dev/github.com/google/uuid)
- [JWT Library](https://pkg.go.dev/github.com/golang-jwt/jwt/v5)

## Requirements
- [Go 1.22 or Latest](https://golang.org/dl/)

## Concepts
This code is using golang best practice to keep everyting simple and clean. All the code is written on Clean Architecture.
### Data Structure
There is 3 entity in this project:
#### Users
```go
type User struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Role     string    `json:"role"`
}
```
#### Assignment
```go
Assignment struct {
		ID        string `json:"id"`
		Subject   string `json:"subject"`
		Title     string `json:"title"`
		Content   string `json:"content"`
		StudentId string `json:"student_id"`
	}
```

#### Grade
```go
Grade struct {
		ID           string `json:"id"`
		AssignmentID string `json:"assignment_id"`
		TeacherID    string `json:"teacher_id"`
		Grade        int    `json:"grade"`
		Feedback     string `json:"feedback"`
	}
```

The relation of this entity is like this:
- User one to many Assignment
- User one to many Grade
- Assignment one to one Grade

This will be support to future development to migrate into relational database. All the ID on that entity is signed using UUID.

### Repository Layer
This layer is used to interact with the database. On this case all the data is stored on hashmap. The ID is used as key, and the entity as the value. This pupose is used to optimize when searching the data.

### Service Layer
This layer is used to interact with the repository layer and also the business logic. This layer is used to keep the code clean and easy to read.

### Handler Layer
This layer is used to interact with the service layer and also the request and response. This layer is used to keep the code clean and easy to read.

## Endpoints
Every endpoint will be return on this structured format:
```json
{
  "status": "success",
  "status_code": 200,
  "message": "Success",
  "data": {}
}
```
`data` will be adjusted based on returned value needed to pass to the frontend.
### POST /login
Used to logged in the user. Only required to pass email because there is no requirement for storing password in secure way.

```bash
curl -X POST http://localhost:8080/login?email=test@gmail.com
```

#### URL Parameter
- email: string

#### Response Object
```json
{
  "status": "success",
  "status_code": 200,
  "message": "Success",
  "data": {
    "token": "",
    "user": {
      "id": "f47ac10b-58cc-0372-8567-0e02b2c3d479",
      "username": "test",
      "email": "test@gmail.com",
      "role": "STUDENT"
    }
  }
}
```

#### Response Object
### POST /users
### POST /assignment
### GET /assignment
### POST /grade
### GET /grade/{student_id}


- POST /user
