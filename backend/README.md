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

### Authorization
Every request to the backend need to pass the token on the header. Only /users and /login is bypass the token requirements.

```typescript
await fetch("http://localhost:8080/users", {
  method: "GET",
  headers: {
    "Authorization": `Bearer ${token}`
  },
)

```

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
### POST /users
Used to create new user.

### Body Parameter
- email: string
- name: string
- role: "STUDENT" | "TEACHER"

```bash
curl -X POST http://localhost:8080/users  -d '{"email": "test@gmail.com", "name": "test", "role": "STUDENT"}'
```

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

### POST /assignment
Used to create new assignment. Only user with STUDENT role is allowed to use this endpoint.

### Body Parameter
- subject: "MATH" | "ENGLISH"
- title: string
- content: string

```bash
curl -X POST http://localhost:8080/assignment  -H "Authorization: Bearer eyouou..." -d '{"subject": "Math", "title": "Assignment 1", "content": "This is the content"}'
```

#### Response Object
```json
{
  "status": "success",
  "status_code": 200,
  "message": "Success",
  "data": {
    "id": "f47ac10b-58cc-0372-8567-0e02b2c3d479",
    "subject": "Math",
    "title": "Assignment 1",
    "content": "This is the content",
    "student_id": "f47ac10b-58cc-0372-8567-0e02b2c3d479"
  }
}
```
### GET /assignment
Used to get all assignment. Only user with TEACHER role is allowed to use this endpoint.

```bash
curl -X GET http://localhost:8080/assignment  -H "Authorization: Bearer eyouou..."
```

#### Response Object
```json
{
  "status": "success",
  "status_code": 200,
  "message": "Success",
  "data": [
    {
      "id": "f47ac10b-58cc-0372-8567-0e02b2c3d479",
      "subject": "Math",
      "title": "Assignment 1",
      "content": "This is the content",
      "student_id": "f47ac10b-58cc-0372-8567-0e02b2c3d479"
    }
  ]
}
```


### POST /grade
Used to create new grade. Only user with TEACHER role is allowed to use this endpoint.

### Body Parameter
- assignment_id: string
- grade: int
- feedback: string

```bash
curl -X POST http://localhost:8080/grade  -H "Authorization: Bearer eyouou..." -d '{"assignment_id": "f47ac10b-58cc-0372-8567-0e02b2c3d479", "grade": 100, "feedback": "Good Job"}'
```

#### Response Object
```json
{
  "status": "success",
  "status_code": 200,
  "message": "Success",
  "data": {
    "id": "f47ac10b-58cc-0372-8567-0e02b2c3d479",
    "assignment_id": "f47ac10b-58cc-0372-8567-0e02b2c3d479",
    "teacher_id": "f47ac10b-58cc-0372-8567-0e02b2c3d479",
    "grade": 100,
    "feedback": "Good Job"
  }
}
```

### GET /grade/{student_id}
Used to get all grade of the student. Only user with STUDENT role is allowed to use this endpoint.

```bash
curl -X GET http://localhost:8080/grade/f47ac10b-58cc-0372-8567-0e02b2c3d479  -H "Authorization: Bearer eyouou..."
```

#### Response Object
```json
{
  "status": "success",
  "status_code": 200,
  "message": "Success",
  "data": [
    {
      "id": "f47ac10b-58cc-0372-8567-0e02b2c3d479",
      "assignment_id": "f47ac10b-58cc-0372-8567-0e02b2c3d479",
      "teacher_id": "f47ac10b-58cc-0372-8567-0e02b2c3d479",
      "grade": 100,
      "feedback": "Good Job"
    }
  ]
}
```
